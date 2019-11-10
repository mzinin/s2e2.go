package operators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorOr_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorOr()
	expectedName := "||"

	assert.Equal(test, expectedName, operator.Name())
}

func TestOperatorOr_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorOr()

	assert.Equal(test, operatorOrPriority, operator.Priority())
}

func TestOperatorOr_Positive_GoorArguments_StackSize(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true, true}
	expectedStackSize := 1

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorOr_Positive_GoorArguments_ResultType(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true, true}

	assert.NoError(test, operator.Invoke(&stack))
	assert.IsType(test, true, stack[0])
}

func TestOperatorOr_Positive_TrueTrue_ResultValue(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true, true}
	expectedValue := true

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorOr_Positive_TrueFalse_ResultValue(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true, false}
	expectedValue := true

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorOr_Positive_FalseTrue_ResultValue(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{false, true}
	expectedValue := true

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorOr_Positive_FalseFalse_ResultValue(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{false, false}
	expectedValue := false

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorOr_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{"ARG", false, true}
	expectedStackSize := 2

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorOr_Negative_FewerArguments(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: not enough arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorOr_Negative_FirstArgumentWrongType(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{"true", true}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorOr_Negative_FirstArgumentNull(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{nil, true}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorOr_Negative_SecondArgumentWrongType(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true, "true"}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorOr_Negative_SecondArgumentNull(test *testing.T) {
	operator := NewOperatorOr()
	stack := []interface{}{true, nil}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}
