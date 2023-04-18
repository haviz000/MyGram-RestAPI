package repository

import "github.com/haviz000/MyGram-RestAPI/model"

type UserRepository interface {
	Create(userRequestData model.User) error
	FindByID(userID string) (model.User, error)
	FindByUsername(username string) (model.User, error)
}

type SocialRepository interface {
	Create(photoReqData model.SocialMedia) error
	FindAll() ([]model.SocialMedia, error)
	FindByID(socialID string) (model.SocialMedia, error)
	FindByUserID(userID string) ([]model.SocialMedia, error)
	Update(socialReqData model.SocialMedia) error
	Delete(photoReqData model.SocialMedia) error
}

type PhotoRepository interface {
	Create(photoReqData model.Photo) error
	FindAll() ([]model.Photo, error)
	FindByID(photoID string) (model.Photo, error)
	FindByUserID(userID string) ([]model.Photo, error)
	Update(photoReqData model.Photo) error
	Delete(photoReqData model.Photo) error
}

type CommentRepository interface {
	Create(commentReqData model.Comment) error
	FindAll() ([]model.Comment, error)
	FindByID(commentID string) (model.Comment, error)
	FindByPhotoID(photoID string) ([]model.Comment, error)
	Update(commentReqData model.Comment) error
	Delete(commentReqData model.Comment) error
}
