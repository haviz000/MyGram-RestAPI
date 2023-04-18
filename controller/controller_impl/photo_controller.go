package controller_impl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haviz000/MyGram-RestAPI/controller"
	"github.com/haviz000/MyGram-RestAPI/helper"
	"github.com/haviz000/MyGram-RestAPI/model"
	"github.com/haviz000/MyGram-RestAPI/service"
)

type PhotoControllerImpl struct {
	photoService service.PhotoService
}

func NewPhotoController(service service.PhotoService) controller.PhotoController {
	return &PhotoControllerImpl{
		photoService: service,
	}
}

// CreatePhoto godoc
//
//	@Summary		create photo
//	@Description	create photo for a particular user
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.PhotoCreateReq	true	"request is required"
//	@Success		200		{object}	model.SuccessResponse{data=model.PhotoCreateRes}
//	@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/photo [post]
func (c *PhotoControllerImpl) CreatePhoto(ctx *gin.Context) {
	var request model.PhotoCreateRequest

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
	validateErrs = helper.PhotoCreateValidator(request)
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

	response, err := c.photoService.Create(request, userID.(string))
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
		Message: "Photo created successfully",
		Data:    response,
	})
}

// GetAllPhoto godoc
//
//	@Summary		get all photo
//	@Description	get all photo
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	model.SuccessResponse{data=[]model.PhotoResponse}
//	@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/photo [get]
func (c *PhotoControllerImpl) GetAll(ctx *gin.Context) {
	response, err := c.photoService.GetAll()
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
		Message: "Get all photo successfully",
		Data:    response,
	})
}

// GetOnePhoto godoc
//
//		@Summary		get one photo
//		@Description	get one photo
//		@Tags			Photo
//		@Accept			json
//		@Produce		json
//	 @Param          photo_id   path      string  true  "PhotoID"
//		@Success		200		{object}	model.SuccessResponse{data=model.PhotoResponse}
//		@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//		@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//		@Security		BearerAuth
//		@Router			/photo/{photoId} [get]
func (c *PhotoControllerImpl) GetOne(ctx *gin.Context) {
	photoID := ctx.Param("photoId")

	response, err := c.photoService.GetOne(photoID)
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
		Message: "Get photo successfully",
		Data:    response,
	})
}

// UpdatePhoto godoc
//
//		@Summary		update photo
//		@Description	update photo
//		@Tags			Photo
//		@Accept			json
//		@Produce		json
//	 @Param          photo_id   path      string  true  "PhotoID"
//		@Param			request	body		model.PhotoUpdateReq	true	"request is required"
//		@Success		200		{object}	model.SuccessResponse{data=model.PhotoUpdateRes}
//		@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//		@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//		@Security		BearerAuth
//		@Router			/photo/{photoId} [put]
func (c *PhotoControllerImpl) UpdatePhoto(ctx *gin.Context) {
	var request model.PhotoUpdateRequest
	photoID := ctx.Param("photoId")

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
	validateErrs = helper.PhotoUpdateValidator(request)
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

	response, err := c.photoService.UpdatePhoto(request, userID.(string), photoID)
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
		Message: "Photo updated successfully",
		Data: model.PhotoUpdateResponse{
			PhotoID: response.PhotoID,
		},
	})
}

// DeletePhoto godoc
//
//		@Summary		delete photo
//		@Description	delete photo
//		@Tags			Photo
//		@Accept			json
//		@Produce		json
//	 @Param          photo_id   path      string  true  "PhotoID"
//		@Success		200		{object}	model.SuccessResponse{data=model.PhotoDeleteRes}
//		@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//		@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//		@Security		BearerAuth
//		@Router			/photos/{photoId} [delete]
func (c *PhotoControllerImpl) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	response, err := c.photoService.Delete(photoID, userID.(string))
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
		Message: "Photo deleted successfully",
		Data: model.PhotoUpdateResponse{
			PhotoID: response.PhotoID,
		},
	})
}
