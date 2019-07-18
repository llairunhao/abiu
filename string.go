package abiu

import (
	"encoding/json"
	"fmt"
)

func NewString(value string) *String {
	return &String{value: value}
}

type String struct {
	value string
}

//-------------------------------------------------------
func (s *String) SetValue(v interface{}) error {
	var ok bool
	if s.value, ok = v.(string); !ok {
		s.value = fmt.Sprintf("%v", v)
	}
	return nil
}

func (s *String) Value() interface{} {
	return s.value
}

func (s *String) ValuePtr() interface{} {
	return &s.value
}

//-------------------------------------------------------
func (s *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

func (s *String) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.value)
}

func (s *String) String() string {
	return s.value
}

//-------------------------------------------------------
func (s *String) IsEmpty() bool {
	if s == nil {
		return true
	}
	return len(s.value) == 0
}
