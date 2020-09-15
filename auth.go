package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/WolvenSpirit/kaimono/protobuf/kaimono"

	"golang.org/x/crypto/bcrypt"
)

func (api apiServiceServer) Handshake(ctx context.Context, rq *kaimono.HandshakeRequest) (*kaimono.HandshakeResponse, error) {
	row := db.QueryRow("select count(*) from $1", authTable)
	var count int
	row.Scan(&count)
	if count >= 1 {
		return &kaimono.HandshakeResponse{ClaimMode: false}, nil
	}
	return &kaimono.HandshakeResponse{ClaimMode: true}, nil
}

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
	if err := row.Scan(&hash); err != nil {
		log.Println(err.Error())
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err.Error())
	}
}
func (api apiServiceServer) Register(ctx context.Context, rq *kaimono.RegistrationRequest) (*kaimono.AuthResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rq.GetPassword()), 11)
	if err != nil {
		log.Println(err.Error())
	}
	ib := dbflavor.NewInsertBuilder()
	ib.Cols("email", "password")
	ib.Values(sanitize(rq.GetEmail()), string(hash))
	q, args := ib.Build()
	if _, err := db.Exec(q, args...); err != nil {
		log.Println(err.Error())
		return &kaimono.AuthResponse{Code: 500}, status.Error(codes.Internal, "500")
	}
	return &kaimono.AuthResponse{Code: 200}, nil
}
func (api apiServiceServer) Login(ctx context.Context, rq *kaimono.LoginRequest) (*kaimono.AuthResponse, error) {
	email := rq.GetEmail()
	password := rq.GetPassword()
	sb := dbflavor.NewSelectBuilder()
	sb.Select("password")
	sb.From(authTable)
	sb.Where(fmt.Sprintf("email=%s", email))
	q, args := sb.Build()
	row := db.QueryRow(q, args...)
	var hash string
	if err := row.Scan(&hash); err != nil {
		log.Println(err.Error())
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err.Error())
		return &kaimono.AuthResponse{Code: 500}, status.Error(codes.Internal, "500")
	}
	return &kaimono.AuthResponse{Code: 200}, nil
}
