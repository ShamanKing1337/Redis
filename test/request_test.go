package test

import (
	"bytes"
	"net/http"
	"testing"
)


//func Test1(t *testing.T){
//	for i:=0;i <4000;i++{
//		http.Get("http://127.0.0.1:4000/KEYS")
//
//	}
//}

func Test1(t *testing.T){
	client := &http.Client{}
	for i:=0;i <2000;i++{
		req, _ := http.NewRequest("GET", "http://127.0.0.1:4000/KEYS", nil)
		req.SetBasicAuth("admin", "admin")
		client.Do(req)
		var jsonStr = []byte(`{"key":"key6", "valueS" : "stringg", "valueD" : { "ddd" : "eee"}}`)
		req1, _ := http.NewRequest("GET", "http://127.0.0.1:4000/GET/{key3}", bytes.NewBuffer(jsonStr))
		req1.SetBasicAuth("admin", "admin")
		client.Do(req1)


	}
}
