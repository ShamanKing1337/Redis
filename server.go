package main

import (
	"Avito/controller"
	"Avito/service"
	"log"
	"net/http"
)

func main() {

	stringCommandsService := service.NewMockStringCommandsService()
	a := controller.NewStringCommandsImplementation(stringCommandsService)
	go stringCommandsService.Check()

	listCommandsService := service.NewMockListCommandsService()
	b := controller.NewListCommandsImplementation(listCommandsService)
	go listCommandsService.Check()

	dictCommandsService := service.NewMockDictCommandsService()
	c := controller.NewDictCommandsImplementation(dictCommandsService)
	go dictCommandsService.Check()

	allCommandsService := service.NewMockAllCommandsService()
	d := controller.NewAllCommandsImplementation(allCommandsService)
	go allCommandsService.Check()

	http.HandleFunc("/SSET", controller.BasicAuth(a.SetString, "Please enter your username and password"))

	http.HandleFunc("/SGET/", controller.BasicAuth(a.GetString, "Please enter your username and password"))
	http.HandleFunc("/SDEL/", controller.BasicAuth(a.DeleteKey, "Please enter your username and password"))
	http.HandleFunc("/SKEYS", controller.BasicAuth(a.GetKeys, "Please enter your username and password"))
	http.HandleFunc("/SSAVE", controller.BasicAuth(a.SaveString, "Please enter your username and password"))

	http.HandleFunc("/LSET", controller.BasicAuth(b.SetList, "Please enter your username and password"))
	http.HandleFunc("/LGET/", controller.BasicAuth(b.GetList, "Please enter your username and password"))
	http.HandleFunc("/LDEL/", controller.BasicAuth(b.DeleteKey, "Please enter your username and password"))
	http.HandleFunc("/LKEYS", controller.BasicAuth(b.GetKeys, "Please enter your username and password"))
	http.HandleFunc("/LSAVE", controller.BasicAuth(b.SaveList, "Please enter your username and password"))

	http.HandleFunc("/HSET", controller.BasicAuth(c.SetDict, "Please enter your username and password"))
	http.HandleFunc("/HGET/", controller.BasicAuth(c.GetDict, "Please enter your username and password"))
	http.HandleFunc("/HDEL/", controller.BasicAuth(c.DeleteKey, "Please enter your username and password"))
	http.HandleFunc("/HKEYS", controller.BasicAuth(c.GetKeys, "Please enter your username and password"))
	http.HandleFunc("/HSAVE", controller.BasicAuth(c.SaveDict, "Please enter your username and password"))

	http.HandleFunc("/SET", controller.BasicAuth(d.SetAll, "Please enter your username and password"))
	http.HandleFunc("/GET/", controller.BasicAuth(d.GetAll, "Please enter your username and password"))
	http.HandleFunc("/DEL/", controller.BasicAuth(d.DeleteKey, "Please enter your username and password"))
	http.HandleFunc("/KEYS", controller.BasicAuth(d.GetKeys, "Please enter your username and password"))
	http.HandleFunc("/SAVE", controller.BasicAuth(d.SaveAll, "Please enter your username and password"))
	http.HandleFunc("/APPEND", controller.BasicAuth(d.AppendAll, "Please enter your username and password"))

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}
