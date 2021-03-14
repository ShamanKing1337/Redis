package controller

import (
	"Avito/remain"
	"Avito/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ListCommandsImplementation interface {
	GetList(w http.ResponseWriter, r *http.Request)
	SetList(w http.ResponseWriter, r *http.Request)
	GetKeys(w http.ResponseWriter, r *http.Request)
	DeleteKey(w http.ResponseWriter, r *http.Request)
	SaveList(w http.ResponseWriter, r *http.Request)
}

type RequestList struct {
	Key   string
	Value string
	Ttl   int64
}

type listCommandsImplementation struct {
	service service.ListCommandsService
}

func NewListCommandsImplementation(service service.ListCommandsService) ListCommandsImplementation {
	return &listCommandsImplementation{service: service}
}

func (s listCommandsImplementation) SetList(w http.ResponseWriter, r *http.Request) {

	if remain.CheckMethod("POST", r) {
		var req RequestList
		err1 := json.NewDecoder(r.Body).Decode(&req)
		if err1 != nil {
			w.WriteHeader(400)
		}

		resp := s.service.Set(req.Key, req.Value, req.Ttl)
		fmt.Fprintf(w, "Resp: %+v", resp)
	} else {
		w.WriteHeader(405)
	}
}

func (s *listCommandsImplementation) SaveList(w http.ResponseWriter, r *http.Request) {

	s.service.Save()

}

func (s *listCommandsImplementation) GetList(w http.ResponseWriter, r *http.Request) {
	if remain.CheckMethod("GET", r) {
		str := strings.Split(fmt.Sprint(r.URL), "B")
		if len(str) != 2 {
			w.WriteHeader(400)
			return
		}
		str1 := strings.Split(str[1], "%")
		if len(str1) != 2 {
			w.WriteHeader(400)
			return
		}

		res := s.service.Get(str1[0])
		if len(res) != 0 {
			fmt.Fprintf(w, "Resp: %+v", res)
		} else {
			fmt.Fprintf(w, "Not Exist")
		}
	} else {
		w.WriteHeader(405)
	}
}

func (s *listCommandsImplementation) DeleteKey(w http.ResponseWriter, r *http.Request) {
	if remain.CheckMethod("DELETE", r) {
		str := strings.Split(fmt.Sprint(r.URL), "B")
		if len(str) != 2 {
			w.WriteHeader(400)
			return
		}
		str1 := strings.Split(str[1], "%")
		if len(str1) != 2 {
			w.WriteHeader(400)
			return
		}
		resp := s.service.Delete(str1[0])

		fmt.Fprintf(w, "Resp: %+v", resp)
	} else {
		w.WriteHeader(405)
	}
}

func (s *listCommandsImplementation) GetKeys(w http.ResponseWriter, r *http.Request) {
	if remain.CheckMethod("GET", r) {
		tmp := s.service.Keys()

		fmt.Fprintf(w, "Resp: %+v", tmp)

	} else {
		w.WriteHeader(405)
	}
}
