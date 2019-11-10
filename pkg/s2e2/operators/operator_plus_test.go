package operators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorPlus_Positive_CreateOperator_Name(test *testing.T) {
	operator := NewOperatorPlus()
	expectedName := "+"

	assert.Equal(test, expectedName, operator.Name())
}

func TestOperatorPlus_Positive_CreateOperator_Priority(test *testing.T) {
	operator := NewOperatorPlus()

	assert.Equal(test, operatorPlusPriority, operator.Priority())
}

func TestOperatorPlus_Positive_StringString_StackSize(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", "B"}
	expectedStackSize := 1

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorPlus_Positive_StringString_ResultType(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", "B"}

	assert.NoError(test, operator.Invoke(&stack))
	assert.IsType(test, "string", stack[0])
}

func TestOperatorPlus_Positive_StringString_ResultValue(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", "B"}
	expectedValue := "AB"

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorPlus_Positive_StringNull_ResultValue(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", nil}
	expectedValue := "A"

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorPlus_Positive_NullString_ResultValue(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{nil, "B"}
	expectedValue := "B"

	assert.NoError(test, operator.Invoke(&stack))

	value, _ := stack[0].(string)
	assert.Equal(test, expectedValue, value)
}

func TestOperatorPlus_Positive_NullNull_ResultValue(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{nil, nil}

	assert.NoError(test, operator.Invoke(&stack))
	assert.Nil(test, stack[0])
}

func TestOperatorPlus_Positive_MoreArguments_StackSize(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", "B", "C"}
	expectedStackSize := 2

	assert.NoError(test, operator.Invoke(&stack))
	assert.Equal(test, expectedStackSize, len(stack))
}

func TestOperatorPlus_Negative_FewerArguments(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A"}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: not enough arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorPlus_Negative_FirstArgumentWrongType(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{5, "B"}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}

func TestOperatorPlus_Negative_SecondArgumentWrongType(test *testing.T) {
	operator := NewOperatorPlus()
	stack := []interface{}{"A", 1}

	err := operator.Invoke(&stack)
	assert.Error(test, err)
	assert.Equal(test, "BaseOperator: invalid arguments for operator "+operator.Name(), err.Error())
}
