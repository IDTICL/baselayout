package auth

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret")

type Claims struct {
	Account string `json:"username"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

func GenerToken(username, role string) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	jwtId := username + strconv.FormatInt(nowTime.Unix(), 10)

	claims := Claims{
		Account: username,
		Role:    role,
		StandardClaims: jwt.StandardClaims{
			Audience:  username,
			ExpiresAt: expireTime.Unix(),
			Id:        jwtId,
			IssuedAt:  nowTime.Unix(),
			Issuer:    "psi_gin_pg",
			NotBefore: nowTime.Add(10 * time.Second).Unix(),
			Subject:   username,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func AuthRequired(c *gin.Context) {
	token := c.GetHeader("Authorization")

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return jwtSecret, nil
		})
	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {

		log.Println("account:", claims.Account)
		log.Println("role:", claims.Role)
		c.Set("account", claims.Account)
		c.Set("role", claims.Role)
		c.Next()
	} else {
		c.Abort()
		return
	}
}
