package functions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFunctionNow_Positive_CreateFunction_Name(test *testing.T) {
	function := NewFunctionNow()
	expectedName := "NOW"

	assert.Equal(test, expectedName, function.Name())
}

func TestFunctionNow_Positive_StackSize(test *testing.T) {
	function := NewFunctionNow()
	stack := []interface{}{}
	expectedStackSize := 1

	assert.NoError(test, function.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestFunctionNow_Positive_ResultType(test *testing.T) {
	function := NewFunctionNow()
	stack := []interface{}{}

	assert.NoError(test, function.Invoke(&stack))
	assert.IsType(test, time.Time{}, stack[0])
}

func TestFunctionNow_Positive_ResultValue(test *testing.T) {
	function := NewFunctionNow()
	stack := []interface{}{}
	maxDifferenceInSeconds := 2.0

	assert.NoError(test, function.Invoke(&stack))

	now := time.Now().UTC()
	functionResult, _ := stack[0].(time.Time)

	assert.False(test, functionResult.After(now))
	assert.LessOrEqual(test, functionResult.Sub(now).Seconds(), maxDifferenceInSeconds)
}

func TestFunctionNow_Positive_MoreArguments_StackSize(test *testing.T) {
	function := NewFunctionNow()
	stack := []interface{}{false, "A", "B"}
	expectedStackSize := 4

	assert.NoError(test, function.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}
