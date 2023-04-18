package service_impl

import (
	"errors"
	"fmt"

	"github.com/haviz000/MyGram-RestAPI/helper"
	"github.com/haviz000/MyGram-RestAPI/model"
	"github.com/haviz000/MyGram-RestAPI/repository"
	"github.com/haviz000/MyGram-RestAPI/service"
)

type UserServiceImpl struct {
	userRepository  repository.UserRepository
	photoRepository repository.PhotoRepository
	socalRepository repository.SocialRepository
}

func NewUserService(userRepo repository.UserRepository, photoRepo repository.PhotoRepository, socialRepo repository.SocialRepository) service.UserService {
	return &UserServiceImpl{
		userRepository:  userRepo,
		photoRepository: photoRepo,
		socalRepository: socialRepo,
	}
}

func (us *UserServiceImpl) Register(userReqData model.UserRegisterRequest) (*model.UserRegisterResponse, error) {
	userID := helper.GenerateID()
	hashedPassword, err := helper.Hash(userReqData.Password)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		UserID:   userID,
		Username: userReqData.Username,
		Email:    userReqData.Email,
		Password: hashedPassword,
		Age:      userReqData.Age,
	}

	err = us.userRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return &model.UserRegisterResponse{
		UserID:   newUser.UserID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
		Age:      newUser.Age,
	}, nil
}

func (us *UserServiceImpl) Login(userReqData model.UserLoginRequest) (*string, error) {
	userResponse, err := us.userRepository.FindByUsername(userReqData.Username)
	if err != nil {
		return nil, err
	}

	isMatch := helper.PasswordIsMatch(userReqData.Password, userResponse.Password)
	if isMatch == false {
		return nil, errors.New(fmt.Sprintf("Invalid username or password"))
	}

	token, err := helper.GenerateAccessToken(userResponse)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (us *UserServiceImpl) GetProfile(userID string) (model.User, error) {
	user, err := us.userRepository.FindByID(userID)
	if err != nil {
		return model.User{}, err
	}

	photos, err := us.photoRepository.FindByUserID(userID)
	if err != nil {
		return model.User{}, err
	}

	socials, err := us.socalRepository.FindByUserID(userID)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		UserID:      userID,
		Username:    user.Username,
		Email:       user.Email,
		Age:         user.Age,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Photos:      photos,
		SocialMedia: socials,
	}, nil
}
