package repositories

import (
	"cake-store/modules/cakes/models/domain"
	"cake-store/utils/database"
	"cake-store/utils/wrapper"
	"context"
)

func (r *RepositoryImpl) Save(ctx context.Context, cake domain.Cake) (*domain.Cake, error) {
	l := r.Logger.LogWithContext(contextName, "Save")

	tx, err := r.DB.Begin()
	if err != nil {
		l.Error(err)
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	SQL := "insert into cakes(title,description,rating,image) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, cake.Title, cake.Description, cake.Rating, cake.Image)
	if err != nil {
		l.Error(err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		l.Error(err)
		return nil, err
	}

	cake.Id = int(id)

	return &cake, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, cake domain.Cake) (*domain.Cake, error) {
	l := r.Logger.LogWithContext(contextName, "Update")

	tx, err := r.DB.Begin()
	if err != nil {
		l.Error(err)
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	SQL := "update cakes set title = ?, description = ?, rating = ?, image = ?, updated_at = ?  where id = ?"
	_, err = tx.ExecContext(ctx, SQL, cake.Title, cake.Description, cake.Rating, cake.Image, cake.UpdatedAt, cake.Id)
	wrapper.PanicIfError(err)

	return &cake, nil
}

func (r *RepositoryImpl) Delete(ctx context.Context, cake domain.Cake) error {
	l := r.Logger.LogWithContext(contextName, "Delete")

	tx, err := r.DB.Begin()
	if err != nil {
		l.Error(err)
		return err
	}
	defer database.CommitOrRollback(tx)

	SQL := "delete from cakes where id = ?"
	_, err = tx.ExecContext(ctx, SQL, cake.Id)
	if err != nil {
		l.Error(err)
		return err
	}
	return nil
}
