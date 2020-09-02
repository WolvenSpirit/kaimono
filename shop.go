package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/huandu/go-sqlbuilder"
)

// internal paths
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
	switch r.URL.Path {
	case "/api":
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err.Error())
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
		request := requestBody{}
		if err = json.Unmarshal(b, &request); err != nil {
			log.Println("Unmarshal error:", err.Error())
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
		switch request.Operation {
		case "select":
			log.Printf("%+v", request)
			sb := sqlbuilder.NewSelectBuilder()
			sb.Select(request.Want...)
			sb.From(request.Table...)
			if len(request.Where) != 0 {
				sb.Where(request.Where...)
			}
			if len(request.Join) != 0 {
				for _, v := range request.Join {
					sb.Join(v.Table, v.On)
				}
			}
			query, args := sb.Build()
			log.Print(query)
			rows, err := db.Query(query, args...)
			if err != nil {
				log.Println(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
			var responseBody []interface{}
			for rows.Next() {
				var resp []interface{} = make([]interface{}, len(request.Want))
				switch len(request.Want) {
				case 1:
					if err := rows.Scan(&resp[0]); err != nil {
						log.Println(err.Error())
					}
				case 2:
					if err := rows.Scan(&resp[0], &resp[1]); err != nil {
						log.Println(err.Error())
					}
				case 3:
					if err := rows.Scan(&resp[0], &resp[1], &resp[2]); err != nil {
						log.Println(err.Error())
					}
				case 4:
					if err := rows.Scan(&resp[0], &resp[1], &resp[2], &resp[3]); err != nil {
						log.Println(err.Error())
					}
				case 5:
					if err := rows.Scan(&resp[0], &resp[1], &resp[2], &resp[3], &resp[4], &resp[5]); err != nil {
						log.Println(err.Error())
					}
				}
				log.Println(resp)
				response := make(map[string]interface{})
				for k, v := range request.Want {
					response[v] = resp[k]
				}
				log.Printf("%+v", response)
				responseBody = append(responseBody, response)
			}
			b, err := json.Marshal(&responseBody)
			if err != nil {
				log.Println(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
			wr.Write(b)
		case "insert":
			ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
			ib.InsertInto(request.Table[0])
			ib.Cols(request.Insert.Cols...)
			ib.Values(request.Insert.Values...)
			query, args := ib.Build()
			log.Println(query)
			_, err := db.Exec(query, args...)
			if err != nil {
				log.Println(err.Error())
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
			if _, err := db.Exec(query, args...); err != nil {
				log.Println(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
		case "delete":
			delb := sqlbuilder.NewDeleteBuilder()
			delb.DeleteFrom(request.Table[0])
			delb.Where(request.Where...)
			query, args := delb.Build()
			if _, err := db.Exec(query, args...); err != nil {
				log.Println(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
		case "create":
			cb := sqlbuilder.NewCreateTableBuilder()
			cb.SetFlavor(sqlbuilder.PostgreSQL)
			cb = cb.CreateTable(request.Table[0]).IfNotExists()
			for _, v := range request.Define {
				cb.Define(v)
			}
			query, _ := cb.Build()
			log.Println(query)
			if _, err := db.Exec(query); err != nil {
				log.Println(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	default:
		log.Println(errors.New("unknown operation").Error())
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}
}
