package abiu

import "encoding/json"

func NewFloat32(value float32) *Float32 {
	return &Float32{value: value}
}

type Float32 struct {
	value float32
}

//-------------------------------------------------------
func (number *Float32) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(float32); !ok {
		number.value = 0
	}
}

func (number *Float32) Value() interface{} {
	return number.value
}

func (number *Float32) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Float32) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Float32) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Float32) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
