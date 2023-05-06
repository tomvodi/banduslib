// Code generated by "enumer -json -yaml -type=Length"; DO NOT EDIT.

package symbols

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _LengthName = "NoLengthWholeHalfQuarterEighthSixteenthThirtysecond"

var _LengthIndex = [...]uint8{0, 8, 13, 17, 24, 30, 39, 51}

const _LengthLowerName = "nolengthwholehalfquartereighthsixteenththirtysecond"

func (i Length) String() string {
	if i >= Length(len(_LengthIndex)-1) {
		return fmt.Sprintf("Length(%d)", i)
	}
	return _LengthName[_LengthIndex[i]:_LengthIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _LengthNoOp() {
	var x [1]struct{}
	_ = x[NoLength-(0)]
	_ = x[Whole-(1)]
	_ = x[Half-(2)]
	_ = x[Quarter-(3)]
	_ = x[Eighth-(4)]
	_ = x[Sixteenth-(5)]
	_ = x[Thirtysecond-(6)]
}

var _LengthValues = []Length{NoLength, Whole, Half, Quarter, Eighth, Sixteenth, Thirtysecond}

var _LengthNameToValueMap = map[string]Length{
	_LengthName[0:8]:        NoLength,
	_LengthLowerName[0:8]:   NoLength,
	_LengthName[8:13]:       Whole,
	_LengthLowerName[8:13]:  Whole,
	_LengthName[13:17]:      Half,
	_LengthLowerName[13:17]: Half,
	_LengthName[17:24]:      Quarter,
	_LengthLowerName[17:24]: Quarter,
	_LengthName[24:30]:      Eighth,
	_LengthLowerName[24:30]: Eighth,
	_LengthName[30:39]:      Sixteenth,
	_LengthLowerName[30:39]: Sixteenth,
	_LengthName[39:51]:      Thirtysecond,
	_LengthLowerName[39:51]: Thirtysecond,
}

var _LengthNames = []string{
	_LengthName[0:8],
	_LengthName[8:13],
	_LengthName[13:17],
	_LengthName[17:24],
	_LengthName[24:30],
	_LengthName[30:39],
	_LengthName[39:51],
}

// LengthString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func LengthString(s string) (Length, error) {
	if val, ok := _LengthNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _LengthNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Length values", s)
}

// LengthValues returns all values of the enum
func LengthValues() []Length {
	return _LengthValues
}

// LengthStrings returns a slice of all String values of the enum
func LengthStrings() []string {
	strs := make([]string, len(_LengthNames))
	copy(strs, _LengthNames)
	return strs
}

// IsALength returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Length) IsALength() bool {
	for _, v := range _LengthValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Length
func (i Length) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Length
func (i *Length) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Length should be a string, got %s", data)
	}

	var err error
	*i, err = LengthString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Length
func (i Length) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Length
func (i *Length) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = LengthString(s)
	return err
}
