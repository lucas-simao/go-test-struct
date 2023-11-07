package tests

import (
	_ "embed"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"testing"
	"text/template"
)

//go:embed logTemplate.txt
var logTemplate string

/*
ValidateStruct is used to check value between two structs.
If the field validator contains a value different from zero this value will be checked.
Example of zero values: https://go.dev/tour/basics/12

	type Person struct {
		Name      string
		Age       int
	}

	var personToCheck = Person{
		Name:      "Lucas",
		Age:       30
	}

	var personWithValuesExpected = Person{
		Name:      "Lucas Simão",
		Age:       31
	}

	ValidateStruct[Person](t, &personToCheck, &personWithValuesExpected)
*/
func ValidateStruct[T any](t *testing.T, obj *T, validation *T) {
	_, file, line, _ := runtime.Caller(1)

	principal := reflect.ValueOf(obj).Elem()
	toCompare := reflect.ValueOf(validation).Elem()

	typeOf := reflect.TypeOf(validation).Elem()

	for i := 0; i < principal.NumField(); i++ {
		if !toCompare.Field(i).IsZero() {
			if !toCompare.Field(i).Equal(principal.Field(i)) {
				output := map[string]interface{}{
					"field":    typeOf.Field(i).Name,
					"type":     typeOf.Field(i).Type,
					"expected": fmt.Sprintf("%v", toCompare.Field(i)),
					"got":      fmt.Sprintf("%v", principal.Field(i)),
					"file":     fmt.Sprintf("%s:%d", file, line),
				}

				tmp, _ := template.New("log").Parse(logTemplate)
				tmp.Execute(os.Stdout, output)
				t.Fail()
				fmt.Print("\n\n")
			}
		}
	}
}
