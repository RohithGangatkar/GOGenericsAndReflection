package GoGenerics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func AddFunction[T constraints.Ordered](a, b T) {

	value := a + b
	fmt.Printf("sum of values:%v \n", value)

}

type Stack[T int | float32] struct {
	Container []T
}

func (s *Stack[T]) AddElement(ele T) {
	s.Container = append(s.Container, ele)
	fmt.Println(s.Container)
}

func (s *Stack[T]) PopElement() {

	s.Container = s.Container[:len(s.Container)-1]
	fmt.Println(s.Container)
}

type UserStru[T any, U any] struct {
	name   T
	salary U
}

func InitStruct(a string, b int) {

	user := UserStru[string, int]{name: a, salary: b}
	fmt.Println(user)
}

func PrintSlice[T any](values []T) {

	for _, each := range values {
		fmt.Println(each)
	}
}
