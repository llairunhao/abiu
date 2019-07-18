package abiu

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type DateFormat int8

const (
	DateTime DateFormat = iota
	DateOnly
	TimeOnly
	Timestamp
)

func NewDate(value time.Time) *Date {
	return &Date{value: value}
}

func CurrentDate() *Date {
	return NewDate(time.Now())
}

type Date struct {
	value  time.Time
	Format DateFormat
}

//-------------------------------------------------------
func (date *Date) SetValue(v interface{}) error {

	var ok bool
	var str string
	var timeSp int64

	date.value, ok = v.(time.Time)
	if ok {
		return nil
	}

	str, ok = v.(string)
	if ok {
		if date.trySetWithString(str) {
			return nil
		}
	}

	timeSp, ok = v.(int64)
	if ok {
		date.value = time.Unix(timeSp, 0)
		return nil
	}

	return typeError(v)
}

func (date *Date) Value() interface{} {
	return date.value
}

func (date *Date) ValuePtr() interface{} {
	return &date.value
}

//-------------------------------------------------------
func (date *Date) MarshalJSON() ([]byte, error) {
	switch date.Format {
	case DateTime:
		fallthrough
	case DateOnly:
		fallthrough
	case TimeOnly:
		fallthrough
	case Timestamp:
		return []byte(fmt.Sprintf("\"%v\"", date.value.String())), nil
	}
	return json.Marshal(date.value)
}

func (date *Date) UnmarshalJSON(b []byte) error {
	var err error
	switch date.Format {
	case DateTime:
		{
			return date.setWithString(string(b), DateTime)
		}
	case DateOnly:
		{
			return date.setWithString(string(b), DateOnly)
		}
	case TimeOnly:
		{
			return date.setWithString(string(b), TimeOnly)
		}
	case Timestamp:
		{
			return date.setWithString(string(b), Timestamp)
		}
	}
	return err
}

func (date *Date) String() string {
	switch date.Format {
	case DateTime:
		return date.value.Format("2006-01-02 15:04:05")
	case DateOnly:
		return date.value.Format("2006-01-02")
	case TimeOnly:
		return date.value.Format("15:04:05")
	case Timestamp:
		return fmt.Sprintf("%v", date.value.Unix())
	}
	return date.value.String()
}

func (date *Date) trySetWithString(str string) bool {
	var err error

	err = date.setWithString(str, DateTime)
	if err == nil {
		return true
	}

	err = date.setWithString(str, Timestamp)
	if err == nil {
		return true
	}

	err = date.setWithString(str, DateOnly)
	if err == nil {
		return true
	}

	err = date.setWithString(str, TimeOnly)
	if err == nil {
		return true
	}

	return false
}

func (date *Date) setWithString(str string, format DateFormat) error {
	var err error
	switch date.Format {
	case DateTime:
		{
			date.value, err = time.Parse("2006-01-02 15:04:05", str)
		}
	case DateOnly:
		{
			length := len(str)
			if length != 8 {
				return date.formatError(str)
			}
			dateString := fmt.Sprintf("%v %v", string(str[1:length-1]), "00:00:00")
			date.value, err = time.Parse("2006-01-02 15:04:05", dateString)
		}
	case TimeOnly:
		{
			length := len(str)
			if length < 8 {
				return date.formatError(str)
			}
			dateString := fmt.Sprintf("%v %v", "1970-01-01", string(str[1:length-1]))
			date.value, err = time.Parse("2006-01-02 15:04:05", dateString)
		}
	case Timestamp:
		{
			timeSp, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return err
			}
			date.value = time.Unix(timeSp, 0)
		}
	}
	return err
}

func (date *Date) formatError(str string) error {
	return errors.New(fmt.Sprintf("错误的时间格式:%v", str))
}
