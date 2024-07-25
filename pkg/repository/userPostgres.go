package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	_ "managep"
	"managep/pkg/model"
	"net/http"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Check(user *model.User) bool {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE full_name=$1 AND email=$2", usersTable)
	var exists bool
	err := r.db.QueryRow(query, user.FullName, user.Email).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true
		}
		log.Fatalf("Error checking user existence: %v", err)
	}
	return false
}

func (r *UserPostgres) GetUser() ([]model.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", usersTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.RegistrationDate, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserPostgres) GetUserById(id string) (model.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, errors.New("user not found")
		}
	}
	defer rows.Close()
	var user model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.RegistrationDate, &user.Role); err != nil {
			return model.User{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserPostgres) CreateUser(user *model.User) (int, error) {
	if r.Check(user) == false {
		return 404, errors.New("user is already registered")
	}
	query := fmt.Sprintf("INSERT INTO %s (full_name, email, registration_date, role) VALUES ($1, $2, $3, $4)", usersTable)
	_, err := r.db.Exec(query, user.FullName, user.Email, user.RegistrationDate, user.Role)
	if err != nil {
		return 404, err
	}
	return 201, nil
}

func (r *UserPostgres) UpdateUser(user *model.User, id string) (int, error) {
	checkQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	err := r.db.QueryRow(checkQuery, id).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 404, errors.New("user not found")
		}
	}
	query := fmt.Sprintf("UPDATE %s SET full_name = $1, email = $2, registration_date = $3, role = $4 WHERE id = $5", usersTable)
	_, err = r.db.Exec(query, user.FullName, user.Email, user.RegistrationDate, user.Role, id)
	if err != nil {
		return 404, err
	}
	return 201, nil
}

func (r *UserPostgres) DeleteUser(id string) (int, error) {
	checkQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	err := r.db.QueryRow(checkQuery, id).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 404, errors.New("user not found")
		}
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return 201, nil
}

func (r *UserPostgres) GetTasksForUser(id string) ([]model.Task, error) {
	checkQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	err := r.db.QueryRow(checkQuery, id).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []model.Task{}, errors.New("user not found")
		}
	}
	return []model.Task{}, nil
}