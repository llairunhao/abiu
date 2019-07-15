package abiu

import "encoding/json"

func NewByte(value byte) *Byte {
	return &Byte{value: value}
}

type Byte struct {
	value byte
}

//-------------------------------------------------------
func (number *Byte) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(byte); !ok {
		number.value = 0
	}
}

func (number *Byte) Value() interface{} {
	return number.value
}

func (number *Byte) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------
func (number *Byte) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Byte) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Byte) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
