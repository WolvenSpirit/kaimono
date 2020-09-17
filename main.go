package main

import (
	//"google.golang.org/genproto/googleapis/api"

	"google.golang.org/grpc"

	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"

	"github.com/WolvenSpirit/kaimono/protobuf/kaimono"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/lib/pq"
)

type apiServiceServer struct {
	kaimono.ApiServiceServer
}

var db *sql.DB
var dbDriver string
var psqlConnectionURL string
var dbflavor sqlbuilder.Flavor
var authTable string
var badWords []string

func init() {
	parseConfig()
	var err error
	badWords, err = getBadWords()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func dbInit() {
	var err error
	authTable = env.DBDriver
	if authTable == "" {
		authTable = yml.DBDriver
	}
	dbDriver = env.DBDriver
	if dbDriver == "" {
		dbDriver = yml.DBDriver
	}
	switch dbDriver {
	case "postgres":
		dbflavor = sqlbuilder.PostgreSQL
	case "mysql":
		dbflavor = sqlbuilder.MySQL
	}
	if envParseErr == nil {
		psqlConnectionURL = strings.Join([]string{dbDriver, "://", env.DBUser, ":", env.DBPassword, "@", env.DBHost, ":", env.DBPort, "/", env.DBName, "?sslmode=", env.DBPqSslMode}, "")
	}
	if ymlParseErr == nil {
		psqlConnectionURL = strings.Join([]string{dbDriver, "://", yml.DBUser, ":", yml.DBPassword, "@", yml.DBHost, ":", yml.DBPort, "/", yml.DBName, "?sslmode=", yml.DBPqSslMode}, "")
	}

	log.Printf("Connection to DB via %s", psqlConnectionURL)
	switch dbDriver {
	case "postgres":
		db, err = sql.Open(dbDriver, psqlConnectionURL)
		if err != nil {
			log.Println("sql.Open", err.Error())
		}
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

func listenGRPC() *grpc.Server {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("GRPC_ADDR"), os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	apiServer := apiServiceServer{}
	s := grpc.NewServer()

	kaimono.RegisterApiServiceServer(s, apiServer)
	go func() {
		s.Serve(listener)
	}()
	return s
}

func main() {
	dbInit()
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	if err := runMigrations(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	grpcServer := listenGRPC()
	log.Println(fmt.Sprintf("Starting server with %s database.", dbDriver))
	<-sigint
	log.Println("Shutting down")
	grpcServer.GracefulStop()
	db.Close()
}
