package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/huandu/go-sqlbuilder"

	_ "github.com/lib/pq"
)

var db *sql.DB
var dbDriver string
var psqlConnectionURL string
var dbflavor sqlbuilder.Flavor
var authTable string
var badWords []string

func init() {
	var err error
	badWords, err = getBadWords()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	authTable = os.Getenv("AUTH_TABLE")
	if authTable == "" {
		authTable = "users"
	}
	dbDriver = os.Getenv("API_DB_DRIVER")
	switch dbDriver {
	case "postgres":
		dbflavor = sqlbuilder.PostgreSQL
	case "mysql":
		dbflavor = sqlbuilder.MySQL
	}
	psqlConnectionURL = strings.Join([]string{"postgres://",
		os.Getenv("API_DB_USER"), ":",
		os.Getenv("API_DB_PASS"),
		"@", os.Getenv("API_DB_HOST"), ":",
		os.Getenv("API_DB_PORT"), "/",
		os.Getenv("API_DB_NAME"),
		"?sslmode=", os.Getenv("API_PQ_SSLMODE")}, "")
	log.Printf("Connection to DB via %s", psqlConnectionURL)
	switch dbDriver {
	case "postgres":
		db, err = sql.Open(dbDriver, psqlConnectionURL)
		if err := db.Ping(); err != nil {
			log.Println(err.Error())
		}
		log.Println(fmt.Sprintf("%+v", db.Stats()))
	case "mysql":

	}

	if err != nil {
		log.Println(err.Error())
	}
}

func cert() *tls.Config {
	cert, err := tls.LoadX509KeyPair(os.Getenv("SHOP_TLSCERT"), os.Getenv("SHOP_TLSKEY"))
	if err != nil {
		log.Println("LoadX509KeyPair:", err.Error())
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}
}

func listen() *http.Server {
	mux := http.ServeMux{}
	server := http.Server{Addr: os.Getenv("SHOP_BIND_ADDR"), TLSConfig: cert()}
	customRoutes(&mux)
	shopPaths(&mux)
	server.Handler = &mux
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()
	return &server
}

func main() {

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	if err := runMigrations(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	s := listen()
	log.Println(fmt.Sprintf("Starting server on %s with %s database.", s.Addr, dbDriver))
	<-sigint
	log.Println("Shutting down")
	s.Shutdown(context.Background())
}
