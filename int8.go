package abiu

import "encoding/json"

func NewInt8(value int8) *Int8 {
	return &Int8{value: value}
}

type Int8 struct {
	value int8
}

//-------------------------------------------------------
func (number *Int8) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(int8); !ok {
		number.value = 0
	}
}

func (number *Int8) Value() interface{} {
	return number.value
}

func (number *Int8) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Int8) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Int8) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
