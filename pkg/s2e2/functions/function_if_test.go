package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionIf_Positive_CreateFunction_Name(test *testing.T) {
	function := NewFunctionIf()
	expectedName := "IF"

	assert.Equal(test, expectedName, function.Name())
}

func TestFunctionIf_Positive_FirstArgumentTrue_StackSize(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{true, "A", "B"}
	expectedStackSize := 1

	assert.NoError(test, function.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestFunctionIf_Positive_FirstArgumentTrue_ResultType(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{true, "A", "B"}

	assert.NoError(test, function.Invoke(&stack))
	assert.IsType(test, "string", stack[0])
}

func TestFunctionIf_Positive_FirstArgumentTrue_ResultValue(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{true, "A", "B"}
	expectedValue := "A"

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestFunctionIf_Positive_FirstArgumentFalse_ResultValue(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{false, "A", "B"}
	expectedValue := "B"

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestFunctionIf_Positive_MoreArguments_StackSize(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{"ARG", false, "A", "B"}
	expectedStackSize := 2

	assert.NoError(test, function.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestFunctionIf_Negative_FewerArguments(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{false}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: not enough arguments for function "+function.Name(), err.Error())
}

func TestFunctionIf_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{"false", "A", "B"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionIf_Negative_FirstArgumentNull(test *testing.T) {
	function := NewFunctionIf()
	stack := []interface{}{nil, "A", "B"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}
