package functions

import "testing"

func TestFunctionIf_Positive_CreateFunction_Name(test *testing.T) {
	function := NewFunctionIf()
	expectedName := "IF"

	if function.Name() != expectedName {
		test.Errorf("Wrong function name %v instead of %v", function.Name(), expectedName)
	}
}

func TestFunctionIf_Positive_FirstArgumentTrue_StackSize(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{true, "A", "B"}
	expectedStackSize := 1

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestFunctionIf_Positive_FirstArgumentTrue_ResultType(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{true, "A", "B"}

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, ok := stack[0].(string); !ok {
		test.Errorf("Wrong result type %T instead of %T", stack[0], "")
	}
}

func TestFunctionIf_Positive_FirstArgumentTrue_ResultValue(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{true, "A", "B"}
	expectedValue := "A"

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestFunctionIf_Positive_FirstArgumentFalse_ResultValue(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{false, "A", "B"}
	expectedValue := "B"

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestFunctionIf_Positive_MoreArguments_StackSize(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{"ARG", false, "A", "B"}
	expectedStackSize := 2

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestFunctionIf_Negative_FewerArguments(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{false}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: not enough arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionIf_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{"false", "A", "B"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionIf_Negative_FirstArgumentNull(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{nil, "A", "B"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
