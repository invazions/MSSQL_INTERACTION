package repositories

import (
	"database/sql"
	"github.com/invazions/MSSQL_INTERACTION/models"
)

type UserRepository interface {
	Create(user models.User) error
	GetByID(id int) (models.User, error)
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user models.User) error {
	_, err := r.db.Exec("INSERT INTO Users (Name, Email) VALUES (?, ?)", user.Name, user.Email)
	return err
}

func (r *userRepository) GetByID(id int) (models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT ID, Name, Email FROM Users WHERE ID = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func (r *userRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM Users WHERE ID = ?", id)
	return err
}
