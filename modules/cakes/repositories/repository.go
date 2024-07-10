package repositories

import (
	"cake-store/modules/cakes/models/domain"
	"cake-store/utils/logger"
	"context"
	"database/sql"
)

type RepositoryImpl struct {
	Logger *logger.Logger
	DB     *sql.DB
}

type Repository interface {
	Save(ctx context.Context, category domain.Cake) (*domain.Cake, error)
}

func NewRepository(logger *logger.Logger, DB *sql.DB) Repository {
	return &RepositoryImpl{
		Logger: logger,
		DB:     DB,
	}
}
