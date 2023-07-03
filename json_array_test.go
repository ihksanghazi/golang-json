package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T){
	customer:=Customer{
		Firstname: "Nur",
		Middlename: "Sandy",
		Lastname: "Ihksan",
		Hobbies: []string{"Game","Reading","Coding"},
	}

	bytes,_:=json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T){
	jsonString := `{"Firstname":"Nur","Middlename":"Sandy","Lastname":"Ihksan","Age":0,"Married":false,"Hobbies":["Game","Reading","Coding"]}`
	jsonBytes:= []byte(jsonString)

	customer:= &Customer{}
	err:=json.Unmarshal(jsonBytes,customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer:=Customer{
		Firstname: "sandy",
		Addresses: []Address{
			{
				Street: "Jl. Panjang Banget",
				Country: "Indonesia",
				PostalCode: "12345",
			},
			{
				Street: "Jl. Pendek Banget",
				Country: "Zimbabwe",
				PostalCode: "54321",
			},
		},
	}
	bytes,_:=json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"Firstname":"sandy","Middlename":"","Lastname":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jl. Panjang Banget","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jl. Pendek Banget","Country":"Zimbabwe","PostalCode":"54321"}]}`
	jsonBytes:= []byte(jsonString)

	customer:= &Customer{}
	err:=json.Unmarshal(jsonBytes,customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Addresses)
}