package abiu

import (
	"encoding/json"
	"fmt"
)

func NewFloat(value float32) *Float {
	return &Float{value: value}
}

type Float struct {
	value float32
}

//-------------------------------------------------------
func (number *Float) SetValue(v interface{}) error {
	switch v.(type) {
	case float32:
		{
			number.value = v.(float32)
		}
	case string:
		{
			str := v.(string)
			value, err := floatValue(str, 32)
			if err != nil {
				return err
			}
			number.value = float32(value)
		}
	default:
		return typeError(v)
	}
	return nil
}

func (number *Float) Value() interface{} {
	return number.value
}

func (number *Float) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Float) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Float) String() string {
	return fmt.Sprintf("%g", number.value)
}

func (number *Float) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
