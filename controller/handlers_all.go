
package controller

import (
	"Avito/remain"
	"Avito/service"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)



type AllCommandsImplementation interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	SetAll(w http.ResponseWriter, r *http.Request)
	GetKeys(w http.ResponseWriter, r *http.Request)
	DeleteKey(w http.ResponseWriter, r *http.Request)
	SaveAll(w http.ResponseWriter, r *http.Request)
	AppendAll(w http.ResponseWriter, r *http.Request)
}



type allCommandsImplementation struct {
	service service.AllCommandsService
}

func NewAllCommandsImplementation(service service.AllCommandsService) AllCommandsImplementation{
	return &allCommandsImplementation{service: service}
}





func(s allCommandsImplementation) SetAll(w http.ResponseWriter, r *http.Request){

	if remain.CheckMethod("PUT",r){


		var req service.RequestAll
		err1 := json.NewDecoder(r.Body).Decode(&req)

		if err1 != nil {
			panic(err1)
		}


		s.service.Set(req)

	} else{
		w.WriteHeader(405)
	}

}

func(s *allCommandsImplementation) SaveAll(w http.ResponseWriter, r *http.Request){

	s.service.Save()

}



func(s *allCommandsImplementation) AppendAll(w http.ResponseWriter, r *http.Request){
	if remain.CheckMethod("POST",r){
		var req service.RequestAll
		err1 := json.NewDecoder(r.Body).Decode(&req)

		if err1 != nil {
			panic(err1)
		}
		s.service.Append(req)
	} else{
		w.WriteHeader(405)
	}
}



func(s *allCommandsImplementation) GetAll(w http.ResponseWriter, r *http.Request){
	if remain.CheckMethod("GET",r){
		fmt.Println(url.Parse(fmt.Sprint(r.URL)))
		str := strings.Split(fmt.Sprint(r.URL), "B")
		str1 := strings.Split(str[1], "%")


		var res *service.Ttl = s.service.Get(str1[0])
		if res.ValueS != "nil"{
			fmt.Fprintf(w, "Resp: %+v", res)
		} else {
			fmt.Fprintf(w, "Not Exist")
		}
	} else{
		w.WriteHeader(405)
	}

}


func(s *allCommandsImplementation) DeleteKey(w http.ResponseWriter, r *http.Request){
	if remain.CheckMethod("DELETE",r){
		fmt.Println(url.Parse(fmt.Sprint(r.URL)))
		str := strings.Split(fmt.Sprint(r.URL), "B")
		str1 := strings.Split(str[1], "%")
		resp := s.service.Delete(str1[0])

		fmt.Fprintf(w, "Resp: %+v", resp)
	} else{
		w.WriteHeader(405)
	}
}


func(s *allCommandsImplementation) GetKeys(w http.ResponseWriter, r *http.Request){
	if remain.CheckMethod("GET",r){

		tmp := s.service.Keys()

		fmt.Fprintf(w, "Resp: %+v", tmp)
	} else{
		w.WriteHeader(405)
	}
}

