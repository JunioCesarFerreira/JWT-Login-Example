package main

import (
	"fmt"
	"log"
	"main/controller"
	"main/data"
	"main/middleware/login"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	port = 8082
)

func main() {
	fmt.Println("initializing go server")

	loginMethod := login.NewJwt()
	respository := data.NewDummy()
	controller := controller.New(loginMethod, respository)

	// Inicialização do roteador do Gorilla Mux
	router := mux.NewRouter()

	router.HandleFunc("/login", controller.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/protected", loginMethod.Authenticate(controller.ProtectedEndpoint)).Methods("GET", "OPTIONS")

	fmt.Printf("listening port %d\n", port)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
