package abiu

import (
	"encoding/json"
	"fmt"
)

func NewUInt16(value uint16) *UInt16 {
	return &UInt16{value: value}
}

type UInt16 struct {
	value uint16
}

//-------------------------------------------------------
func (number *UInt16) SetValue(v interface{}) error {
	switch v.(type) {
	case uint16:
		{
			number.value = v.(uint16)
		}
	case string:
		{
			str := v.(string)
			value, err := uintValue(str, 16)
			if err != nil {
				return err
			}
			number.value = uint16(value)
		}
	default:
		return typeError(v)
	}
	return nil
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

func (number *UInt16) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *UInt16) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
