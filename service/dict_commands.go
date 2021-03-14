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

type DictCommandsService interface {
	Get(key string) map[string]string
	Set(key string, value Dict, ttl int64) string
	Delete(key string) string
	Keys() []string
	Check()
	Save()
}

type Dict struct {
	KeyDict   string
	ValueDict string
}

type TtlDict struct {
	value map[string]string
	ttl   int64
}

type dictCommandsService struct {
	data  map[string]*TtlDict
	mutex sync.Mutex
}

func (s *dictCommandsService) Get(key string) map[string]string {
	s.mutex.Lock()
	tmp, ok := s.data[key]
	s.mutex.Unlock()
	if ok {
		return tmp.value
	} else {
		return map[string]string{"nil": "nil"}
	}

}

func (s *dictCommandsService) Set(key string, value Dict, ttl int64) string {
	s.mutex.Lock()

	if s.data[key] == nil {
		s.data[key] = &TtlDict{}
		s.data[key].value = map[string]string{}
	}

	tmp := s.data[key].value

	tmp[value.KeyDict] = value.ValueDict

	if ttl > 0 {
		s.data[key] = &TtlDict{value: tmp, ttl: time.Now().Add(time.Second * time.Duration(ttl)).Unix()}
	} else {
		s.data[key] = &TtlDict{value: tmp, ttl: -1}
	}
	s.mutex.Unlock()
	return "OK"
}

func (s *dictCommandsService) Delete(key string) string {
	s.mutex.Lock()
	delete(s.data, key)
	s.mutex.Unlock()
	return "OK"
}

func (s *dictCommandsService) Keys() []string {
	var tmp = []string{}
	s.mutex.Lock()
	for k := range s.data {
		tmp = append(tmp, k)
	}
	s.mutex.Unlock()
	return tmp
}

func (s *dictCommandsService) Check() {

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

func (s *dictCommandsService) Save() {
	tmp := make(map[string]TtlDict)
	s.mutex.Lock()
	for k := range s.data {
		tmp[k] = *s.data[k]
	}
	s.mutex.Unlock()

	fo, err := os.Create("outputDICT.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	for k := range s.data {
		var str string

		for j := range tmp[k].value {
			str = str + "\t" + j + ": " + tmp[k].value[j] + "\n "
		}

		fo.WriteString(k + ": \n" + str + "\t" + "ttl: " + strconv.FormatInt(tmp[k].ttl, 10) + "\n")
	}
}

func NewDictCommandsService() DictCommandsService {
	return &dictCommandsService{data: make(map[string]*TtlDict)}
}

func NewMockDictCommandsService() DictCommandsService {
	return &dictCommandsService{data: map[string]*TtlDict{"key1": {map[string]string{"value1": "dsad"}, time.Now().Add(time.Second * 20).Unix()}, "key2": {map[string]string{"value2": "dsadasdsad"}, time.Now().Add(time.Minute * 3).Unix()}, "key3": {map[string]string{"value3": "1111111", "value2": "jjj"}, time.Now().Add(time.Minute * 6).Unix()}}}
}
