// Code generated by "enumer -type=Type -transform=snake -output=type_string.go -json"; DO NOT EDIT.

package action

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TypeName = "cursor_movecamera_zoomsummon_unitmove_unitremove_unitsteal_liveplace_towerselect_towerselected_towerselected_tower_invaliddeselect_towerincome_ticktower_attackunit_killedwindow_resizingjoin_roomadd_playerupdate_state"

var _TypeIndex = [...]uint8{0, 11, 22, 33, 42, 53, 63, 74, 86, 100, 122, 136, 147, 159, 170, 185, 194, 204, 216}

const _TypeLowerName = "cursor_movecamera_zoomsummon_unitmove_unitremove_unitsteal_liveplace_towerselect_towerselected_towerselected_tower_invaliddeselect_towerincome_ticktower_attackunit_killedwindow_resizingjoin_roomadd_playerupdate_state"

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
	_ = x[MoveUnit-(3)]
	_ = x[RemoveUnit-(4)]
	_ = x[StealLive-(5)]
	_ = x[PlaceTower-(6)]
	_ = x[SelectTower-(7)]
	_ = x[SelectedTower-(8)]
	_ = x[SelectedTowerInvalid-(9)]
	_ = x[DeselectTower-(10)]
	_ = x[IncomeTick-(11)]
	_ = x[TowerAttack-(12)]
	_ = x[UnitKilled-(13)]
	_ = x[WindowResizing-(14)]
	_ = x[JoinRoom-(15)]
	_ = x[AddPlayer-(16)]
	_ = x[UpdateState-(17)]
}

var _TypeValues = []Type{CursorMove, CameraZoom, SummonUnit, MoveUnit, RemoveUnit, StealLive, PlaceTower, SelectTower, SelectedTower, SelectedTowerInvalid, DeselectTower, IncomeTick, TowerAttack, UnitKilled, WindowResizing, JoinRoom, AddPlayer, UpdateState}

var _TypeNameToValueMap = map[string]Type{
	_TypeName[0:11]:         CursorMove,
	_TypeLowerName[0:11]:    CursorMove,
	_TypeName[11:22]:        CameraZoom,
	_TypeLowerName[11:22]:   CameraZoom,
	_TypeName[22:33]:        SummonUnit,
	_TypeLowerName[22:33]:   SummonUnit,
	_TypeName[33:42]:        MoveUnit,
	_TypeLowerName[33:42]:   MoveUnit,
	_TypeName[42:53]:        RemoveUnit,
	_TypeLowerName[42:53]:   RemoveUnit,
	_TypeName[53:63]:        StealLive,
	_TypeLowerName[53:63]:   StealLive,
	_TypeName[63:74]:        PlaceTower,
	_TypeLowerName[63:74]:   PlaceTower,
	_TypeName[74:86]:        SelectTower,
	_TypeLowerName[74:86]:   SelectTower,
	_TypeName[86:100]:       SelectedTower,
	_TypeLowerName[86:100]:  SelectedTower,
	_TypeName[100:122]:      SelectedTowerInvalid,
	_TypeLowerName[100:122]: SelectedTowerInvalid,
	_TypeName[122:136]:      DeselectTower,
	_TypeLowerName[122:136]: DeselectTower,
	_TypeName[136:147]:      IncomeTick,
	_TypeLowerName[136:147]: IncomeTick,
	_TypeName[147:159]:      TowerAttack,
	_TypeLowerName[147:159]: TowerAttack,
	_TypeName[159:170]:      UnitKilled,
	_TypeLowerName[159:170]: UnitKilled,
	_TypeName[170:185]:      WindowResizing,
	_TypeLowerName[170:185]: WindowResizing,
	_TypeName[185:194]:      JoinRoom,
	_TypeLowerName[185:194]: JoinRoom,
	_TypeName[194:204]:      AddPlayer,
	_TypeLowerName[194:204]: AddPlayer,
	_TypeName[204:216]:      UpdateState,
	_TypeLowerName[204:216]: UpdateState,
}

var _TypeNames = []string{
	_TypeName[0:11],
	_TypeName[11:22],
	_TypeName[22:33],
	_TypeName[33:42],
	_TypeName[42:53],
	_TypeName[53:63],
	_TypeName[63:74],
	_TypeName[74:86],
	_TypeName[86:100],
	_TypeName[100:122],
	_TypeName[122:136],
	_TypeName[136:147],
	_TypeName[147:159],
	_TypeName[159:170],
	_TypeName[170:185],
	_TypeName[185:194],
	_TypeName[194:204],
	_TypeName[204:216],
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
