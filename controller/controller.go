package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
}

type SocialController interface {
	CreateSocial(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type PhotoController interface {
	CreatePhoto(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}
