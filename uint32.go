package abiu

import (
	"encoding/json"
	"fmt"
)

func NewUInt32(value uint32) *UInt32 {
	return &UInt32{value: value}
}

type UInt32 struct {
	value uint32
}

//-------------------------------------------------------
func (number *UInt32) SetValue(v interface{}) error {
	switch v.(type) {
	case uint32:
		{
			number.value = v.(uint32)
		}
	case string:
		{
			str := v.(string)
			value, err := uintValue(str, 32)
			if err != nil {
				return err
			}
			number.value = uint32(value)
		}
	default:
		return typeError(v)
	}
	return nil
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

func (number *UInt32) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *UInt32) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
