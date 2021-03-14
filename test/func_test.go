package test

import (
	"Avito/service"
	"testing"
)



var allCommandsService = service.NewAllCommandsService()


func CheckList(a []string, b []string) bool{
	if len(a) == len(b){
		for i:=0; i < len(a); i++{
			if a[i] != b[i]{
				return false
			}
		}
	} else{
		return false
	}
	return true
}


func CheckDict(a map[string]string, b map[string]string) bool{
	if len(a) == len(b){
		for k := range b {
			if a[k] != b[k]{
				return false
			}
		}
	} else{
		return false
	}
	return true
}



func Test(t *testing.T) {

	req := service.RequestAll{Key: "key5", ValueS: "string", ValueL: []string{"dsad","ewe"}, ValueD: map[string]string{"key1" : "value1", "key2" : "value2"}}
	allCommandsService.Set(req)

	check := allCommandsService.Get(req.Key)
	if check.ValueS == req.ValueS && CheckDict(check.ValueD, req.ValueD) && CheckList(check.ValueL, req.ValueL)  {

	} else {
		t.Errorf("error on Get and Set")
	}

	keys := allCommandsService.Keys()

	if keys[0] == "key5"{

	} else {
		t.Errorf("error on Keys")
	}


	allCommandsService.Append(service.RequestAll{Key: "key5", ValueL: []string{"222","11"}, ValueD: map[string]string{"key3" : "value3"}})
	check = allCommandsService.Get(req.Key)

	if check.ValueS == "string" && ( check.ValueL[0] == "dsad" && check.ValueL[1] == "ewe" && check.ValueL[2] == "222" && check.ValueL[3] == "11") && ( check.ValueD["key1"] == "value1" && check.ValueD["key2"] == "value2" && check.ValueD["key3"] == "value3") {
	} else {
		t.Errorf("error on Append")
	}


	allCommandsService.Delete("key5")

	check1 := allCommandsService.Get("key5")

	if check1.ValueS == "nil"{
	} else {
		t.Errorf("error on Delete")
	}

}

