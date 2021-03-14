package service

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type StringCommandsService interface {
	Get(key string) string
	Set(key string, value string, ttl int64) string
	Delete(key string) string
	Keys() []string
	Check()
	Save()
}

type TtlStr struct {
	value string
	ttl   int64
}

type stringCommandsService struct {
	data  map[string]*TtlStr
	mutex sync.Mutex
}

func (s *stringCommandsService) Get(key string) string {
	s.mutex.Lock()
	tmp, ok := s.data[key]
	s.mutex.Unlock()
	if ok {
		return tmp.value
	} else {
		return "(nil)"
	}

}

func (s *stringCommandsService) Set(key string, value string, ttl int64) string {
	s.mutex.Lock()
	if ttl > 0 {
		s.data[key] = &TtlStr{value: value, ttl: time.Now().Add(time.Second * time.Duration(ttl)).Unix()}
	} else {
		s.data[key] = &TtlStr{value: value, ttl: -1}
	}
	s.mutex.Unlock()
	return "OK"
}

func (s *stringCommandsService) Delete(key string) string {
	s.mutex.Lock()
	delete(s.data, key)
	s.mutex.Unlock()
	return "OK"
}

func (s *stringCommandsService) Keys() []string {
	var tmp = []string{}
	s.mutex.Lock()
	for k := range s.data {
		tmp = append(tmp, k)
	}
	s.mutex.Unlock()
	return tmp
}

func (s *stringCommandsService) Check() {

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

func (s *stringCommandsService) Save() {
	tmp := make(map[string]TtlStr)
	s.mutex.Lock()
	for k := range s.data {
		tmp[k] = *s.data[k]
	}
	s.mutex.Unlock()

	fo, err := os.Create("outputSTR.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	for k := range s.data {
		fo.WriteString(k + ": " + tmp[k].value + ", " + strconv.FormatInt(tmp[k].ttl, 10) + "\n")
	}

}

func NewStringCommandsService() StringCommandsService {
	return &stringCommandsService{data: make(map[string]*TtlStr)}
}

func NewMockStringCommandsService() StringCommandsService {
	return &stringCommandsService{data: map[string]*TtlStr{"key1": {"value1", time.Now().Add(time.Second * 20).Unix()}, "key2": {"value2", time.Now().Add(time.Minute * 3).Unix()}, "key3": {"value3", time.Now().Add(time.Minute * 6).Unix()}}}
}
