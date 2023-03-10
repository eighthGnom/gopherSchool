package apiserver

import (
	"net/http"

	"github.com/eighthGnom/http-rest-api/storage"
	"github.com/sirupsen/logrus"
)

func (api *APIServer) configStorage() error {
	storage := storage.New(api.config.StorageConfig)
	err := storage.Open()
	if err != nil {
		return err
	}
	api.storage = storage
	return nil
}

func (api *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(level)
	return nil
}

func (api *APIServer) configRouter() {
	api.router.HandleFunc("/health-check", api.healthCheck())
}

func (api *APIServer) healthCheck() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Okay"))
	}
}
