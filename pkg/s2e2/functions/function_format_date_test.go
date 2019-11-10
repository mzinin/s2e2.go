package functions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFunctionFormatDate_Positive_CreateFunction_Name(test *testing.T) {
	function := NewFunctionFormatDate()
	expectedName := "FORMAT_DATE"

	assert.Equal(test, expectedName, function.Name())
}

func TestFunctionFormatDate_Positive_GoodArguments_StackSize(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Now().UTC(), "2006-01-02"}
	expectedStackSize := 1

	assert.NoError(test, function.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestFunctionFormatDate_Positive_GoodArguments_ResultType(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Date(2019, 7, 13, 12, 15, 0, 0, time.UTC), "2006-01-02 15:04:05"}

	assert.NoError(test, function.Invoke(&stack))
	assert.IsType(test, "string", stack[0])
}

func TestFunctionFormatDate_Positive_GoodArguments_ResultValue(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Date(2019, 7, 13, 12, 15, 0, 0, time.UTC), "2006-01-02 15:04:05"}
	expectedValue := "2019-07-13 12:15:00"

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestFunctionFormatDate_Positive_MoreArguments_StackSize(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{false, time.Now().UTC(), "2006-01-02"}
	expectedStackSize := 2

	assert.NoError(test, function.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestFunctionFormatDate_Negative_FewerArguments(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Now().UTC()}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: not enough arguments for function "+function.Name(), err.Error())
}

func TestFunctionFormatDate_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{"2019-07-13", "2006-01-02"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionFormatDate_Negative_FirstArgumentNull(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{nil, "2006-01-02"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionFormatDate_Negative_SecondArgumentWrongType(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Now().UTC(), 15}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionFormatDate_Positive_SecondArgumentWrongValue(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Now().UTC(), "year-month-day"}
	expectedValue := "year-month-day"

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestFunctionFormatDate_Negative_SecondArgumentNull(test *testing.T) {
	function := NewFunctionFormatDate()
	stack := []interface{}{time.Now().UTC(), nil}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}
