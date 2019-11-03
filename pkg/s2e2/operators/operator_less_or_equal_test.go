package operators

import "testing"

func TestOperatorLessOrEqual_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	expectedName := "<="

	if operator.Name() != expectedName {
		test.Errorf("Wrong operator name %v instead of %v", operator.Name(), expectedName)
	}
}

func TestOperatorLessOrEqual_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorLessOrEqual()

	if operator.Priority() != operatorLessOrEqualPriority {
		test.Errorf("Wrong operator priority %v instead of %v", operator.Name(), operatorLessOrEqualPriority)
	}
}

func TestOperatorLessOrEqual_Positive_GoorArguments_StackSize(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", "String2"}
	expectedStackSize := 1

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestOperatorLessOrEqual_Positive_GoorArguments_ResultType(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", "String2"}

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, ok := stack[0].(bool); !ok {
		test.Errorf("Wrong result type %T instead of %T", stack[0], "")
	}
}

func TestOperatorLessOrEqual_Positive_EqualStrings_ResultType(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", "String1"}
	expectedValue := true

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorLessOrEqual_Positive_FirstArgumentLess_ResultType(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", "String2"}
	expectedValue := true

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorLessOrEqual_Positive_SecondArgumentLess_ResultType(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String2", "String1"}
	expectedValue := false

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorLessOrEqual_Positive_BothArgumentsNull_ResultType(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{nil, nil}
	expectedValue := true

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorLessOrEqual_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", "String2", "String3"}
	expectedStackSize := 2

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestOperatorLessOrEqual_Negative_FewerArguments(test *testing.T) {
	function := NewOperatorLessOrEqual()
	stack := []interface{}{"String1"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: not enough arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorLessOrEqual_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewOperatorLessOrEqual()
	stack := []interface{}{5, "5"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorLessOrEqual_Negative_SecondArgumentWrongType(test *testing.T) {
	function := NewOperatorLessOrEqual()
	stack := []interface{}{"1", 2}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorLessOrEqual_Negative_BothArgumentsWrongType(test *testing.T) {
	function := NewOperatorLessOrEqual()
	stack := []interface{}{1, 2}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorLessOrEqual_Negative_FirstArgumentNull(test *testing.T) {
	function := NewOperatorLessOrEqual()
	stack := []interface{}{nil, "String2"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorLessOrEqual_Negative_SecondArgumentNull(test *testing.T) {
	function := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", nil}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
