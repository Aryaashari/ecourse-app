package repository

import (
	"context"
	"database/sql"
	"ecourse-app/helper"
	"ecourse-app/model/domain"
	"errors"
)

type UserRepository interface {
	FindAll(ctx context.Context, transaction *sql.Tx) []domain.User
	FindById(ctx context.Context, transaction *sql.Tx, id int64) (domain.User, error)
	FindByEmail(ctx context.Context, transaction *sql.Tx, email string) (domain.User, error)
	Insert(ctx context.Context, transaction *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, transaction *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, transaction *sql.Tx, id int64)
}

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, transaction *sql.Tx) []domain.User {
	rows, err := transaction.QueryContext(ctx, "SELECT id,name,email,password FROM users")
	helper.PanicError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		helper.PanicError(err)

		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, transaction *sql.Tx, id int64) (domain.User, error) {
	rows, err := transaction.QueryContext(ctx, "SELECT id,name,email,password FROM users WHERE id=?", id)
	helper.PanicError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		helper.PanicError(err)
	} else {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, transaction *sql.Tx, email string) (domain.User, error) {
	rows, err := transaction.QueryContext(ctx, "SELECT id,name,email,password FROM users WHERE email=?", email)
	helper.PanicError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		helper.PanicError(err)
	} else {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, transaction *sql.Tx, user domain.User) domain.User {
	result, err := transaction.ExecContext(ctx, "INSERT INTO users(name,email,password) VALUES(?,?,?)", user.Name, user.Email, user.Password)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	user.Id = id

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, transaction *sql.Tx, user domain.User) domain.User {
	_, err := transaction.ExecContext(ctx, "UPDATE users SET name=? WHERE id=?", user.Name, user.Id)
	helper.PanicError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, transaction *sql.Tx, id int64) {
	_, err := transaction.ExecContext(ctx, "DELETE FROM users WHERE id=?", id)
	helper.PanicError(err)
}
