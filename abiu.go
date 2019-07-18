package abiu

import (
	"errors"
	"fmt"
	"strconv"
)

type ABType interface {
	SetValue(v interface{}) error

	Value() interface{}

	ValuePtr() interface{}
}

func typeError(v interface{}) error {
	return errors.New(fmt.Sprintf("未知类型：%v", v))
}

func uintValue(str string, bitSize int) (uint64, error) {
	if str == "" {
		return 0, nil
	}
	return strconv.ParseUint(str, 10, bitSize)
}

func intValue(str string, bitSize int) (int64, error) {
	if str == "" {
		return 0, nil
	}
	return strconv.ParseInt(str, 10, bitSize)
}

func floatValue(str string, bitSize int) (float64, error) {
	if str == "" {
		return 0, nil
	}
	return strconv.ParseFloat(str, bitSize)
}
