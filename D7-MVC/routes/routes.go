package routes

import (
	"MVC/controllers"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {

	e.GET("/users", controllers.GetUsersController)          //get all users
	e.GET("/users/:id", controllers.GetUserController)       //get user by id
	e.POST("/users", controllers.CreateUserController)       //Create new user
	e.PUT("/users/:id", controllers.UpdateUserController)    //update user
	e.DELETE("/users/:id", controllers.DeleteUserController) //delete user

}
