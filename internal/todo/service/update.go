package service

import (
	"context"
	"errors"
	"go-clean-arch/internal/todo/model"
	"go-clean-arch/pkg/erru"

	"github.com/asaskevich/govalidator"
)

type UpdateParams struct {
	ID          int `valid:"required"`
	Name        *string
	Description *string
	Status      *model.Status
}

func (s Service) Update(ctx context.Context, params UpdateParams) error {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return erru.ErrArgument{Wrapped: err}
	}

	todo, err := s.Get(ctx, params.ID)
	if err != nil {
		return err
	}

	if params.Name != nil {
		todo.Name = *params.Name
	}
	if params.Description != nil {
		todo.Description = *params.Description
	}
	if params.Status != nil {
		if !params.Status.IsValid() {
			return erru.ErrArgument{Wrapped: errors.New("Given status invalid")}
		}
		todo.Status = *params.Status
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.repo.Update(ctx, todo)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
