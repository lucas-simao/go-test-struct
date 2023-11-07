package tests

import (
	"testing"
	"time"
)

type Person struct {
	Name       string
	Age        int
	Salary     float64
	Active     bool
	Properties []string
	CreatedAt  time.Time
}

func TestValidateStruct(t *testing.T) {
	var p = Person{
		Name:       "Lucas Simão",
		Age:        31,
		Properties: []string{"casa", "carro"},
	}

	t.Run("Should return correct struct", func(t *testing.T) {
		ValidateStruct[Person](t, &p, &Person{
			Name:       "Lucas Simão",
			Age:        31,
			Properties: []string{"casa", "carro", "moto"},
		})
	})
}
