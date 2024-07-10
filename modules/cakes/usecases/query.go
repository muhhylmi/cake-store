package usecases

import (
	"cake-store/modules/cakes/models/web"
	"cake-store/utils/wrapper"
	"context"
)

func (u *UsecaseImpl) FindById(ctx context.Context, req *web.CakeGetRequest) web.CakeResponse {
	l := u.Logger.LogWithContext(contextName, "FindById")

	cake, err := u.Repository.FindById(ctx, req.CakeId)
	if err != nil {
		l.Error(err)
		panic(wrapper.NewNotFoundError(err.Error()))
	}

	return web.ToModelResponse(cake)
}
