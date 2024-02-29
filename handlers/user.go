package handlers

import (
	"fmt"
	"net/http"

	"crypto/sha256"
	"encoding/hex"
	"um/models"
	"um/services"

	"github.com/dgrijalva/jwt-go"

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

func checkPassword(storedPassword, inputPassword string) bool {
	// Hash the input password using SHA-256
	hashedStoredPassword := hashPassword(storedPassword)
	inputPasswordHash := hashPassword(inputPassword)
	return hashedStoredPassword == inputPasswordHash
}

func hashPassword(password string) string {
	// Hash the password using SHA-256
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword
}

func generateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		// Add any additional claims as needed
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Validate loginReq fields if needed

	user, err := h.UserService.GetUserByUsername(loginReq.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if !checkPassword(user.Password, loginReq.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) GetMe(c *gin.Context) {
	// ... retrieve and return authenticated user details
}
