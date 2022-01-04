package repository

import (
	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/database"
)

type SubjectRepositoryInterface interface {
	FindById(id, userId int) (*data.Subject, error)
	FindAll(userId int) ([]*data.Subject, error)
	Create(name string, userId int) error
}

type subjectRepository struct {
}

func NewSubjectRepository() *subjectRepository {
	return &subjectRepository{}
}

func (sr *subjectRepository) FindById(id, userId int) (*data.Subject, error) {

	row := database.GetDb().QueryRow(`
		SELECT id, subject_name
		FROM tbl_subject s
		WHERE s.id = $1
		AND s.user_id = $2
	`, id, userId)

	subject := &data.Subject{}

	err := row.Scan(&subject.Id, &subject.Name)

	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (sr *subjectRepository) FindAll(userId int) ([]*data.Subject, error) {

	rows, err := database.GetDb().Query(`
		SELECT id, subject_name
		FROM tbl_subject s
		WHERE s.user_id = $1
	`, userId)

	if err != nil {
		return nil, err
	}

	subjects := []*data.Subject{}

	for rows.Next() {
		subject := &data.Subject{}

		err := rows.Scan(
			&subject.Id,
			&subject.Name,
		)

		if err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func (sr *subjectRepository) Create(name string, userId int) error {

	_, err := database.GetDb().Exec(`
		INSERT INTO tbl_subject
			(user_id, subject_name)
		VALUES
			($1, $2)		
	`, userId, name)

	return err
}
