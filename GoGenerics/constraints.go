package GoGenerics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// The any Constraint
func AnyConstraints[T any](value T) {
	fmt.Println(value)
}

// The comparable Constraint
func ComparableCon[T comparable](value T, typeValues []T) {
	for _, eh := range typeValues {
		if value == eh {
			return
		}
	}
}

// Built-in Interface Constraints
// A generic function that works with ordered types (integers and floats)
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Union of Specific Types (Type Sets)
func UnionConstraint[T int | float64](a, b T) T {
	return a + b
}

// Custom Interface Constraints
type Stringer interface {
	Pstrings() string
}

type person struct {
	Name string
}

type product struct {
	Name  string
	Pcode int
}

func (p person) Pstrings() string {
	return fmt.Sprintf("Pesron name : %s\n", p.Name)
}

func (p product) Pstrings() string {
	return fmt.Sprintf("Product name : %s and product code: %d \n", p.Name, p.Pcode)
}

func PrintString[T Stringer](s T) {
	fmt.Println(s.Pstrings())
}

func InitalizeReq() {
	per := person{Name: "Name"}
	pro := product{Name: "Name", Pcode: 1000}
	PrintString(per)
	PrintString(pro)
}
