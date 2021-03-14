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



type StringCommandsImplementation interface {
	GetString(w http.ResponseWriter, r *http.Request)
	SetString(w http.ResponseWriter, r *http.Request)
	GetKeys(w http.ResponseWriter, r *http.Request)
	DeleteKey(w http.ResponseWriter, r *http.Request)
	SaveString(w http.ResponseWriter, r *http.Request)
}

type RequestString struct{
	Key string
	Value string
	Ttl int64
}

type stringCommandsImplementation struct {
	service service.StringCommandsService
}

func NewStringCommandsImplementation(service service.StringCommandsService) StringCommandsImplementation{
	return &stringCommandsImplementation{service: service}
}





func(s *stringCommandsImplementation) SetString(w http.ResponseWriter, r *http.Request){


	if remain.CheckMethod("PUT",r){
		var req RequestString
		err1 := json.NewDecoder(r.Body).Decode(&req)
		if err1 != nil {
			panic(err1)
		}

		s.service.Set(req.Key, req.Value, req.Ttl)
	} else{
		w.WriteHeader(405)
	}




}

func(s *stringCommandsImplementation) SaveString(w http.ResponseWriter, r *http.Request){

	s.service.Save()

}







func(s *stringCommandsImplementation) GetString(w http.ResponseWriter, r *http.Request){
	if remain.CheckMethod("GET",r){
		fmt.Println(url.Parse(fmt.Sprint(r.URL)))
		str := strings.Split(fmt.Sprint(r.URL), "B")
		str1 := strings.Split(str[1], "%")

		res := s.service.Get(str1[0])
		if (len(res)!= 0){
			fmt.Fprintf(w, "Resp: %+v", res)
		} else {
			fmt.Fprintf(w, "Not Exist")
		}
	} else{
		w.WriteHeader(405)
	}


}


func(s *stringCommandsImplementation) DeleteKey(w http.ResponseWriter, r *http.Request){
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


func(s *stringCommandsImplementation) GetKeys(w http.ResponseWriter, r *http.Request){
	if remain.CheckMethod("GET",r){
		tmp := s.service.Keys()

		fmt.Fprintf(w, "Resp: %+v", tmp)
	} else{
		w.WriteHeader(405)
	}
}
