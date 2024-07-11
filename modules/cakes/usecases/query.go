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

func (u *UsecaseImpl) List(ctx context.Context, req *web.CakeListRequest) []web.CakeResponse {
	l := u.Logger.LogWithContext(contextName, "List")
	responses := []web.CakeResponse{}

	cakes, err := u.Repository.List(ctx, req)
	if err != nil {
		l.Error(err)
	}

	for _, cake := range cakes {
		responses = append(responses, web.CakeResponse{
			Id:          cake.Id,
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		})
	}
	return responses
}
