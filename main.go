package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/AlanL2/rss-scraper/internal/database"
	
	_ "github.com/lib/pq"
)

type apiConfig struct { // holds connection to db
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("ERROR: PORT NOT FOUND IN ENVIRONMENT")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("ERROR: DB_URL NOT FOUND IN ENVIRONMENT")
	}

	db_conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error: database cannot be connected to", err)
	}
	
	queries := database.New(db_conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()

	//handlers
	v1Router.Get("/healthz", handlerReadiness) // Get scopes handler to only fire on Get, otherwise use handleFunc
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.handlerGetUser)

	router.Mount("/v1", v1Router) // nesting v1 router to /v1 path
	// full path is /v1/healthz, responds if server is alive and running

	srv := &http.Server{
		Handler: router, 
		Addr: ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port: ", portString)
}