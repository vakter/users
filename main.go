package main

import (
    "github.com/labstack/echo/v4"
    "net/http"
    "users-microservice/actions"
    "github.com/vakter/users/actions"
)

func main() {
    e := echo.New()

    e.POST("/users", actions.CreateUser)
    e.GET("/users/:id", actions.GetUser)
    e.PUT("/users/:id", actions.UpdateUser)
    e.DELETE("/users/:id", actions.DeleteUser)
    e.GET("/users", actions.GetUsers)

    e.Logger.Fatal(e.Start(":8080"))
}
