package web

import (
	"cake-store/modules/cakes/models/domain"
	"time"
)

type CakeUpdateRequest struct {
	Id          int      `validate:"required" params:"cakeId"`
	Title       *string  `validate:"omitempty,min=1,max=255" json:"title"`
	Description *string  `validate:"omitempty,min=1,max=255" json:"description"`
	Rating      *float64 `validate:"omitempty" json:"rating"`
	Image       *string  `validate:"omitempty" json:"image"`
}

func (r CakeUpdateRequest) ToModelUpdate(c *domain.Cake) *domain.Cake {
	if r.Title != nil {
		c.Title = *r.Title
	}

	if r.Description != nil {
		c.Description = *r.Description
	}

	if r.Rating != nil {
		c.Rating = *r.Rating
	}

	if r.Image != nil {
		c.Image = *r.Image
	}

	c.UpdatedAt = time.Now()

	return c
}
