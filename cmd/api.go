package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/julianh99/goexpenses/model"
	"github.com/julianh99/goexpenses/store"
	"github.com/julianh99/goexpenses/util"
)

type ApiServer struct {
	host  string
	port  int
	store store.Store
}

func NewApiServer(host string, port int, s store.Store) *ApiServer {

	s.Initialize()

	return &ApiServer{
		host:  host,
		port:  port,
		store: s,
	}
}

func (server *ApiServer) Run() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	r.HandleFunc("/", server.Index)
	r.HandleFunc("/movement/add", server.NewMovement)

	srv := &http.Server{
		Handler: r,
		Addr:    server.host + ":" + strconv.Itoa(server.port),
	}
	fmt.Println("Server running")

	log.Fatal(srv.ListenAndServe())
}

func (server ApiServer) Index(w http.ResponseWriter, r *http.Request) {

	movements, err := server.store.GetMovements()

	if err != nil {
		log.Fatal(err)
	}

	data := struct{ Movements []*model.Movement }{Movements: movements}

	util.RenderTemplate(util.TemplateOptions{TemplateName: "index"}, data, w)

}

func (server ApiServer) NewMovement(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	value, _ := strconv.Atoi(r.FormValue("value"))
	movementType := r.FormValue("type")

	movement := model.Movement{Value: value, MovementType: model.ToMovementType(movementType), Date: time.Now()}

	err := server.store.CreateMovement(movement)

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)

}
