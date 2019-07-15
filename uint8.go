package abiu

import "encoding/json"

func NewUInt8(value uint8) *UInt8 {
	return &UInt8{value: value}
}

type UInt8 struct {
	value uint8
}

//-------------------------------------------------------
func (number *UInt8) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(uint8); !ok {
		number.value = 0
	}
}

func (number *UInt8) Value() interface{} {
	return number.value
}

func (number *UInt8) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------
func (number *UInt8) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *UInt8) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *UInt8) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
