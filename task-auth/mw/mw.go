package mw

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vladqstrn/tasker-auth/task-auth/config"
)

func CORSMiddleware(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Origins},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Set-Cookie", "Cookie"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

}
