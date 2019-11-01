package functions

import (
	"testing"
	"time"
)

const (
	secondsPerDay int64 = 86400
)

func TestFunctionAddDays_Positive_CreateFunction_Name(test *testing.T) {
	function := NewFunctionAddDays()
	expectedName := "ADD_DAYS"

	if function.Name() != expectedName {
		test.Errorf("Wrong function name %v instead of %v", function.Name(), expectedName)
	}
}

func TestFunctionAddDays_Positive_GoodArguments_StackSize(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC(), "1"}
	expectedStackSize := 1

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestFunctionAddDays_Positive_GoodArguments_ResultType(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC(), "1"}

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, ok := stack[0].(time.Time); !ok {
		test.Errorf("Wrong result type %T instead of %T", stack[0], "")
	}
}

func TestFunctionAddDays_Positive_SecondArgumentPositive_ResultType(test *testing.T) {
	function := NewFunctionAddDays()
	firstArgument := time.Date(2019, 7, 13, 12, 15, 0, 0, time.UTC)
	stack := []interface{}{firstArgument, "1"}

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	result, _ := stack[0].(time.Time)

	if result.Unix()-firstArgument.Unix() != secondsPerDay {
		test.Errorf("Wrong result value %v", result)
	}
}

func TestFunctionAddDays_Positive_SecondArgumentZero_ResultType(test *testing.T) {
	function := NewFunctionAddDays()
	firstArgument := time.Date(2019, 7, 13, 12, 15, 0, 0, time.UTC)
	stack := []interface{}{firstArgument, "0"}

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	result, _ := stack[0].(time.Time)

	if result != firstArgument {
		test.Errorf("Wrong result value %v instead of %v", result, firstArgument)
	}
}

func TestFunctionAddDays_Positive_SecondArgumentNegative_ResultType(test *testing.T) {
	function := NewFunctionAddDays()
	firstArgument := time.Date(2019, 7, 13, 12, 15, 0, 0, time.UTC)
	stack := []interface{}{firstArgument, "-1"}

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	result, _ := stack[0].(time.Time)

	if firstArgument.Unix()-result.Unix() != secondsPerDay {
		test.Errorf("Wrong result value %v", result)
	}
}

func TestFunctionAddDays_Positive_MoreArguments_StackSize(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{"ARG", time.Now().UTC(), "1"}
	expectedStackSize := 2

	if err := function.Invoke(&stack); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(stack) != expectedStackSize {
		test.Errorf("Wrong stack size %v instead of %v", len(stack), expectedStackSize)
	}
}

func TestFunctionAddDays_Negative_FewerArguments(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC()}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: not enough arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionAddDays_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{"2019-07-13 00:00:00", "1"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionAddDays_Negative_FirstArgumentNull(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{nil, "1"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionAddDays_Negative_SecondArgumentWrongType(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC(), 1}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionAddDays_Negative_SecondArgumentWrongValue(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC(), "A"}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestFunctionAddDays_Negative_SecondArgumentNull(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC(), nil}

	if err := function.Invoke(&stack); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "BaseFunction: invalid arguments for function "+function.Name() {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
