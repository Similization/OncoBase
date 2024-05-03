package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db *sqlx.DB
}

func NewCourseRepository(db *sqlx.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

// Create course in database and get him from database
func (r *CourseRepository) CreateCourse(course model.Course) (model.Course, error) {
	var createdCourse model.Course
	query := fmt.Sprintf("INSERT INTO %s (id, period, frequency, dose, drug, measure_code) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *", courseTable)
	err := r.db.Get(&createdCourse, query,
		course.Id,
		course.Period,
		course.Frequency,
		course.Dose,
		course.Drug,
		course.MeasureCode,
	)
	return createdCourse, err
}

// Get course list from database
func (r *CourseRepository) GetCourseList() ([]model.Course, error) {
	var courseList []model.Course
	query := fmt.Sprintf("SELECT * FROM %s", courseTable)
	err := r.db.Select(&courseList, query)
	return courseList, err
}

// Get course from database by ID
func (r *CourseRepository) GetCourseById(id string) (model.Course, error) {
	var course model.Course
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", courseTable)
	err := r.db.Get(&course, query, id)
	return course, err
}

// Update course data in database
func (r *CourseRepository) UpdateCourse(course model.Course) (model.Course, error) {
	var updatedCourse model.Course
	query := fmt.Sprintf("UPDATE %s SET period=$2, frequency=$3, dose=$4, drug=$5, measure_code=$6 WHERE id=$1 RETURNING *", courseTable)
	err := r.db.Get(&updatedCourse, query,
		course.Id,
		course.Period,
		course.Frequency,
		course.Dose,
		course.Drug,
		course.MeasureCode,
	)
	return updatedCourse, err
}

// Delete course data from database
func (r *CourseRepository) DeleteCourse(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", courseTable)
	_, err := r.db.Exec(query, id)
	return err
}
