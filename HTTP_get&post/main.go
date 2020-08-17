package main

import (
	"fmt"
	"net/http"
)
var calcTempl=`
<html>
   <body>
      <form action="/" method="post">
        <input type="text" name="email" placeholder="email">
        <input type="password" name="pass" placeholder="password">
        <input type="submit" value="Login">
      </form>
   </body>
</html>
`
func main() {
	http.HandleFunc("/getPage",getPage)
	http.HandleFunc("/",postPage)

	http.ListenAndServe(":1309",nil)
}

func getPage(w http.ResponseWriter,r *http.Request){
	arg1:=r.URL.Query().Get("arg1")
	if arg1!=""{
		fmt.Fprintln(w,"first argument:",arg1)

	}
	arg2:=r.FormValue("arg2")
	if arg2!=""{
		fmt.Fprintln(w,"second argument:",arg2)
	}
}
func postPage(w http.ResponseWriter,r *http.Request){
	if r.Method!=http.MethodPost{
		w.Write([]byte(calcTempl))
		return
	}
	r.ParseForm()
	val1:=r.Form["email"][0]
	val2:=r.Form["pass"][0]


	fmt.Fprintln(w,"email:",val1,"\npassword:",val2)
}
