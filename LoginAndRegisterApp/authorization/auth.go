package authorization

import (
	"LoginAndRegisterApp/helpers"
	user "LoginAndRegisterApp/models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	count "LoginAndRegisterApp/counters"
)
var users = count.Start()

func RenderRegistrationPage(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "register.html")
}
func RenderLoginPage(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "login.html")
}

func Register(c *gin.Context) {
	pwdConfirm := ""
	user := user.User{}
	err := c.Request.ParseForm()
	if err != nil {
		log.Println("some error happened")
		return
	}
	user.Username = c.Request.FormValue("username")
	user.ID = GenerateNum()
	user.Email = c.Request.FormValue("email")
	user.Password = c.Request.FormValue("password")
	pwdConfirm = c.Request.FormValue("confirm")
	//Empty data checking
	uNameCheck := helpers.IsEmpty(user.Username)

	emailCheck := helpers.IsEmpty(user.Email)

	pwdCheck := helpers.IsEmpty(user.Password)

	pwdConfirmCheck := helpers.IsEmpty(pwdConfirm)

	if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
		fmt.Fprintf(c.Writer, "ErrorCode is -10 : There is empty data.")
		return
	} else {
		users.InitCounter(user.Password, user)
		c.Writer.WriteHeader(http.StatusOK)
		p, err := json.Marshal(users.V)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Fprintln(c.Writer, "Registration successful\n")
		c.Writer.Write(p)

	}

}
func CreateToken(userId uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
func Login(c *gin.Context) {
	email := c.Request.FormValue("email")  // Data from the form
	pwd := c.Request.FormValue("password") // Data from the form

	// Empty data checking
	emailCheck := helpers.IsEmpty(email)
	pwdCheck := helpers.IsEmpty(pwd)

	if emailCheck || pwdCheck {
		fmt.Fprintf(c.Writer, "ErrorCode is -10 : There is empty data.")
		return
	}
	for k, v := range users.V {

		if pwd == k && email == v.Email {
			//router.POST(r.URL.Path,Login)
			fmt.Fprintln(c.Writer, "Login succesful!")
			token, err := CreateToken(v.ID)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}
			//fmt.Println("token:", token)
			c.JSON(http.StatusOK, token)
		} else {
			c.JSON(http.StatusUnauthorized, "Please provide valid login details")

			fmt.Fprintln(c.Writer, "Login failed!")
			return
		}
	}
}
func GenerateNum() uint64 {
	rand.Seed(time.Now().UnixNano())
	randomNum := uint64random(1, 30)
	return randomNum
}
func uint64random(value1, value2 uint64) uint64 {
	return value1 + value2 + rand.Uint64()
}