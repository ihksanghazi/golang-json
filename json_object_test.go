package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct{
	Firstname string
	Middlename string
	Lastname string
	Age int
	Married bool
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