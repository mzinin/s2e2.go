package functions

import (
	"testing"
	"time"
)

func TestFunctionNow_Positive_CreateFunction_Name(test *testing.T) {
	function := NewFunctionNow()
	expectedName := "NOW"

	if function.Name() != expectedName {
		test.Errorf("Wrong function name %v instead of %v", function.Name(), expectedName)
	}
}

func TestFunctionNow_Positive_StackSize(test *testing.T) {
	function := NewFunctionNow()
	stack := []interface{}{}
	expectedStackSize := 1

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestFunctionNow_Positive_ResultType(test *testing.T) {
	function := NewFunctionNow()
	stack := []interface{}{}

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, ok := stack[0].(time.Time); !ok {
		test.Errorf("Wrong result type %T instead of %T", stack[0], "")
	}
}

func TestFunctionNow_Positive_ResultValue(test *testing.T) {
	function := NewFunctionNow()
	stack := []interface{}{}
	maxDifferenceInSeconds := 2.0

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	now := time.Now().UTC()
	functionResult, _ := stack[0].(time.Time)

	if functionResult.After(now) {
		test.Errorf("Function result is after now")
	}
	if functionResult.Sub(now).Seconds() > maxDifferenceInSeconds {
		test.Errorf("Function result is before now")
	}
}

func TestFunctionNow_Positive_MoreArguments_StackSize(test *testing.T) {
	function := NewFunctionNow()
	stack := []interface{}{false, "A", "B"}
	expectedStackSize := 4

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}
