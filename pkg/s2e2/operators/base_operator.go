package operators

import "fmt"

// BaseOperator is the base structure for all standard operators.
type BaseOperator struct {
	derived   DerivedOperator // Derived operator i.e. concrete implementation.
	name      string          // Operator's name.
	priority  int             // Operator's priority.
	arguments []interface{}   // List of arguments.
}

// MakeBaseOperator creates an instance of base operator.
func MakeBaseOperator(derived DerivedOperator, name string, priority int, numberOfArguments int) BaseOperator {
	return BaseOperator{derived, name, priority, make([]interface{}, numberOfArguments)}
}

// Name gets the name of the operator.
func (o *BaseOperator) Name() string {
	return o.name
}

// Priority gets the priotity of the operator.
func (o *BaseOperator) Priority() int {
	return o.priority
}

// Invoke calls the operator i.e. pops all its arguments from the stack and puts result in.
// Returns error in case of wrong number or type of arguments.
func (o *BaseOperator) Invoke(stackPointer *[]interface{}) error {
	stack := *stackPointer

	if len(stack) < len(o.arguments) {
		return fmt.Errorf("BaseOperator: not enough arguments for operator %v", o.name)
	}

	for i := 0; i < len(o.arguments); i++ {
		o.arguments[i] = stack[len(stack)-len(o.arguments)+i]
	}
	stack = stack[:len(stack)-len(o.arguments)]

	if !o.derived.CheckArguments(o.arguments) {
		return fmt.Errorf("BaseOperator: invalid arguments for operator %v", o.name)
	}

	result := o.derived.Result(o.arguments)
	*stackPointer = append(stack, result)

	return nil
}
