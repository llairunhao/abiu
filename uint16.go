package abiu

import "encoding/json"

func NewUInt16(value uint16) *UInt16 {
	return &UInt16{value: value}
}

type UInt16 struct {
	value uint16
}

//-------------------------------------------------------
func (number *UInt16) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(uint16); !ok {
		number.value = 0
	}
}

func (number *UInt16) Value() interface{} {
	return number.value
}

func (number *UInt16) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------
func (number *UInt16) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *UInt16) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *UInt16) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
