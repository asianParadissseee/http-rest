package api_server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type API_SERVER struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// new
func New(config *Config) *API_SERVER {
	return &API_SERVER{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// start
func (s *API_SERVER) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Staring api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *API_SERVER) configureLogger() error {
	lever, err := logrus.ParseLevel(s.config.LogLever)
	if err != nil {
		return err
	}
	s.logger.SetLevel(lever)

	return nil
}

func (s *API_SERVER) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *API_SERVER) handleHello() http.HandlerFunc {

	type request struct {
		name string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
