package repo

import (
	"bez/internal/entity"
	"bez/pkg/postgres"
	"context"
	"fmt"
)

type ServiceRepo struct {
	*postgres.Postgres
}

func NewServiceRepo(pg *postgres.Postgres) *ServiceRepo {
	return &ServiceRepo{pg}
}

func (s *ServiceRepo) GetAllServices(ctx context.Context) ([]entity.ServiceAccount, error) {
	query := `SELECT * FROM service_account`

	rows, err := s.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	var services []entity.ServiceAccount

	for rows.Next() {
		var service entity.ServiceAccount
		err = rows.Scan(
			&service.ID,
			&service.Email,
			&service.RefreshToken,
			&service.StorageLimit,
			&service.StorageUsage)
		if err != nil {
			return nil, fmt.Errorf("cannot parse service entity: %v", err)
		}
		services = append(services, service)
	}
	return services, err
}
