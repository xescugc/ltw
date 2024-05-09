// Code generated by "enumer -type=Type -transform=snake -output=type_string.go -json"; DO NOT EDIT.

package action

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TypeName = "cursor_movecamera_zoomsummon_unitupdate_unitupdate_towertpsremove_unitsteal_liveplace_towerremove_towerselect_towerselected_towerselected_tower_invaliddeselect_towerincome_ticktower_attackunit_killedwindow_resizingnavigate_tostart_gameopen_tower_menuclose_tower_menugo_homechange_unit_linesign_up_erroruser_sign_upuser_sign_inuser_sign_outjoin_waiting_roomexit_waiting_roomstart_roomtoggle_statsversion_errorcreate_lobbydelete_lobbyjoin_lobbyadd_lobbiesselect_lobbyleave_lobbyupdate_lobbystart_lobbyadd_playerremove_playersync_statesync_userswait_room_countdown_ticksync_waiting_room"

var _TypeIndex = [...]uint16{0, 11, 22, 33, 44, 56, 59, 70, 80, 91, 103, 115, 129, 151, 165, 176, 188, 199, 214, 225, 235, 250, 266, 273, 289, 302, 314, 326, 339, 356, 373, 383, 395, 408, 420, 432, 442, 453, 465, 476, 488, 499, 509, 522, 532, 542, 566, 583}

const _TypeLowerName = "cursor_movecamera_zoomsummon_unitupdate_unitupdate_towertpsremove_unitsteal_liveplace_towerremove_towerselect_towerselected_towerselected_tower_invaliddeselect_towerincome_ticktower_attackunit_killedwindow_resizingnavigate_tostart_gameopen_tower_menuclose_tower_menugo_homechange_unit_linesign_up_erroruser_sign_upuser_sign_inuser_sign_outjoin_waiting_roomexit_waiting_roomstart_roomtoggle_statsversion_errorcreate_lobbydelete_lobbyjoin_lobbyadd_lobbiesselect_lobbyleave_lobbyupdate_lobbystart_lobbyadd_playerremove_playersync_statesync_userswait_room_countdown_ticksync_waiting_room"

