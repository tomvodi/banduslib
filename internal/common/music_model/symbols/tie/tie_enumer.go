// Code generated by "enumer -json -yaml -type=Tie"; DO NOT EDIT.

package tie

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TieName = "NoTieStartEnd"

var _TieIndex = [...]uint8{0, 5, 10, 13}

const _TieLowerName = "notiestartend"

func (i Tie) String() string {
	if i >= Tie(len(_TieIndex)-1) {
		return fmt.Sprintf("Tie(%d)", i)
	}
	return _TieName[_TieIndex[i]:_TieIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TieNoOp() {
	var x [1]struct{}
	_ = x[NoTie-(0)]
	_ = x[Start-(1)]
	_ = x[End-(2)]
}

var _TieValues = []Tie{NoTie, Start, End}

var _TieNameToValueMap = map[string]Tie{
	_TieName[0:5]:        NoTie,
	_TieLowerName[0:5]:   NoTie,
	_TieName[5:10]:       Start,
	_TieLowerName[5:10]:  Start,
	_TieName[10:13]:      End,
	_TieLowerName[10:13]: End,
}

var _TieNames = []string{
	_TieName[0:5],
	_TieName[5:10],
	_TieName[10:13],
}

// TieString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TieString(s string) (Tie, error) {
	if val, ok := _TieNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TieNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Tie values", s)
}

// TieValues returns all values of the enum
func TieValues() []Tie {
	return _TieValues
}

// TieStrings returns a slice of all String values of the enum
func TieStrings() []string {
	strs := make([]string, len(_TieNames))
	copy(strs, _TieNames)
	return strs
}

// IsATie returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Tie) IsATie() bool {
	for _, v := range _TieValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Tie
func (i Tie) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Tie
func (i *Tie) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Tie should be a string, got %s", data)
	}

	var err error
	*i, err = TieString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Tie
func (i Tie) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Tie
func (i *Tie) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = TieString(s)
	return err
}
