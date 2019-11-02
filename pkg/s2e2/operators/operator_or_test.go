package operators

import "testing"

func TestOperatorOr_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorOr()
	expectedName := "||"

	if operator.Name() != expectedName {
		test.Errorf("Wrong operator name %v instead of %v", operator.Name(), expectedName)
	}
}

func TestOperatorOr_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorOr()

	if operator.Priority() != operatorOrPriority {
		test.Errorf("Wrong operator priority %v instead of %v", operator.Name(), operatorAndPriority)
	}
}

func TestOperatorOr_Positive_GoorArguments_StackSize(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true, true}
	expectedStackSize := 1

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestOperatorOr_Positive_GoorArguments_ResultType(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true, true}

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, ok := stack[0].(bool); !ok {
		test.Errorf("Wrong result type %T instead of %T", stack[0], "")
	}
}

func TestOperatorOr_Positive_TrueTrue_ResultValue(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true, true}
	expectedValue := true

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorOr_Positive_TrueFalse_ResultValue(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true, false}
	expectedValue := true

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorOr_Positive_FalseTrue_ResultValue(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{false, true}
	expectedValue := true

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorOr_Positive_FalseFalse_ResultValue(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{false, false}
	expectedValue := false

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(bool); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestOperatorOr_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{"ARG", false, true}
	expectedStackSize := 2

	if err := operator.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestOperatorOr_Negative_FewerArguments(test *testing.T) {
	function := NewOperatorOr()
	stack := []interface{}{true}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: not enough arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorOr_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewOperatorOr()
	stack := []interface{}{"true", true}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorOr_Negative_FirstArgumentNull(test *testing.T) {
	function := NewOperatorOr()
	stack := []interface{}{nil, true}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorOr_Negative_SecondArgumentWrongType(test *testing.T) {
	function := NewOperatorOr()
	stack := []interface{}{true, "true"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestOperatorOr_Negative_SecondArgumentNull(test *testing.T) {
	function := NewOperatorOr()
	stack := []interface{}{true, nil}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseOperator: invalid arguments for operator "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
