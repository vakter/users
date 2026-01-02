package actions

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// In-memory mock database for demonstration purposes
var users = make(map[string]User)

// CreateUser creates a new user
func CreateUser(c echo.Context) error {
    user := User{}
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    users[user.ID] = user
    return c.JSON(http.StatusCreated, user)
}

// GetUser retrieves a user by ID
func GetUser(c echo.Context) error {
    id := c.Param("id")
    user, exists := users[id]
    if !exists {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
    }
    return c.JSON(http.StatusOK, user)
}

// UpdateUser updates an existing user by ID
func UpdateUser(c echo.Context) error {
    id := c.Param("id")
    userUpdate := User{}
    if err := c.Bind(&userUpdate); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    if _, exists := users[id]; !exists {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
    }
    userUpdate.ID = id // Maintain the same ID
    users[id] = userUpdate
    return c.JSON(http.StatusOK, userUpdate)
}

// DeleteUser deletes a user by ID
func DeleteUser(c echo.Context) error {
    id := c.Param("id")
    if _, exists := users[id]; !exists {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
    }
    delete(users, id)
    return c.JSON(http.StatusNoContent, nil)
}

// GetUsers retrieves all users
func GetUsers(c echo.Context) error {
    userList := []User{}
    for _, user := range users {
        userList = append(userList, user)
    }
    return c.JSON(http.StatusOK, userList)
}
