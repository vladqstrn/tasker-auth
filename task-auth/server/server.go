package server

import (
	"github.com/gin-gonic/gin"
	cfg "github.com/vladqstrn/tasker-auth/task-auth/config"
	"github.com/vladqstrn/tasker-auth/task-auth/mw"
	"github.com/vladqstrn/tasker-auth/task-auth/tasker/delivery/rest/routes"
	"github.com/vladqstrn/tasker-auth/task-auth/tasker/repo"
	"github.com/vladqstrn/tasker-auth/task-auth/tasker/usecase"
)

type App struct {
	r    *gin.Engine
	uc   *usecase.AuthUsecase
	repo *repo.UserRepository
}

func New(r *gin.Engine, uc *usecase.AuthUsecase, repo *repo.UserRepository) *App {
	return &App{
		r:    r,
		uc:   uc,
		repo: repo,
	}
}

func (a *App) Run() {
	mw.CORSMiddleware(a.r)
	routes.AuthRoutes(a.r, a.uc)

	a.r.Run(cfg.Domain + ":" + cfg.AppPort)
}
