package repository

import (
	"context"
	"database/sql"
	"ecourse-app/helper"
	"ecourse-app/model/domain"
	"errors"
)

type CourseCategoryRepository interface {
	FindById(ctx context.Context, transaction *sql.Tx, id int64) (domain.CourseCategory, error)
	FindAll(ctx context.Context, transaction *sql.Tx) []domain.CourseCategory
	Insert(ctx context.Context, transaction *sql.Tx, courseCategory domain.CourseCategory) domain.CourseCategory
	Update(ctx context.Context, transaction *sql.Tx, courseCategory domain.CourseCategory) domain.CourseCategory
	Delete(ctx context.Context, transaction *sql.Tx, id int64)
}

type CourseCategoryRepositoryImpl struct{}

func NewCourseCategoryRepository() CourseCategoryRepository {
	return &CourseCategoryRepositoryImpl{}
}

func (repository *CourseCategoryRepositoryImpl) FindById(ctx context.Context, transaction *sql.Tx, id int64) (domain.CourseCategory, error) {

	rows, err := transaction.QueryContext(ctx, "SELECT id,name FROM course_categories WHERE id=?", id)
	helper.PanicError(err)
	defer rows.Close()

	var courseCategory domain.CourseCategory
	if rows.Next() {
		err := rows.Scan(&courseCategory.Id, &courseCategory.Name)
		helper.PanicError(err)
	} else {
		return courseCategory, errors.New("course category not found")
	}

	return courseCategory, nil

}

func (repository *CourseCategoryRepositoryImpl) FindAll(ctx context.Context, transaction *sql.Tx) []domain.CourseCategory {

	rows, err := transaction.QueryContext(ctx, "SELECT id,name FROM course_categories")
	helper.PanicError(err)
	defer rows.Close()

	var courseCategories []domain.CourseCategory
	for rows.Next() {
		var courseCategory domain.CourseCategory
		err := rows.Scan(&courseCategory.Id, &courseCategory.Name)
		helper.PanicError(err)

		courseCategories = append(courseCategories, courseCategory)
	}

	return courseCategories

}

func (repository *CourseCategoryRepositoryImpl) Insert(ctx context.Context, transaction *sql.Tx, courseCategory domain.CourseCategory) domain.CourseCategory {

	result, err := transaction.ExecContext(ctx, "INSERT INTO course_categories(name) VALUES(?)", courseCategory.Name)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	courseCategory.Id = id

	return courseCategory
}

func (repository *CourseCategoryRepositoryImpl) Update(ctx context.Context, transaction *sql.Tx, courseCategory domain.CourseCategory) domain.CourseCategory {

	_, err := transaction.ExecContext(ctx, "UPDATE course_categories SET name=? WHERE id=?", courseCategory.Name, courseCategory.Id)
	helper.PanicError(err)

	return courseCategory

}

func (repository *CourseCategoryRepositoryImpl) Delete(ctx context.Context, transaction *sql.Tx, id int64) {

	_, err := transaction.ExecContext(ctx, "DELETE FROM course_categories WHERE id=?", id)
	helper.PanicError(err)

}
