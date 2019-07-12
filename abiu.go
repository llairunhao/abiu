package abiu

import "encoding/json"

type String struct {
	Value string
}

func (s *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Value)
}

func (s *String) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.Value)
}

func (s *String) IsEmpty() bool {
	if s == nil {
		return true
	}
	return len(s.Value) == 0
}

//---------------------------
type Byte struct {
	Value byte
}

func (number *Byte) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *Byte) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *Byte) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type UInt8 struct {
	Value uint8
}

func (number *UInt8) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *UInt8) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *UInt8) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type UInt16 struct {
	Value uint16
}

func (number *UInt16) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *UInt16) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *UInt16) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type UInt32 struct {
	Value uint8
}

func (number *UInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *UInt32) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *UInt32) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type UInt64 struct {
	Value uint64
}

func (number *UInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *UInt64) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *UInt64) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type UInt struct {
	Value uint
}

func (number *UInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *UInt) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *UInt) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type Int8 struct {
	Value int8
}

func (number *Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *Int8) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *Int8) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type Int16 struct {
	Value int16
}

func (number *Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *Int16) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *Int16) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type Int32 struct {
	Value int32
}

func (number *Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *Int32) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *Int32) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type Int64 struct {
	Value int64
}

func (number *Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *Int64) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *Int64) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type Int struct {
	Value int
}

func (number *Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *Int) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *Int) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type Float struct {
	Value float32
}

func (number *Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *Float) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *Float) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}

//---------------------------
type Double struct {
	Value float64
}

func (number *Double) MarshalJSON() ([]byte, error) {
	return json.Marshal(number.Value)
}

func (number *Double) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &number.Value)
}

func (number *Double) IsZero() bool {
	if number == nil {
		return true
	}
	return number.Value == 0
}
