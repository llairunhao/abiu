package abiu

import (
	"encoding/json"
	"fmt"
)

func NewUInt64(value uint64) *UInt64 {
	return &UInt64{value: value}
}

type UInt64 struct {
	value uint64
}

//-------------------------------------------------------
func (number *UInt64) SetValue(v interface{}) error {
	switch v.(type) {
	case uint64:
		{
			number.value = v.(uint64)
		}
	case string:
		{
			str := v.(string)
			value, err := uintValue(str, 64)
			if err != nil {
				return err
			}
			number.value = uint64(value)
		}
	default:
		return typeError(v)
	}
	return nil
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

func (number *UInt64) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *UInt64) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
