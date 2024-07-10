package web

import "cake-store/modules/cakes/models/domain"

func ToModelResponse(c *domain.Cake) CakeResponse {
	return CakeResponse{
		Id:          c.Id,
		Title:       c.Title,
		Description: c.Description,
		Rating:      c.Rating,
		Image:       c.Image,
	}
}
