package repositories

import (
	"cake-store/modules/cakes/models/domain"
	"cake-store/utils/database"
	"cake-store/utils/wrapper"
	"context"
)

func (r *RepositoryImpl) Save(ctx context.Context, cake domain.Cake) (*domain.Cake, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	SQL := "insert into cakes(title,description,rating,image) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, cake.Title, cake.Description, cake.Rating, cake.Image)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	cake.Id = int(id)

	return &cake, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, cake domain.Cake) (*domain.Cake, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	SQL := "update cakes set title = ?, description = ?, rating = ?, image = ?  where id = ?"
	_, err = tx.ExecContext(ctx, SQL, cake.Title, cake.Description, cake.Rating, cake.Image, cake.Id)
	wrapper.PanicIfError(err)

	return &cake, nil
}
