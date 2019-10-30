package functions

// DerivedFunction is the interface for expression functions which use BaseFunction
// as part or their implementation
type DerivedFunction interface {

	// CheckArguments checks if all arguments are correct.
	CheckArguments(arguments []interface{}) bool

	// Result calculates result of the function for given arguments.
	Result(arguments []interface{}) interface{}
}
