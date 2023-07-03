package golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T){
	writer,_:=os.Create("CustomerOut.json")
	encoder:=json.NewEncoder(writer)

	customer:=Customer{
		Firstname: "Nur",
		Middlename: "Sandy",
		Lastname: "Ihksan",
	}

	encoder.Encode(customer)
}