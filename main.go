package main

import (
    "github.com/labstack/echo/v4"
    "os"
    "log"
    "github.com/vakter/users/actions"
)

func main() {
    e := echo.New()

    e.POST("/users", actions.CreateUser)
    e.GET("/users/:id", actions.GetUser)
    e.PUT("/users/:id", actions.UpdateUser)
    e.DELETE("/users/:id", actions.DeleteUser)
    e.GET("/users", actions.GetUsers)
    
    

    port := os.Getenv("SERVER_PORT")
    if port == ""{
        port = ":8080" // Default port
    }

    // Start server
    log.Printf("Starting server on :%s...", port)
    e.Logger.Fatal(e.Start(":" + port))
}

