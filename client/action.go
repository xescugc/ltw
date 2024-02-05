package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/xescugc/go-flux"
	"github.com/xescugc/maze-wars/action"
	"github.com/xescugc/maze-wars/store"
	"github.com/xescugc/maze-wars/utils"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// ActionDispatcher is in charge of dispatching actions to the
// application dispatcher
type ActionDispatcher struct {
	dispatcher *flux.Dispatcher
	opt        Options
	store      *store.Store
	logger     *slog.Logger
}

// NewActionDispatcher initializes the action dispatcher
// with the give dispatcher
func NewActionDispatcher(d *flux.Dispatcher, s *store.Store, l *slog.Logger, opt Options) *ActionDispatcher {
	return &ActionDispatcher{
		dispatcher: d,
		opt:        opt,
		store:      s,
		logger:     l,
	}
}

// Dispatch is a helper to access to the internal dispatch directly with an action.
// This should only be used from the WS Handler to forward server actions directly
func (ac *ActionDispatcher) Dispatch(a *action.Action) {
	b := time.Now()
	ac.dispatcher.Dispatch(a)
	d := time.Now().Sub(b)
	if d > time.Millisecond {
		ac.logger.Info("action dispatched", "type", a.Type, "time", d)
	}
	ac.logger.Debug("action dispatched", "type", a.Type, "time", d)
}

// CursorMove dispatches an action of moving the Cursor
// to the new x,y coordinates
func (ac *ActionDispatcher) CursorMove(x, y int) {
	cma := action.NewCursorMove(x, y)
	ac.Dispatch(cma)
}

// SummonUnit summons the 'unit' from the player id 'pid' to the line
// 'plid' and with the current line id 'clid'
func (ac *ActionDispatcher) SummonUnit(unit, pid string, plid, clid int) {
	sua := action.NewSummonUnit(unit, pid, plid, clid)
	wsSend(sua)
	//ac.Dispatch(sua)
}

// TPS is the call for every TPS event
func (ac *ActionDispatcher) TPS() {
	tpsa := action.NewTPS()
	ac.Dispatch(tpsa)
}

// RemoveUnit removes the unit with the id 'uid'
func (ac *ActionDispatcher) RemoveUnit(uid string) {
	rua := action.NewRemoveUnit(uid)
	wsSend(rua)
	ac.Dispatch(rua)
}

// StealLive removes one live from the player with id 'fpid' and
// adds it to the player with id 'tpid'
func (ac *ActionDispatcher) StealLive(fpid, tpid string) {
	sla := action.NewStealLive(fpid, tpid)
	wsSend(sla)
	ac.Dispatch(sla)
}

// CameraZoom zooms the camera the direction 'd'
func (ac *ActionDispatcher) CameraZoom(d int) {
	cza := action.NewCameraZoom(d)
	ac.Dispatch(cza)
}

// PlaceTower places the tower 't' on the position X and Y of the player pid
func (ac *ActionDispatcher) PlaceTower(t, pid string, x, y int) {
	units := ac.store.Units.List()
	to := utils.Object{X: float64(x), Y: float64(y), H: 32, W: 32}
	canPlace := true
	for _, u := range units {
		if u.IsColliding(to) {
			canPlace = false
			break
		}
	}
	if canPlace {
		pta := action.NewPlaceTower(t, pid, x, y)
		wsSend(pta)
	}
	//ac.Dispatch(pta)
}

// RemoveTower removes the tower tid
func (ac *ActionDispatcher) RemoveTower(pid, tid, tt string) {
	rta := action.NewRemoveTower(pid, tid, tt)
	wsSend(rta)
	//ac.Dispatch(rta)
}

// SelectTower selects the tower 't' on the position x, y
func (ac *ActionDispatcher) SelectTower(t string, x, y int) {
	sta := action.NewSelectTower(t, x, y)
	ac.Dispatch(sta)
}

// SelectTower selects the tower 't' on the position x, y
func (ac *ActionDispatcher) SelectedTowerInvalid(i bool) {
	sta := action.NewSelectedTowerInvalid(i)
	ac.Dispatch(sta)
}

// DeelectTower cleans the current selected tower
func (ac *ActionDispatcher) DeselectTower(t string) {
	dsta := action.NewDeselectTower(t)
	ac.Dispatch(dsta)
}

// IncomeTick a new tick for the income
func (ac *ActionDispatcher) IncomeTick() {
	it := action.NewIncomeTick()
	ac.Dispatch(it)
}

