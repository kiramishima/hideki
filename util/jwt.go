package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"hideki/internal/core/domain"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

// GenerateJWT generate JWT token
func GenerateJWT(user domain.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	//log.Println(time.Now())
	//log.Println(time.Now().Add(time.Second * time.Duration(tokenTTL)))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.RoleID,
		"iat":  time.Now().Unix(),
		"eat":  time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

// ValidateJWT validate JWT token
func ValidateJWT(req *http.Request) error {
	token, err := getToken(req)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

// ValidateAdminRoleJWT validate Admin role
func ValidateAdminRoleJWT(req *http.Request) error {
	token, err := getToken(req)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 1 {
		return nil
	}
	return errors.New("invalid admin token provided")
}

// ValidateCustomerRoleJWT validate Customer role
func ValidateCustomerRoleJWT(req *http.Request) error {
	token, err := getToken(req)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 2 || userRole == 1 {
		return nil
	}
	return errors.New("invalid customer or admin token provided")
}

// CurrentUser fetch user details from the token
/*func CurrentUser(req *http.Request) domain.User {
	err := ValidateJWT(req)
	if err != nil {
		return domain.User{}
	}
	token, _ := getToken(req)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	user, err := domain.GetUserById(userId)
	if err != nil {
		return domain.User{}
	}
	return user
}*/

// getToken check token validity
func getToken(req *http.Request) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(req)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

// getTokenFromRequest extract token from request Authorization header
func getTokenFromRequest(req *http.Request) string {
	bearerToken := req.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
