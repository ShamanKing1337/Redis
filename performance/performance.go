package performance

import (
	"bytes"
	"net/http"
)

//func Test1(){
//	for i:=0;i <4000;i++{
//		http.Get("http://127.0.0.1:4000/KEYS")
//
//	}
//}

func ReqTest2() {
	client := &http.Client{}
	for i := 0; i < 2000; i++ {
		req, _ := http.NewRequest("GET", "http://127.0.0.1:4000/KEYS", nil)
		req.SetBasicAuth("admin", "admin")
		client.Do(req)
		var jsonStr = []byte(`{"key":"key6", "valueS" : "stringg", "valueD" : { "ddd" : "eee"}}`)
		req1, _ := http.NewRequest("GET", "http://127.0.0.1:4000/SET/{key3}", bytes.NewBuffer(jsonStr))
		req1.SetBasicAuth("admin", "admin")
		client.Do(req1)
	}
}

func main() {
	ReqTest2()
}
