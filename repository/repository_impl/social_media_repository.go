package repository_impl

import (
	"errors"

	"github.com/haviz000/MyGram-RestAPI/model"
	"github.com/haviz000/MyGram-RestAPI/repository"
	"gorm.io/gorm"
)

type SocialRepositoryImpl struct {
	DB *gorm.DB
}

func NewSocialRepository(db *gorm.DB) repository.SocialRepository {
	return &SocialRepositoryImpl{
		DB: db,
	}
}

func (sr *SocialRepositoryImpl) Create(socialReqData model.SocialMedia) error {
	err := sr.DB.Create(&socialReqData).Error
	if err != nil {
		return err
	}
	return nil
}

func (sr *SocialRepositoryImpl) FindAll() ([]model.SocialMedia, error) {
	socials := []model.SocialMedia{}

	err := sr.DB.Find(&socials).Error
	if err != nil {
		return []model.SocialMedia{}, err
	}
	return socials, nil
}

func (sr *SocialRepositoryImpl) FindByID(socialID string) (model.SocialMedia, error) {
	social := model.SocialMedia{}

	err := sr.DB.Debug().Where("id = ?", socialID).Take(&social).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SocialMedia{}, err
		}

		return model.SocialMedia{}, err
	}

	return social, nil
}

func (sr *SocialRepositoryImpl) FindByUserID(userID string) ([]model.SocialMedia, error) {
	socials := []model.SocialMedia{}

	err := sr.DB.Debug().Where("user_id = ?", userID).Find(&socials).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.SocialMedia{}, err
		}

		return []model.SocialMedia{}, err
	}

	return socials, nil
}

func (sr *SocialRepositoryImpl) Update(socialReqData model.SocialMedia) error {
	err := sr.DB.Save(&model.SocialMedia{
		ID:             socialReqData.ID,
		Name:           socialReqData.Name,
		SocialMediaURL: socialReqData.SocialMediaURL,
		UserID:         socialReqData.UserID,
		UpdatedAt:      socialReqData.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (sr *SocialRepositoryImpl) Delete(socialReqData model.SocialMedia) error {
	err := sr.DB.Delete(&socialReqData).Error
	if err != nil {
		return err
	}
	return nil
}
