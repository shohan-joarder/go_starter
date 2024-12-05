package repositories

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shohan-joarder/go_pos/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

// NewAuthRepository initializes a new AuthRepository
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

// Login authenticates the user and generates a JWT token
func (r *AuthRepository) Login(user *models.LoginUser) (string, error) {
	var dbUser models.User

	// Check if the user exists
	err := r.DB.First(&dbUser, "email = ?", user.Email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	// Verify the password
	if err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	// Generate the JWT token
	token, err := createToken(dbUser)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func createToken(userDetails models.User) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = userDetails.ID
	claims["name"] = userDetails.Name
	claims["email"] = userDetails.Email
	claims["phone"] = userDetails.Phone
	claims["role"] = userDetails.RoleID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
