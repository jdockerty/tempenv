package tempenv

import (
	"fmt"
	"os"
	"reflect"
)

// Set is used to perform a function call with the setting and unsetting of the target environment variable
// handled inbetween. This returns a slice of reflect values, which are the outputs, if any, of the function
// that you passed. An error is returned if the action you pass is not a function, this is done via the
// type of reflect.Func from the reflect package.
func Set(actionFunc interface{}, targetVar, tempVar, resetVar string, actionFuncArgs ...interface{}) (*[]reflect.Value, error) {

	switch t := reflect.TypeOf(actionFunc).Kind(); t {
	case reflect.Func:
		os.Setenv(targetVar, tempVar)
		defer os.Setenv(targetVar, resetVar)

		// Collect the passed arguments into a desirable slice that we can use.
		fnArguments := make([]reflect.Value, len(actionFuncArgs))
		for i, _ := range actionFuncArgs {
			fnArguments[i] = reflect.ValueOf(actionFuncArgs[i])
		}

		funcResult := reflect.ValueOf(actionFunc).Call(fnArguments)

		return &funcResult, nil

	default:
		return nil, fmt.Errorf("the passed action was not of type func, it is not callable")
	}

}
