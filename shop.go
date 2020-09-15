package main

import (
	"google.golang.org/grpc/codes"
	//"google.golang.org/genproto/googleapis/rpc/status"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/WolvenSpirit/kaimono/protobuf/kaimono"
	"github.com/huandu/go-sqlbuilder"
	"google.golang.org/grpc/status"
)

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

// Legacy function non-GRPC
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
			sb := dbflavor.NewSelectBuilder()
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
					if err := rows.Scan(&resp[0], &resp[1], &resp[2], &resp[3], &resp[4]); err != nil {
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
			ib := dbflavor.NewInsertBuilder()
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
			ib := dbflavor.NewUpdateBuilder()
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
			delb := dbflavor.NewDeleteBuilder()
			delb.DeleteFrom(request.Table[0])
			delb.Where(request.Where...)
			query, args := delb.Build()
			if _, err := db.Exec(query, args...); err != nil {
				log.Println(err.Error())
				wr.WriteHeader(http.StatusInternalServerError)
				return
			}
		case "create":
			cb := dbflavor.NewCreateTableBuilder()
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

func (api apiServiceServer) Select(ctx context.Context, rq *kaimono.RequestBody) (*kaimono.ResponseBody, error) {
	sb := dbflavor.NewSelectBuilder()
	sb.Select(rq.Want...)
	sb.From(rq.Table...)
	if len(rq.Where) != 0 {
		sb.Where(rq.Where...)
	}
	if len(rq.Join) != 0 {
		for _, v := range rq.Join {
			sb.Join(v.Table, v.On)
		}
	}
	query, args := sb.Build()
	log.Print(query)
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err.Error())
		return nil, status.Error(codes.Internal, "500")
	}
	var responseBody []interface{}
	for rows.Next() {
		var resp []interface{} = make([]interface{}, len(rq.Want))
		switch len(rq.Want) {
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
		for k, v := range rq.Want {
			response[v] = resp[k]
		}
		log.Printf("%+v", response)
		responseBody = append(responseBody, response)
	}
	bd := &kaimono.ResponseBody{}
	bd.Payload, err = json.Marshal(&responseBody)
	if err != nil {
		log.Println(err.Error())
		return &kaimono.ResponseBody{StatusCode: 500}, status.Error(codes.Internal, "500")
	}
	return bd, nil
}
func (api apiServiceServer) Insert(ctx context.Context, rq *kaimono.RequestBody) (*kaimono.ResponseBody, error) {
	ib := dbflavor.NewInsertBuilder()
	ib.InsertInto(rq.Table[0])
	ib.Cols(rq.Insert.Cols...)
	var values []interface{}
	for _, v := range rq.Insert.Values {
		values = append(values, v)
	}
	ib.Values(values...)
	query, args := ib.Build()
	log.Println(query)
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err.Error())
		return &kaimono.ResponseBody{StatusCode: 500}, status.Error(codes.Internal, "500")
	}
	return &kaimono.ResponseBody{StatusCode: 200}, nil
}
func (api apiServiceServer) Update(ctx context.Context, rq *kaimono.RequestBody) (*kaimono.ResponseBody, error) {
	ib := dbflavor.NewUpdateBuilder()
	ib.Update(rq.Table[0])
	ib.Set(rq.UpdateAssignment)
	ib.Where(rq.Where...)
	query, args := ib.Build()
	if _, err := db.Exec(query, args...); err != nil {
		log.Println(err.Error())
		return &kaimono.ResponseBody{StatusCode: 500}, status.Error(codes.Internal, "500")
	}
	return &kaimono.ResponseBody{StatusCode: 200}, nil
}
func (api apiServiceServer) Delete(ctx context.Context, rq *kaimono.RequestBody) (*kaimono.ResponseBody, error) {
	delb := dbflavor.NewDeleteBuilder()
	delb.DeleteFrom(rq.Table[0])
	delb.Where(rq.Where...)
	query, args := delb.Build()
	if _, err := db.Exec(query, args...); err != nil {
		log.Println(err.Error())
		return &kaimono.ResponseBody{StatusCode: 500}, status.Error(codes.Internal, "500")
	}
	return &kaimono.ResponseBody{StatusCode: 200}, nil
}
func (api apiServiceServer) Create(ctx context.Context, rq *kaimono.RequestBody) (*kaimono.ResponseBody, error) {
	cb := dbflavor.NewCreateTableBuilder()
	cb = cb.CreateTable(rq.Table[0]).IfNotExists()
	for _, v := range rq.Define {
		cb.Define(v)
	}
	query, _ := cb.Build()
	log.Println(query)
	if _, err := db.Exec(query); err != nil {
		log.Println(err.Error())
		return &kaimono.ResponseBody{StatusCode: 500}, status.Error(codes.Internal, "500")
	}
	return &kaimono.ResponseBody{StatusCode: 200}, nil
}

func (api apiServiceServer) PublicResources(ctx context.Context, rq *kaimono.ResourceRequest) (*kaimono.ResourceOverviewResponse, error) {
	selectTableQuery, err := ioutil.ReadFile("sql/select_tables.sql")
	if err != nil {
		log.Println(err.Error())
	}
	selectColumnsQuery, err := ioutil.ReadFile("sql/select_columns.sql")
	if err != nil {
		log.Println(err.Error())
	}
	rows, err := db.Query(string(selectTableQuery))
	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()
	var tables []string
	response := make(map[string][]string)
	for rows.Next() {
		var t string
		rows.Scan(&t)
		tables = append(tables, t)
	}
	for _, t := range tables {
		rows, err := db.Query(string(selectColumnsQuery), t)
		if err != nil {
			log.Println(err.Error())
		}
		for rows.Next() {
			c, _ := rows.Columns()
			columns := make([]interface{}, len(c))
			rows.Scan(columns...)
			var cols []string
			for _, c := range columns {
				cols = append(cols, c.(string))
			}
			response[t] = cols
		}
	}
	b, _ := json.Marshal(&response)
	return &kaimono.ResourceOverviewResponse{Resources: b}, nil
}
