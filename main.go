package main

import (
	"fmt"
	"net/http"

	"kilvish.io/cerebro/server"
)

func main() {
	fmt.Println("Hello world!")
	server.Init()
	router := server.GetRouter()
	http.ListenAndServe(":8080", router)
}

// func main() {
// 	fmt.Println("Hello world!")
// 	kp := service.KafkaPublisher{}
// 	kp.Publish("kilvish-test", "gfjkwnljllfjwf")
// }
