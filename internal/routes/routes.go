package routes

import (
	"log"
	"net/http"
	"os"
	"pos-acen/internal/helper"
	"pos-acen/internal/modules/users/handler/rest"
	"pos-acen/pkg/config"
	"time"

	"github.com/gorilla/mux"
)

type Routes struct {
	Router *mux.Router
	User   *rest.UserHandler
}

func (r *Routes) Run(port string) {
	r.SetupRouter()

	log.Printf("[HTTP SRV] clients on localhost port :%s", port)
	srv := &http.Server{
		Handler:      r.Router,
		Addr:         "localhost:" + port,
		WriteTimeout: config.WriteTimeout() * time.Second,
		ReadTimeout:  config.ReadTimeout() * time.Second,
	}

	log.Panic(srv.ListenAndServe())
}

func (r *Routes) SetupRouter() {
	r.Router = mux.NewRouter()
	r.Router.Use(helper.EnabledCors, helper.LoggerMiddleware())

	r.SetupBaseURL()
	r.SetupUser()
}

func (r *Routes) SetupBaseURL() {
	baseURL := os.Getenv("BASE_URL_PATH")
	if baseURL != "" && baseURL != "/" {
		r.Router.PathPrefix(baseURL).HandlerFunc(helper.URLRewriter(r.Router, baseURL))
	}
}

func (r *Routes) SetupUser() {
	userRoutes := r.Router.PathPrefix("/users").Subrouter()

	userRoutes.HandleFunc("/signup", r.User.Register).Methods(http.MethodPost, http.MethodOptions)
}
