package abiu

import "encoding/json"

func NewFloat64(value float64) *Float64 {
	return &Float64{value: value}
}

type Float64 struct {
	value float64
}

//-------------------------------------------------------
func (number *Float64) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(float64); !ok {
		number.value = 0
	}
}

func (number *Float64) Value() interface{} {
	return number.value
}

func (number *Float64) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Float64) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Float64) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
