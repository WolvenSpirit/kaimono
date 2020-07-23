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

	"github.com/azer/logger"
	_ "github.com/lib/pq"
)

var db *sql.DB
var l = logger.New("shop")
var dbDriver string
var psqlConnectionURL string

func init() {
	dbDriver = os.Getenv("API_DB_DRIVER")
	psqlConnectionURL = strings.Join([]string{"postgres://",
		os.Getenv("API_DB_USER"), ":",
		os.Getenv("API_DB_PASS"),
		"@", os.Getenv("API_DB_HOST"), ":",
		os.Getenv("API_DB_PORT"), "/",
		os.Getenv("API_DB_DBNAME"),
		"?sslmode=", os.Getenv("API_PQ_SSLMODE")}, "")

	var err error
	switch dbDriver {
	case "postgres":
		db, err = sql.Open(dbDriver, psqlConnectionURL)
		if err := db.Ping(); err != nil {
			l.Error(err.Error())
		}
		l.Info(fmt.Sprintf("%+v", db.Stats()))
	case "mysql":

	}

	if err != nil {
		l.Error(err.Error())
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
	s := listen()
	l.Info(fmt.Sprintf("Starting server on %s with %s database.", s.Addr, dbDriver))
	<-sigint
	l.Info("Shutting down")
	s.Shutdown(context.Background())
}
