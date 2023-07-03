package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}){
	byte,err:=json.Marshal(data)
	if err != nil {
		panic(err)
	}	
	fmt.Println(string(byte))
}

func TestEncode(t *testing.T){
	logJson("Sandy")
	logJson(1)
	logJson(true)
	logJson([]string{"Nur","Sandy","Ihksan"})
}
