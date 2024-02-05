package client

import (
	"bytes"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xescugc/maze-wars/assets"
	"github.com/xescugc/maze-wars/store"
	"github.com/xescugc/maze-wars/unit"
	"github.com/xescugc/maze-wars/utils"
)

type Units struct {
	game *Game

	lifeBarProgress image.Image
	lifeBarUnder    image.Image
}

var (
	directionToTile = map[utils.Direction]int{
		utils.Down:  0,
		utils.Up:    1,
		utils.Left:  2,
		utils.Right: 3,
	}
)

func NewUnits(g *Game) (*Units, error) {
	lbpi, _, err := image.Decode(bytes.NewReader(assets.LifeBarMiniProgress_png))
	if err != nil {
		return nil, err
	}

	lbui, _, err := image.Decode(bytes.NewReader(assets.LifeBarMiniUnder_png))
	if err != nil {
		return nil, err
	}

	us := &Units{
		game:            g,
		lifeBarProgress: ebiten.NewImageFromImage(lbpi),
		lifeBarUnder:    ebiten.NewImageFromImage(lbui),
	}

	return us, nil
}

func (us *Units) Update() error {
	b := time.Now()
	defer utils.LogTime(us.game.Logger, b, "units update")

	cp := us.game.Store.Players.FindCurrent()

	for _, u := range us.game.Store.Units.List() {
		// Only do the events as the owner of the unit if not the actionDispatcher
		// will also dispatch it to the server and the event will be done len(players)
		// amount of times
		if cp.ID == u.PlayerID {
			if u.Health == 0 {
				p := us.game.Store.Players.FindByLineID(u.CurrentLineID)
				actionDispatcher.UnitKilled(p.ID, u.Type)
				actionDispatcher.RemoveUnit(u.ID)
				continue
			}
			if us.game.Store.Map.IsAtTheEnd(u.Object, u.CurrentLineID) {
				p := us.game.Store.Players.FindByLineID(u.CurrentLineID)
				actionDispatcher.StealLive(p.ID, u.PlayerID)
				nlid := us.game.Store.Map.GetNextLineID(u.CurrentLineID)
				if nlid == u.PlayerLineID {
					actionDispatcher.RemoveUnit(u.ID)
				} else {
					actionDispatcher.ChangeUnitLine(u.ID)
				}
			}
		}
	}

	return nil
}

func (us *Units) Draw(screen *ebiten.Image) {
	b := time.Now()
	defer utils.LogTime(us.game.Logger, b, "units draw")

	for _, u := range us.game.Store.Units.List() {
		us.DrawUnit(screen, us.game.Camera, u)
	}
}

func (us *Units) DrawUnit(screen *ebiten.Image, c *CameraStore, u *store.Unit) {
	cs := c.GetState().(CameraState)
	// This is to display the full unit calculated path as a line
	// used for testing visually the path
	//for _, s := range u.Path {
	//screen.Set(int(s.X-cs.X), int(s.Y-cs.Y), color.Black)
	//}
	if !u.IsColliding(cs.Object) {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(u.X-cs.X, u.Y-cs.Y)
	op.GeoM.Scale(cs.Zoom, cs.Zoom)
	sx := directionToTile[u.Facing] * int(u.W)
	i := (u.MovingCount / 5) % 4
	sy := i * int(u.H)
	screen.DrawImage(ebiten.NewImageFromImage(u.Sprite()).SubImage(image.Rect(sx, sy, sx+int(u.W), sy+int(u.H))).(*ebiten.Image), op)

	// Only draw the Health bar if the unit has been hit
	h := unit.Units[u.Type].Health
	if unit.Units[u.Type].Health != u.Health {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(u.X-cs.X, u.Y-cs.Y-float64(us.lifeBarUnder.Bounds().Dy()))
		screen.DrawImage(us.lifeBarUnder.(*ebiten.Image), op)

		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(u.X-cs.X, u.Y-cs.Y-float64(us.lifeBarProgress.Bounds().Dy()))
		screen.DrawImage(us.lifeBarProgress.(*ebiten.Image).SubImage(image.Rect(0, 0, int(float64(us.lifeBarProgress.Bounds().Dx())*(u.Health/h)), us.lifeBarProgress.Bounds().Dy())).(*ebiten.Image), op)
	}
}
