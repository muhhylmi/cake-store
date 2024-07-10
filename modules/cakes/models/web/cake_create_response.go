package web

import "cake-store/modules/cakes/models/domain"

type CakeCreateResponse struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
}

func ToModelResponse(c *domain.Cake) CakeCreateResponse {
	return CakeCreateResponse{
		Id:          c.Id,
		Title:       c.Title,
		Description: c.Description,
		Rating:      c.Rating,
		Image:       c.Image,
	}
}
