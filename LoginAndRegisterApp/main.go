package main

import (
	auth "LoginAndRegisterApp/authorization"
	count "LoginAndRegisterApp/counters"
	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)
var users = count.Start()

func main() {
	router.Any("/startPage", auth.RenderRegistrationPage)
	//Signup
	router.POST("/register", auth.Register)
	//Login
	router.Any("/loginPage", auth.RenderLoginPage)
	router.POST("/login", auth.Login)

	log.Fatal(router.Run(":8080"))
}
