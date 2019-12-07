package server

import (
	"fmt"

	"github.com/gorilla/mux"
	"kilvish.io/cerebro/server/handler"
)

var r *mux.Router

func Init() {
	// all initiliazation go here
	fmt.Println("initializing handlers!")
	r = mux.NewRouter()
	mapHandlers(r)
}

func mapHandlers(r *mux.Router) {
	r.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	aah := handler.NewAutoAssignHandler()
	r.HandleFunc("/auto-assign", aah.ServeHTTP).Methods("GET")
	//SwaggerUI Handler
	// fs := http.FileServer(http.Dir("./swaggerui/"))
	// r.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
	// r.HandleFunc("/swagger.yml", func(wr http.ResponseWriter, rd *http.Request) {
	// 	http.ServeFile(wr, rd, path.Join(config.SwaggerPath(), "swagger.yml"))
	// })

}

func GetRouter() *mux.Router {
	return r
}
