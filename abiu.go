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
