package abiu

import "encoding/json"

func NewUInt(value uint) *UInt {
	return &UInt{value: value}
}

type UInt struct {
	value uint
}

//-------------------------------------------------------
func (number *UInt) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(uint); !ok {
		number.value = 0
	}
}

func (number *UInt) Value() interface{} {
	return number.value
}

func (number *UInt) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *UInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *UInt) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *UInt) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
