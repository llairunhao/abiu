package abiu

import (
	"encoding/json"
	"fmt"
)

func NewUInt8(value uint8) *UInt8 {
	return &UInt8{value: value}
}

type UInt8 struct {
	value uint8
}

//-------------------------------------------------------
func (number *UInt8) SetValue(v interface{}) error {
	switch v.(type) {
	case uint8:
		{
			number.value = v.(uint8)
		}
	case string:
		{
			str := v.(string)
			value, err := uintValue(str, 8)
			if err != nil {
				return err
			}
			number.value = uint8(value)
		}
	default:
		return typeError(v)
	}
	return nil
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

func (number *UInt8) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *UInt8) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
