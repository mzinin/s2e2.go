package operators

// OperatorGreaterOrEqual is operator >=
// Lexicographically compares two strings.
type OperatorGreaterOrEqual struct {
	BaseOperator
}

// NewOperatorGreaterOrEqual creates an instance of OperatorGreaterOrEqual.
func NewOperatorGreaterOrEqual() *OperatorGreaterOrEqual {
	result := &OperatorGreaterOrEqual{MakeBaseOperator(nil, ">=", operatorGreaterOrEqualPriority, 2)}
	result.SetDerived(result)
	return result
}

// CheckArguments checks if all arguments are correct.
func (o *OperatorGreaterOrEqual) CheckArguments(arguments []interface{}) bool {
	if arguments[0] == nil && arguments[1] == nil {
		return true
	}

	_, ok1 := arguments[0].(string)
	_, ok2 := arguments[1].(string)
	return ok1 && ok2
}

// Result calculates result of the function for given arguments.
func (o *OperatorGreaterOrEqual) Result(arguments []interface{}) interface{} {
	if arguments[0] == nil {
		return arguments[1] == nil
	}

	if arguments[1] == nil {
		return arguments[0] == nil
	}

	arg1, _ := arguments[0].(string)
	arg2, _ := arguments[1].(string)
	return arg1 >= arg2
}
