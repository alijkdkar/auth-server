package main

import (
	"github.com/alijkdkar/auth-server-go/Controller"
	"github.com/alijkdkar/auth-server-go/Initilzer"
	"github.com/alijkdkar/auth-server-go/Middelware"
	"github.com/gin-gonic/gin"
)

func main() {
	Initilzer.InitilzerMethod()
	r := gin.Default()
	r.POST("/SignUp", Controller.SignUP)
	r.POST("/Login", Controller.Login)
	r.GET("/validation", Middelware.CheckJWT, Controller.Validation)
	r.Run() // listen and serve on 0.0.0.0:8080

}
