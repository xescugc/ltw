package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/xescugc/go-flux"
	"github.com/xescugc/ltw/action"
	"github.com/xescugc/ltw/assets"
	"github.com/xescugc/ltw/store"
	"github.com/xescugc/ltw/utils"
)

// HUDStore is in charge of keeping track of all the elements
// on the player HUD that are static and always seen
type HUDStore struct {
	*flux.ReduceStore

	game *Game

	cyclopeFacesetImage image.Image
	tilesetHouseImage   image.Image
}

// HUDState stores the HUD state
type HUDState struct {
	CyclopeButton utils.Object
	SoldierButton utils.Object

	SelectedTower *SelectedTower

	LastCursorPosition utils.Object
}

type SelectedTower struct {
	store.Tower

	Invalid bool
}

// NewHUDStore creates a new HUDStore with the Dispatcher d and the Game g
func NewHUDStore(d *flux.Dispatcher, g *Game) (*HUDStore, error) {
	fi, _, err := image.Decode(bytes.NewReader(assets.CyclopeFaceset_png))
	if err != nil {
		return nil, err
	}

	thi, _, err := image.Decode(bytes.NewReader(assets.TilesetHouse_png))
	if err != nil {
		return nil, err
	}

	hs := &HUDStore{
		game: g,

		cyclopeFacesetImage: ebiten.NewImageFromImage(fi),
		tilesetHouseImage:   ebiten.NewImageFromImage(thi).SubImage(image.Rect(5*16, 17*16, 5*16+16*2, 17*16+16*2)),
	}
	cs := g.Camera.GetState().(CameraState)
	hs.ReduceStore = flux.NewReduceStore(d, hs.Reduce, HUDState{
		CyclopeButton: utils.Object{
			X: float64(cs.W - float64(hs.cyclopeFacesetImage.Bounds().Dx())),
			Y: float64(cs.H - float64(hs.cyclopeFacesetImage.Bounds().Dy())),
			W: float64(hs.cyclopeFacesetImage.Bounds().Dx()),
			H: float64(hs.cyclopeFacesetImage.Bounds().Dy()),
		},
		SoldierButton: utils.Object{
			X: 0,
			Y: float64(cs.H - 16*2),
			W: float64(16 * 2),
			H: float64(16 * 2),
		},
	})

	return hs, nil
}

func (hs *HUDStore) Update() error {
	hst := hs.GetState().(HUDState)
	x, y := ebiten.CursorPosition()
	cp := hs.game.Players.GetCurrentPlayer()
	// Only send a CursorMove when the curso has actually moved
	if hst.LastCursorPosition.X != float64(x) || hst.LastCursorPosition.Y != float64(y) {
		actionDispatcher.CursorMove(x, y)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		obj := utils.Object{
			X: float64(x),
			Y: float64(y),
			W: 1, H: 1,
		}
		// Check what the user has just clicked
		if cp.Gold >= unitGold && hst.CyclopeButton.IsColliding(obj) {
			actionDispatcher.SummonUnit("cyclope", cp.ID, cp.LineID, hs.game.Map.GetNextLineID(cp.LineID))
			return nil
		}
		if cp.Gold >= towerGold && hst.SoldierButton.IsColliding(obj) {
			actionDispatcher.SelectTower("soldier", x, y)
			return nil
		}

		if hst.SelectedTower != nil && !hst.SelectedTower.Invalid {
			cs := hs.game.Camera.GetState().(CameraState)
			actionDispatcher.PlaceTower(hst.SelectedTower.Type, cp.ID, int(hst.SelectedTower.X+cs.X), int(hst.SelectedTower.Y+cs.Y))
		}
	}
	if cp.Gold >= towerGold && inpututil.IsKeyJustPressed(ebiten.KeyT) {
		actionDispatcher.SelectTower("soldier", x, y)
		return nil
	}
	if hst.SelectedTower != nil {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			actionDispatcher.DeselectTower(hst.SelectedTower.Type)
		} else {
			ts := hs.game.Store.Towers.GetState().(store.TowersState)
			cs := hs.game.Camera.GetState().(CameraState)
			var invalid bool
			for _, t := range ts.Towers {
				// The t.Object has the X and Y relative to the map
				// and the hst.SelectedTower has them relative to the
				// screen so we need to port the t.Object to the same
				// relative values
				neo := t.Object
				neo.X -= cs.X
				neo.Y -= cs.Y
				if hst.SelectedTower.IsColliding(neo) {
					invalid = true
					break
				}
			}
			neo := hst.SelectedTower.Object
			neo.X += cs.X
			neo.Y += cs.Y
			if !hs.game.Map.IsInValidBuildingZone(neo, hst.SelectedTower.LineID) {
				invalid = true
			}
			if invalid != hst.SelectedTower.Invalid {
				actionDispatcher.SelectedTowerInvalid(invalid)
			}
		}
	}

	return nil
}

