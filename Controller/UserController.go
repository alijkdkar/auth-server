package Controller

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"time"

	initilzer "github.com/alijkdkar/auth-server-go/Initilzer"
	"github.com/alijkdkar/auth-server-go/Models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SignUP(c *gin.Context) {
	type body struct {
		Email    string
		Password string
	}
	bod := body{}
	err := c.ShouldBindJSON(&bod)
	if err != nil {
		log.Println("Error To Bind RequstBody")
	}
	hashedPass := sha256.Sum256([]byte(bod.Password))
	m := Models.MyUser{Email: &bod.Email, Password: fmt.Sprintf("%x", hashedPass)}

	res := initilzer.DB.Create(&m)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message:": res.Error,
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message: row Affected:": res.RowsAffected,
	})
}

func Login(c *gin.Context) {
	type body struct {
		Email    string
		Password string
	}
	bod := body{}
	err := c.ShouldBindJSON(&bod)
	if err != nil {
		log.Println("Error To Bind RequstBody")
	}

	user := Models.MyUser{}
	initilzer.DB.First(&user, "Email = ?", string(bod.Email))

	inputHashPass := sha256.Sum256([]byte(bod.Password))

	if user.ID == 0 || user.Password != fmt.Sprintf("%x", inputHashPass) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "UserName or password is wrong",
		})
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(initilzer.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error to create Token" + err.Error(),
		})
		return
	}
	c.SetCookie("authToken", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "You Logined in Success",
	})

}

func Validation(c *gin.Context) {
	println("begin Valida")
	user, exists := c.Get("onlineUser")

	if exists {
		c.JSON(http.StatusAccepted, gin.H{"message": user})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"message": "User not Found!"})
	}

}
