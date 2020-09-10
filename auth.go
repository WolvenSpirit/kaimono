package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func getBadWords() ([]string, error) {
	fl, err := os.Open("sanitize_check/words.txt")
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(fl)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(b), ","), nil
}

func sanitize(input string) string {
	for _, v := range badWords {
		if strings.Contains(input, v) {
			return ""
		}
	}
	return strings.TrimSpace(input)
}

func register(wr http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	hash, err := bcrypt.GenerateFromPassword([]byte(sanitize("password")), 11)
	if err != nil {
		log.Println(err.Error())
	}
	ib := dbflavor.NewInsertBuilder()
	ib.Cols("email", "password", "name", "surname")
	ib.Values(sanitize(r.FormValue("email")), string(hash), sanitize(r.FormValue("surname")), sanitize(r.FormValue("name")))
	q, args := ib.Build()
	if _, err := db.Exec(q, args...); err != nil {
		log.Println(err.Error())
	}
}

func login(wr http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := sanitize(r.FormValue("email"))
	password := sanitize(r.FormValue("password"))
	sb := dbflavor.NewSelectBuilder()
	sb.Select("password")
	sb.From(authTable)
	sb.Where(fmt.Sprintf("email=%s", email))
	q, args := sb.Build()
	row := db.QueryRow(q, args...)
	var hash string
	if err := row.Scan(&hash);err!= nil {
		log.Println(err.Error())
	}	
	if err:= bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password);err!= nil {
		log.Println(err.Error())
	}
}
