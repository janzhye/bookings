package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	Street   string
	District string
	City     string
}

func (addr *Address) Display() {
	fmt.Println(addr.Street, addr.District, addr.City)
}

type Employee struct {
	Firstname string `json:"first_name" kbtg:"fname"`
	Lastname  string `json:"last_name" kbtg:"lname"`
	Address   `json:"address" kbtg:"addr"`
}

func main() {
	emp := Employee{
		Firstname: "Werasak",
		Lastname:  "Chongnguluam",
		Address: Address{
			Street:   "ออเงิน",
			District: "สายไหม",
			City:     "Bangkok",
		},
	}
	fmt.Printf("%+v\n", emp)

	fmt.Println(emp.Street)
	fmt.Println(emp.District)
	fmt.Println(emp.City)
	emp.Display()

	typ := reflect.TypeOf(emp)
	fmt.Println(typ.Name())
	fmt.Println(typ.Kind())
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		fmt.Println(fld.Name)
	}

	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	fmt.Println(reflect.DeepEqual(a, b))
	//typ := reflect.TypeOf(emp)
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		fmt.Println(fld.Name, fld.Tag.Get("kbtg"))
	}
}
