package operators

// OperatorLess is operator <
// Lexicographically compares two strings.
type OperatorLess struct {
	BaseOperator
}

// NewOperatorLess creates an instance of OperatorLess.
func NewOperatorLess() *OperatorLess {
	result := &OperatorLess{MakeBaseOperator(nil, "<", operatorLessPriority, 2)}
	result.SetDerived(result)
	return result
}

// CheckArguments checks if all arguments are correct.
func (o *OperatorLess) CheckArguments(arguments []interface{}) bool {
	_, ok1 := arguments[0].(string)
	_, ok2 := arguments[1].(string)
	return ok1 && ok2
}

// Result calculates result of the function for given arguments.
func (o *OperatorLess) Result(arguments []interface{}) interface{} {
	arg1, _ := arguments[0].(string)
	arg2, _ := arguments[1].(string)
	return arg1 < arg2
}
