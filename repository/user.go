package repository

import (
	"database/sql"

	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/database"
)

type UserRepositoryInterface interface {
	FindById(id int) (*data.User, error)
	FindByCredentials(email, password string) (*data.User, error)
	Create(user *data.User) error
}

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (ur *userRepository) FindById(id int) (*data.User, error) {

	row := database.GetDb().QueryRow(`
		SELECT id, email 
		FROM tbl_user u 
		WHERE u.id = $1	`, id,
	)

	user := &data.User{}

	err := row.Scan(&user.Id, &user.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindByCredentials(email, password string) (*data.User, error) {

	row := database.GetDb().QueryRow(`
		SELECT id, email
		FROM tbl_user u 
		WHERE u.email = $1
		AND u.password = $2
	`, email, password)

	user := &data.User{}

	err := row.Scan(&user.Id, &user.Email)

	if err == sql.ErrNoRows {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) Create(user *data.User) error {

	_, err := database.GetDb().Exec(`
		INSERT INTO
			tbl_user(email, password)
		VALUES
			($1, $2)
	`, user.Email, user.Password)

	return err
}
