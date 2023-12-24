package handlers

import (
	todoRepo "go-clean-arch/internal/todo/repository"
	todoService "go-clean-arch/internal/todo/service"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger      *logrus.Logger
	router      *mux.Router
	todoService todoService.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:      lg,
		todoService: todoService.NewService(todoRepo.NewRepository(db)),
	}
}
