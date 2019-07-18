package abiu

import (
	"encoding/json"
	"fmt"
)

func NewInt(value int) *Int {
	return &Int{value: value}
}

type Int struct {
	value int
}

//-------------------------------------------------------
func (number *Int) SetValue(v interface{}) error {
	switch v.(type) {
	case int:
		{
			number.value = v.(int)
		}
	case string:
		{
			str := v.(string)
			value, err := intValue(str, 0)
			if err != nil {
				return err
			}
			number.value = int(value)
		}
	default:
		return typeError(v)
	}
	return nil
}

func (number *Int) Value() interface{} {
	return number.value
}

func (number *Int) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Int) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Int) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *Int) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
