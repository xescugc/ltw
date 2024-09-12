// Code generated by "enumer -type=Type -transform=snake -output=type_string.go -json"; DO NOT EDIT.

package action

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TypeName = "cursor_movecamera_zoomsummon_unitupdate_unitupdate_towertpsplace_towerremove_towerselect_towerselected_towerselected_tower_invaliddeselect_towerincome_tickwindow_resizingnavigate_tostart_gameopen_tower_menuopen_unit_menuclose_tower_menuclose_unit_menugo_homesign_up_erroruser_sign_upuser_sign_inuser_sign_outjoin_vs6_waiting_roomexit_vs6_waiting_roomjoin_vs1_waiting_roomexit_vs1_waiting_roomstart_roomversion_errorsetup_gamefind_gameexit_searching_gameaccept_waiting_gamecancel_waiting_gameshow_scoreboardcreate_lobbydelete_lobbyjoin_lobbyadd_lobbiesselect_lobbyleave_lobbyupdate_lobbystart_lobbyseen_lobbiesadd_playerremove_playersync_statewait_room_countdown_ticksync_vs6_waiting_roomsync_vs1_waiting_roomsync_searching_roomsync_waiting_room"

var _TypeIndex = [...]uint16{0, 11, 22, 33, 44, 56, 59, 70, 82, 94, 108, 130, 144, 155, 170, 181, 191, 206, 220, 236, 251, 258, 271, 283, 295, 308, 329, 350, 371, 392, 402, 415, 425, 434, 453, 472, 491, 506, 518, 530, 540, 551, 563, 574, 586, 597, 609, 619, 632, 642, 666, 687, 708, 727, 744}

const _TypeLowerName = "cursor_movecamera_zoomsummon_unitupdate_unitupdate_towertpsplace_towerremove_towerselect_towerselected_towerselected_tower_invaliddeselect_towerincome_tickwindow_resizingnavigate_tostart_gameopen_tower_menuopen_unit_menuclose_tower_menuclose_unit_menugo_homesign_up_erroruser_sign_upuser_sign_inuser_sign_outjoin_vs6_waiting_roomexit_vs6_waiting_roomjoin_vs1_waiting_roomexit_vs1_waiting_roomstart_roomversion_errorsetup_gamefind_gameexit_searching_gameaccept_waiting_gamecancel_waiting_gameshow_scoreboardcreate_lobbydelete_lobbyjoin_lobbyadd_lobbiesselect_lobbyleave_lobbyupdate_lobbystart_lobbyseen_lobbiesadd_playerremove_playersync_statewait_room_countdown_ticksync_vs6_waiting_roomsync_vs1_waiting_roomsync_searching_roomsync_waiting_room"

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
	_ = x[PlaceTower-(6)]
	_ = x[RemoveTower-(7)]
	_ = x[SelectTower-(8)]
	_ = x[SelectedTower-(9)]
	_ = x[SelectedTowerInvalid-(10)]
	_ = x[DeselectTower-(11)]
	_ = x[IncomeTick-(12)]
	_ = x[WindowResizing-(13)]
	_ = x[NavigateTo-(14)]
	_ = x[StartGame-(15)]
	_ = x[OpenTowerMenu-(16)]
	_ = x[OpenUnitMenu-(17)]
	_ = x[CloseTowerMenu-(18)]
	_ = x[CloseUnitMenu-(19)]
	_ = x[GoHome-(20)]
	_ = x[SignUpError-(21)]
	_ = x[UserSignUp-(22)]
	_ = x[UserSignIn-(23)]
	_ = x[UserSignOut-(24)]
	_ = x[JoinVs6WaitingRoom-(25)]
	_ = x[ExitVs6WaitingRoom-(26)]
	_ = x[JoinVs1WaitingRoom-(27)]
	_ = x[ExitVs1WaitingRoom-(28)]
	_ = x[StartRoom-(29)]
	_ = x[VersionError-(30)]
	_ = x[SetupGame-(31)]
	_ = x[FindGame-(32)]
	_ = x[ExitSearchingGame-(33)]
	_ = x[AcceptWaitingGame-(34)]
	_ = x[CancelWaitingGame-(35)]
	_ = x[ShowScoreboard-(36)]
	_ = x[CreateLobby-(37)]
	_ = x[DeleteLobby-(38)]
	_ = x[JoinLobby-(39)]
	_ = x[AddLobbies-(40)]
	_ = x[SelectLobby-(41)]
	_ = x[LeaveLobby-(42)]
	_ = x[UpdateLobby-(43)]
	_ = x[StartLobby-(44)]
	_ = x[SeenLobbies-(45)]
	_ = x[AddPlayer-(46)]
	_ = x[RemovePlayer-(47)]
	_ = x[SyncState-(48)]
	_ = x[WaitRoomCountdownTick-(49)]
	_ = x[SyncVs6WaitingRoom-(50)]
	_ = x[SyncVs1WaitingRoom-(51)]
	_ = x[SyncSearchingRoom-(52)]
	_ = x[SyncWaitingRoom-(53)]
}

