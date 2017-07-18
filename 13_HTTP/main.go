package main

import (
	"net/http"
	"io"
	"fmt"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res, fmt.Sprintf(
		`<DOCTYPE html>
	<html>
	  <head>
		  <title>Hello World</title>
	  </head>
	  <body>
		  Hello World! Request URI: %s
	  </body>
	</html>`, req.RequestURI),
	)
}
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}
