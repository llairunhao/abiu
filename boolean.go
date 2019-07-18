package abiu

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func NewBoolean(value bool) *Boolean {
	return &Boolean{value: value}
}

type Boolean struct {
	value bool
}

//-------------------------------------------------------
func (number *Boolean) SetValue(v interface{}) error {
	switch v.(type) {
	case bool:
		{
			number.value = v.(bool)
		}
	case string:
		{
			str := v.(string)
			if str == "" {
				number.value = false
				break
			}
			value, err := strconv.ParseBool(str)
			if err != nil {
				return err
			}
			number.value = value
		}
	default:
		return typeError(v)
	}
	return nil
}

func (number *Boolean) Value() interface{} {
	return number.value
}

func (number *Boolean) ValuePtr() interface{} {
	return &number.value
}

//-------------------------------------------------------
func (number *Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.value)
}

func (number *Boolean) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.value)
}

func (number *Boolean) String() string {
	return fmt.Sprintf("%t", number.value)
}
