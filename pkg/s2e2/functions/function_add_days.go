package functions

import (
	"strconv"
	"time"
)

const (
	hoursPerDay = 24
)

// FunctionAddDays is ADD_DAYS(<datetime>, <days>)
// Adds days number of days to datetime.
type FunctionAddDays struct {
	BaseFunction
}

// NewFunctionAddDays creates an instance of FunctionAddDays.
func NewFunctionAddDays() *FunctionAddDays {
	result := &FunctionAddDays{MakeBaseFunction(nil, "ADD_DAYS", 2)}
	result.derived = result
	return result
}

// CheckArguments checks if all arguments are correct.
func (f *FunctionAddDays) CheckArguments(arguments []interface{}) bool {
	// check 1st argument
	if _, ok := arguments[0].(time.Time); !ok {
		return false
	}

	// check 2nd argument
	if _, ok := arguments[1].(string); !ok {
		return false
	}
	if _, err := strconv.Atoi(arguments[1].(string)); err != nil {
		return false
	}

	return true
}

// Result calculates result of the function for given arguments.
func (f *FunctionAddDays) Result(arguments []interface{}) interface{} {
	datetime, _ := arguments[0].(time.Time)
	days, _ := strconv.Atoi(arguments[1].(string))
	return datetime.Add(time.Hour * time.Duration(hoursPerDay*days))
}
