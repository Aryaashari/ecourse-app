package repository

import (
	"context"
	"database/sql"
	"ecourse-app/model/domain"
)

type CourseCategoryRepository interface {
	FindById(ctx context.Context, transaction *sql.Tx, id int64) (domain.CourseCategory, error)
	FindAll(ctx context.Context, transaction *sql.Tx) []domain.CourseCategory
	Insert(ctx context.Context, transaction *sql.Tx, courseCategory domain.CourseCategory) domain.CourseCategory
	Update(ctx context.Context, transaction *sql.Tx, courseCategory domain.CourseCategory) domain.CourseCategory
	Delete(ctx context.Context, transaction *sql.Tx, id int64)
}
