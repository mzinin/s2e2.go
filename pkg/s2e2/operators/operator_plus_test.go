package operators

import "testing"

func TestOperatorPlus_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorPlus()
	expectedName := "+"

	if operator.Name() != expectedName {
		test.Errorf("Wrong operator name %v instead of %v", operator.Name(), expectedName)
	}
}

func TestOperatorPlus_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorPlus()

	if operator.Priority() != operatorPlusPriority {
		test.Errorf("Wrong operator priority %v instead of %v", operator.Name(), operatorAndPriority)
	}
}

func TestOperatorPlus_Positive_StringString_StackSize(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", "B"}
	expectedStackSize := 1

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestOperatorPlus_Positive_StringString_ResultType(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", "B"}

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, ok := stack[0].(string); !ok {
		test.Errorf("Wrong result type %T instead of %T", stack[0], "")
	}
}

func TestOperatorPlus_Positive_StringString_ResultValue(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", "B"}
	expectedValue := "AB"

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorPlus_Positive_StringNull_ResultValue(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", nil}
	expectedValue := "A"

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorPlus_Positive_NullString_ResultValue(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{nil, "B"}
	expectedValue := "B"

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorPlus_Positive_NullNull_ResultValue(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{nil, nil}
	var expectedValue interface{}

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if stack[0] != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", stack[0], expectedValue)
	}
}

func TestOperatorPlus_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", "B", "C"}
	expectedStackSize := 2

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestOperatorPlus_Negative_FewerArguments(test *testing.T) {
	function := NewOperatorPlus()
	stack := []interface{}{"A"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: not enough arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorPlus_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewOperatorPlus()
	stack := []interface{}{5, "B"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorPlus_Negative_SecondArgumentWrongType(test *testing.T) {
	function := NewOperatorPlus()
	stack := []interface{}{"A", 1}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
