package main

import (
	"encoding/json"
	"errors"
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
type insert struct {
	Cols   []string
	Values []interface{}
}

type requestBody struct {
	Operation        string
	Table            []string
	Want             []string
	Where            []string
	Join             []join
	Define           []string
	Insert           insert
	UpdateAssignment string
}

func shopMain(wr http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/api":
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			l.Error(err.Error())
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
		request := requestBody{}
		if err = json.Unmarshal(b, &request); err != nil {
			l.Error(err.Error())
			wr.WriteHeader(http.StatusInternalServerError)
			return
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
				wr.WriteHeader(http.StatusInternalServerError)
				return
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
			b, err := json.Marshal(&responseBody)
			if err != nil {
				l.Error(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
			wr.Write(b)
		case "insert":
			ib := sqlbuilder.NewInsertBuilder()
			ib.InsertInto(request.Table[0])
			ib.Cols(request.Insert.Cols...)
			ib.Values(request.Insert.Values...)
			query, args := ib.Build()
			_, err := db.Exec(query, args)
			if err != nil {
				l.Error(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		case "update":
			ib := sqlbuilder.NewUpdateBuilder()
			ib.Update(request.Table[0])
			ib.Set(request.UpdateAssignment)
			ib.Where(request.Where...)
			query, args := ib.Build()
			if _, err := db.Exec(query, args); err != nil {
				l.Error(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
		case "delete":
			delb := sqlbuilder.NewDeleteBuilder()
			delb.DeleteFrom(request.Table[0])
			delb.Where(request.Where...)
			query, args := delb.Build()
			if _, err := db.Exec(query, args); err != nil {
				l.Error(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
		case "create":
			cb := sqlbuilder.NewCreateTableBuilder()
			cb.CreateTable(request.Table[0])
			cb.IfNotExists()
			cb.Define(request.Define...)
			query, args := cb.Build()
			if _, err := db.Exec(query, args); err != nil {
				l.Error(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	default:
		l.Error(errors.New("unknown operation").Error())
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}
}
