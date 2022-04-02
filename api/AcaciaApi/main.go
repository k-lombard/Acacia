package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/k-lombard/Acacia/AcaciaApi/database"
	"github.com/k-lombard/Acacia/AcaciaApi/handler"
	_ "github.com/lib/pq"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	m := melody.New()
	dbUserName, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := database.InitializeDatabase(dbUserName, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	httpHandler := handler.RouteHandler(database, m)
	listener := httpHandler.Run(":8080")
	log.Fatal(listener)
	log.Printf("Started server on %s", "8080")
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	// log.Println(fmt.Sprint(<-ch))
	// log.Println("Stopping API server.")
}
