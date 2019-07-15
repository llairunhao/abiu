package abiu

import "encoding/json"

func NewInt(value int) *Int {
	return &Int{value: value}
}

type Int struct {
	value int
}

//-------------------------------------------------------
func (number *Int) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(int); !ok {
		number.value = 0
	}
}

func (number *Int) Value() interface{} {
	return number.value
}

func (number *Int) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Int) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Int) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
