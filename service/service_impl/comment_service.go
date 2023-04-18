package service_impl

import (
	"errors"
	"time"

	"github.com/haviz000/MyGram-RestAPI/helper"
	"github.com/haviz000/MyGram-RestAPI/model"
	"github.com/haviz000/MyGram-RestAPI/repository"
	"github.com/haviz000/MyGram-RestAPI/service"
)

type CommentServiceIml struct {
	commentRepository repository.CommentRepository
	photoRepository   repository.PhotoRepository
}

func NewCommentService(commentRepo repository.CommentRepository, photoRepo repository.PhotoRepository) service.CommentService {
	return &CommentServiceIml{
		commentRepository: commentRepo,
		photoRepository:   photoRepo,
	}
}

func (cs *CommentServiceIml) Create(commentReqData model.CommentCreateRequest, userID string, photoID string) (*model.CommentResponse, error) {
	_, err := cs.photoRepository.FindByID(photoID)
	if err != nil {
		return nil, err
	}

	commentID := helper.GenerateID()
	newComment := model.Comment{
		CommentID: commentID,
		Message:   commentReqData.Message,
		PhotoID:   photoID,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = cs.commentRepository.Create(newComment)
	if err != nil {
		return nil, err
	}

	return &model.CommentResponse{
		CommentID: newComment.CommentID,
		Message:   newComment.Message,
		PhotoID:   newComment.PhotoID,
		UserID:    newComment.UserID,
		CreatedAt: newComment.CreatedAt,
		UpdatedAt: newComment.UpdatedAt,
	}, nil
}

func (cs *CommentServiceIml) GetAll() ([]model.CommentResponse, error) {
	commentsResult, err := cs.commentRepository.FindAll()
	if err != nil {
		return []model.CommentResponse{}, err
	}

	commentsResponse := []model.CommentResponse{}
	for _, commentRes := range commentsResult {
		commentsResponse = append(commentsResponse, model.CommentResponse(commentRes))
	}

	return commentsResponse, nil
}

func (cs *CommentServiceIml) GetOne(commentID string) (model.CommentResponse, error) {
	commentResult, err := cs.commentRepository.FindByID(commentID)
	if err != nil {
		return model.CommentResponse{}, err
	}

	return model.CommentResponse(commentResult), nil
}

func (cs *CommentServiceIml) UpdateComment(commentReqData model.CommentUpdateRequest, userID string, commentID string) (*model.CommentResponse, error) {
	findCommentResponse, err := cs.commentRepository.FindByID(commentID)
	if err != nil {
		return nil, err
	}

	if userID != findCommentResponse.UserID {
		return nil, errors.New("Unauthorized")
	}

	updatedCommentReq := model.Comment{
		CommentID: findCommentResponse.CommentID,
		Message:   commentReqData.Message,
		PhotoID:   findCommentResponse.PhotoID,
		UserID:    userID,
		UpdatedAt: time.Now(),
	}

	err = cs.commentRepository.Update(updatedCommentReq)
	if err != nil {
		return nil, err
	}

	return &model.CommentResponse{
		CommentID: commentID,
	}, nil
}

func (cs *CommentServiceIml) Delete(commentlID string, userID string) (model.CommentDeleteResponse, error) {
	findCommentResponse, err := cs.commentRepository.FindByID(commentlID)
	if err != nil {
		return model.CommentDeleteResponse{}, err
	}

	if userID != findCommentResponse.UserID {
		return model.CommentDeleteResponse{}, errors.New("Unauthorized")
	}

	err = cs.commentRepository.Delete(model.Comment{CommentID: commentlID})
	if err != nil {
		return model.CommentDeleteResponse{}, err
	}

	return model.CommentDeleteResponse{
		CommentID: commentlID,
	}, nil
}
