package d6gorm

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

type User struct {
	gorm.Model
	Name      string `json:"name" form:"name" gorm:"not null"`
	Email     string `json:"email" form:"email" gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type BaseResponse struct {
	Status   string      `json:"status"`
	Messages string      `json:"messages"`
	Data     interface{} `json:"data"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

func getUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"messages": "success get all users",
		"users":    users,
	})
}

func getUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"messages": "success get user",
		"user":     user,
	})
}

func createUserController(c echo.Context) error {
	user := new(User)
	c.Bind(user)
	if err := DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"messages": "success create user",
		"user":     user,
	})
}

func updateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user User

	// Find the user by ID
	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	// Bind the request data to the user
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	// Update the user in the database
	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "User updated successfully",
		"user":    user,
	})
}

func deleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DB.Delete(&User{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"messages": "success delete user",
	})
}

func main() {
	e := echo.New()

	e.GET("/users", getUsersController)
	e.GET("/users/:id", getUserController)
	e.POST("/users", createUserController)
	e.PUT("/users/:id", updateUserController)
	e.DELETE("/users/:id", deleteUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
