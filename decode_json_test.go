package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T){
	jsonString := `{"Firstname":"Nur","Middlename":"Sandy","Lastname":"Ihksan","Age":21,"Married":false}`
	jsonBytes:= []byte(jsonString)

	customer:= &Customer{}
	err:=json.Unmarshal(jsonBytes,customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)

}
