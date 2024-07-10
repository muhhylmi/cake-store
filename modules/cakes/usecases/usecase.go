package usecases

import (
	"cake-store/modules/cakes/models/web"
	"cake-store/modules/cakes/repositories"
	"cake-store/utils/config"
	"cake-store/utils/logger"
	"context"
	"database/sql"
)

const contextName = "modules.cake.usecase"

type UsecaseImpl struct {
	Logger     *logger.Logger
	Repository repositories.Repository
	Config     *config.Configurations
	DB         *sql.DB
}

type Usecases interface {
	Create(ctx context.Context, payload *web.CakeCreateRequest) web.CakeCreateResponse
}

func NewUsecase(config *config.Configurations, logger *logger.Logger, repository repositories.Repository) Usecases {
	return &UsecaseImpl{
		Logger:     logger,
		Repository: repository,
		Config:     config,
	}
}