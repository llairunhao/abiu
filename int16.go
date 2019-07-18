package abiu

import (
	"encoding/json"
	"fmt"
)

func NewInt16(value int16) *Int16 {
	return &Int16{value: value}
}

type Int16 struct {
	value int16
}

//-------------------------------------------------------
func (number *Int16) SetValue(v interface{}) error {
	switch v.(type) {
	case int16:
		{
			number.value = v.(int16)
		}
	case string:
		{
			str := v.(string)
			value, err := intValue(str, 16)
			if err != nil {
				return err
			}
			number.value = int16(value)
		}
	default:
		return typeError(v)
	}
	return nil
}

func (number *Int16) Value() interface{} {
	return number.value
}

func (number *Int16) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Int16) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Int16) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *Int16) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
