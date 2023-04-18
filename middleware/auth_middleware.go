package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/haviz000/MyGram-RestAPI/helper"
	"github.com/haviz000/MyGram-RestAPI/model"
)

func AuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if bearerIsExist := strings.Contains(auth, "Bearer"); !bearerIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: "Unauthorized",
		})
		return
	}

	token := strings.Split(auth, " ")
	if len(token) < 2 {
		err := errors.New("Must provide Authorization header with format `Bearer {token}`")

		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	claims, err := helper.VerifyAccessToken(token[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	c.Set("username", claims.Username)
	c.Set("userID", claims.UserID)

	c.Next()
}

func AuthRefreshMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	token := strings.Split(auth, " ")[1]

	jwtToken, err := helper.VerifyRefreshToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Server Error",
			Errors: err.Error(),
		})
		return
	}

	v, ok := claims["refresh"]
	if !ok || !v.(bool) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.Set("user_id", claims["user_id"])

	ctx.Next()
}