func (i Type) String() string {
	if i < 0 || i >= Type(len(_TypeIndex)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _TypeName[_TypeIndex[i]:_TypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TypeNoOp() {
	var x [1]struct{}
	_ = x[CursorMove-(0)]
	_ = x[CameraZoom-(1)]
	_ = x[SummonUnit-(2)]
	_ = x[UpdateUnit-(3)]
	_ = x[UpdateTower-(4)]
	_ = x[TPS-(5)]
	_ = x[RemoveUnit-(6)]
	_ = x[StealLive-(7)]
	_ = x[PlaceTower-(8)]
	_ = x[RemoveTower-(9)]
	_ = x[SelectTower-(10)]
	_ = x[SelectedTower-(11)]
	_ = x[SelectedTowerInvalid-(12)]
	_ = x[DeselectTower-(13)]
	_ = x[IncomeTick-(14)]
	_ = x[TowerAttack-(15)]
	_ = x[UnitKilled-(16)]
	_ = x[WindowResizing-(17)]
	_ = x[NavigateTo-(18)]
	_ = x[StartGame-(19)]
	_ = x[OpenTowerMenu-(20)]
	_ = x[CloseTowerMenu-(21)]
	_ = x[GoHome-(22)]
	_ = x[ChangeUnitLine-(23)]
	_ = x[SignUpError-(24)]
	_ = x[UserSignUp-(25)]
	_ = x[UserSignIn-(26)]
	_ = x[UserSignOut-(27)]
	_ = x[JoinWaitingRoom-(28)]
	_ = x[ExitWaitingRoom-(29)]
	_ = x[StartRoom-(30)]
	_ = x[ToggleStats-(31)]
	_ = x[VersionError-(32)]
	_ = x[CreateLobby-(33)]
	_ = x[DeleteLobby-(34)]
	_ = x[JoinLobby-(35)]
	_ = x[AddLobbies-(36)]
	_ = x[SelectLobby-(37)]
	_ = x[LeaveLobby-(38)]
	_ = x[UpdateLobby-(39)]
	_ = x[StartLobby-(40)]
	_ = x[AddPlayer-(41)]
	_ = x[RemovePlayer-(42)]
	_ = x[SyncState-(43)]
	_ = x[SyncUsers-(44)]
	_ = x[WaitRoomCountdownTick-(45)]
	_ = x[SyncWaitingRoom-(46)]
}

var _TypeValues = []Type{CursorMove, CameraZoom, SummonUnit, UpdateUnit, UpdateTower, TPS, RemoveUnit, StealLive, PlaceTower, RemoveTower, SelectTower, SelectedTower, SelectedTowerInvalid, DeselectTower, IncomeTick, TowerAttack, UnitKilled, WindowResizing, NavigateTo, StartGame, OpenTowerMenu, CloseTowerMenu, GoHome, ChangeUnitLine, SignUpError, UserSignUp, UserSignIn, UserSignOut, JoinWaitingRoom, ExitWaitingRoom, StartRoom, ToggleStats, VersionError, CreateLobby, DeleteLobby, JoinLobby, AddLobbies, SelectLobby, LeaveLobby, UpdateLobby, StartLobby, AddPlayer, RemovePlayer, SyncState, SyncUsers, WaitRoomCountdownTick, SyncWaitingRoom}

var _TypeNameToValueMap = map[string]Type{
	_TypeName[0:11]:         CursorMove,
	_TypeLowerName[0:11]:    CursorMove,
	_TypeName[11:22]:        CameraZoom,
	_TypeLowerName[11:22]:   CameraZoom,
	_TypeName[22:33]:        SummonUnit,
	_TypeLowerName[22:33]:   SummonUnit,
	_TypeName[33:44]:        UpdateUnit,
	_TypeLowerName[33:44]:   UpdateUnit,
	_TypeName[44:56]:        UpdateTower,
	_TypeLowerName[44:56]:   UpdateTower,
	_TypeName[56:59]:        TPS,
	_TypeLowerName[56:59]:   TPS,
	_TypeName[59:70]:        RemoveUnit,
	_TypeLowerName[59:70]:   RemoveUnit,
	_TypeName[70:80]:        StealLive,
	_TypeLowerName[70:80]:   StealLive,
	_TypeName[80:91]:        PlaceTower,
	_TypeLowerName[80:91]:   PlaceTower,
	_TypeName[91:103]:       RemoveTower,
	_TypeLowerName[91:103]:  RemoveTower,
	_TypeName[103:115]:      SelectTower,
	_TypeLowerName[103:115]: SelectTower,
	_TypeName[115:129]:      SelectedTower,
	_TypeLowerName[115:129]: SelectedTower,
	_TypeName[129:151]:      SelectedTowerInvalid,
	_TypeLowerName[129:151]: SelectedTowerInvalid,
	_TypeName[151:165]:      DeselectTower,
	_TypeLowerName[151:165]: DeselectTower,
	_TypeName[165:176]:      IncomeTick,
	_TypeLowerName[165:176]: IncomeTick,
	_TypeName[176:188]:      TowerAttack,
	_TypeLowerName[176:188]: TowerAttack,
	_TypeName[188:199]:      UnitKilled,
	_TypeLowerName[188:199]: UnitKilled,
	_TypeName[199:214]:      WindowResizing,
	_TypeLowerName[199:214]: WindowResizing,
	_TypeName[214:225]:      NavigateTo,
	_TypeLowerName[214:225]: NavigateTo,
	_TypeName[225:235]:      StartGame,
	_TypeLowerName[225:235]: StartGame,
	_TypeName[235:250]:      OpenTowerMenu,
	_TypeLowerName[235:250]: OpenTowerMenu,
	_TypeName[250:266]:      CloseTowerMenu,
	_TypeLowerName[250:266]: CloseTowerMenu,
	_TypeName[266:273]:      GoHome,
	_TypeLowerName[266:273]: GoHome,
	_TypeName[273:289]:      ChangeUnitLine,
	_TypeLowerName[273:289]: ChangeUnitLine,
	_TypeName[289:302]:      SignUpError,
	_TypeLowerName[289:302]: SignUpError,
	_TypeName[302:314]:      UserSignUp,
	_TypeLowerName[302:314]: UserSignUp,
	_TypeName[314:326]:      UserSignIn,
	_TypeLowerName[314:326]: UserSignIn,
	_TypeName[326:339]:      UserSignOut,
	_TypeLowerName[326:339]: UserSignOut,
	_TypeName[339:356]:      JoinWaitingRoom,
	_TypeLowerName[339:356]: JoinWaitingRoom,
	_TypeName[356:373]:      ExitWaitingRoom,
	_TypeLowerName[356:373]: ExitWaitingRoom,
	_TypeName[373:383]:      StartRoom,
	_TypeLowerName[373:383]: StartRoom,
	_TypeName[383:395]:      ToggleStats,
	_TypeLowerName[383:395]: ToggleStats,
	_TypeName[395:408]:      VersionError,
	_TypeLowerName[395:408]: VersionError,
	_TypeName[408:420]:      CreateLobby,
	_TypeLowerName[408:420]: CreateLobby,
	_TypeName[420:432]:      DeleteLobby,
	_TypeLowerName[420:432]: DeleteLobby,
	_TypeName[432:442]:      JoinLobby,
	_TypeLowerName[432:442]: JoinLobby,
	_TypeName[442:453]:      AddLobbies,
	_TypeLowerName[442:453]: AddLobbies,
	_TypeName[453:465]:      SelectLobby,
	_TypeLowerName[453:465]: SelectLobby,
	_TypeName[465:476]:      LeaveLobby,
	_TypeLowerName[465:476]: LeaveLobby,
	_TypeName[476:488]:      UpdateLobby,
	_TypeLowerName[476:488]: UpdateLobby,
	_TypeName[488:499]:      StartLobby,
	_TypeLowerName[488:499]: StartLobby,
	_TypeName[499:509]:      AddPlayer,
	_TypeLowerName[499:509]: AddPlayer,
	_TypeName[509:522]:      RemovePlayer,
	_TypeLowerName[509:522]: RemovePlayer,
	_TypeName[522:532]:      SyncState,
	_TypeLowerName[522:532]: SyncState,
	_TypeName[532:542]:      SyncUsers,
	_TypeLowerName[532:542]: SyncUsers,
	_TypeName[542:566]:      WaitRoomCountdownTick,
	_TypeLowerName[542:566]: WaitRoomCountdownTick,
	_TypeName[566:583]:      SyncWaitingRoom,
	_TypeLowerName[566:583]: SyncWaitingRoom,
}

var _TypeNames = []string{
	_TypeName[0:11],
	_TypeName[11:22],
	_TypeName[22:33],
	_TypeName[33:44],
	_TypeName[44:56],
	_TypeName[56:59],
	_TypeName[59:70],
	_TypeName[70:80],
	_TypeName[80:91],
	_TypeName[91:103],
	_TypeName[103:115],
	_TypeName[115:129],
	_TypeName[129:151],
	_TypeName[151:165],
	_TypeName[165:176],
	_TypeName[176:188],
	_TypeName[188:199],
	_TypeName[199:214],
	_TypeName[214:225],
	_TypeName[225:235],
	_TypeName[235:250],
	_TypeName[250:266],
	_TypeName[266:273],
	_TypeName[273:289],
	_TypeName[289:302],
	_TypeName[302:314],
	_TypeName[314:326],
	_TypeName[326:339],
	_TypeName[339:356],
	_TypeName[356:373],
	_TypeName[373:383],
	_TypeName[383:395],
	_TypeName[395:408],
	_TypeName[408:420],
	_TypeName[420:432],
	_TypeName[432:442],
	_TypeName[442:453],
	_TypeName[453:465],
	_TypeName[465:476],
	_TypeName[476:488],
	_TypeName[488:499],
	_TypeName[499:509],
	_TypeName[509:522],
	_TypeName[522:532],
	_TypeName[532:542],
	_TypeName[542:566],
	_TypeName[566:583],
}

// TypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TypeString(s string) (Type, error) {
	if val, ok := _TypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Type values", s)
}

// TypeValues returns all values of the enum
func TypeValues() []Type {
	return _TypeValues
}

// TypeStrings returns a slice of all String values of the enum
func TypeStrings() []string {
	strs := make([]string, len(_TypeNames))
	copy(strs, _TypeNames)
	return strs
}

// IsAType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Type) IsAType() bool {
	for _, v := range _TypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Type
func (i Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Type
func (i *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Type should be a string, got %s", data)
	}

	var err error
	*i, err = TypeString(s)
	return err
}
