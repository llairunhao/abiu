package abiu

import (
	"encoding/json"
	"fmt"
)

func NewUInt(value uint) *UInt {
	return &UInt{value: value}
}

type UInt struct {
	value uint
}

//-------------------------------------------------------
func (number *UInt) SetValue(v interface{}) error {
	switch v.(type) {
	case uint:
		{
			number.value = v.(uint)
		}
	case string:
		{
			str := v.(string)
			value, err := uintValue(str, 0)
			if err != nil {
				return err
			}
			number.value = uint(value)
		}
	default:
		return typeError(v)
	}
	return nil
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

func (number *UInt) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *UInt) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