var _TypeValues = []Type{CursorMove, CameraZoom, SummonUnit, UpdateUnit, UpdateTower, TPS, PlaceTower, RemoveTower, SelectTower, SelectedTower, SelectedTowerInvalid, DeselectTower, IncomeTick, WindowResizing, NavigateTo, StartGame, OpenTowerMenu, OpenUnitMenu, CloseTowerMenu, CloseUnitMenu, GoHome, SignUpError, UserSignUp, UserSignIn, UserSignOut, JoinVs6WaitingRoom, ExitVs6WaitingRoom, JoinVs1WaitingRoom, ExitVs1WaitingRoom, StartRoom, VersionError, SetupGame, FindGame, ExitSearchingGame, AcceptWaitingGame, CancelWaitingGame, ShowScoreboard, CreateLobby, DeleteLobby, JoinLobby, AddLobbies, SelectLobby, LeaveLobby, UpdateLobby, StartLobby, SeenLobbies, AddPlayer, RemovePlayer, SyncState, WaitRoomCountdownTick, SyncVs6WaitingRoom, SyncVs1WaitingRoom, SyncSearchingRoom, SyncWaitingRoom}

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
	_TypeName[59:70]:        PlaceTower,
	_TypeLowerName[59:70]:   PlaceTower,
	_TypeName[70:82]:        RemoveTower,
	_TypeLowerName[70:82]:   RemoveTower,
	_TypeName[82:94]:        SelectTower,
	_TypeLowerName[82:94]:   SelectTower,
	_TypeName[94:108]:       SelectedTower,
	_TypeLowerName[94:108]:  SelectedTower,
	_TypeName[108:130]:      SelectedTowerInvalid,
	_TypeLowerName[108:130]: SelectedTowerInvalid,
	_TypeName[130:144]:      DeselectTower,
	_TypeLowerName[130:144]: DeselectTower,
	_TypeName[144:155]:      IncomeTick,
	_TypeLowerName[144:155]: IncomeTick,
	_TypeName[155:170]:      WindowResizing,
	_TypeLowerName[155:170]: WindowResizing,
	_TypeName[170:181]:      NavigateTo,
	_TypeLowerName[170:181]: NavigateTo,
	_TypeName[181:191]:      StartGame,
	_TypeLowerName[181:191]: StartGame,
	_TypeName[191:206]:      OpenTowerMenu,
	_TypeLowerName[191:206]: OpenTowerMenu,
	_TypeName[206:220]:      OpenUnitMenu,
	_TypeLowerName[206:220]: OpenUnitMenu,
	_TypeName[220:236]:      CloseTowerMenu,
	_TypeLowerName[220:236]: CloseTowerMenu,
	_TypeName[236:251]:      CloseUnitMenu,
	_TypeLowerName[236:251]: CloseUnitMenu,
	_TypeName[251:258]:      GoHome,
	_TypeLowerName[251:258]: GoHome,
	_TypeName[258:271]:      SignUpError,
	_TypeLowerName[258:271]: SignUpError,
	_TypeName[271:283]:      UserSignUp,
	_TypeLowerName[271:283]: UserSignUp,
	_TypeName[283:295]:      UserSignIn,
	_TypeLowerName[283:295]: UserSignIn,
	_TypeName[295:308]:      UserSignOut,
	_TypeLowerName[295:308]: UserSignOut,
	_TypeName[308:329]:      JoinVs6WaitingRoom,
	_TypeLowerName[308:329]: JoinVs6WaitingRoom,
	_TypeName[329:350]:      ExitVs6WaitingRoom,
	_TypeLowerName[329:350]: ExitVs6WaitingRoom,
	_TypeName[350:371]:      JoinVs1WaitingRoom,
	_TypeLowerName[350:371]: JoinVs1WaitingRoom,
	_TypeName[371:392]:      ExitVs1WaitingRoom,
	_TypeLowerName[371:392]: ExitVs1WaitingRoom,
	_TypeName[392:402]:      StartRoom,
	_TypeLowerName[392:402]: StartRoom,
	_TypeName[402:415]:      VersionError,
	_TypeLowerName[402:415]: VersionError,
	_TypeName[415:425]:      SetupGame,
	_TypeLowerName[415:425]: SetupGame,
	_TypeName[425:434]:      FindGame,
	_TypeLowerName[425:434]: FindGame,
	_TypeName[434:453]:      ExitSearchingGame,
	_TypeLowerName[434:453]: ExitSearchingGame,
	_TypeName[453:472]:      AcceptWaitingGame,
	_TypeLowerName[453:472]: AcceptWaitingGame,
	_TypeName[472:491]:      CancelWaitingGame,
	_TypeLowerName[472:491]: CancelWaitingGame,
	_TypeName[491:506]:      ShowScoreboard,
	_TypeLowerName[491:506]: ShowScoreboard,
	_TypeName[506:518]:      CreateLobby,
	_TypeLowerName[506:518]: CreateLobby,
	_TypeName[518:530]:      DeleteLobby,
	_TypeLowerName[518:530]: DeleteLobby,
	_TypeName[530:540]:      JoinLobby,
	_TypeLowerName[530:540]: JoinLobby,
	_TypeName[540:551]:      AddLobbies,
	_TypeLowerName[540:551]: AddLobbies,
	_TypeName[551:563]:      SelectLobby,
	_TypeLowerName[551:563]: SelectLobby,
	_TypeName[563:574]:      LeaveLobby,
	_TypeLowerName[563:574]: LeaveLobby,
	_TypeName[574:586]:      UpdateLobby,
	_TypeLowerName[574:586]: UpdateLobby,
	_TypeName[586:597]:      StartLobby,
	_TypeLowerName[586:597]: StartLobby,
	_TypeName[597:609]:      SeenLobbies,
	_TypeLowerName[597:609]: SeenLobbies,
	_TypeName[609:619]:      AddPlayer,
	_TypeLowerName[609:619]: AddPlayer,
	_TypeName[619:632]:      RemovePlayer,
	_TypeLowerName[619:632]: RemovePlayer,
	_TypeName[632:642]:      SyncState,
	_TypeLowerName[632:642]: SyncState,
	_TypeName[642:666]:      WaitRoomCountdownTick,
	_TypeLowerName[642:666]: WaitRoomCountdownTick,
	_TypeName[666:687]:      SyncVs6WaitingRoom,
	_TypeLowerName[666:687]: SyncVs6WaitingRoom,
	_TypeName[687:708]:      SyncVs1WaitingRoom,
	_TypeLowerName[687:708]: SyncVs1WaitingRoom,
	_TypeName[708:727]:      SyncSearchingRoom,
	_TypeLowerName[708:727]: SyncSearchingRoom,
	_TypeName[727:744]:      SyncWaitingRoom,
	_TypeLowerName[727:744]: SyncWaitingRoom,
}

