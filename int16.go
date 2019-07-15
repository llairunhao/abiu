package abiu

import "encoding/json"

func NewInt16(value int16) *Int16 {
	return &Int16{value: value}
}

type Int16 struct {
	value int16
}

//-------------------------------------------------------
func (number *Int16) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(int16); !ok {
		number.value = 0
	}
}

func (number *Int16) Value() interface{} {
	return number.value
}

func (number *Int16) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Int16) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Int16) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
