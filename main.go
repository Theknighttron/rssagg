package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
    "github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

    // Load the environment variable
    godotenv.Load(".env")

    portString := os.Getenv("PORT")
    if portString == " " {
        log.Fatal("PORT is not found in the environment variable")
    }

    router := chi.NewRouter()

    // Cors configurations
    router.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{"https://*", "http://*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: false,
        MaxAge: 300,
    }))

    v1Router := chi.NewRouter()
    v1Router.Get("/healthz", handlerReadiness)
    v1Router.Get("/err", handlerError)

    // Mount all v1Router to /v1 path
    router.Mount("/v1", v1Router)

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
