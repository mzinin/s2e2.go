package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionReplace_Positive_CreateFunction_Name(test *testing.T) {
	function := NewFunctionReplace()
	expectedName := "REPLACE"

	assert.Equal(test, expectedName, function.Name())
}

func TestFunctionReplace_Positive_StringReplace_StackSize(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "A", "B"}
	expectedStackSize := 1

	assert.NoError(test, function.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestFunctionReplace_Positive_StringReplace_ResultType(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "A", "B"}

	assert.NoError(test, function.Invoke(&stack))
	assert.IsType(test, "string", stack[0])
}

func TestFunctionReplace_Positive_StringReplace_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "A", "B"}
	expectedValue := "BBB"

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestFunctionReplace_Positive_RegexReplace_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABCABA", "A.*?C", "D"}
	expectedValue := "DABA"

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestFunctionReplace_Positive_SpecialSymbolReplace_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"A * B == C", "\\*", "+"}
	expectedValue := "A + B == C"

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestFunctionReplace_Positive_FirstArgumentNull_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{nil, "A", "B"}

	assert.NoError(test, function.Invoke(&stack))
	assert.Nil(test, stack[0])
}

func TestFunctionReplace_Positive_FirstArgumentEmptyString_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"", "A", "B"}
	expectedValue := ""

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestFunctionReplace_Negative_FirstArgumentWrongType(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{5, "A", "B"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionReplace_Negative_SecondArgumentEmptyString(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "", "B"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionReplace_Negative_SecondArgumentNull(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", nil, "B"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionReplace_Negative_SecondArgumentWrongType(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"AB5", 5, "B"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionReplace_Positive_ThirdArgumentEmptyString_ResultValue(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "B", ""}
	expectedValue := "AA"

	assert.NoError(test, function.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestFunctionReplace_Negative_ThirdArgumentNull(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "B", nil}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionReplace_Negative_ThirdArgumentWrongType(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "B", 5}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: invalid arguments for function "+function.Name(), err.Error())
}

func TestFunctionReplace_Positive_MoreArguments_StackSize(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{false, "ABA", "A", "B"}
	expectedStackSize := 2

	assert.NoError(test, function.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestFunctionReplace_Negative_FewerArguments(test *testing.T) {
	function := NewFunctionReplace()
	stack := []interface{}{"ABA", "A"}

	err := function.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseFunction: not enough arguments for function "+function.Name(), err.Error())
}
