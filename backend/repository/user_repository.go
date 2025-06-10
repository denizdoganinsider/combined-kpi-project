package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"myapp-backend/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (userRepository *UserRepository) CreateUser(user domain.User) error {
	_, err := userRepository.db.Exec("INSERT INTO users(first_name, last_name, email, hashed_password) VALUES (?, ?, ?, ?)", user.Name, user.Surname, user.Email, user.HashedPassword)

	if err != nil {
		return fmt.Errorf("user not added to db %s", err)
	}
	return nil
}

func (userRepository *UserRepository) FindByEmail(email string) (*domain.User, error) {
	query := "SELECT first_name, last_name, email, hashed_password FROM users WHERE email = ?"

	row := userRepository.db.QueryRow(query, email)

	var user domain.User

	err := row.Scan(&user.Name, &user.Surname, &user.Email, &user.HashedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("error occured while gettint the user data")
	}

	return &user, nil
}
