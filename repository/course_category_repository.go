package repository

import (
	"context"
	"database/sql"
	"ecourse-app/model/domain"
)

type CourseCategoryRepository interface {
	FindById(context context.Context, transaction *sql.Tx, id int64, courses bool)
	FindAll(context context.Context, transaction *sql.Tx, courses bool)
	Insert(context context.Context, transaction *sql.Tx, courseCategory domain.CourseCategory)
	Update(context context.Context, transaction *sql.Tx, courseCategory domain.CourseCategory)
	Delete(context context.Context, transaction *sql.Tx, id int64)
}
