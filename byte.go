package abiu

import (
	"encoding/json"
	"fmt"
)

func NewByte(value byte) *Byte {
	return &Byte{value: value}
}

type Byte struct {
	value byte
}

//-------------------------------------------------------
func (number *Byte) SetValue(v interface{}) error {
	switch v.(type) {
	case byte:
		{
			number.value = v.(byte)
		}
	case string:
		{
			str := v.(string)
			value, err := uintValue(str, 8)
			if err != nil {
				return err
			}
			number.value = byte(value)
		}
	default:
		return typeError(v)
	}
	return nil
}

func (number *Byte) Value() interface{} {
	return number.value
}

func (number *Byte) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------
func (number *Byte) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Byte) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Byte) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *Byte) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
