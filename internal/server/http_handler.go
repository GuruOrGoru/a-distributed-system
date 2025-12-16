package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpApp struct {
	Log *Log
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceRespond struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

func (app *httpApp) handleProduce() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ProduceRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("Invalid request type, can't decode %v to %+v: %v", r.Body, req, err.Error()), http.StatusBadRequest)
			return
		}
		offset, err := app.Log.Append(req.Record)
		if err != nil {
			http.Error(w, fmt.Sprintf("Opps! something bad happened while appending records to log: %v", err.Error()), http.StatusInternalServerError)
			return
		}

		response := &ProduceRespond{
			Offset: offset,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %v", err.Error()), http.StatusInternalServerError)
			return
		}
	}
}

func (app *httpApp) handleConsume() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ConsumeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("Invalid request type, can't decode %v to %+v: %v", r.Body, req, err.Error()), http.StatusBadRequest)
			return
		}
		rec, err := app.Log.Read(req.Offset)
		if err == ErrOffsetNotFound {
			http.Error(w, fmt.Sprintf("Opps! something bad happened while reading from log records: %v", err.Error()), http.StatusInternalServerError)
			return
		}

		response := &ConsumeResponse{
			Record: rec,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %v", err.Error()), http.StatusInternalServerError)
			return
		}
	}
}

func newHttpServer() *httpApp {
	return &httpApp{
		Log: NewLog(),
	}
}
