package controller_impl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haviz000/MyGram-RestAPI/controller"
	"github.com/haviz000/MyGram-RestAPI/helper"
	"github.com/haviz000/MyGram-RestAPI/model"
	"github.com/haviz000/MyGram-RestAPI/service"
	"gorm.io/gorm"
)

type CommentControllerImpl struct {
	commentService service.CommentService
}

func NewCommentController(service service.CommentService) controller.CommentController {
	return &CommentControllerImpl{
		commentService: service,
	}
}

// CreateComment godoc
//
//		@Summary		create comment
//		@Description	create comment for a particular user
//		@Tags			Comment
//		@Accept			json
//		@Produce		json
//	 @Param          photo_id   path      string  true  "photo_id"
//		@Param			request	body		model.CommentCreateReq	true	"request is required"
//		@Success		200		{object}	model.SuccessResponse{data=model.CommentResponse}
//		@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//		@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//		@Security		BearerAuth
//		@Router			/comments/{photo_id} [post]
func (c *CommentControllerImpl) CreateComment(ctx *gin.Context) {
	var request model.CommentCreateRequest
	photoID := ctx.Param("photo_id")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.CommentCreateValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := c.commentService.Create(request, userID.(string), photoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Errors: errors.New("Photo not found").Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Comment created successfully",
		Data:    response,
	})
}

// GetAllComment godoc
//
//	@Summary		get all comment
//	@Description	get all comment
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	model.SuccessResponse{data=[]model.CommentResponse}
//	@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/comments [get]
func (c *CommentControllerImpl) GetAll(ctx *gin.Context) {
	response, err := c.commentService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Get all comment successfully",
		Data:    response,
	})
}

// GetOneComment godoc
//
//		@Summary		get one comment
//		@Description	get one comment
//		@Tags			Comment
//		@Accept			json
//		@Produce		json
//	 @Param          comment_id   path      string  true  "comment_id"
//		@Success		200		{object}	model.SuccessResponse{data=model.CommentResponse}
//		@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//		@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//		@Security		BearerAuth
//		@Router			/comments/{comment_id} [get]
func (c *CommentControllerImpl) GetOne(ctx *gin.Context) {
	commentID := ctx.Param("comment_id")

	response, err := c.commentService.GetOne(commentID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Get comment successfully",
		Data:    response,
	})
}

// UpdateComment godoc
//
//		@Summary		update comment
//		@Description	update comment
//		@Tags			Comment
//		@Accept			json
//		@Produce		json
//	 @Param          comment_id   path      string  true  "comment_id"
//		@Param			request	body		model.CommentUpdateReq	true	"request is required"
//		@Success		200		{object}	model.SuccessResponse{data=model.CommentUpdateRes}
//		@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//		@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//		@Security		BearerAuth
//		@Router			/comments/{comment_id} [put]
func (c *CommentControllerImpl) UpdateComment(ctx *gin.Context) {
	var request model.CommentUpdateRequest
	commentID := ctx.Param("comment_id")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.CommentUpdateValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := c.commentService.UpdateComment(request, userID.(string), commentID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Comment updated successfully",
		Data: model.CommentUpdateResponse{
			CommentID: response.CommentID,
		},
	})
}

// DeleteComment godoc
//
//		@Summary		delete comment
//		@Description	delete comment
//		@Tags			Comment
//		@Accept			json
//		@Produce		json
//	 @Param          comment_id   path      string  true  "comment_id"
//		@Success		200		{object}	model.SuccessResponse{data=model.CommentDeleteRes}
//		@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//		@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//		@Security		BearerAuth
//		@Router			/comments/{comment_id} [delete]
func (c *CommentControllerImpl) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("comment_id")

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	response, err := c.commentService.Delete(commentID, userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Comment deleted successfully",
		Data: model.CommentDeleteResponse{
			CommentID: response.CommentID,
		},
	})
}
