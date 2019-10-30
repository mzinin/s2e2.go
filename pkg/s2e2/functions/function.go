package functions

// Function is the interface for all expression functions.
type Function interface {

	// Name gets name of the function.
	Name() string

	// Invoke calls the function i.e. pops all its arguments from the stack and puts result in.
	// Returns error in case of wrong number or type of arguments.
	Invoke(stack *[]interface{}) error
}
