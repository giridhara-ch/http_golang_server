package main

import (
	"encoding/json"
	"errors"
	"giridhara-ch/http_golang_server/internal/database"
	"net/http"
	"time"
)

const addr = "localhost:8080"

func main() {
	servMux := http.NewServeMux()
	servMux.HandleFunc("/", handler)
	servMux.HandleFunc("/err", testErrHandler)
	srv := http.Server{
		Handler:      servMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	srv.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, database.User{Email: "test@email.com", CreatedAt: time.Now()})
}

func testErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, errors.New("sample error"))
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, 500, err)
	} else {
		w.Write(response)
	}
}

type errorBody struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	respondWithJson(w, code, errorBody{Error: err.Error()})
}
