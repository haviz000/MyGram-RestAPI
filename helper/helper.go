package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/haviz000/MyGram-RestAPI/model"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

type authCustomClaims struct {
	Username string
	UserID   string
	refresh  bool
	jwt.StandardClaims
}

func GenerateID() string {
	return uuid.New().String()
}

func Hash(plain string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plain), 8)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func PasswordIsMatch(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

func GenerateAccessToken(userData model.User) (string, error) {
	claims := &authCustomClaims{
		Username: userData.Username,
		UserID:   userData.UserID,
		refresh:  false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "my_gram",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func GenerateRefreshToken(userData model.User) (string, error) {
	claims := &authCustomClaims{
		Username: userData.Username,
		UserID:   userData.UserID,
		refresh:  true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "my_gram",
			IssuedAt:  time.Now().Unix(),
		},
	}

	return generateToken(claims, os.Getenv("SECRET_KEY"))
}

func generateToken(claims *authCustomClaims, secret string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := jwtToken.SignedString([]byte(secret))
	return tokenString, err
}

func VerifyAccessToken(tokenStr string) (*authCustomClaims, error) {
	claims := &authCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token is invalid")
	}

	claims, ok := token.Claims.(*authCustomClaims)

	if !ok {
		return nil, fmt.Errorf("Couldn't parse claims")
	}

	return claims, nil
}

func VerifyRefreshToken(token string) (*jwt.Token, error) {

	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Token is Invalid")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	return jwtToken, err
}
