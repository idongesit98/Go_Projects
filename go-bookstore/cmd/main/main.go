package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/idongesit98/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/",router)
	fmt.Print("Server starting at Port 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010",router))
}