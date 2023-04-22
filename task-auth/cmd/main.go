package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vladqstrn/tasker-auth/task-auth/config"
	"github.com/vladqstrn/tasker-auth/task-auth/database"
	"github.com/vladqstrn/tasker-auth/task-auth/server"
	"github.com/vladqstrn/tasker-auth/task-auth/tasker/repo"
	"github.com/vladqstrn/tasker-auth/task-auth/tasker/usecase"
)

func main() {
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	if err := config.InitConfig(); err != nil {
		panic("failed to initialize config")
	}

	r := gin.Default()

	db := database.InitDB()
	repo := repo.NewUserRepository(db)
	uc := usecase.NewAuthUsecase(repo)

	app := server.New(r, uc, repo)
	app.Run()

}
