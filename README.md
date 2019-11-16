# s2e2.go

This library provides Go implementation of Simple String Expression Evaluator a.k.a. `s2e2`. The Evaluator returns value of an input expression. Unlike commonly known mathematical expression evaluators this one treats all parts of the expression as a strings and its output value is also a string.

For example:
* the value of the expression `A + B` is `AB`
* the value of `REPLACE("The cat is black", cat, dog)` is `The dog is black`

This is how one can use Evaluator to get value of some expression:
```go
import "github.com/mzinin/s2e2.go/pkg/s2e2"

evaluator := s2e2.NewEvaluator()

evaluator.AddStandardFunctions()
evaluator.AddStandardOperators()

expression := "A + B"
result, err := evaluator.Evaluate(expression)
```

## Supported expressions

Supported expressions consist of the following tokens: string literals, operators (unary and binary), functions, predefined constants, round brackets for function's arguments denoting, commas for function's arguments separation and double quotes for characters escaping. 

The difference between a function and an operator is that a function is always followed by a pair of round brackets with a list of function's arguments (probably empty) in between, while an operator does not use brackets and, if it is a binary operator, sticks between its operands. Also operators can have different priorities a.k.a. precedence.

For example:
* this is a function of 2 arguments: `FUNC(Arg1, Arg2)`
* and this is a binary operator: `Arg1 OP Arg2`


## Constants

There is only one predefined constant - `NULL` - which corresponds to an `nil` value in Go. It can be used to check if some sub-expression is evaluated into some result: `IF(SUBEXPR(Arg1, Arg2) == NULL, NULL, Value)`


## Functions

`s2e2` provides a small set of predefined functions. They are:

* Function `IF(Condition, Value1, Value2)`
  
  Returns `Value1` if `Condition` is true, and `Value2` otherwise. `Condition` must be a boolean value.

* Function `REPLACE(Source, Regex, Replacement)`

  Returns copy of `Source` with all matches of `Regex` replaced by `Replacement`. All three arguments are strings, `Regex` cannot be `NULL` or an empty string, `Replacement` cannot be `NULL`.

* Function `NOW()`

  Returns current UTC datetime. The result is of `time.Time` type.

* Function `ADD_DAYS(Datetime, NumberOfDays)`

  Adds days to the provided datetime. `Datetime` must be of `time.Time` type and not `NULL`. `NumberOfDays` is a not `NULL` string parsable into an any integer. The result is of `time.Time` type.

* Function `FORMAT_DATE(Datetime, Format)`

  Converts `Datetime` into a string according to `Format`. `Datetime` must be of `time.Time` type and not `NULL`. `Format` is a not `NULL` string.
  

### Custom functions

It is possible to create and use any custom function. Here is a simple example:
```go
import (
    "fmt"
    "github.com/mzinin/s2e2.go/pkg/s2e2"
    "github.com/mzinin/s2e2.go/pkg/s2e2/functions"
)

type CustomFunction struct {
    functions.BaseFunction
    set map[string]bool
}

func NewCustomFunction(set map[string]bool) *CustomFunction {
	result := &CustomFunction{functions.MakeBaseFunction(nil, "CONTAINS", 1), set}
    result.SetDerived(result)
	return result
}

func (f *CustomFunction) CheckArguments(arguments []interface{}) bool {
	_, ok := arguments[0].(string)
	return ok
}

func (f *CustomFunction) Result(arguments []interface{}) interface{} {
	if f.set == nil {
		return false
    }

    key, _ := arguments[0].(string)
    _, ok := f.set[key]
	return ok
}

func customFunctionExample() {
    evaluator := s2e2.NewEvaluator()
    evaluator.AddStandardFunctions()
    evaluator.AddStandardOperators()

    set := make(map[string]bool)
    set["key1"] = true
    set["key2"] = true

    customFunction := NewCustomFunction(set)
    evaluator.AddFunction(customFunction)

    expression := "IF(CONTAINS(key1), YES, NO)"
    result, err := evaluator.Evaluate(expression)
    if err != nil {
	    fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Value: %v\n", *result)
    }
}
```

## Operators

As it was mentioned before, every operator has a priority. Within `s2e2` the range of priorities is from 1 to 999. A set of predefined operators is provided. They are:

* Binary operator `+`, priority `500`

  Concatenates two strings. Every operand can be either a `NULL` or a string. The result is a string.

* Binary operator `==`, priority `300`

  Compares two strings, including `NULL`. If both operands are `NULL` the result is `true`. The type of the result is boolean.

* Binary operator `!=`, priority `300`

  The same as `==`, but checks objects for inequality. 

* Binary operator `>`, priority `400`

  Compares two strings lexicographically. None of the operands can be `NULL`. The result is a boolean.

* Binary operator `>=`, priority `400`

  Compares two string lexicographically as well. Both operands must be not `NULL` or both must be `NULL`. In the latter case the result is `true`.

* Binary operator `<`, priority `400`

  Same as `>`, but checks if first operand is less that the second one.

* Binary operator `<=`, priority `400`

  Same as `>=`, but checks if first operand is less or equal that the second one.

* Binary operator `&&`, priority `200`

  Computes logical conjunction of two boolean values. Both arguments are boolean, not `NULL` value. The result is a boolean.

* Binary operator `||`, priority `100`

  Computes logical disjunction of two boolean values. Both arguments are boolean, not `NULL` value. The result is a boolean.

* Unary operator `!`, priority `600`

  Negates boolean value. Operand cannot be `NULL`. The result is a boolean.


### Custom operators

It is possible to create and use any custom operator. Here is a simple example:
```go
import (
    "fmt"
    "github.com/mzinin/s2e2.go/pkg/s2e2"
    "github.com/mzinin/s2e2.go/pkg/s2e2/operators"
)

type CustomOperator struct {
    operators.BaseOperator
}

func NewCustomOperator() *CustomOperator {
    // ~ - symbols of the custom operator
    // 600 - priority
    // 1 - number of arguments
	result := &CustomOperator{operators.MakeBaseOperator(nil, "~", 600, 1)}
    result.SetDerived(result)
	return result
}

func (o *CustomOperator) CheckArguments(arguments []interface{}) bool {
	_, ok := arguments[0].(string)
	return ok
}

func (o *CustomOperator) Result(arguments []interface{}) interface{} {
    runes := []rune(arguments[0].(string))
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func customOperatorExample() {
    evaluator := s2e2.NewEvaluator()
    evaluator.AddStandardFunctions()
    evaluator.AddStandardOperators()

    customOperator := NewCustomOperator()
    evaluator.AddOperator(customOperator)

    expression := "~Foo"
    result, err := evaluator.Evaluate(expression)
    if err != nil {
	    fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Value: %v\n", *result)
    }
}
```

## Getting Started

### Prerequisites

To compile this project one would need:
* [Go](https://golang.org/) >= 1.6 (minimal tested version)

`testify` package is used for unit testing and will be downloaded by the corresponding script.


### Get library

If go modules are not available (i.e. Go <= 1.10) or not used:
```
go get github.com/mzinin/s2e2.go
```
With go modules:
```
go mod init <package name>
go build
```


### Run tests

On Linux:
```
./test.sh
```
`testify` package will be downloaded and used.
Since running this script requires a symlink, it is not possible to run tests on Windows.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
