package service

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type ListCommandsService interface {
	Get(key string) []string
	Set(key string, value string, ttl int64) string
	Delete(key string) string
	Keys() []string
	Check()
	Save()
}

type TtlList struct {
	value []string
	ttl   int64
}

type listCommandsService struct {
	data  map[string]*TtlList
	mutex sync.Mutex
}

func (s *listCommandsService) Get(key string) []string {
	s.mutex.Lock()
	tmp, ok := s.data[key]
	s.mutex.Unlock()
	if ok {
		return tmp.value
	} else {

		return []string{"nil"}
	}

}

func (s *listCommandsService) Set(key string, value string, ttl int64) string {
	s.mutex.Lock()
	if s.data[key] == nil {

		s.data[key] = &TtlList{}
	}
	if ttl > 0 {
		s.data[key] = &TtlList{value: append(s.data[key].value, value), ttl: time.Now().Add(time.Second * time.Duration(ttl)).Unix()}
	} else {
		s.data[key] = &TtlList{value: append(s.data[key].value, value), ttl: -1}
	}
	s.mutex.Unlock()
	return "OK"
}

func (s *listCommandsService) Delete(key string) string {
	s.mutex.Lock()
	delete(s.data, key)
	s.mutex.Unlock()
	return "OK"
}

func (s *listCommandsService) Keys() []string {
	var tmp = []string{}
	s.mutex.Lock()
	for k := range s.data {
		tmp = append(tmp, k)
	}
	s.mutex.Unlock()
	return tmp
}

func (s *listCommandsService) Check() {

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

func (s *listCommandsService) Save() {
	tmp := make(map[string]TtlList)
	s.mutex.Lock()
	for k := range s.data {
		tmp[k] = *s.data[k]
	}
	s.mutex.Unlock()

	fo, err := os.Create("outputLIST.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	for k := range s.data {
		var str string

		for i := 0; i < len(tmp[k].value); i++ {
			str = str + tmp[k].value[i] + ", "
		}
		fo.WriteString(k + ": " + str + ", " + strconv.FormatInt(tmp[k].ttl, 10) + "\n")
	}

}

func NewListCommandsService() ListCommandsService {
	return &listCommandsService{data: make(map[string]*TtlList)}
}

func NewMockListCommandsService() ListCommandsService {
	return &listCommandsService{data: map[string]*TtlList{"key1": {[]string{"value1", "dsad"}, time.Now().Add(time.Second * 20).Unix()}, "key2": {[]string{"value2", "ds123123ad"}, time.Now().Add(time.Minute * 3).Unix()}, "key3": {[]string{"value3", "dsad", "dlllllll"}, time.Now().Add(time.Minute * 6).Unix()}}}
}
