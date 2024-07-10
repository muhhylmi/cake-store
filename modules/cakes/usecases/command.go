package usecases

import (
	"cake-store/modules/cakes/models/web"
	"cake-store/utils/wrapper"
	"context"
)

func (usecase *UsecaseImpl) Create(ctx context.Context, payload *web.CakeCreateRequest) web.CakeResponse {
	log := usecase.Logger.LogWithContext(contextName, "Create")
	cakeData := payload.ToModel()

	cake, err := usecase.Repository.Save(ctx, *cakeData)

	if err != nil {
		log.Error("Cannot Create Cart:  " + err.Error())
		panic(wrapper.NewConflictError(err.Error()))
	}
	return web.ToModelResponse(cake)
}

func (usecase *UsecaseImpl) Update(ctx context.Context, request *web.CakeUpdateRequest) web.CakeResponse {
	l := usecase.Logger.LogWithContext(contextName, "Update")

	cake, err := usecase.Repository.FindById(ctx, request.Id)
	if err != nil {
		l.Error(err)
		panic(wrapper.NewNotFoundError(err.Error()))
	}

	cake = request.ToModelUpdate(cake)
	cake, err = usecase.Repository.Update(ctx, *cake)
	if err != nil {
		l.Error(err)
		panic(wrapper.NewConflictError(err.Error()))
	}

	return web.ToModelResponse(cake)
}
