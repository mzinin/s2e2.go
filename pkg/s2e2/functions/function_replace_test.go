package functions

import "testing"

func TestFunctionReplace_Positive_CreateFunction_Name(test *testing.T) {
	function := NewFunctionReplace()
	expectedName := "REPLACE"

	if function.Name() != expectedName {
		test.Errorf("Wrong function name %v instead of %v", function.Name(), expectedName)
	}
}

func TestFunctionReplace_Positive_StringReplace_StackSize(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "A", "B"}
	expectedStackSize := 1

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestFunctionReplace_Positive_StringReplace_ResultType(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "A", "B"}

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, ok := stack[0].(string); !ok {
		test.Errorf("Wrong result type %T instead of %T", stack[0], "")
	}
}

func TestFunctionReplace_Positive_StringReplace_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "A", "B"}
	expectedValue := "BBB"

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestFunctionReplace_Positive_RegexReplace_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABCABA", "A.*?C", "D"}
	expectedValue := "DABA"

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestFunctionReplace_Positive_SpecialSymbolReplace_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"A * B == C", "\\*", "+"}
	expectedValue := "A + B == C"

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestFunctionReplace_Positive_FirstArgumentNull_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{nil, "A", "B"}
	var expectedValue interface{}

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if stack[0] != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", stack[0], expectedValue)
	}
}

func TestFunctionReplace_Positive_FirstArgumentEmptyString_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"", "A", "B"}
	expectedValue := ""

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestFunctionReplace_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{5, "A", "B"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionReplace_Negative_SecondArgumentEmptyString(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "", "B"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionReplace_Negative_SecondArgumentNull(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", nil, "B"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionReplace_Negative_SecondArgumentWrongType(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"AB5", 5, "B"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionReplace_Positive_ThirdArgumentEmptyString_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "B", ""}
	expectedValue := "AA"

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestFunctionReplace_Negative_ThirdArgumentNull(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "B", nil}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionReplace_Negative_ThirdArgumentWrongType(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "B", 5}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionReplace_Positive_MoreArguments_StackSize(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{false, "ABA", "A", "B"}
	expectedStackSize := 2

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestFunctionReplace_Negative_FewerArguments(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "A"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: not enough arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
