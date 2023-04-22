package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	cfg "github.com/vladqstrn/tasker-auth/task-auth/config"
	"github.com/vladqstrn/tasker-auth/task-auth/models"
	"github.com/vladqstrn/tasker-auth/task-auth/tasker/usecase"
	"github.com/vladqstrn/tasker-auth/task-auth/utils"
)

type UserHandler struct {
	useCase usecase.Auth
}

func NewTaskHandler(userService usecase.Auth) *UserHandler {
	return &UserHandler{useCase: userService}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	if err := h.useCase.Register(&user); err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if err := h.useCase.Login(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	aToken, rToken := utils.GenarateTokens(user.Username)

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Authorization", "Bearer "+aToken)

	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie(cfg.CookieName, rToken, cfg.CookieMaxAge, cfg.CookiePath, cfg.CookieDomain, cfg.CookieSecure, cfg.CookieHttpOnly)
	_, err := utils.ParseJWTWithClaims(aToken)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (h *UserHandler) CheckAuth(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		ctx.Abort()
		return
	}

	if tokenString != "" {
		_, err := utils.ParseJWTWithClaims(tokenString[7:])
		if err != nil {

			switch err.Error() {
			case "token is expired":
				ctx.Redirect(http.StatusPermanentRedirect, "/user/updatetokens")
			default:

				ctx.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}

}

func (h *UserHandler) UpdateExpiredTokens(ctx *gin.Context) {

	rToken, err := ctx.Cookie("auth")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := utils.ValidateToken(rToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	usr, err := h.useCase.GetUser(user)
	if err != nil {
		return
	}

	aToken, rToken := utils.GenarateTokens(usr.Username)

	ctx.Header("Authorization", "Bearer "+aToken)
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie(cfg.CookieName, rToken, cfg.CookieMaxAge, cfg.CookiePath, cfg.CookieDomain, cfg.CookieSecure, cfg.CookieHttpOnly)
}
