package repositories

import (
	"cake-store/modules/cakes/models/domain"
	"cake-store/utils/database"
	"cake-store/utils/wrapper"
	"context"
	"errors"
)

func (r *RepositoryImpl) FindById(ctx context.Context, cakeId int) (*domain.Cake, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	SQL := "select id, title, description, rating, image from cakes where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, cakeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cake := domain.Cake{}
	if rows.Next() {
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image)
		if err != nil {
			return nil, err
		}
		return &cake, nil
	} else {
		return &cake, errors.New("category is not found")
	}
}

func (r *RepositoryImpl) List(ctx context.Context) ([]domain.Cake, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	SQL := "select id, title, description, rating, image from cakes"
	rows, err := tx.QueryContext(ctx, SQL)
	wrapper.PanicIfError(err)
	defer rows.Close()

	var cakes []domain.Cake
	for rows.Next() {
		cake := domain.Cake{}
		rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image)
		cakes = append(cakes, cake)
	}

	return cakes, nil
}
