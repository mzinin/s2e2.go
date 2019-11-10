package operators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorAnd_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorAnd()
	expectedName := "&&"

	assert.Equal(test, expectedName, operator.Name())
}

func TestOperatorAnd_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorAnd()

	assert.Equal(test, operatorAndPriority, operator.Priority())
}

func TestOperatorAnd_Positive_TrueTrue_StackSize(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{true, true}
	expectedStackSize := 1

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorAnd_Positive_TrueTrue_ResultType(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{true, true}

	assert.NoError(test, operator.Invoke(&stack))
	assert.IsType(test, true, stack[0])
}

func TestOperatorAnd_Positive_TrueTrue_ResultValue(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{true, true}
	expectedValue := true

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorAnd_Positive_TrueFalse_ResultValue(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{true, false}
	expectedValue := false

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorAnd_Positive_FalseTrue_ResultValue(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{false, true}
	expectedValue := false

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorAnd_Positive_FalseFalse_ResultValue(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{false, false}
	expectedValue := false

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(bool)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorAnd_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{"ARG", false, true}
	expectedStackSize := 2

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorAnd_Negative_FewerArguments(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{true}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: not enough arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorAnd_Negative_FirstArgumentWrongType(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{"true", true}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorAnd_Negative_FirstArgumentNull(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{nil, true}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorAnd_Negative_SecondArgumentWrongType(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{true, "true"}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorAnd_Negative_SecondArgumentNull(test *testing.T) {
	operator := NewOperatorAnd()
	stack := []interface{}{true, nil}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}
