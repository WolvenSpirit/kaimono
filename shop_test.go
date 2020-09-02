package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func init() {
	os.Setenv("API_DB_DRIVER", "postgres")
	os.Setenv("API_DB_USER", "postgres")
	os.Setenv("API_DB_PASS", "")
	os.Setenv("API_DB_HOST", "localhost")
	os.Setenv("API_DB_PORT", "5432")
	os.Setenv("API_DB_NAME", "test01")
	os.Setenv("API_PQ_SSLMODE", "disable")

	dbDriver = os.Getenv("API_DB_DRIVER")
	psqlConnectionURL = strings.Join([]string{"postgres://",
		os.Getenv("API_DB_USER"), ":",
		os.Getenv("API_DB_PASS"),
		"@", os.Getenv("API_DB_HOST"), ":",
		os.Getenv("API_DB_PORT"), "/",
		os.Getenv("API_DB_NAME"),
		"?sslmode=", os.Getenv("API_PQ_SSLMODE")}, "")
	log.Printf("Connection to DB via %s", psqlConnectionURL)
	var err error
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
func Test_shopMain(t *testing.T) {

	var recorders = []*httptest.ResponseRecorder{}
	var requests = []*http.Request{}
	var testTable = "test_table_test"
	for i := 0; i < 6; i++ {
		rqb := requestBody{}
		switch i {
		case 0: // create
			rqb.Operation = "create"
			rqb.Table = []string{testTable}
			rqb.Define = []string{"id serial primary key", "column1 text", "column2 integer"}
		case 1: // insert
			rqb.Operation = "insert"
			rqb.Table = []string{testTable}
			rqb.Insert.Cols = []string{"column1"}
			rqb.Insert.Values = append(rqb.Insert.Values, "test")
		case 2: // select
			rqb.Operation = "select"
			rqb.Want = []string{"column1"}
			rqb.Table = []string{testTable}
			rqb.Where = []string{"column1 = 'test'"}
		case 3: // update
			rqb.Operation = "update"
			rqb.Table = []string{testTable}
			rqb.UpdateAssignment = "column1='updateTest'"
			rqb.Where = []string{"column1='test'"}
		case 4: // delete
			rqb.Operation = "delete"
			rqb.Table = []string{testTable}
			//rqb.Where = []string{"id=1"}
		case 5: // unknown

		}
		b, e := json.Marshal(&rqb)
		if e != nil {
			t.Error(e.Error())
		}
		body := bytes.Buffer{}
		body.Write(b)
		rec := httptest.NewRecorder()
		rq, e := http.NewRequest("POST", "/api", &body)
		if e != nil {
			t.Error(e.Error())
		}
		recorders = append(recorders, rec)
		requests = append(requests, rq)
	}
	type args struct {
		wr http.ResponseWriter
		r  *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "create",
			args: args{wr: recorders[0], r: requests[0]},
		},
		{
			name: "insert",
			args: args{wr: recorders[1], r: requests[1]},
		},
		{
			name: "select",
			args: args{wr: recorders[2], r: requests[2]},
		},
		{
			name: "update",
			args: args{wr: recorders[3], r: requests[3]},
		},
		{
			name: "delete",
			args: args{wr: recorders[4], r: requests[4]},
		},
		{
			name: "unknown",
			args: args{wr: recorders[5], r: requests[5]},
		},
	}
	for tk, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shopMain(tt.args.wr, tt.args.r)
			if recorders[tk].Result().StatusCode != 200 {
				t.Error(fmt.Sprintf("%+v", recorders[tk].Result()))
			}
		})
	}
}
