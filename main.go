package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

var username string
var password string
var realm string

func main() {
	username = os.Getenv("FS_USERNAME")
	password = os.Getenv("FS_PASSWORD")
	realm = os.Getenv("FS_REALM")

	if len(username) < 3 || len(password) < 8 || len(realm) < 4 {
		log.Fatal(errors.New("invalid username/pass length"))
	}

	fmt.Println("Starting server on :9812")
	http.Handle("/", http.StripPrefix("/", authentication(http.FileServer(http.Dir("www")))))
	err := http.ListenAndServe(":9812", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing authentication")
		w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
		u, p, ok := r.BasicAuth()
		if !ok {
			log.Println("Error parsing basic auth")
			w.WriteHeader(401)
			return
		}
		if u != username {
			fmt.Printf("Username provided is incorrect: %s\n", u)
			w.WriteHeader(401)
			return
		}
		if p != password {
			fmt.Printf("Password provided is incorrect: %s\n", u)
			w.WriteHeader(401)
			return
		}

		next.ServeHTTP(w, r)
	})
}