func (hs *HUDStore) Draw(screen *ebiten.Image) {
	hst := hs.GetState().(HUDState)
	cs := hs.game.Camera.GetState().(CameraState)
	cp := hs.game.Players.GetCurrentPlayer()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(hst.CyclopeButton.X, hst.CyclopeButton.Y)
	if cp.Gold < unitGold {
		op.ColorM.Scale(2, 0.5, 0.5, 0.9)
	}
	screen.DrawImage(hs.cyclopeFacesetImage.(*ebiten.Image), op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(hst.SoldierButton.X, hst.SoldierButton.Y)
	if cp.Gold < towerGold {
		op.ColorM.Scale(2, 0.5, 0.5, 0.9)
	} else if hst.SelectedTower != nil && hst.SelectedTower.Type == "soldier" {
		// Once the tower is selected we gray it out
		op.ColorM.Scale(0.5, 0.5, 0.5, 0.5)
	}
	screen.DrawImage(hs.tilesetHouseImage.(*ebiten.Image), op)

	if hst.SelectedTower != nil {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(hst.SelectedTower.X/cs.Zoom, hst.SelectedTower.Y/cs.Zoom)
		op.GeoM.Scale(cs.Zoom, cs.Zoom)

		if hst.SelectedTower != nil && hst.SelectedTower.Invalid {
			op.ColorM.Scale(2, 0.5, 0.5, 0.9)
		}

		screen.DrawImage(hst.SelectedTower.Image().(*ebiten.Image), op)
	}

	ps := hs.game.Store.Players.GetState().(store.PlayersState)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Income Timer: %ds", ps.IncomeTimer), 0, 0)
	var pcount = 1
	var sortedPlayers = make([]*store.Player, 0, 0)
	for _, p := range ps.Players {
		sortedPlayers = append(sortedPlayers, p)
	}
	sort.Slice(sortedPlayers, func(i, j int) bool { return sortedPlayers[i].LineID < sortedPlayers[j].LineID })
	for _, p := range sortedPlayers {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Name: %s, Lives: %d, Gold: %d, Income: %d", p.Name, p.Lives, p.Gold, p.Income), 0, 15*pcount)
		pcount++
	}
	if verbose {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("(X: %d, Y: %d)", int(hst.LastCursorPosition.X+cs.X), int(hst.LastCursorPosition.Y+cs.Y)), 0, 15*(pcount))
	}
}

func (hs *HUDStore) Reduce(state, a interface{}) interface{} {
	act, ok := a.(*action.Action)
	if !ok {
		return state
	}

	hstate, ok := state.(HUDState)
	if !ok {
		return state
	}

	switch act.Type {
	case action.WindowResizing:
		hs.GetDispatcher().WaitFor(hs.game.Camera.GetDispatcherToken())
		cs := hs.game.Camera.GetState().(CameraState)
		hstate.CyclopeButton = utils.Object{
			X: float64(cs.W - float64(hs.cyclopeFacesetImage.Bounds().Dx())),
			Y: float64(cs.H - float64(hs.cyclopeFacesetImage.Bounds().Dy())),
			W: float64(hs.cyclopeFacesetImage.Bounds().Dx()),
			H: float64(hs.cyclopeFacesetImage.Bounds().Dy()),
		}
		hstate.SoldierButton = utils.Object{
			X: 0,
			Y: float64(cs.H - 16*2),
			W: float64(16 * 2),
			H: float64(16 * 2),
		}
	case action.SelectTower:
		hs.GetDispatcher().WaitFor(hs.game.Store.Players.GetDispatcherToken())
		cp := hs.game.Players.GetCurrentPlayer()
		// TODO: Insead of hardcoding the image and W, H we should
		// use the Type on the action to select the right image
		hstate.SelectedTower = &SelectedTower{
			Tower: store.Tower{
				Object: utils.Object{
					X: float64(act.SelectTower.X) - (hstate.SoldierButton.W / 2),
					Y: float64(act.SelectTower.Y) - (hstate.SoldierButton.H / 2),
					W: hstate.SoldierButton.W,
					H: hstate.SoldierButton.H,
				},
				Type:   act.SelectTower.Type,
				LineID: cp.LineID,
			},
		}
	case action.CursorMove:
		// We update the last seen cursor position to not resend unnecessary events
		hstate.LastCursorPosition.X = float64(act.CursorMove.X)
		hstate.LastCursorPosition.Y = float64(act.CursorMove.Y)

		if hstate.SelectedTower != nil {
			// We find the closes multiple in case the cursor moves too fast, between FPS reloads,
			// and lands in a position not 'multiple' which means the position of the SelectedTower
			// is not updated and the result is the cursor far away from the Drawing of the SelectedTower
			// as it has stayed on the previous position
			var multiple int = 8
			if act.CursorMove.X%multiple == 0 {
				hstate.SelectedTower.X = float64(act.CursorMove.X) - (hstate.SoldierButton.W / 2)
			} else if math.Abs(float64(act.CursorMove.X)-hstate.SelectedTower.X) > float64(multiple) {
				hstate.SelectedTower.X = float64(closestMultiple(act.CursorMove.X, multiple)) - (hstate.SoldierButton.W / 2)
			}
			if act.CursorMove.Y%multiple == 0 {
				hstate.SelectedTower.Y = float64(act.CursorMove.Y) - (hstate.SoldierButton.H / 2)
			} else if math.Abs(float64(act.CursorMove.Y)-hstate.SelectedTower.Y) > float64(multiple) {
				hstate.SelectedTower.Y = float64(closestMultiple(act.CursorMove.Y, multiple)) - (hstate.SoldierButton.H / 2)
			}
		}
	case action.PlaceTower, action.DeselectTower:
		hstate.SelectedTower = nil
	case action.SelectedTowerInvalid:
		if hstate.SelectedTower != nil {
			hstate.SelectedTower.Invalid = act.SelectedTowerInvalid.Invalid
		}
	default:
	}

	return hstate
}

// closestMultiple finds the coses multiple of 'b' for the number 'a'
func closestMultiple(a, b int) int {
	a = a + b/2
	a = a - (a % b)
	return a
}
