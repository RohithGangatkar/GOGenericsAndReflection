package GorReflection

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

type Name string

func Reflection(i, j interface{}) {

	itype := reflect.TypeOf(i)
	ivalue := reflect.ValueOf(i)
	iKindvalue := reflect.TypeOf(i).Kind()

	// jtype := reflect.TypeOf(j)
	// jvalue := reflect.ValueOf(j)
	// jKindvalue := reflect.TypeOf(j).Kind()

	fmt.Println("Type of first interface:", itype)
	fmt.Println("Kind of first interface:", iKindvalue)
	fmt.Println("value of first interface:", ivalue)

	//NOTE:[
	//if we add %#v - GorReflection.User{Name:"name", Age:10}
	// if we add %v - {name 10}
	//]
	// fmt.Sprintf("i type of : %T and  value of : %v \n", itype, ivalue)
	// fmt.Println(iKindvalue)

	// fmt.Sprintf("j type of : %T and  value of : %v \n", jtype, jvalue)
	// fmt.Println(jKindvalue)
	numFiled := ivalue.NumField()
	if ivalue.Kind() == reflect.Struct {
		for k := 0; k < numFiled; k++ {
			fmt.Sprintf(" FiledNo : %d ,  type of : %T and  value of : %#v \n", (k + 1), ivalue.Field(k), ivalue.Field(k))
		}
	}

}

func CreateQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		qVlaue := reflect.ValueOf(q)
		numField := qVlaue.NumField()
		tname := reflect.TypeOf(q).Name()

		query := fmt.Sprintf(" insert into %s values (", tname)

		for i := 0; i < numField; i++ {

			if i == 0 {
				if qVlaue.Field(i).Kind() == reflect.Int {
					query = fmt.Sprintf("%s %d", query, qVlaue.Field(i).Int())
				} else if qVlaue.Field(i).Kind() == reflect.String {
					query = fmt.Sprintf("%s \"%s\"", query, qVlaue.Field(i))
				} else {
					fmt.Println("Unsupported type")
					return
				}
			} else {
				if qVlaue.Field(i).Kind() == reflect.Int {
					query = fmt.Sprintf("%s , %d", query, qVlaue.Field(i).Int())
				} else if qVlaue.Field(i).Kind() == reflect.String {
					query = fmt.Sprintf("%s , \"%s\"", query, qVlaue.Field(i))
				} else {
					fmt.Println("Unsupported type")
					return
				}
			}
		}
		query = fmt.Sprintf("%s )", query)

		fmt.Println(query)
	} else {
		fmt.Println("Unsupported type")
		return
	}

}

func ModfiyByReflection(e interface{}) {

	element := reflect.ValueOf(e).Elem()

	nameElement := element.FieldByName("Name")
	if nameElement.CanSet() {
		nameElement.SetString("Emp")
	}

	salaryElement := element.FieldByName("Salary")
	if salaryElement.CanSet() {
		salaryElement.SetInt(200000)
	}

	fmt.Println("After modifying struct with refelection")
	fmt.Println(e)

}

// Mo
func ModifyByReflection(e interface{}) {
	element := reflect.ValueOf(e).Elem()

	nameElement := element.FieldByName("Name")
	if nameElement.CanSet() {
		nameElement.SetString("Emp")
	}

	salaryElement := element.FieldByName("Salary")
	if salaryElement.CanSet() {
		salaryElement.SetInt(200000)
	}
}
