package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
http.HandleFunc("/login",loginPage)
http.HandleFunc("/logout",logoutPage)
http.HandleFunc("/",mainPage)
http.ListenAndServe(":1708",nil)
}
func mainPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("id")
	loggedIn := err != http.ErrNoCookie
	if loggedIn {
		fmt.Fprintln(w,`<a href="/logout">logout</a>`)
		fmt.Fprintln(w,"Welcome,",session.Value)
	} else {
		fmt.Fprintln(w,`<a href="/login">login</a>`)
		fmt.Fprintln(w,"You have to login!")
	}

}
func loginPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(10 * time.Hour)

	session := http.Cookie{
		Name:    "id",
		Value:   "phoenix",
		Expires: expiration,
	}
	http.SetCookie(w, &session)
	http.Redirect(w, r, "/", http.StatusFound)
}
func logoutPage(w http.ResponseWriter, r *http.Request) {
session,err:=r.Cookie("id")
if err==http.ErrNoCookie{
	http.Redirect(w,r,"/",http.StatusFound)
	return
}
session.Expires=time.Now().AddDate(0,0,-1)
http.SetCookie(w,session)
http.Redirect(w,r,"/",http.StatusFound)
}
