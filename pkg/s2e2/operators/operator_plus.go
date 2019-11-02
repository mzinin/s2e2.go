package operators

// OperatorPlus is operator +
// Concatenates two strings.
type OperatorPlus struct {
	BaseOperator
}

// NewOperatorPlus creates an instance of OperatorPlus.
func NewOperatorPlus() *OperatorPlus {
	result := &OperatorPlus{MakeBaseOperator(nil, "+", operatorPlusPriority, 2)}
	result.derived = result
	return result
}

// CheckArguments checks if all arguments are correct.
func (o *OperatorPlus) CheckArguments(arguments []interface{}) bool {
	ok1 := arguments[0] == nil
	if !ok1 {
		_, ok1 = arguments[0].(string)
	}

	ok2 := arguments[1] == nil
	if !ok2 {
		_, ok2 = arguments[1].(string)
	}

	return ok1 && ok2
}

// Result calculates result of the function for given arguments.
func (o *OperatorPlus) Result(arguments []interface{}) interface{} {
	if arguments[0] == nil && arguments[1] == nil {
		return nil
	}

	arg1 := ""
	if arguments[0] != nil {
		arg1, _ = arguments[0].(string)
	}

	arg2 := ""
	if arguments[1] != nil {
		arg2, _ = arguments[1].(string)
	}

	return arg1 + arg2
}
