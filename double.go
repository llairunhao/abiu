package abiu

import (
	"encoding/json"
	"fmt"
)

func NewDouble(value float64) *Double {
	return &Double{value: value}
}

type Double struct {
	value float64
}

//-------------------------------------------------------
func (number *Double) SetValue(v interface{}) error {
	switch v.(type) {
	case float64:
		{
			number.value = v.(float64)
		}
	case string:
		{
			str := v.(string)
			value, err := floatValue(str, 64)
			if err != nil {
				return err
			}
			number.value = float64(value)
		}
	default:
		return typeError(v)
	}
	return nil
}

func (number *Double) Value() interface{} {
	return number.value
}

func (number *Double) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Double) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Double) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Double) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *Double) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
