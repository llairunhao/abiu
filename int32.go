package abiu

import "encoding/json"

func NewInt32(value int32) *Int32 {
	return &Int32{value: value}
}

type Int32 struct {
	value int32
}

//-------------------------------------------------------
func (number *Int32) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(int32); !ok {
		number.value = 0
	}
}

func (number *Int32) Value() interface{} {
	return number.value
}

func (number *Int32) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Int32) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Int32) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
