# go-test-struct

This package is used to validate any struct, you should pass two structs to verify, and t like type testing.T

This func only will check values is different from zero. [Example of zero values](https://go.dev/tour/basics/12)

```go
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
```

```bash
field:     Name
type:      string
expected:  "Lucas Simão"
got:       "Lucas"
file:      /Users/lucas/workspace/lucas/go-test-struct/tests_test.go:23

field:     Age
type:      int
expected:  "31"
got:       "30"
file:      /Users/lucas/workspace/lucas/go-test-struct/tests_test.go:23

--- FAIL: TestValidateStruct (0.00s)
    --- FAIL: TestValidateStruct/Should_return_correct_struct (0.00s)
FAIL
FAIL    github.com/lucas-simao/go-test-struct   0.232s
FAIL
```
