package Middelware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alijkdkar/auth-server-go/Initilzer"
	"github.com/alijkdkar/auth-server-go/Models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckJWT(c *gin.Context) {

	token1, err := c.Cookie("authToken")
	if err != nil && token1 == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(token1, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", "token.Header[alg]")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return Initilzer.SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		println("in Ok")
		// fmt.Println(claims["foo"], claims["nbf"])
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		user := Models.MyUser{}
		Initilzer.DB.First(&user, claims["ID"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("onlineUser", user)

		c.Next()
	} else {
		println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}
