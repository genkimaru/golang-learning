package tools

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"text/template"
)

func Eval(s string, funcs template.FuncMap) (string, error) {
	args, name := parseArgs(s)
	fn, ok := funcs[name]
	if !ok {
		return "", fmt.Errorf("function %s is not defined", name)
	}
	v := reflect.ValueOf(fn)
	t := v.Type()
	if len(args) != t.NumIn() {
		return "", fmt.Errorf("invalid number of arguments. got: %v, want: %v", len(args), t.NumIn())
	}
	argv := make([]reflect.Value, len(args))
	// go over the arguments, validate and build them.
	// note that we support only int, and string in this simple example.
	for i := range argv {
		var argType reflect.Kind
		// if the argument "i" is string.
		if strings.HasPrefix(args[i], "\"") {
			argType = reflect.String
			argv[i] = reflect.ValueOf(strings.Trim(args[i], "\""))
		} else {
			argType = reflect.Int
			// assume that the input is valid.
			num, _ := strconv.Atoi(args[i])
			argv[i] = reflect.ValueOf(num)
		}
		if t.In(i).Kind() != argType {
			return "", fmt.Errorf("Invalid argument. got: %v, want: %v", argType, t.In(i).Kind())
		}
	}
	result := v.Call(argv)
	// in real-world code, we validate it before executing the function,
	// using the v.NumOut() method. similiar to the text/template package.
	if len(result) != 1 || result[0].Kind() != reflect.String {
		return "", fmt.Errorf("function %s must return a one string value", name)
	}
	return result[0].String(), nil
}

// parseArgs is an auxiliary function, that extract the function and its
// parameter from the given expression.
func parseArgs(s string) ([]string, string) {
	args := strings.Split(strings.Trim(s, "{ }"), "|")
	return strings.Split(strings.Trim(args[0], " "), " "), strings.Trim(args[len(args)-1], " ")
}
