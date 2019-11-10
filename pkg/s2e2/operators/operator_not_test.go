package operators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorNot_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorNot()
	expectedName := "!"

	assert.Equal(test, expectedName, operator.Name())
}

func TestOperatorNot_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorNot()

	assert.Equal(test, operatorNotPriority, operator.Priority())
}

func TestOperatorNot_Positive_GoorArguments_StackSize(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{true}
	expectedStackSize := 1

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorNot_Positive_GoorArguments_ResultType(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{true}

	assert.NoError(test, operator.Invoke(&stack))
	assert.IsType(test, true, stack[0])
}

func TestOperatorNot_Positive_ArgumentTrue_ResultValue(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{true}
	expectedValue := false

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorNot_Positive_ArgumentFalse_ResultValue(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{false}
	expectedValue := true

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorNot_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{false, true}
	expectedStackSize := 2

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorNot_Negative_FewerArguments(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: not enough arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorNot_Negative_ArgumentWrongType(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{"true"}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorNot_Negative_ArgumentNull(test *testing.T) {
	operator := NewOperatorNot()
	stack := []interface{}{nil}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}
