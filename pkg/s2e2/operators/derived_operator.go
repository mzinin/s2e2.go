package operators

// DerivedOperator is the interface for expression operators which use BaseOperator
// as part or their implementation
type DerivedOperator interface {

	// CheckArguments checks if all arguments are correct.
	CheckArguments(arguments []interface{}) bool

	// Result calculates result of the function for given arguments.
	Result(arguments []interface{}) interface{}
}
