package operators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorEqual_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorEqual()
	expectedName := "=="

	assert.Equal(test, expectedName, operator.Name())
}

func TestOperatorEqual_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorEqual()

	assert.Equal(test, operatorEqualPriority, operator.Priority())
}

func TestOperatorEqual_Positive_GoorArguments_StackSize(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", "String2"}
	expectedStackSize := 1

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorEqual_Positive_GoorArguments_ResultType(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", "String2"}

	assert.NoError(test, operator.Invoke(&stack))
	assert.IsType(test, true, stack[0])
}

func TestOperatorEqual_Positive_EqualStrings_ResultValue(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", "String1"}
	expectedValue := true

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorEqual_Positive_DifferentStrings_ResultValue(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", "String2"}
	expectedValue := false

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorEqual_Positive_FirstArgumentNull_ResultValue(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{nil, "String2"}
	expectedValue := false

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorEqual_Positive_SecondArgumentNull_ResultValue(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", nil}
	expectedValue := false

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorEqual_Positive_BothArgumentsNull_ResultValue(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{nil, nil}
	expectedValue := true

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorEqual_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1", "String2", "String3"}
	expectedStackSize := 2

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorEqual_Negative_FewerArguments(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"String1"}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: not enough arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorEqual_Negative_FirstArgumentWrongType(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{5, "5"}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorEqual_Negative_SecondArgumentWrongType(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{"1", 1}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorEqual_Negative_BothArgumentsWrongType(test *testing.T) {
	operator := NewOperatorEqual()
	stack := []interface{}{1, 1}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}
