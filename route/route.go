package route

import (
	"github.com/gin-gonic/gin"
	"github.com/haviz000/MyGram-RestAPI/controller/controller_impl"
	"github.com/haviz000/MyGram-RestAPI/middleware"
	"github.com/haviz000/MyGram-RestAPI/repository/repository_impl"
	"github.com/haviz000/MyGram-RestAPI/service/service_impl"

	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {

	userRepository := repository_impl.NewUserRepository(db)
	photoRepository := repository_impl.NewPhotoRepository(db)
	socialRepository := repository_impl.NewSocialRepository(db)
	commentRepsitory := repository_impl.NewCommentRepository(db)

	userService := service_impl.NewUserService(userRepository, photoRepository, socialRepository)
	socialService := service_impl.NewSocialService(socialRepository)
	photoService := service_impl.NewPhotoService(photoRepository, commentRepsitory)
	commentService := service_impl.NewCommentService(commentRepsitory, photoRepository)

	userController := controller_impl.NewUserController(userService)
	socialController := controller_impl.NewSocialController(socialService)
	photoController := controller_impl.NewPhotoController(photoService)
	commentController := controller_impl.NewCommentController(commentService)

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	authUser := router.Group("/User", middleware.AuthMiddleware)
	{
		authUser.GET("/profile", userController.GetProfile)
	}

	authSocial := router.Group("/social", middleware.AuthMiddleware)
	{
		authSocial.POST("", socialController.CreateSocial)
		authSocial.GET("", socialController.GetAll)
		authSocial.GET("/:socialId", socialController.GetOne)
		authSocial.PUT("/:socialId", socialController.UpdateSocialMedia)
		authSocial.DELETE("/:socialId", socialController.DeleteSocialMedia)
	}

	authPhoto := router.Group("/photo", middleware.AuthMiddleware)
	{
		authPhoto.POST("", photoController.CreatePhoto)
		authPhoto.GET("", photoController.GetAll)
		authPhoto.GET("/:photoId", photoController.GetOne)
		authPhoto.PUT("/:photoId", photoController.UpdatePhoto)
		authPhoto.DELETE("/:photoId", photoController.DeletePhoto)
	}

	authComment := router.Group("/comment", middleware.AuthMiddleware)
	{
		authComment.POST("/:photoId", commentController.CreateComment)
		authComment.GET("", commentController.GetAll)
		authComment.GET("/:commentId", commentController.GetOne)
		authComment.PUT("/:commentId", commentController.UpdateComment)
		authComment.DELETE("/:commentId", commentController.DeleteComment)
	}
}
