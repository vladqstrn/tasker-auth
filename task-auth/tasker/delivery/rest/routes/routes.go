package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vladqstrn/tasker-auth/task-auth/tasker/delivery/rest/handler"
	"github.com/vladqstrn/tasker-auth/task-auth/tasker/usecase"
)

func AuthRoutes(r *gin.Engine, useCase usecase.Auth) {
	userHandler := handler.NewTaskHandler(useCase)
	userGroup := r.Group("/user")

	userGroup.POST("/register", userHandler.Register)
	userGroup.POST("/login", userHandler.Login)
	userGroup.POST("/auth", userHandler.CheckAuth)
	userGroup.POST("/updatetokens", userHandler.UpdateExpiredTokens)
}
