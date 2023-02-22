package repository

import (
	"context"
	"database/sql"
	"ecourse-app/helper"
	"ecourse-app/model/domain"
	"errors"
)

type CourseRepository interface {
	FindAll(ctx context.Context, transaction *sql.Tx) ([]domain.Course, []domain.CourseCategory)
	FindById(ctx context.Context, transaction *sql.Tx, id int64) (domain.Course, error)
	FindByCategoryId(ctx context.Context, transaction *sql.Tx, categoryId int64) (domain.Course, error)
	Insert(ctx context.Context, transaction *sql.Tx, course domain.Course) domain.Course
	Update(ctx context.Context, transaction *sql.Tx, course domain.Course) domain.Course
	Delete(ctx context.Context, transaction *sql.Tx, id int64)
}

type CourseRepositoryImpl struct{}

func NewCourseRepository() CourseRepository {
	return &CourseRepositoryImpl{}
}

func (repository *CourseRepositoryImpl) FindAll(ctx context.Context, transaction *sql.Tx) ([]domain.Course, []domain.CourseCategory) {
	rows, err := transaction.QueryContext(ctx, "SELECT courses.id,courses.title,course_category_id,course_categories.name AS course_category_name FROM courses INNER JOIN course_categories ON courses.course_category_id=course_categories.id")
	helper.PanicError(err)
	defer rows.Close()

	var courses []domain.Course
	var courseCategories []domain.CourseCategory
	for rows.Next() {
		var course domain.Course
		var courseCategory domain.CourseCategory
		err := rows.Scan(&course.Id, &course.Title, &course.CourseCategoryId, &courseCategory.Name)
		helper.PanicError(err)

		courseCategory.Id = course.CourseCategoryId
		courses = append(courses, course)
		courseCategories = append(courseCategories, courseCategory)
	}

	return courses, courseCategories
}

func (repository *CourseRepositoryImpl) FindById(ctx context.Context, transaction *sql.Tx, id int64) (domain.Course, error) {
	rows, err := transaction.QueryContext(ctx, "SELECT id,title,course_category_id FROM courses WHERE id=?", id)
	helper.PanicError(err)
	defer rows.Close()

	var course domain.Course
	if rows.Next() {
		err := rows.Scan(&course.Id, &course.Title, &course.CourseCategoryId)
		helper.PanicError(err)
	} else {
		return course, errors.New("course not found")
	}

	return course, nil
}

func (repository *CourseRepositoryImpl) FindByCategoryId(ctx context.Context, transaction *sql.Tx, courseCategoryId int64) (domain.Course, error) {
	rows, err := transaction.QueryContext(ctx, "SELECT id,title,course_category_id FROM courses WHERE course_category_id=?", courseCategoryId)
	helper.PanicError(err)
	defer rows.Close()

	var course domain.Course
	if rows.Next() {
		err := rows.Scan(&course.Id, &course.Title, &course.CourseCategoryId)
		helper.PanicError(err)
	} else {
		return course, errors.New("course not found")
	}

	return course, nil
}

func (repository *CourseRepositoryImpl) Insert(ctx context.Context, transaction *sql.Tx, course domain.Course) domain.Course {
	result, err := transaction.ExecContext(ctx, "INSERT INTO courses(title,course_category_id) VALUES(?,?)", course.Title, course.CourseCategoryId)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	course.Id = id

	return course
}

func (repository *CourseRepositoryImpl) Update(ctx context.Context, transaction *sql.Tx, course domain.Course) domain.Course {
	_, err := transaction.ExecContext(ctx, "UPDATE courses SET title=?,course_category_id=? WHERE id=?", course.Title, course.CourseCategoryId, course.Id)
	helper.PanicError(err)

	return course
}

func (repository *CourseRepositoryImpl) Delete(ctx context.Context, transaction *sql.Tx, id int64) {
	_, err := transaction.ExecContext(ctx, "DELETE FROM courses WHERE id=?", id)
	helper.PanicError(err)
}
