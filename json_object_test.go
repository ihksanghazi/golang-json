package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct{
	Street string
	Country string
	PostalCode string
}
type Customer struct{
	Firstname string
	Middlename string
	Lastname string
	Age int
	Married bool
	Hobbies []string
	Addresses []Address
}

func TestJsonObject(t *testing.T) {
	customer:= Customer{
		Firstname: "Nur",
		Middlename: "Sandy",
		Lastname: "Ihksan",
		Age: 21,
		Married: false,
	}	

	byte,err:=json.Marshal(customer)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(byte))
}