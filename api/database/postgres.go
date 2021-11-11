package database

import (
	"bytes"
	"fmt"
	"log"
	"oloapi/api/models"
	"os"
	"os/exec"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB represents a Database instance
var DB *gorm.DB

// ConnectToDB connects the server with database
func ConnectToDB() {

	// creating database
	// when PGPASSWORD is set we dont need to provide a password interactively
	os.Setenv("PGPASSWORD", os.Getenv("PSQL_PASS"))
	cmd := exec.Command("createdb", "-p", os.Getenv("PSQL_PORT"), "-h", os.Getenv("PSQL_IP"), "-U", os.Getenv("PSQL_USER"), "-e", os.Getenv("PSQL_DBNAME"))
	var out bytes.Buffer
	//stores output of cmd after run in out buffer so that we can print it afterwards
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		// %v checks if the value (in this case err) includes the fmt.Stringer interface (which is a single String() string method)
		log.Printf("Error (Database already exists): %v\n", err)
	}
	// %q like %s but safely escapes a string and puts quotes to it
	log.Printf("Database create output: %q", out.String())

	//connecting to postgres database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("PSQL_IP"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PORT"))

	log.Print("Connecting to Postgres DB...")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)

	}
	log.Println("connected")

	// turned on the loger on info mode
	DB.Logger = logger.Default.LogMode(logger.Silent)

	log.Print("Running the migrations...")
	DB.AutoMigrate(&models.User{}, &models.Location{})
}
