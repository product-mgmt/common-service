package middleware

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/product-mgmt/common-service/constants/messages"
	"github.com/product-mgmt/common-service/storage"
	"github.com/product-mgmt/common-service/utils/commfunc"
)

type Storage struct {
	logger     *logrus.Logger
	mysqlStore storage.MySQLStorage
}

func New(logger *logrus.Logger, mysqlStore storage.MySQLStorage) *Storage {
	return &Storage{
		logger:     logger,
		mysqlStore: mysqlStore,
	}
}

func (s *Storage) PathNotFoundHanler(w http.ResponseWriter, r *http.Request) error {

	resp := commfunc.ApiError{
		Error: fmt.Errorf(messages.PATHNOTFOUND, r.URL.Path).Error(),
	}

	return commfunc.WriteJSON(w, http.StatusNotFound, resp)
}

func (s *Storage) MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf(messages.METHODNOTALLOWED, r.Method)
}

func (s *Storage) RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Request: %s %s", r.Method, r.RequestURI)
		s.logger.Info(msg)
		next.ServeHTTP(w, r)
	})
}
