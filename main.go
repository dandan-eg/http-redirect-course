package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	//serving index.html
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/seeother", seeOther)
	http.HandleFunc("/movedPermanently", movedPermanently)
	http.HandleFunc("/temporary", temporary)
	http.HandleFunc("/hey", hey)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func seeOther(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Location", "/hey")
	//w.WriteHeader(http.StatusSeeOther)

	fmt.Println("seeOther's method :", req.Method)

	http.Redirect(w, req, "/hey", http.StatusSeeOther)
}

func movedPermanently(w http.ResponseWriter, req *http.Request) {
	fmt.Println("movedPermanently's method :", req.Method)

	http.Redirect(w, req, "/hey", http.StatusMovedPermanently)

}

func temporary(w http.ResponseWriter, req *http.Request) {
	fmt.Println("temporary's method : ", req.Method)

	http.Redirect(w, req, "hey", http.StatusTemporaryRedirect)

}

func hey(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hey's method :", req.Method)

	io.WriteString(w, req.Method)
}
