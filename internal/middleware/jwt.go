package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

// MySigningKey ...
var MySigningKey []byte

// JWTExpireTime ...
var JWTExpireTime int

// MyCustomClaims ...
type MyCustomClaims struct {
	ID    uint64 `json:"Id"`
	Email string `json:"Email"`
	jwt.StandardClaims
}

// AuthID - Access details
var AuthID uint64

// JWT ...
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		val := c.Request.Header.Get("Authorization")
		if len(val) == 0 || !strings.Contains(val, "Bearer ") {
			log.Info("no vals or no Bearer found")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		vals := strings.Split(val, " ")
		if len(vals) != 2 {
			log.Info("result split not valid")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(vals[1], &MyCustomClaims{}, validateJWT)

		if err != nil {
			log.WithError(err).Error("error parsing JWT", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			AuthID = claims.ID
		}
	}
}

// validateJWT ...
func validateJWT(token *jwt.Token) (interface{}, error) {
	// log.Println("try to parse the JWT")
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// log.Println("error parsing JWT")
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return MySigningKey, nil
}

// GetJWT ...
func GetJWT(id uint64, email string) (string, error) {
	// Create the Claims
	claims := MyCustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(JWTExpireTime)).Unix(),
			Issuer:    "GoRest API",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtValue, err := token.SignedString(MySigningKey)
	if err != nil {
		return "", err
	}
	return jwtValue, nil
}
