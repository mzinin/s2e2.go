package operators

// OperatorNotEqual is operator !=
// Compares two strings.
type OperatorNotEqual struct {
	BaseOperator
}

// NewOperatorNotEqual creates an instance of OperatorNotEqual.
func NewOperatorNotEqual() *OperatorNotEqual {
	result := &OperatorNotEqual{MakeBaseOperator(nil, "!=", operatorNotEqualPriority, 2)}
	result.derived = result
	return result
}

// CheckArguments checks if all arguments are correct.
func (o *OperatorNotEqual) CheckArguments(arguments []interface{}) bool {
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
func (o *OperatorNotEqual) Result(arguments []interface{}) interface{} {
	if arguments[0] == nil {
		return arguments[1] != nil
	}

	if arguments[1] == nil {
		return arguments[0] != nil
	}

	arg1, _ := arguments[0].(string)
	arg2, _ := arguments[1].(string)
	return arg1 != arg2
}
