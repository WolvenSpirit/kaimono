package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/huandu/go-sqlbuilder"
)

func shopPaths(mux *http.ServeMux) {
	mux.HandleFunc("/", shopMain)
}

type join struct {
	Table string
	On    string
}

type requestBody struct {
	Operation string
	Table     []string
	Want      []string
	Where     []string
	Join      []join
	Define    []string
}

func shopMain(wr http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/api":
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			l.Error(err.Error())
		}
		request := requestBody{}
		if err = json.Unmarshal(b, &request); err != nil {
			l.Error(err.Error())
		}
		switch request.Operation {
		case "select":
			sb := sqlbuilder.NewSelectBuilder()
			sb.Select(request.Want...)
			sb.From(request.Table...)
			sb.Where(request.Where...)
			if len(request.Join) != 0 {
				for _, v := range request.Join {
					sb.Join(v.Table, v.On)
				}
			}
			query, args := sb.Build()
			rows, err := db.Query(query, args...)
			if err != nil {
				l.Error(err.Error())
			}
			var responseBody []interface{}
			for rows.Next() {
				var resp []interface{} = make([]interface{}, len(request.Want))
				rows.Scan(resp...)
				response := make(map[string]interface{})
				for k, v := range request.Want {
					response[v] = resp[k]
				}
				responseBody = append(responseBody, response)
			}

		case "insert":

		case "update":

		case "delete":

		case "create":
			cb := sqlbuilder.CreateTableBuilder{}
			cb.CreateTable(request.Table[0])
			cb.IfNotExists()
			cb.Define(request.Define...)
			query := cb.String()
		}
	default:

	}
}
