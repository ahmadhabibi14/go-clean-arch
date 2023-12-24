package service

import (
	"context"
	"errors"
	"go-clean-arch/internal/todo/model"
	"go-clean-arch/pkg/db"
	"go-clean-arch/pkg/erru"
)

func (s Service) Get(ctx context.Context, id int) (model.ToDo, error) {
	todo, err := s.repo.Find(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.ToDo{}, erru.ErrArgument{Wrapped: errors.New("Todo object not found")}
	default:
		return model.ToDo{}, err
	}

	return todo, nil
}
