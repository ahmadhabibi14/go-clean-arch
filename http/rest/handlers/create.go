package handlers

import (
	"go-clean-arch/internal/todo/model"
	todoService "go-clean-arch/internal/todo/service"
	"net/http"
)

func (s service) Create() http.HandlerFunc {
	type request struct {
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
	}

	type response struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := s.decode(r, &req)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		id, err := s.todoService.Create(r.Context(), todoService.CreateParams{
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
		})
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{ID: id}, http.StatusOK)
	}
}
