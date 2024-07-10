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
	//command
	Save(ctx context.Context, cake domain.Cake) (*domain.Cake, error)
	Update(ctx context.Context, cake domain.Cake) (*domain.Cake, error)

	//Query
	FindById(ctx context.Context, id int) (*domain.Cake, error)
	List(ctx context.Context) ([]domain.Cake, error)
}

func NewRepository(logger *logger.Logger, DB *sql.DB) Repository {
	return &RepositoryImpl{
		Logger: logger,
		DB:     DB,
	}
}
