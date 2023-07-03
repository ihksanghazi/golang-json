package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T){
	jsonString:= `{"id":"P0001","name":"Apple Mac Book Pro","image_url":"http://example.com/image.png"}`
	jsonBytes:= []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {
	product:=map[string]interface{}{
		"id":"P0001",
		"name":"Kopi Kapal Api",
		"price":"Rp 1500",
	}
	bytes,_:=json.Marshal(product)
	fmt.Println(string(bytes))
}