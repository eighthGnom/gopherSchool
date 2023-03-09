package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *APIServer) Start() error {
	api.configRouter()
	err := api.configureLogger()
	if err != nil {
		return err
	}
	api.logger.Infof("Starting server at port %s, with loggin level %s", api.config.BindAddr, api.config.LoggerLevel)
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
