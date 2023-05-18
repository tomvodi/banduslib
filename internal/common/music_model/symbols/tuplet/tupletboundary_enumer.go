// Code generated by "enumer -json -yaml -type=TupletBoundary"; DO NOT EDIT.

package tuplet

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TupletBoundaryName = "NoBoundaryStartEnd"

var _TupletBoundaryIndex = [...]uint8{0, 10, 15, 18}

const _TupletBoundaryLowerName = "noboundarystartend"

func (i TupletBoundary) String() string {
	if i >= TupletBoundary(len(_TupletBoundaryIndex)-1) {
		return fmt.Sprintf("TupletBoundary(%d)", i)
	}
	return _TupletBoundaryName[_TupletBoundaryIndex[i]:_TupletBoundaryIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TupletBoundaryNoOp() {
	var x [1]struct{}
	_ = x[NoBoundary-(0)]
	_ = x[Start-(1)]
	_ = x[End-(2)]
}

var _TupletBoundaryValues = []TupletBoundary{NoBoundary, Start, End}

var _TupletBoundaryNameToValueMap = map[string]TupletBoundary{
	_TupletBoundaryName[0:10]:       NoBoundary,
	_TupletBoundaryLowerName[0:10]:  NoBoundary,
	_TupletBoundaryName[10:15]:      Start,
	_TupletBoundaryLowerName[10:15]: Start,
	_TupletBoundaryName[15:18]:      End,
	_TupletBoundaryLowerName[15:18]: End,
}

var _TupletBoundaryNames = []string{
	_TupletBoundaryName[0:10],
	_TupletBoundaryName[10:15],
	_TupletBoundaryName[15:18],
}

// TupletBoundaryString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TupletBoundaryString(s string) (TupletBoundary, error) {
	if val, ok := _TupletBoundaryNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TupletBoundaryNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TupletBoundary values", s)
}

// TupletBoundaryValues returns all values of the enum
func TupletBoundaryValues() []TupletBoundary {
	return _TupletBoundaryValues
}

// TupletBoundaryStrings returns a slice of all String values of the enum
func TupletBoundaryStrings() []string {
	strs := make([]string, len(_TupletBoundaryNames))
	copy(strs, _TupletBoundaryNames)
	return strs
}

// IsATupletBoundary returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TupletBoundary) IsATupletBoundary() bool {
	for _, v := range _TupletBoundaryValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for TupletBoundary
func (i TupletBoundary) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TupletBoundary
func (i *TupletBoundary) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TupletBoundary should be a string, got %s", data)
	}

	var err error
	*i, err = TupletBoundaryString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for TupletBoundary
func (i TupletBoundary) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for TupletBoundary
func (i *TupletBoundary) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = TupletBoundaryString(s)
	return err
}