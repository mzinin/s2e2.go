package operators

// OperatorNot is operator !
//  Negates boolean value.
type OperatorNot struct {
	BaseOperator
}

// NewOperatorNot creates an instance of OperatorNot.
func NewOperatorNot() *OperatorNot {
	result := &OperatorNot{MakeBaseOperator(nil, "!", operatorNotPriority, 1)}
	result.derived = result
	return result
}

// CheckArguments checks if all arguments are correct.
func (o *OperatorNot) CheckArguments(arguments []interface{}) bool {
	_, ok := arguments[0].(bool)
	return ok
}

// Result calculates result of the function for given arguments.
func (o *OperatorNot) Result(arguments []interface{}) interface{} {
	value, _ := arguments[0].(bool)
	return !value
}
