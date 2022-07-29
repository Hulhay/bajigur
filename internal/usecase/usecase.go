package usecase

import (
	"hulhay-mall/internal/repositories"
)

type useCase struct {
	repo repositories.Repositories
}

type UseCase interface {
}

func NewUseCase(r repositories.Repositories) UseCase {
	return &useCase{repo: r}
}
