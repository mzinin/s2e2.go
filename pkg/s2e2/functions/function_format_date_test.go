package functions

import (
	"testing"
	"time"
)

func TestFunctionFormatDate_Positive_CreateFunction_Name(test *testing.T) {
	function := NewFunctionFormatDate()
	expectedName := "FORMAT_DATE"

	if function.Name() != expectedName {
		test.Errorf("Wrong function name %v instead of %v", function.Name(), expectedName)
	}
}

func TestFunctionFormatDate_Positive_GoodArguments_StackSize(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Now().UTC(), "2006-01-02"}
	expectedStackSize := 1

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestFunctionFormatDate_Positive_GoodArguments_ResultType(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Date(2019, 7, 13, 12, 15, 0, 0, time.UTC), "2006-01-02 15:04:05"}

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, ok := stack[0].(string); !ok {
		test.Errorf("Wrong result type %T instead of %T", stack[0], "")
	}
}

func TestFunctionFormatDate_Positive_GoodArguments_ResultValue(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Date(2019, 7, 13, 12, 15, 0, 0, time.UTC), "2006-01-02 15:04:05"}
	expectedValue := "2019-07-13 12:15:00"

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestFunctionFormatDate_Positive_MoreArguments_StackSize(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{false, time.Now().UTC(), "2006-01-02"}
	expectedStackSize := 2

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestFunctionFormatDate_Negative_FewerArguments(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Now().UTC()}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: not enough arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionFormatDate_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{"2019-07-13", "2006-01-02"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionFormatDate_Negative_FirstArgumentNull(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{nil, "2006-01-02"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionFormatDate_Negative_SecondArgumentWrongType(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Now().UTC(), 15}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionFormatDate_Positive_SecondArgumentWrongValue(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Now().UTC(), "year-month-day"}
	expectedValue := "year-month-day"

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value, _ := stack[0].(string); value != expectedValue {
		test.Errorf("Wrong result value %v instead of %v", value, expectedValue)
	}
}

func TestFunctionFormatDate_Negative_SecondArgumentNull(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Now().UTC(), nil}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
