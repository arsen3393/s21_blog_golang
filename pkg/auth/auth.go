package auth

//
//import (
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//const secretKey = "my_secret_key"
//
//func authMiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenString := c.GetHeader("Authorization")
//		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, http.ErrAbortHandler
//			}
//			return []byte(secretKey), nil
//		})
//		if err != nil || !token.Valid {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//			c.Abort()
//			return
//		}
//		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//			c.Set("claims", claims)
//		} else {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var jwtKey = []byte("secret")

func GenerateToken(userID int, password string) (string, error) {
	userIDString := strconv.Itoa(userID)
	claims := jwt.StandardClaims{
		Issuer:    "blog-app",
		Subject:   userIDString,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
