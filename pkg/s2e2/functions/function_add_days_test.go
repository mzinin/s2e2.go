package functions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	secondsPerDay int64 = 86400
)

func TestFunctionAddDays_Positive_CreateFunction_Name(test *testing.T) {
	function := NewFunctionAddDays()
	expectedName := "ADD_DAYS"

	assert.Equal(test, expectedName, function.Name())
}

func TestFunctionAddDays_Positive_GoodArguments_StackSize(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC(), "1"}
	expectedStackSize := 1

	assert.NoError(test, function.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestFunctionAddDays_Positive_GoodArguments_ResultType(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC(), "1"}

	assert.NoError(test, function.Invoke(&stack))
	assert.IsType(test, time.Time{}, stack[0])
}

func TestFunctionAddDays_Positive_SecondArgumentPositive_ResultValue(test *testing.T) {
	function := NewFunctionAddDays()
	firstArgument := time.Date(2019, 7, 13, 12, 15, 0, 0, time.UTC)
	stack := []interface{}{firstArgument, "1"}

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(time.Time)
	assert.Equal(test, secondsPerDay, value.Unix()-firstArgument.Unix())
}

func TestFunctionAddDays_Positive_SecondArgumentZero_ResultValue(test *testing.T) {
	function := NewFunctionAddDays()
	firstArgument := time.Date(2019, 7, 13, 12, 15, 0, 0, time.UTC)
	stack := []interface{}{firstArgument, "0"}

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(time.Time)
	assert.Equal(test, firstArgument, value)
}

func TestFunctionAddDays_Positive_SecondArgumentNegative_ResultValue(test *testing.T) {
	function := NewFunctionAddDays()
	firstArgument := time.Date(2019, 7, 13, 12, 15, 0, 0, time.UTC)
	stack := []interface{}{firstArgument, "-1"}

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(time.Time)
	assert.Equal(test, secondsPerDay, firstArgument.Unix()-value.Unix())
}

func TestFunctionAddDays_Positive_MoreArguments_StackSize(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{"ARG", time.Now().UTC(), "1"}
	expectedStackSize := 2

	assert.NoError(test, function.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestFunctionAddDays_Negative_FewerArguments(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC()}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: not enough arguments for function "+function.Name(), err.Error())
}

func TestFunctionAddDays_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{"2019-07-13 00:00:00", "1"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionAddDays_Negative_FirstArgumentNull(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{nil, "1"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionAddDays_Negative_SecondArgumentWrongType(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC(), 1}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionAddDays_Negative_SecondArgumentWrongValue(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC(), "A"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionAddDays_Negative_SecondArgumentNull(test *testing.T) {
	function := NewFunctionAddDays()
	stack := []interface{}{time.Now().UTC(), nil}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}
