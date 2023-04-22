package usecase

import (
	"time"

	"github.com/vladqstrn/tasker-auth/task-auth/custom_err"
	"github.com/vladqstrn/tasker-auth/task-auth/models"
	"github.com/vladqstrn/tasker-auth/task-auth/tasker/repo"
	"github.com/vladqstrn/tasker-auth/task-auth/utils"
)

type Auth interface {
	Register(user *models.User) error
	Login(user *models.User) error
	GetUser(user string) (*models.User, error)
}

type AuthUsecase struct {
	auth Auth
	repo *repo.UserRepository
}

func NewAuthUsecase(repo *repo.UserRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (a *AuthUsecase) Register(user *models.User) error {
	userExists, err := a.repo.GetUserByUsername(user.Username)

	if err == custom_err.ErrUserNotFound {

		h, err := utils.HashPassword(user.Password)
		if err != nil {

			return err
		}
		user.Password = h

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		err = a.repo.CreateUser(user)
		if err != nil {
			return err
		}
	} else {
		if err != nil {
			return err
		}
	}
	if userExists != nil {
		return custom_err.ExsistsUser
	}

	return nil
}

func (a *AuthUsecase) Login(user *models.User) error {
	usr, err := a.repo.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}

	err = utils.CheckPassHash(user.Password, usr.Password)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthUsecase) GetUser(user string) (*models.User, error) {
	usr, err := a.repo.GetUserByUsername(user)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
