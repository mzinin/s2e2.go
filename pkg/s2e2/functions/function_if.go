package functions

// FunctionIf is IF(<conition>, <value1>, <value2>)
// Returns value1 if boolean condition is true, and value2 otherwise.
type FunctionIf struct {
	BaseFunction
}

// NewFunctionIf creates an instance of FunctionIf.
func NewFunctionIf() *FunctionIf {
	result := &FunctionIf{MakeBaseFunction(nil, "IF", 3)}
	result.derived = result
	return result
}

// CheckArguments checks if all arguments are correct.
func (f *FunctionIf) CheckArguments(arguments []interface{}) bool {
	_, ok := arguments[0].(bool)
	return ok
}

// Result calculates result of the function for given arguments.
func (f *FunctionIf) Result(arguments []interface{}) interface{} {
	if arguments[0].(bool) {
		return arguments[1]
	}
	return arguments[2]
}
