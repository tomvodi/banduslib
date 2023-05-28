// Code generated by "enumer -json -yaml -type=Type"; DO NOT EDIT.

package movement

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TypeName = "NoMovementCadenceEmbariEndariChedariHedariDiliTraEdreDareCheCheReGripDedaEnbainOtroOdroAdedaEchoBeatDarodoHiharinRodinChelalhoDinLemluathTaorluathCrunluathTriplings"

var _TypeIndex = [...]uint8{0, 10, 17, 23, 29, 36, 42, 46, 49, 53, 57, 65, 69, 73, 79, 83, 87, 92, 100, 106, 113, 118, 126, 129, 137, 146, 155, 164}

const _TypeLowerName = "nomovementcadenceembariendarichedarihedaridilitraedredarechecheregripdedaenbainotroodroadedaechobeatdarodohiharinrodinchelalhodinlemluathtaorluathcrunluathtriplings"

func (i Type) String() string {
	if i >= Type(len(_TypeIndex)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _TypeName[_TypeIndex[i]:_TypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TypeNoOp() {
	var x [1]struct{}
	_ = x[NoMovement-(0)]
	_ = x[Cadence-(1)]
	_ = x[Embari-(2)]
	_ = x[Endari-(3)]
	_ = x[Chedari-(4)]
	_ = x[Hedari-(5)]
	_ = x[Dili-(6)]
	_ = x[Tra-(7)]
	_ = x[Edre-(8)]
	_ = x[Dare-(9)]
	_ = x[CheCheRe-(10)]
	_ = x[Grip-(11)]
	_ = x[Deda-(12)]
	_ = x[Enbain-(13)]
	_ = x[Otro-(14)]
	_ = x[Odro-(15)]
	_ = x[Adeda-(16)]
	_ = x[EchoBeat-(17)]
	_ = x[Darodo-(18)]
	_ = x[Hiharin-(19)]
	_ = x[Rodin-(20)]
	_ = x[Chelalho-(21)]
	_ = x[Din-(22)]
	_ = x[Lemluath-(23)]
	_ = x[Taorluath-(24)]
	_ = x[Crunluath-(25)]
	_ = x[Tripling-(26)]
}

var _TypeValues = []Type{NoMovement, Cadence, Embari, Endari, Chedari, Hedari, Dili, Tra, Edre, Dare, CheCheRe, Grip, Deda, Enbain, Otro, Odro, Adeda, EchoBeat, Darodo, Hiharin, Rodin, Chelalho, Din, Lemluath, Taorluath, Crunluath, Tripling}

var _TypeNameToValueMap = map[string]Type{
	_TypeName[0:10]:         NoMovement,
	_TypeLowerName[0:10]:    NoMovement,
	_TypeName[10:17]:        Cadence,
	_TypeLowerName[10:17]:   Cadence,
	_TypeName[17:23]:        Embari,
	_TypeLowerName[17:23]:   Embari,
	_TypeName[23:29]:        Endari,
	_TypeLowerName[23:29]:   Endari,
	_TypeName[29:36]:        Chedari,
	_TypeLowerName[29:36]:   Chedari,
	_TypeName[36:42]:        Hedari,
	_TypeLowerName[36:42]:   Hedari,
	_TypeName[42:46]:        Dili,
	_TypeLowerName[42:46]:   Dili,
	_TypeName[46:49]:        Tra,
	_TypeLowerName[46:49]:   Tra,
	_TypeName[49:53]:        Edre,
	_TypeLowerName[49:53]:   Edre,
	_TypeName[53:57]:        Dare,
	_TypeLowerName[53:57]:   Dare,
	_TypeName[57:65]:        CheCheRe,
	_TypeLowerName[57:65]:   CheCheRe,
	_TypeName[65:69]:        Grip,
	_TypeLowerName[65:69]:   Grip,
	_TypeName[69:73]:        Deda,
	_TypeLowerName[69:73]:   Deda,
	_TypeName[73:79]:        Enbain,
	_TypeLowerName[73:79]:   Enbain,
	_TypeName[79:83]:        Otro,
	_TypeLowerName[79:83]:   Otro,
	_TypeName[83:87]:        Odro,
	_TypeLowerName[83:87]:   Odro,
	_TypeName[87:92]:        Adeda,
	_TypeLowerName[87:92]:   Adeda,
	_TypeName[92:100]:       EchoBeat,
	_TypeLowerName[92:100]:  EchoBeat,
	_TypeName[100:106]:      Darodo,
	_TypeLowerName[100:106]: Darodo,
	_TypeName[106:113]:      Hiharin,
	_TypeLowerName[106:113]: Hiharin,
	_TypeName[113:118]:      Rodin,
	_TypeLowerName[113:118]: Rodin,
	_TypeName[118:126]:      Chelalho,
	_TypeLowerName[118:126]: Chelalho,
	_TypeName[126:129]:      Din,
	_TypeLowerName[126:129]: Din,
	_TypeName[129:137]:      Lemluath,
	_TypeLowerName[129:137]: Lemluath,
	_TypeName[137:146]:      Taorluath,
	_TypeLowerName[137:146]: Taorluath,
	_TypeName[146:155]:      Crunluath,
	_TypeLowerName[146:155]: Crunluath,
	_TypeName[155:164]:      Tripling,
	_TypeLowerName[155:164]: Tripling,
}

var _TypeNames = []string{
	_TypeName[0:10],
	_TypeName[10:17],
	_TypeName[17:23],
	_TypeName[23:29],
	_TypeName[29:36],
	_TypeName[36:42],
	_TypeName[42:46],
	_TypeName[46:49],
	_TypeName[49:53],
	_TypeName[53:57],
	_TypeName[57:65],
	_TypeName[65:69],
	_TypeName[69:73],
	_TypeName[73:79],
	_TypeName[79:83],
	_TypeName[83:87],
	_TypeName[87:92],
	_TypeName[92:100],
	_TypeName[100:106],
	_TypeName[106:113],
	_TypeName[113:118],
	_TypeName[118:126],
	_TypeName[126:129],
	_TypeName[129:137],
	_TypeName[137:146],
	_TypeName[146:155],
	_TypeName[155:164],
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

// MarshalYAML implements a YAML Marshaler for Type
func (i Type) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Type
func (i *Type) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = TypeString(s)
	return err
}