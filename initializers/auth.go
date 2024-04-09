package initializers

import (
	"errors"
	"fmt"
	"github.com/RakhimovAns/FinalYandexTask/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"net/url"
)

// Define Authorized middleware with a closure to pass channel parameter
func Authorized() gin.HandlerFunc {
	// Return the actual middleware function
	return func(c *gin.Context) {
		auth, err := ReadCookie("token", c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			c.Abort()
			return
		}

		token, id, err := ParseToken(auth)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			c.Abort()
			return
		}
		// Set the user ID in the context
		c.Set("userID", id)
		c.Next()
	}
}

// Other functions remain unchanged

func ParseToken(accessToken string) (*jwt.Token, int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing methon: %v ", token.Header["alg"])
		}
		return []byte("My Key"), nil
	})
	if err != nil {
		return nil, -1, err
	}
	claims, ok := token.Claims.(*models.TokenClaim)
	if !ok {
		return nil, -1, err
	}
	return token, claims.UserID, err
}

func ReadCookie(name string, r *http.Request) (string, error) {
	if name == "" {
		return "", errors.New("you are trying to read empty cookie")
	}
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	str := cookie.Value
	value, _ := url.QueryUnescape(str)
	return value, nil
}
