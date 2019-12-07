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
