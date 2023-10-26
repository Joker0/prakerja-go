package configs

import (
	"MVC/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Host:     "localhost", // PostgreSQL server address
		DB_Port:     "5432",      // PostgreSQL port
		DB_Username: "postgres",  // Your PostgreSQL username
		DB_Password: "secret",    // Your PostgreSQL password
		DB_Name:     "crud_go",   // Your database name
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DB_Host,
		config.DB_Port,
		config.DB_Username,
		config.DB_Password,
		config.DB_Name)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
}
