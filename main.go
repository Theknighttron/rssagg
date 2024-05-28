package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TheKnighttron/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
    DB *database.Queries
}

func main() {
	fmt.Println("Hello, World!")

    // Load the environment variable
    godotenv.Load(".env")

    portString := os.Getenv("PORT")
    if portString == " " {
        log.Fatal("PORT is not found in the environment variable")
    }

    dbUrl := os.Getenv("DB_URL")
    if dbUrl == " " {
        log.Fatal("DB_URL is not found in the environment variable")
    }


    // Connect to the database
    conn, err := sql.Open("postgres", dbUrl)
    if err != nil {
        log.Fatal("Can't connect to the database")
    }

    // Convert to db query
    queries := database.New(conn)
    apiCfg := apiConfig{
        DB: queries,
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
    v1Router.Post("/users", apiCfg.handlerCreateUser)
    v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
    v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
    v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
    v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
    v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))

    // Mount all v1Router to /v1 path
    router.Mount("/v1", v1Router)

    serve := &http.Server{
        Handler: router,
        Addr: ":" + portString,
    }

    log.Printf("Server is running of port: %v ", portString)
    err = serve.ListenAndServe()
    if err != nil {
        log.Fatal("An error occured: ", err)
    }


    fmt.Println("PORT: ", portString)
}
