package main

import (
	"GOGenericsAndReflection/GoGenerics"
	GorReflection "GOGenericsAndReflection/GoReflection"
	"fmt"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	Name    string
	id      int
	address string
	Salary  int
	country string
}

func main() {

	//---------------------------------Reflection----------------------------
	// name := GorReflection.Name("12345678")
	// var user = GorReflection.User{
	// 	Name: "name",
	// 	Age:  10,
	// }

	// GorReflection.Reflection(user, name)

	o := order{
		ordId:      456,
		customerId: 56,
	}
	GorReflection.CreateQuery(o)

	e := employee{
		Name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		Salary:  90000,
		country: "India",
	}
	GorReflection.CreateQuery(e)
	i := 90
	GorReflection.CreateQuery(i)

	//---------------------------------Generics----------------------------
	GoGenerics.AddFunction(10, 20)

	// s := GoGenerics.Stack[int]{Container: []int{1, 2, 3}}
	s := GoGenerics.Stack[int]{}
	s.AddElement(1)
	s.AddElement(2)
	s.PopElement()

	GoGenerics.InitStruct("emp", 20000)

	GoGenerics.PrintSlice([]int{1, 2, 3, 4, 5})
	GoGenerics.PrintSlice([]string{"first", "second"})
	fmt.Println(GoGenerics.UnionConstraint(10, 20.5))
	GoGenerics.InitalizeReq()

	//---------------------------------Generics Modification-----------------

	// GorReflection.ModfiyByReflection(e)

	//The original error occurred because .Elem() was being called on a non-pointer value.
	//Elem() is only valid on pointers, as it dereferences the pointer to access the actual value.
	// emp1 := employee{Name: "John", Salary: 100000}
	// GorReflection.ModifyByReflection(emp1)
	// fmt.Println("emp1 after modification:", emp1) // Output: emp1 after modification: {John 100000}

	emp2 := &employee{Name: "Doe", Salary: 150000}
	GorReflection.ModifyByReflection(emp2)
	fmt.Println("emp2 after modification:", emp2) // Output: emp2 after modification: &{Emp 200000}
}
