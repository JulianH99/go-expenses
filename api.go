package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	host string
	port int
}

func NewApiServer(host string, port int) *ApiServer {
	return &ApiServer{
		host: host,
		port: port,
	}
}

func (server *ApiServer) Run() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	r.HandleFunc("/", Index)
	r.HandleFunc("/movement/add", NewMovement)

	srv := &http.Server{
		Handler: r,
		Addr:    server.host + ":" + strconv.Itoa(server.port),
	}

	log.Fatal(srv.ListenAndServe())
}

func Index(w http.ResponseWriter, r *http.Request) {

	movements := []Movement{{MovementType: Income, Value: 1000, Date: time.Now()}, {MovementType: Expense, Value: 1000, Date: time.Now()}}

	data := struct{ Movements []Movement }{Movements: movements}

	renderTemplate(TemplateOptions{templateName: "index"}, data, w)

}

func NewMovement(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	value := r.FormValue("value")

	fmt.Println(value)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}
