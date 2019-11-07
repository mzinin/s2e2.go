package functions

import "fmt"

// BaseFunction is the base structure for all standard functions.
type BaseFunction struct {
	derived   DerivedFunction // Derived function i.e. concrete implementation.
	name      string          // Function's name.
	arguments []interface{}   // List of arguments.
}

// MakeBaseFunction creates an instance of base function.
func MakeBaseFunction(derived DerivedFunction, name string, numberOfArguments int) BaseFunction {
	return BaseFunction{derived, name, make([]interface{}, numberOfArguments)}
}

// SetDerived sets derived function.
func (f *BaseFunction) SetDerived(derived DerivedFunction) {
	f.derived = derived
}

// Name gets name of the function.
func (f *BaseFunction) Name() string {
	return f.name
}

// Invoke calls the function i.e. pops all its arguments from the stack and puts result in.
// Returns error in case of wrong number or type of arguments.
func (f *BaseFunction) Invoke(stackPointer *[]interface{}) error {
	stack := *stackPointer

	if len(stack) < len(f.arguments) {
		return fmt.Errorf("BaseFunction: not enough arguments for function %v", f.name)
	}

	for i := 0; i < len(f.arguments); i++ {
		f.arguments[i] = stack[len(stack)-len(f.arguments)+i]
	}
	stack = stack[:len(stack)-len(f.arguments)]

	if !f.derived.CheckArguments(f.arguments) {
		return fmt.Errorf("BaseFunction: invalid arguments for function %v", f.name)
	}

	result := f.derived.Result(f.arguments)
	*stackPointer = append(stack, result)

	return nil
}
