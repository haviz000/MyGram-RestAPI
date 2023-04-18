package repository_impl

import (
	"errors"
	"fmt"

	"github.com/haviz000/MyGram-RestAPI/model"
	"github.com/haviz000/MyGram-RestAPI/repository"
	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &CommentRepositoryImpl{
		DB: db,
	}
}

func (r *CommentRepositoryImpl) Create(commentReqData model.Comment) error {
	err := r.DB.Create(&commentReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommentRepositoryImpl) FindAll() ([]model.Comment, error) {
	comments := []model.Comment{}

	err := cr.DB.Find(&comments).Error
	if err != nil {
		return []model.Comment{}, err
	}

	return comments, nil
}

func (cr *CommentRepositoryImpl) FindByID(commentID string) (model.Comment, error) {
	comment := model.Comment{}

	err := cr.DB.Debug().Where("comment_id = ?", commentID).Take(&comment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Comment{}, err
		}

		return model.Comment{}, err
	}

	return comment, nil
}

func (cr *CommentRepositoryImpl) FindByPhotoID(photoID string) ([]model.Comment, error) {
	comments := []model.Comment{}

	err := cr.DB.Where("photo_id = ?", photoID).Find(&comments).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Comment{}, err
		}

		return []model.Comment{}, err
	}

	fmt.Println("comments: ", comments)

	return comments, nil
}

func (cr *CommentRepositoryImpl) Update(commentReqData model.Comment) error {
	err := cr.DB.Save(&model.Comment{
		CommentID: commentReqData.CommentID,
		Message:   commentReqData.Message,
		PhotoID:   commentReqData.PhotoID,
		UserID:    commentReqData.UserID,
		UpdatedAt: commentReqData.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *CommentRepositoryImpl) Delete(commentReqData model.Comment) error {
	err := cr.DB.Delete(&commentReqData).Error
	if err != nil {
		return err
	}

	return nil
}
