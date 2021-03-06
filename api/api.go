package api

import (
	"net/http"

	"github.com/OleksiiPyvovar/companies-crud/pkg/app"
	"github.com/OleksiiPyvovar/companies-crud/pkg/auth"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ServerAddr string
	APISecret  string

	DBUser    string
	DBPwd     string
	DBTCPHost string
	DBPort    string
	DBName    string

	DefaultListLimit int
}

type API struct {
	router *httprouter.Router
	server *http.Server

	Logger           *log.Logger
	CompaniesService app.Service

	Authorizer auth.Interface
	Config     *Config
}

func (a *API) init() {
	a.router.GET("/api/v1/companies", a.CompanyListHandler)
	a.router.GET("/api/v1/companies/:id", a.CompanyGetByIDHandler)
	a.router.DELETE("/api/v1/companies/:id", a.MiddlewareAuthentication(a.CompanyDeleteByIDHandler))
	a.router.POST("/api/v1/companies/", a.MiddlewareAuthentication(a.CompanyCreateHandler))
	a.router.PUT("/api/v1/companies/", a.MiddlewareAuthentication(a.CompanyUpdateHandler))

	a.router.ServeFiles("/docs/*filepath", http.Dir("static/swaggerui"))
}

func (a *API) Run() {
	a.init()
	a.Logger.Fatal(a.server.ListenAndServe())
}

func NewAPI(conf *Config, cs app.Service) *API {
	router := httprouter.New()
	server := &http.Server{Addr: conf.ServerAddr, Handler: router}

	return &API{
		router:     router,
		server:     server,
		Config:     conf,
		Logger:     log.New(),
		Authorizer: auth.New(conf.APISecret),

		CompaniesService: cs,
	}
}
