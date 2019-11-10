package operators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorLessOrEqual_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	expectedName := "<="

	assert.Equal(test, expectedName, operator.Name())
}

func TestOperatorLessOrEqual_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorLessOrEqual()

	assert.Equal(test, operatorLessOrEqualPriority, operator.Priority())
}

func TestOperatorLessOrEqual_Positive_GoorArguments_StackSize(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", "String2"}
	expectedStackSize := 1

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorLessOrEqual_Positive_GoorArguments_ResultType(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", "String2"}

	assert.NoError(test, operator.Invoke(&stack))
	assert.IsType(test, true, stack[0])
}

func TestOperatorLessOrEqual_Positive_EqualStrings_ResultValue(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", "String1"}
	expectedValue := true

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorLessOrEqual_Positive_FirstArgumentLess_ResultValue(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", "String2"}
	expectedValue := true

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorLessOrEqual_Positive_SecondArgumentLess_ResultValue(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String2", "String1"}
	expectedValue := false

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorLessOrEqual_Positive_BothArgumentsNull_ResultValue(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{nil, nil}
	expectedValue := true

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorLessOrEqual_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", "String2", "String3"}
	expectedStackSize := 2

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorLessOrEqual_Negative_FewerArguments(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1"}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: not enough arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorLessOrEqual_Negative_FirstArgumentWrongType(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{5, "5"}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorLessOrEqual_Negative_SecondArgumentWrongType(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"1", 2}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorLessOrEqual_Negative_BothArgumentsWrongType(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{1, 2}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorLessOrEqual_Negative_FirstArgumentNull(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{nil, "String2"}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorLessOrEqual_Negative_SecondArgumentNull(test *testing.T) {
	operator := NewOperatorLessOrEqual()
	stack := []interface{}{"String1", nil}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}
