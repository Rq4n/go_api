package main

import (
	"fmt"
	"log"
	"minha-primeira-api/handler"
	"minha-primeira-api/internal/models/database"
	"net/http"
)

func main() {
	database.Connect()
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/protected", handler.ProtectedHandler)
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
