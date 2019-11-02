package operators

// OperatorOr is operator ||
// Computes disjunction of two boolean values.
type OperatorOr struct {
	BaseOperator
}

// NewOperatorOr creates an instance of OperatorOr.
func NewOperatorOr() *OperatorOr {
	result := &OperatorOr{MakeBaseOperator(nil, "||", operatorOrPriority, 2)}
	result.derived = result
	return result
}

// CheckArguments checks if all arguments are correct.
func (o *OperatorOr) CheckArguments(arguments []interface{}) bool {
	_, ok1 := arguments[0].(bool)
	_, ok2 := arguments[1].(bool)
	return ok1 && ok2
}

// Result calculates result of the function for given arguments.
func (o *OperatorOr) Result(arguments []interface{}) interface{} {
	value1, _ := arguments[0].(bool)
	value2, _ := arguments[1].(bool)
	return value1 || value2
}
