package usecase

import (
	"bez/internal/entity"
	"context"
)

type ServiceUseCase struct {
	repo ServiceRp
}

func NewServiceUseCase(repo ServiceRp) *ServiceUseCase {
	return &ServiceUseCase{repo: repo}
}

func (s *ServiceUseCase) GetServices(ctx context.Context) ([]entity.ServiceAccount, error) {
	return s.repo.GetAllServices(ctx)
}
