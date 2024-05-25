package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

    godotenv.Load(".env")

    portString := os.Getenv("PORT")
    if portString == " " {
        log.Fatal("PORT is not found in the environment variable")
    }

    router := chi.NewRouter()

    serve := &http.Server{
        Handler: router,
        Addr: ":" + portString,
    }

    log.Printf("Server is running of port: %v ", portString)
    err := serve.ListenAndServe()
    if err != nil {
        log.Fatal("An error occured: ", err)
    }


    fmt.Println("PORT: ", portString)
}
