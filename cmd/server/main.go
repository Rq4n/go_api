package main

import (
	"fmt"
	"log"
	"minha-primeira-api/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/getUser", handler.GetUser)
	mux.HandleFunc("/api/createUser", handler.CreateUser)
	mux.HandleFunc("/api/deleteUser", handler.DeleteUser)
	mux.HandleFunc("/api/updateUser", handler.UpdateUser)

	fmt.Println("HTTP server running!")

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
