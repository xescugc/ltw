// Code generated by "enumer -type=Ability -transform=lower -json -transform=snake -output=ability_string.go"; DO NOT EDIT.

package ability

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _AbilityName = "splitburrowresurrectionhybridcamouflageattack"

var _AbilityIndex = [...]uint8{0, 5, 11, 23, 29, 39, 45}

const _AbilityLowerName = "splitburrowresurrectionhybridcamouflageattack"

func (i Ability) String() string {
	if i < 0 || i >= Ability(len(_AbilityIndex)-1) {
		return fmt.Sprintf("Ability(%d)", i)
	}
	return _AbilityName[_AbilityIndex[i]:_AbilityIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _AbilityNoOp() {
	var x [1]struct{}
	_ = x[Split-(0)]
	_ = x[Burrow-(1)]
	_ = x[Resurrection-(2)]
	_ = x[Hybrid-(3)]
	_ = x[Camouflage-(4)]
	_ = x[Attack-(5)]
}

var _AbilityValues = []Ability{Split, Burrow, Resurrection, Hybrid, Camouflage, Attack}

var _AbilityNameToValueMap = map[string]Ability{
	_AbilityName[0:5]:        Split,
	_AbilityLowerName[0:5]:   Split,
	_AbilityName[5:11]:       Burrow,
	_AbilityLowerName[5:11]:  Burrow,
	_AbilityName[11:23]:      Resurrection,
	_AbilityLowerName[11:23]: Resurrection,
	_AbilityName[23:29]:      Hybrid,
	_AbilityLowerName[23:29]: Hybrid,
	_AbilityName[29:39]:      Camouflage,
	_AbilityLowerName[29:39]: Camouflage,
	_AbilityName[39:45]:      Attack,
	_AbilityLowerName[39:45]: Attack,
}

var _AbilityNames = []string{
	_AbilityName[0:5],
	_AbilityName[5:11],
	_AbilityName[11:23],
	_AbilityName[23:29],
	_AbilityName[29:39],
	_AbilityName[39:45],
}

// AbilityString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AbilityString(s string) (Ability, error) {
	if val, ok := _AbilityNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _AbilityNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Ability values", s)
}

// AbilityValues returns all values of the enum
func AbilityValues() []Ability {
	return _AbilityValues
}

// AbilityStrings returns a slice of all String values of the enum
func AbilityStrings() []string {
	strs := make([]string, len(_AbilityNames))
	copy(strs, _AbilityNames)
	return strs
}

// IsAAbility returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Ability) IsAAbility() bool {
	for _, v := range _AbilityValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Ability
func (i Ability) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Ability
func (i *Ability) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Ability should be a string, got %s", data)
	}

	var err error
	*i, err = AbilityString(s)
	return err
}
