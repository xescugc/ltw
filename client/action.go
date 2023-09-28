package main

import (
	"github.com/xescugc/go-flux"
	"github.com/xescugc/ltw/action"
)

// ActionDispatcher is in charge of dispatching actions to the
// application dispatcher
type ActionDispatcher struct {
	dispatcher *flux.Dispatcher
}

// NewActionDispatcher initializes the action dispatcher
// with the give dispatcher
func NewActionDispatcher(d *flux.Dispatcher) *ActionDispatcher {
	return &ActionDispatcher{
		dispatcher: d,
	}
}

// Dispatch is a helper to access to the internal dispatch directly with an action.
// This should only be used from the WS Handler to forward server actions directly
func (ac *ActionDispatcher) Dispatch(a *action.Action) {
	ac.dispatcher.Dispatch(a)
}

// CursorMove dispatches an action of moving the Cursor
// to the new x,y coordinates
func (ac *ActionDispatcher) CursorMove(x, y int) {
	cma := action.NewCursorMove(x, y)
	ac.dispatcher.Dispatch(cma)
}

// SummonUnit summons the 'unit' from the player id 'pid' to the line
// 'plid' and with the current line id 'clid'
func (ac *ActionDispatcher) SummonUnit(unit, pid string, plid, clid int) {
	sua := action.NewSummonUnit(unit, pid, plid, clid)
	wsSend(sua)
	//ac.dispatcher.Dispatch(sua)
}

// MoveUnit moves all the units
func (ac *ActionDispatcher) MoveUnit() {
	mua := action.NewMoveUnit()
	ac.dispatcher.Dispatch(mua)
}

// RemoveUnit removes the unit with the id 'uid'
func (ac *ActionDispatcher) RemoveUnit(uid string) {
	rua := action.NewRemoveUnit(uid)
	wsSend(rua)
	ac.dispatcher.Dispatch(rua)
}

// StealLive removes one live from the player with id 'fpid' and
// adds it to the player with id 'tpid'
func (ac *ActionDispatcher) StealLive(fpid, tpid string) {
	sla := action.NewStealLive(fpid, tpid)
	wsSend(sla)
	ac.dispatcher.Dispatch(sla)
}

// CameraZoom zooms the camera the direction 'd'
func (ac *ActionDispatcher) CameraZoom(d int) {
	cza := action.NewCameraZoom(d)
	ac.dispatcher.Dispatch(cza)
}

// PlaceTower places the tower 't' on the position X and Y of the player pid
func (ac *ActionDispatcher) PlaceTower(t, pid string, x, y int) {
	bta := action.NewPlaceTower(t, pid, x, y)
	wsSend(bta)
	ac.dispatcher.Dispatch(bta)
}

// SelectTower selects the tower 't' on the position x, y
func (ac *ActionDispatcher) SelectTower(t string, x, y int) {
	sta := action.NewSelectTower(t, x, y)
	ac.dispatcher.Dispatch(sta)
}

// SelectTower selects the tower 't' on the position x, y
func (ac *ActionDispatcher) SelectedTowerInvalid(i bool) {
	sta := action.NewSelectedTowerInvalid(i)
	ac.dispatcher.Dispatch(sta)
}

// DeelectTower cleans the current selected tower
func (ac *ActionDispatcher) DeselectTower(t string) {
	dsta := action.NewDeselectTower(t)
	ac.dispatcher.Dispatch(dsta)
}

// IncomeTick a new tick for the income
func (ac *ActionDispatcher) IncomeTick() {
	it := action.NewIncomeTick()
	ac.dispatcher.Dispatch(it)
}

// TowerAttack issues a attack to the Unit with uid
func (ac *ActionDispatcher) TowerAttack(uid, tt string) {
	ta := action.NewTowerAttack(uid, tt)
	ac.dispatcher.Dispatch(ta)
}

// UnitKilled adds gold to the user
func (ac *ActionDispatcher) UnitKilled(pid, ut string) {
	uk := action.NewUnitKilled(pid, ut)
	wsSend(uk)
	ac.dispatcher.Dispatch(uk)
}

// WindowResizing new sizes of the window
func (ac *ActionDispatcher) WindowResizing(w, h int) {
	wr := action.NewWindowResizing(w, h)
	ac.dispatcher.Dispatch(wr)
}

// JoinRoom new sizes of the window
func (ac *ActionDispatcher) JoinRoom(room, name string) {
	jr := action.NewJoinRoom(room, name)
	wsSend(jr)
	ac.dispatcher.Dispatch(jr)
}
