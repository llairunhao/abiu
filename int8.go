package abiu

import (
	"encoding/json"
	"fmt"
)

func NewInt8(value int8) *Int8 {
	return &Int8{value: value}
}

type Int8 struct {
	value int8
}

//-------------------------------------------------------
func (number *Int8) SetValue(v interface{}) error {
	switch v.(type) {
	case int8:
		{
			number.value = v.(int8)
		}
	case string:
		{
			str := v.(string)
			value, err := intValue(str, 8)
			if err != nil {
				return err
			}
			number.value = int8(value)
		}
	default:
		return typeError(v)
	}
	return nil
}

func (number *Int8) Value() interface{} {
	return number.value
}

func (number *Int8) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Int8) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Int8) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *Int8) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
