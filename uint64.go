package abiu

import "encoding/json"

func NewUInt64(value uint64) *UInt64 {
	return &UInt64{value: value}
}

type UInt64 struct {
	value uint64
}

//-------------------------------------------------------
func (number *UInt64) SetValue(v interface{}) {
	var ok bool
	if number.value, ok = v.(uint64); !ok {
		number.value = 0
	}
}

func (number *UInt64) Value() interface{} {
	return number.value
}

func (number *UInt64) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *UInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *UInt64) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *UInt64) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
