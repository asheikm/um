package handlers

import (
	//"encoding/json"
	"fmt"
	"net/http"

	// "io"
	"um/models"
	"um/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		fmt.Println(err)
		return
	}

	if h.UserService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserService is not initialized"})
		fmt.Println("UserService is not initialized")
		return
	}

	// TODO: Validate user data if needed (e.g., check for empty fields)
	fmt.Println("Creating user...")
	if err := h.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

/*
func (h *UserHandler) Signup(c *gin.Context) {
	var user models.User

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error reading body
		return
	}

	fmt.Println("Received JSON data:", string(body))

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// TODO: Validate user data if needed (e.g., check for empty fields)
	fmt.Println("Creating user...")
	err = h.UserService.CreateUser(&user)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}*/

func (h *UserHandler) Login(c *gin.Context) {
	// ... handle user login
}

func (h *UserHandler) GetMe(c *gin.Context) {
	// ... retrieve and return authenticated user details
}
