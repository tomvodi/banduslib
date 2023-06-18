// Code generated by "enumer -json -yaml -type=DacapoType"; DO NOT EDIT.

package barline

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _DacapoTypeName = "NoDacapoTypeFineDacapoAlFine"

var _DacapoTypeIndex = [...]uint8{0, 12, 16, 28}

const _DacapoTypeLowerName = "nodacapotypefinedacapoalfine"

func (i DacapoType) String() string {
	if i >= DacapoType(len(_DacapoTypeIndex)-1) {
		return fmt.Sprintf("DacapoType(%d)", i)
	}
	return _DacapoTypeName[_DacapoTypeIndex[i]:_DacapoTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _DacapoTypeNoOp() {
	var x [1]struct{}
	_ = x[NoDacapoType-(0)]
	_ = x[Fine-(1)]
	_ = x[DacapoAlFine-(2)]
}

var _DacapoTypeValues = []DacapoType{NoDacapoType, Fine, DacapoAlFine}

var _DacapoTypeNameToValueMap = map[string]DacapoType{
	_DacapoTypeName[0:12]:       NoDacapoType,
	_DacapoTypeLowerName[0:12]:  NoDacapoType,
	_DacapoTypeName[12:16]:      Fine,
	_DacapoTypeLowerName[12:16]: Fine,
	_DacapoTypeName[16:28]:      DacapoAlFine,
	_DacapoTypeLowerName[16:28]: DacapoAlFine,
}

var _DacapoTypeNames = []string{
	_DacapoTypeName[0:12],
	_DacapoTypeName[12:16],
	_DacapoTypeName[16:28],
}

// DacapoTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DacapoTypeString(s string) (DacapoType, error) {
	if val, ok := _DacapoTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _DacapoTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DacapoType values", s)
}

// DacapoTypeValues returns all values of the enum
func DacapoTypeValues() []DacapoType {
	return _DacapoTypeValues
}

// DacapoTypeStrings returns a slice of all String values of the enum
func DacapoTypeStrings() []string {
	strs := make([]string, len(_DacapoTypeNames))
	copy(strs, _DacapoTypeNames)
	return strs
}

// IsADacapoType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DacapoType) IsADacapoType() bool {
	for _, v := range _DacapoTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for DacapoType
func (i DacapoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for DacapoType
func (i *DacapoType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("DacapoType should be a string, got %s", data)
	}

	var err error
	*i, err = DacapoTypeString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for DacapoType
func (i DacapoType) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for DacapoType
func (i *DacapoType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = DacapoTypeString(s)
	return err
}