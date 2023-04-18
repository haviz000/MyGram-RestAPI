package service

import "github.com/haviz000/MyGram-RestAPI/model"

type UserService interface {
	Register(userReqData model.UserRegisterRequest) (*model.UserRegisterResponse, error)
	Login(userReqData model.UserLoginRequest) (*string, error)
	GetProfile(userID string) (model.User, error)
}

type SocialService interface {
	Create(socialReqData model.SocialMediaCreateRequest, userID string) (*model.SocialMediaResponse, error)
	GetAll() ([]model.SocialMediaResponse, error)
	GetOne(socialID string) (model.SocialMediaResponse, error)
	UpdateSocialMedia(socialReqData model.SocialMediaUpdateRequest, userID string, socialID string) (*model.SocialMediaResponse, error)
	Delete(socialID string, userID string) (model.SocialMediaResponse, error)
}

type PhotoService interface {
	Create(photoReqData model.PhotoCreateRequest, userID string) (*model.PhotoCreateResponse, error)
	GetAll() ([]model.PhotoResponse, error)
	GetOne(photoID string) (model.PhotoResponse, error)
	UpdatePhoto(photoReqData model.PhotoUpdateRequest, userID string, photoID string) (*model.PhotoResponse, error)
	Delete(photoID string, userID string) (model.PhotoResponse, error)
}

type CommentService interface {
	Create(commentReqData model.CommentCreateRequest, userID string, photoID string) (*model.CommentResponse, error)
	GetAll() ([]model.CommentResponse, error)
	GetOne(commentID string) (model.CommentResponse, error)
	UpdateComment(commentReqData model.CommentUpdateRequest, userID string, commentID string) (*model.CommentResponse, error)
	Delete(commentID string, userID string) (model.CommentDeleteResponse, error)
}
