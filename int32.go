package abiu

import (
	"encoding/json"
	"fmt"
)

func NewInt32(value int32) *Int32 {
	return &Int32{value: value}
}

type Int32 struct {
	value int32
}

//-------------------------------------------------------
func (number *Int32) SetValue(v interface{}) error {
	switch v.(type) {
	case int32:
		{
			number.value = v.(int32)
		}
	case string:
		{
			str := v.(string)
			value, err := intValue(str, 32)
			if err != nil {
				return err
			}
			number.value = int32(value)
		}
	default:
		return typeError(v)
	}
	return nil
}

func (number *Int32) Value() interface{} {
	return number.value
}

func (number *Int32) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Int32) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Int32) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *Int32) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
