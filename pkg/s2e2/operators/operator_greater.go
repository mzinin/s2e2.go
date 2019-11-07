package operators

// OperatorGreater is operator >
// Lexicographically compares two strings.
type OperatorGreater struct {
	BaseOperator
}

// NewOperatorGreater creates an instance of OperatorGreater.
func NewOperatorGreater() *OperatorGreater {
	result := &OperatorGreater{MakeBaseOperator(nil, ">", operatorGreaterPriority, 2)}
	result.SetDerived(result)
	return result
}

// CheckArguments checks if all arguments are correct.
func (o *OperatorGreater) CheckArguments(arguments []interface{}) bool {
	_, ok1 := arguments[0].(string)
	_, ok2 := arguments[1].(string)
	return ok1 && ok2
}

// Result calculates result of the function for given arguments.
func (o *OperatorGreater) Result(arguments []interface{}) interface{} {
	arg1, _ := arguments[0].(string)
	arg2, _ := arguments[1].(string)
	return arg1 > arg2
}
