package handlers

import (
	"net/http"

	"fmt"
	"io"
	"um/models"
	"um/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepository repositories.UserRepository // Add UserRepository dependency
}

func NewUserHandler(userRepository repositories.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
	}
}

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

	err = h.UserRepository.CreateUser(&user)
	if err != nil {
		// Handle error
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
	// ... handle user login
}

func (h *UserHandler) GetMe(c *gin.Context) {
	// ... retrieve and return authenticated user details
}
