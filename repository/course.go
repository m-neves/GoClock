package repository

import (
	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/database"
)

type CourseRepositoryInterface interface {
	Create(course *data.Course, userId int) error
	FindAll(userId int) ([]*data.Course, error)
}

type courseRepository struct{}

func NewCourseRepository() CourseRepositoryInterface {
	return &courseRepository{}
}

func (cr *courseRepository) Create(course *data.Course, userId int) error {
	_, err := database.GetDb().Exec(`
		INSERT INTO 
			goclock.tbl_course(course_name, duration, user_id)
		VALUES
			($1, $2, $3)
	`, course.Name, course.Duration, userId)

	return err
}

func (cr *courseRepository) FindAll(userId int) ([]*data.Course, error) {

	rows, err := database.GetDb().Query(`
		SELECT id, course_name, duration
		FROM tbl_course c
		WHERE c.user_id = $1
	`, userId)

	if err != nil {
		return nil, err
	}

	var courses []*data.Course

	for rows.Next() {
		course := &data.Course{}

		err := rows.Scan(
			&course.Id,
			&course.Name,
			&course.Duration,
		)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}
