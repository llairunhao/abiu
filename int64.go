package abiu

import (
	"encoding/json"
	"fmt"
)

func NewInt64(value int64) *Int64 {
	return &Int64{value: value}
}

type Int64 struct {
	value int64
}

//-------------------------------------------------------
func (number *Int64) SetValue(v interface{}) error {
	switch v.(type) {
	case int64:
		{
			number.value = v.(int64)
		}
	case string:
		{
			str := v.(string)
			value, err := intValue(str, 64)
			if err != nil {
				return err
			}
			number.value = int64(value)
		}
	default:
		return typeError(v)
	}
	return nil
}

func (number *Int64) Value() interface{} {
	return number.value
}

func (number *Int64) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------

func (number *Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Int64) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Int64) String() string {
	return fmt.Sprintf("%d", number.value)
}

func (number *Int64) IsZero() bool {
	if number == nil {
		return true
	}
	return number.value == 0
}
