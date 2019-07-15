package abiu

import "encoding/json"

func NewUInt32(value uint32) *UInt32 {
	return &UInt32{value: value}
}

type UInt32 struct {
	value uint32
}

//-------------------------------------------------------
func (number *UInt32) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(uint32); !ok {
		number.value = 0
	}
}

func (number *UInt32) Value() interface{} {
	return number.value
}

func (number *UInt32) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *UInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *UInt32) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *UInt32) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
