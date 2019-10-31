package functions

import (
	"regexp"
	"strings"
)

// FunctionReplace is REPLACE(<source>, <regex>, <replacement>)
type FunctionReplace struct {
	BaseFunction
}

// NewFunctionReplace creates an instance of FunctionReplace.
func NewFunctionReplace() *FunctionReplace {
	result := &FunctionReplace{MakeBaseFunction(nil, "REPLACE", 3)}
	result.derived = result
	return result
}

// CheckArguments checks if all arguments are correct.
func (f *FunctionReplace) CheckArguments(arguments []interface{}) bool {
	// check 1st argument
	if arguments[0] != nil {
		if _, ok := arguments[0].(string); !ok {
			return false
		}
	}

	// check 2nd argument
	if _, ok := arguments[1].(string); !ok {
		return false
	}
	if len(arguments[1].(string)) == 0 {
		return false
	}

	// check 3rd argument
	if _, ok := arguments[2].(string); !ok {
		return false
	}

	return true
}

// Result calculates result of the function for given arguments.
func (f *FunctionReplace) Result(arguments []interface{}) interface{} {
	if arguments[0] == nil {
		return nil
	}

	source := arguments[0].(string)
	regex := arguments[1].(string)
	replacement := arguments[2].(string)

	re, err := regexp.Compile(regex)
	if err != nil {
		return strings.ReplaceAll(source, regex, replacement)
	}
	return re.ReplaceAllString(source, replacement)
}
