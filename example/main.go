package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func handler(w http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("template/index.html")
	tmpl.Execute(w, nil)
}

func main() {
	var port string

	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	fmt.Println("Starting app on port: " + port)

	http.ListenAndServe(":"+port, nil)
}