// TowerAttack issues a attack to the Unit with uid
func (ac *ActionDispatcher) TowerAttack(uid, tt string) {
	ta := action.NewTowerAttack(uid, tt)
	wsSend(ta)
	ac.Dispatch(ta)
}

// UnitKilled adds gold to the user
func (ac *ActionDispatcher) UnitKilled(pid, ut string) {
	uk := action.NewUnitKilled(pid, ut)
	wsSend(uk)
	ac.Dispatch(uk)
}

func (ac *ActionDispatcher) RemovePlayer(pid string) {
	rpa := action.NewRemovePlayer(pid)
	wsSend(rpa)
	ac.Dispatch(rpa)
	ac.Dispatch(action.NewNavigateTo(LobbyRoute))
}

// WindowResizing new sizes of the window
func (ac *ActionDispatcher) WindowResizing(w, h int) {
	wr := action.NewWindowResizing(w, h)
	ac.Dispatch(wr)
}

// NavigateTo navigates to the given route
func (ac *ActionDispatcher) NavigateTo(route string) {
	nt := action.NewNavigateTo(route)
	ac.Dispatch(nt)
}

// OpenTowerMenu when a tower is clicked and the menu of
// the tower is displayed
func (ac *ActionDispatcher) OpenTowerMenu(tid string) {
	otm := action.NewOpenTowerMenu(tid)
	ac.Dispatch(otm)
}

// CloseTowerMenu when a tower menu needs to be closed
func (ac *ActionDispatcher) CloseTowerMenu() {
	ctm := action.NewCloseTowerMenu()
	ac.Dispatch(ctm)
}

// GoHome will move the camera to the current player home line
func (ac *ActionDispatcher) GoHome() {
	gha := action.NewGoHome()
	ac.Dispatch(gha)
}

// CheckedPath will set the value of the path checked
func (ac *ActionDispatcher) CheckedPath(cp bool) {
	cpa := action.NewCheckedPath(cp)
	ac.Dispatch(cpa)
}

// ChangeUnitLine will move the unit to the next line
func (ac *ActionDispatcher) ChangeUnitLine(uid string) {
	cula := action.NewChangeUnitLine(uid)
	wsSend(cula)
	ac.Dispatch(cula)
}

func (ac *ActionDispatcher) SignUpSubmit(un string) {
	httpu, _ := url.Parse(ac.opt.HostURL)
	httpu.Path = "/users"
	resp, err := http.Post(httpu.String(), "application/json", bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":"%s"}`, un))))
	if err != nil {
		ac.Dispatch(action.NewSignUpError(err.Error()))
		return
	}
	body := struct {
		Error string `json:"error"`
	}{}
	if resp.StatusCode != http.StatusCreated {
		err = json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			ac.Dispatch(action.NewSignUpError(err.Error()))
			return
		}
		ac.Dispatch(action.NewSignUpError(body.Error))
		return
	}

	ac.Dispatch(action.NewSignUpError(""))

	ctx := context.Background()

	// We manually clone it to then change it for the WS
	chttpu := *httpu
	wsu := &chttpu
	wsu.Scheme = "ws"
	if httpu.Scheme == "https" {
		wsu.Scheme = "wss"
	}
	wsu.Path = "/ws"

	wsc, _, err = websocket.Dial(ctx, wsu.String(), nil)
	if err != nil {
		panic(fmt.Errorf("failed to dial the server %q: %w", wsu.String(), err))
	}

	wsc.SetReadLimit(-1)

	usia := action.NewUserSignIn(un)
	err = wsjson.Write(ctx, wsc, usia)
	if err != nil {
		panic(fmt.Errorf("failed to write JSON: %w", err))
	}

	ac.Dispatch(usia)

	go wsHandler(ctx)

	ac.Dispatch(action.NewNavigateTo(LobbyRoute))
}

func (ac *ActionDispatcher) JoinWaitingRoom(un string) {
	jwra := action.NewJoinWaitingRoom(un)
	wsSend(jwra)

	ac.Dispatch(action.NewNavigateTo(WaitingRoomRoute))
}

func (ac *ActionDispatcher) ExitWaitingRoom(un string) {
	ewra := action.NewExitWaitingRoom(un)
	wsSend(ewra)

	ac.Dispatch(action.NewNavigateTo(LobbyRoute))
}

func (ac *ActionDispatcher) ToggleStats() {
	tsa := action.NewToggleStats()

	ac.Dispatch(tsa)
}
