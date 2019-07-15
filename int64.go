package abiu

import "encoding/json"

func NewInt64(value int64) *Int64 {
	return &Int64{value: value}
}

type Int64 struct {
	value int64
}

//-------------------------------------------------------
func (number *Int64) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(int64); !ok {
		number.value = 0
	}
}

func (number *Int64) Value() interface{} {
	return number.value
}

func (number *Int64) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Int64) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Int64) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
