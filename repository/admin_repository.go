package repository

import (
	"context"
	"database/sql"
	"ecourse-app/helper"
	"ecourse-app/model/domain"
	"errors"
)

type AdminRepository interface {
	FindByEmail(ctx context.Context, transaction *sql.Tx, email string) (domain.Admin, error)
	Insert(ctx context.Context, transaction *sql.Tx, admin domain.Admin) domain.Admin
}

type AdminRepositoryImpl struct{}

func NewAdminRepository() AdminRepository {
	return &AdminRepositoryImpl{}
}
func (repository *AdminRepositoryImpl) FindByEmail(ctx context.Context, transaction *sql.Tx, email string) (domain.Admin, error) {
	rows, err := transaction.QueryContext(ctx, "SELECT id,name,email,password FROM admin WHERE email=?", email)
	helper.PanicError(err)
	defer rows.Close()

	var admin domain.Admin
	if rows.Next() {
		err := rows.Scan(&admin.Id, &admin.Name, &admin.Email, &admin.Password)
		helper.PanicError(err)

		return admin, nil
	} else {
		return admin, errors.New("admin not found")
	}
}

func (repository *AdminRepositoryImpl) Insert(ctx context.Context, transaction *sql.Tx, admin domain.Admin) domain.Admin {

	result, err := transaction.ExecContext(ctx, "INSERT INTO admin(name,email,password) VALUES(?,?,?)", admin.Name, admin.Email, admin.Password)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	admin.Id = id

	return admin

}
