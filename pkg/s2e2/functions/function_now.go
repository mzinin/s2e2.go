package functions

import "time"

// FunctionNow is NOW()
// Returns current UTC datetime.
type FunctionNow struct {
	BaseFunction
}

// NewFunctionNow creates an instance of FunctionNow.
func NewFunctionNow() *FunctionNow {
	result := &FunctionNow{MakeBaseFunction(nil, "NOW", 0)}
	result.derived = result
	return result
}

// CheckArguments checks if all arguments are correct.
func (f *FunctionNow) CheckArguments(arguments []interface{}) bool {
	return true
}

// Result calculates result of the function for given arguments.
func (f *FunctionNow) Result(arguments []interface{}) interface{} {
	return time.Now().UTC()
}
