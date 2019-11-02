package operators

// OperatorAnd is operator &&
// Computes conjunction of two boolean values.
type OperatorAnd struct {
	BaseOperator
}

// NewOperatorAnd creates an instance of OperatorAnd.
func NewOperatorAnd() *OperatorAnd {
	result := &OperatorAnd{MakeBaseOperator(nil, "&&", operatorAndPriority, 2)}
	result.derived = result
	return result
}

// CheckArguments checks if all arguments are correct.
func (o *OperatorAnd) CheckArguments(arguments []interface{}) bool {
	_, ok1 := arguments[0].(bool)
	_, ok2 := arguments[1].(bool)
	return ok1 && ok2
}

// Result calculates result of the function for given arguments.
func (o *OperatorAnd) Result(arguments []interface{}) interface{} {
	value1, _ := arguments[0].(bool)
	value2, _ := arguments[1].(bool)
	return value1 && value2
}
