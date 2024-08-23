package handler

import (
	internalError "api/internal/errors"
	"api/internal/models"
	utils "api/pkg/httputils"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("my_secret_key")

type RegisterCredentials struct {
	Username *string `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type UserInfo struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Signup(c *gin.Context) {
	var creds RegisterCredentials

	if err := c.BindJSON(&creds); err != nil {
		internalError.SetJsonBindingErrorResponse(c, err)
		return
	}

	if (creds.Email == "") || (creds.Password == "") {
		internalError.SetUnprocessableEntity(c, "Email address and password are required")

		return
	}

	if !utils.EmailValidator(creds.Email) {
		internalError.SetUnprocessableEntity(c, "User email address is invalid")

		return
	}

	connection := utils.GetConnection()

	if connection.Where("email = ?", creds.Email).First(&models.User{}).RowsAffected > 0 {
		internalError.SetUnprocessableEntity(c, "A user with this email address already exists")

		return
	}

	if creds.Username != nil && connection.Where("username = ?", *creds.Username).First(&models.User{}).RowsAffected > 0 {
		internalError.SetUnprocessableEntity(c, "A user with this username already exists")

		return
	}

	hashedPassword, err := utils.HashPassword(creds.Password)

	if err != nil {
		internalError.SetInternalServerError(c, "Internal server error while hashing password")

		return
	}

	var user models.User
	var errorOccuredInTransaction error

	connection.Transaction(func(transaction *gorm.DB) error {

		user = models.User{
			Username: creds.Username,
			Email:    creds.Email,
			Password: string(hashedPassword),
		}

		result := transaction.Create(&user)

		if result.Error != nil {
			errorOccuredInTransaction = result.Error

			return errorOccuredInTransaction
		}

		vehicles := []models.Vehicle{}

		transaction.Find(&vehicles)

		if len(vehicles) == 0 {

			errorOccuredInTransaction = errors.New(
				"cannot register a user because no vehicles are available in the database",
			)

			return errorOccuredInTransaction
		}

		for _, vehicle := range vehicles {
			_, createVehiculeStateError := vehicle.InitVehicleState(&user, transaction)

			if createVehiculeStateError != nil {
				errorOccuredInTransaction = errors.New(
					createVehiculeStateError.Error(),
				)

				return createVehiculeStateError
			}
		}
		return nil
	})

	if errorOccuredInTransaction != nil {
		internalError.SetInternalServerError(c, errorOccuredInTransaction.Error())

		return
	}

	utils.SetCreated(c, "User created", user)
}

func Login(c *gin.Context) {
	var creds LoginCredentials

	if err := c.BindJSON(&creds); err != nil {
		internalError.SetJsonBindingErrorResponse(c, err)
		return
	}

	if (creds.Email == "") || (creds.Password == "") {
		internalError.SetUnprocessableEntity(c, "Email address and password are required")

		return
	}

	if !utils.EmailValidator(creds.Email) {
		internalError.SetUnprocessableEntity(c, "User email address is invalid")

		return
	}

	var user models.User

	query := utils.GetConnection().Where("email = ?", creds.Email).Find(&user)

	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			internalError.SetNotFound(c, "Invalid credentials")
			return
		}

		internalError.SetInternalServerError(c, "Internal server error")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		internalError.SetUnprocessableEntity(c, "Invalid credentials")
		return
	}

	expirationTime := time.Now().Add(3 * time.Hour)

	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		internalError.SetInternalServerError(c, "Internal server error")
		return
	}

	userInfo := UserInfo{
		ID:        user.ID,
		Email:     user.Email,
		Username:  *user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"token":     tokenString,
		"expire_at": expirationTime,
		"user":      userInfo,
	})
}

func Welcome(c *gin.Context) {
	claims, err := GetClaimsFromToken(c)

	if err != nil {
		internalError.SetUnauthorized(c, "Unauthorized")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome " + claims.Email})
}