var _TypeNames = []string{
	_TypeName[0:11],
	_TypeName[11:22],
	_TypeName[22:33],
	_TypeName[33:44],
	_TypeName[44:56],
	_TypeName[56:59],
	_TypeName[59:70],
	_TypeName[70:82],
	_TypeName[82:94],
	_TypeName[94:108],
	_TypeName[108:130],
	_TypeName[130:144],
	_TypeName[144:155],
	_TypeName[155:170],
	_TypeName[170:181],
	_TypeName[181:191],
	_TypeName[191:206],
	_TypeName[206:220],
	_TypeName[220:236],
	_TypeName[236:251],
	_TypeName[251:258],
	_TypeName[258:271],
	_TypeName[271:283],
	_TypeName[283:295],
	_TypeName[295:308],
	_TypeName[308:329],
	_TypeName[329:350],
	_TypeName[350:371],
	_TypeName[371:392],
	_TypeName[392:402],
	_TypeName[402:415],
	_TypeName[415:425],
	_TypeName[425:434],
	_TypeName[434:453],
	_TypeName[453:472],
	_TypeName[472:491],
	_TypeName[491:506],
	_TypeName[506:518],
	_TypeName[518:530],
	_TypeName[530:540],
	_TypeName[540:551],
	_TypeName[551:563],
	_TypeName[563:574],
	_TypeName[574:586],
	_TypeName[586:597],
	_TypeName[597:609],
	_TypeName[609:619],
	_TypeName[619:632],
	_TypeName[632:642],
	_TypeName[642:666],
	_TypeName[666:687],
	_TypeName[687:708],
	_TypeName[708:727],
	_TypeName[727:744],
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
