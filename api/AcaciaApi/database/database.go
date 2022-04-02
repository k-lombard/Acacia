package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST = "db"
	PORT = 5433
)

type Database struct {
	Conn *gorm.DB
}

var ErrNoMatch = fmt.Errorf("Error: no matching table record")

var ErrDuplicate = fmt.Errorf("Error: table record already exists")

func InitializeDatabase(username, password, database string) (Database, error) {
	db := Database{}
	instanceConnection := os.Getenv("INSTANCE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		instanceConnection,
		username,
		password,
		database)
	conn, err := gorm.Open(postgres.New(postgres.Config{DriverName: "cloudsqlpostgres", DSN: dsn}), &gorm.Config{})
	if err != nil {
		return db, err
	}
	db.Conn = conn
	log.Println("Database connection established")
	return db, nil
}
