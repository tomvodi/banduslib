// Code generated by "enumer -json -yaml -type=Fix"; DO NOT EDIT.

package import_message

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _FixName = "NoFixSkipSymbol"

var _FixIndex = [...]uint8{0, 5, 15}

const _FixLowerName = "nofixskipsymbol"

func (i Fix) String() string {
	if i >= Fix(len(_FixIndex)-1) {
		return fmt.Sprintf("Fix(%d)", i)
	}
	return _FixName[_FixIndex[i]:_FixIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FixNoOp() {
	var x [1]struct{}
	_ = x[NoFix-(0)]
	_ = x[SkipSymbol-(1)]
}

var _FixValues = []Fix{NoFix, SkipSymbol}

var _FixNameToValueMap = map[string]Fix{
	_FixName[0:5]:       NoFix,
	_FixLowerName[0:5]:  NoFix,
	_FixName[5:15]:      SkipSymbol,
	_FixLowerName[5:15]: SkipSymbol,
}

var _FixNames = []string{
	_FixName[0:5],
	_FixName[5:15],
}

// FixString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FixString(s string) (Fix, error) {
	if val, ok := _FixNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FixNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Fix values", s)
}

// FixValues returns all values of the enum
func FixValues() []Fix {
	return _FixValues
}

// FixStrings returns a slice of all String values of the enum
func FixStrings() []string {
	strs := make([]string, len(_FixNames))
	copy(strs, _FixNames)
	return strs
}

// IsAFix returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Fix) IsAFix() bool {
	for _, v := range _FixValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Fix
func (i Fix) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Fix
func (i *Fix) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Fix should be a string, got %s", data)
	}

	var err error
	*i, err = FixString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Fix
func (i Fix) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Fix
func (i *Fix) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = FixString(s)
	return err
}