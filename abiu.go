package abiu

type ABType interface {
	SetValue(v interface{})

	Value() interface{}

	ValuePtr() interface{}
}
