package operators

import "testing"

func TestOperatorNot_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorNot()
	expectedName := "!"

	if operator.Name() != expectedName {
		test.Errorf("Wrong operator name %v instead of %v", operator.Name(), expectedName)
	}
}

func TestOperatorNot_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorNot()

	if operator.Priority() != operatorNotPriority {
		test.Errorf("Wrong operator priority %v instead of %v", operator.Name(), operatorAndPriority)
	}
}

func TestOperatorNot_Positive_GoorArguments_StackSize(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{true}
	expectedStackSize := 1

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestOperatorNot_Positive_GoorArguments_ResultType(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{true}

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, ok := stack[0].(bool); !ok {
		test.Errorf("Wrong result type %T instead of %T", stack[0], "")
	}
}

func TestOperatorNot_Positive_ArgumentTrue_ResultValue(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{true}
	expectedValue := false

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorNot_Positive_ArgumentFalse_ResultValue(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{false}
	expectedValue := true

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorNot_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{false, true}
	expectedStackSize := 2

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestOperatorNot_Negative_FewerArguments(test *testing.T) {
	function := NewOperatorNot()
	stack := []interface{}{}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: not enough arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorNot_Negative_ArgumentWrongType(test *testing.T) {
	function := NewOperatorNot()
	stack := []interface{}{"true"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorNot_Negative_ArgumentNull(test *testing.T) {
	function := NewOperatorNot()
	stack := []interface{}{nil}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
