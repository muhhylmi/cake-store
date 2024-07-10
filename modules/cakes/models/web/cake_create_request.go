package web

import "cake-store/modules/cakes/models/domain"

type CakeCreateRequest struct {
	Title       string  `validate:"required,min=1,max=255" json:"title"`
	Description string  `validate:"omitempty,min=1,max=255" json:"description"`
	Rating      float64 `validate:"required" json:"rating"`
	Image       string  `validate:"omitempty" json:"image"`
}

func (r *CakeCreateRequest) ToModel() *domain.Cake {
	return &domain.Cake{
		Title:       r.Title,
		Description: r.Description,
		Rating:      r.Rating,
		Image:       r.Image,
	}
}
