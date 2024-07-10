package usecases

import (
	"cake-store/modules/cakes/models/web"
	"cake-store/utils/wrapper"
	"context"
)

func (usecase *UsecaseImpl) Create(ctx context.Context, payload *web.CakeCreateRequest) web.CakeCreateResponse {
	log := usecase.Logger.LogWithContext(contextName, "CreateCart")
	cakeData := payload.ToModel()

	cake, err := usecase.Repository.Save(ctx, *cakeData)

	if err != nil {
		log.Error("Cannot Create Cart:  " + err.Error())
		panic(wrapper.NewConflictError(err.Error()))
	}
	return web.ToModelResponse(cake)
}
