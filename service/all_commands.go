package service

import (
	"fmt"
	"os"
	"strconv"

	//"os"
	//"strconv"
	"sync"
	"time"
)

type AllCommandsService interface {
	Get(key string) *Ttl
	Set(req RequestAll) string
	Delete(key string) string
	Keys() []string
	Check()
	Save()
	Append(req RequestAll) string
}

type RequestAll struct {
	Key    string
	ValueS string
	ValueL []string
	ValueD map[string]string

	Ttl int64
}

type Ttl struct {
	ValueD map[string]string
	ValueL []string
	ValueS string
	ttl    int64
}

type allCommandsService struct {
	data  map[string]*Ttl
	mutex sync.Mutex
}

func (s *allCommandsService) Get(key string) *Ttl {
	s.mutex.Lock()
	tmp, ok := s.data[key]
	s.mutex.Unlock()
	if ok {
		return tmp
	} else {

		return &Ttl{ValueS: "nil"}
	}

}

func (s *allCommandsService) Set(req RequestAll) string {
	s.mutex.Lock()

	if req.Ttl > 0 {
		s.data[req.Key] = &Ttl{ValueS: req.ValueS, ValueL: req.ValueL, ValueD: req.ValueD, ttl: time.Now().Add(time.Second * time.Duration(req.Ttl)).Unix()}
	} else {
		s.data[req.Key] = &Ttl{ValueS: req.ValueS, ValueL: req.ValueL, ValueD: req.ValueD, ttl: -1}
	}
	s.mutex.Unlock()
	return "OK"
}

func (s *allCommandsService) Delete(key string) string {
	s.mutex.Lock()
	delete(s.data, key)
	s.mutex.Unlock()
	return "OK"
}

func MapConctenation(a map[string]string, b map[string]string) map[string]string {
	for k := range b {
		a[k] = b[k]
	}
	return a
}

func ListConctenation(a []string, b []string) []string {
	for i := 0; i < len(b); i++ {
		a = append(a, b[i])
	}
	return a
}

func (s *allCommandsService) Append(req RequestAll) string {
	s.mutex.Lock()

	s.data[req.Key] = &Ttl{ValueS: s.data[req.Key].ValueS, ValueL: ListConctenation(s.data[req.Key].ValueL, req.ValueL), ValueD: MapConctenation(s.data[req.Key].ValueD, req.ValueD), ttl: s.data[req.Key].ttl}

	s.mutex.Unlock()
	return "OK"
}

func (s *allCommandsService) Keys() []string {
	var tmp = []string{}
	s.mutex.Lock()
	for k := range s.data {
		tmp = append(tmp, k)
	}
	s.mutex.Unlock()
	return tmp
}

func (s *allCommandsService) Check() {

	for {
		for k := range s.data {
			if s.data[k].ttl <= time.Now().Unix() && s.data[k].ttl != -1 {
				fmt.Println("delete: ", k)
				delete(s.data, k)
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func (s *allCommandsService) Save() {
	tmp := make(map[string]Ttl)
	s.mutex.Lock()
	for k := range s.data {
		tmp[k] = *s.data[k]
	}
	s.mutex.Unlock()

	fo, err := os.Create("outputALL.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	for k := range s.data {
		var strL string
		var strD string
		for i := 0; i < len(tmp[k].ValueL); i++ {
			strL = strL + tmp[k].ValueL[i] + ", "
		}
		for j := range tmp[k].ValueD {
			strD = strD + "\t" + j + ": " + tmp[k].ValueD[j] + "\n "
		}
		fo.WriteString(k + ": " + tmp[k].ValueS + "\n" + strL + "\n" + strD + "\nttl:" + strconv.FormatInt(tmp[k].ttl, 10) + "\n")
	}

}

func NewAllCommandsService() AllCommandsService {
	return &allCommandsService{data: make(map[string]*Ttl)}
}

func NewMockAllCommandsService() AllCommandsService {
	return &allCommandsService{data: map[string]*Ttl{"key1": {ValueL: []string{"value1", "dsad"}, ValueS: "llll", ttl: time.Now().Add(time.Second * 80).Unix()}, "key2": {ValueL: []string{"value2", "ds123123ad"}, ttl: time.Now().Add(time.Minute * 3).Unix()}, "key3": {ValueL: []string{"value3", "dsad", "dlllllll"}, ValueD: map[string]string{"value3": "1111111", "value2": "jjj"}, ttl: time.Now().Add(time.Minute * 6).Unix()}}}
}
