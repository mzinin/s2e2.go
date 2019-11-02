package operators

import "testing"

func TestOperatorEqual_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorEqual()
	expectedName := "=="

	if operator.Name() != expectedName {
		test.Errorf("Wrong operator name %v instead of %v", operator.Name(), expectedName)
	}
}

func TestOperatorEqual_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorEqual()

	if operator.Priority() != operatorEqualPriority {
		test.Errorf("Wrong operator priority %v instead of %v", operator.Name(), operatorEqualPriority)
	}
}

func TestOperatorEqual_Positive_GoorArguments_StackSize(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", "String2"}
	expectedStackSize := 1

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestOperatorEqual_Positive_GoorArguments_ResultType(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", "String2"}

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, ok := stack[0].(bool); !ok {
		test.Errorf("Wrong result type %T instead of %T", stack[0], "")
	}
}

func TestOperatorEqual_Positive_EqualStrings_ResultType(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", "String1"}
	expectedValue := true

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorEqual_Positive_DifferentStrings_ResultType(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", "String2"}
	expectedValue := false

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorEqual_Positive_FirstArgumentNull_ResultType(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{nil, "String2"}
	expectedValue := false

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorEqual_Positive_SecondArgumentNull_ResultType(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", nil}
	expectedValue := false

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorEqual_Positive_BothArgumentsNull_ResultType(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{nil, nil}
	expectedValue := true

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorEqual_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", "String2", "String3"}
	expectedStackSize := 2

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestOperatorEqual_Negative_FewerArguments(test *testing.T) {
	function := NewOperatorEqual()
	stack := []interface{}{"String1"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: not enough arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorEqual_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewOperatorEqual()
	stack := []interface{}{5, "5"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorEqual_Negative_SecondArgumentWrongType(test *testing.T) {
	function := NewOperatorEqual()
	stack := []interface{}{"1", 1}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorEqual_Negative_BothArgumentsWrongType(test *testing.T) {
	function := NewOperatorEqual()
	stack := []interface{}{1, 1}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
