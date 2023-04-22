package repo

import (
	"errors"

	"github.com/go-pg/pg"
	"github.com/vladqstrn/tasker-auth/task-auth/models"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.db.Model(user).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(user).Where("username = ?", username).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil

}
