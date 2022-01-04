package repository

import (
	"context"

	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/database"
)

type StudyPlanRepositoryInterface interface {
	Create(studyPlan *data.StudyPlan) error
	AddCourse(studyPlan *data.StudyPlanCourse) error
	FindById(id, user_id uint) (*data.StudyPlanCourses, error)
}

type studyPlanRepository struct{}

func NewStudyPlanRepository() StudyPlanRepositoryInterface {
	return &studyPlanRepository{}
}

func (spr *studyPlanRepository) Create(studyPlan *data.StudyPlan) error {

	_, err := database.GetDb().Exec(`
		INSERT INTO goclock.tbl_study_plan
			(study_plan_name, user_id)
		VALUES
			($1, $2)
	`, studyPlan.Name, studyPlan.UserId)

	return err
}

func (spr *studyPlanRepository) AddCourse(studyPlan *data.StudyPlanCourse) error {
	tx, err := database.GetDb().BeginTx(context.Background(), nil)

	if err != nil {
		return err
	}

	row := tx.QueryRow(`
		SELECT COALESCE(MAX(spc.study_plan_course_order), 0) max_order
		FROM goclock.tbl_study_plan_course spc
		WHERE spc.study_plan_id = $1 
	`, studyPlan.StudyPlanId)

	var maxOrder uint
	err = row.Scan(&maxOrder)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO goclock.tbl_study_plan_course
			(course_id, study_plan_id, study_plan_course_order)
		VALUES
			($1, $2, $3)
	`, studyPlan.CourseId, studyPlan.StudyPlanId, maxOrder)

	if err != nil {
		return err
	}

	return nil
}

func (spr *studyPlanRepository) FindById(id, user_id uint) (*data.StudyPlanCourses, error) {

	rows, err := database.GetDb().Query(`
		SELECT
			tsp.id, tsp.study_plan_name, tsp.current_course, tc.course_name, tc.duration 
		FROM
			goclock.tbl_study_plan_course tspc 
		INNER JOIN goclock.tbl_study_plan tsp 
			ON tsp.id = tspc.study_plan_id 
		INNER JOIN goclock.tbl_course tc
			ON tc.id = tspc.course_id 
		WHERE
			tspc.study_plan_id = $1
		AND tsp.user_id = $2
	`, id, user_id)

	if err != nil {
		return nil, err
	}

	studyPlanCourses := &data.StudyPlanCourses{}

	for rows.Next() {
		course := &data.Course{}

		rows.Scan(
			&studyPlanCourses.StudyPlan.Id,
			&studyPlanCourses.StudyPlan.Name,
			&studyPlanCourses.StudyPlan.CurrentCourse,
			&course.Name,
			&course.Duration,
		)

		studyPlanCourses.Courses = append(studyPlanCourses.Courses, course)
	}

	return studyPlanCourses, nil
}
