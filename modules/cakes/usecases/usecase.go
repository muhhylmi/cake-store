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
	//Command
	Create(ctx context.Context, payload *web.CakeCreateRequest) web.CakeResponse
	Update(ctx context.Context, payload *web.CakeUpdateRequest) web.CakeResponse
	Delete(ctx context.Context, cakeId int) bool

	//Query
	FindById(ctx context.Context, payload *web.CakeGetRequest) web.CakeResponse
	List(ctx context.Context, req *web.CakeListRequest) []web.CakeResponse
}

func NewUsecase(logger *logger.Logger, repository repositories.Repository) Usecases {
	return &UsecaseImpl{
		Logger:     logger,
		Repository: repository,
	}
}
