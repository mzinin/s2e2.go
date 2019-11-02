package operators

// Operator is the interface for all expression operators.
type Operator interface {

	// Name gets the name of the operator.
	Name() string

	// Priority gets the priotity of the operator.
	Priority() int

	// Invoke calls the operator i.e. pops all its arguments from the stack and puts result in.
	// Returns error in case of wrong number or type of arguments.
	Invoke(stack *[]interface{}) error
}
