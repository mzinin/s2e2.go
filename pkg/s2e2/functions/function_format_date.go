package functions

import (
	"time"
)

// FunctionFormatDate is FORMAT_DATE(<datetime>, <format>)
// Converts datetime to string according to format.
type FunctionFormatDate struct {
	BaseFunction
}

// NewFunctionFormatDate creates an instance of FunctionFormatDate.
func NewFunctionFormatDate() *FunctionFormatDate {
	result := &FunctionFormatDate{MakeBaseFunction(nil, "FORMAT_DATE", 2)}
	result.derived = result
	return result
}

// CheckArguments checks if all arguments are correct.
func (f *FunctionFormatDate) CheckArguments(arguments []interface{}) bool {
	// check 1st argument
	if _, ok := arguments[0].(time.Time); !ok {
		return false
	}

	// check 2nd argument
	if _, ok := arguments[1].(string); !ok {
		return false
	}

	return true
}

// Result calculates result of the function for given arguments.
func (f *FunctionFormatDate) Result(arguments []interface{}) interface{} {
	datetime, _ := arguments[0].(time.Time)
	format, _ := arguments[1].(string)
	return datetime.Format(format)
}
