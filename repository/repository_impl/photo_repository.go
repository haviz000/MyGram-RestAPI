package repository_impl

import (
	"errors"

	"github.com/haviz000/MyGram-RestAPI/model"
	"github.com/haviz000/MyGram-RestAPI/repository"
	"gorm.io/gorm"
)

type PhotoRepositoryImpl struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) repository.PhotoRepository {
	return &PhotoRepositoryImpl{
		DB: db,
	}
}

func (pr *PhotoRepositoryImpl) Create(photoReqData model.Photo) error {
	err := pr.DB.Create(&photoReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (pr *PhotoRepositoryImpl) FindAll() ([]model.Photo, error) {
	photos := []model.Photo{}

	err := pr.DB.Find(&photos).Error
	if err != nil {
		return []model.Photo{}, err
	}

	return photos, nil
}

func (pr *PhotoRepositoryImpl) FindByID(photoID string) (model.Photo, error) {
	photo := model.Photo{}

	err := pr.DB.Debug().Where("photo_id = ?", photoID).Take(&photo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Photo{}, err
		}

		return model.Photo{}, err
	}

	return photo, nil
}

func (pr *PhotoRepositoryImpl) FindByUserID(userID string) ([]model.Photo, error) {
	photos := []model.Photo{}

	err := pr.DB.Where("user_id = ?", userID).Find(&photos).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Photo{}, err
		}

		return []model.Photo{}, err
	}

	return photos, nil
}

func (pr *PhotoRepositoryImpl) Update(photoReqData model.Photo) error {
	err := pr.DB.Save(&model.Photo{
		PhotoID:   photoReqData.PhotoID,
		Title:     photoReqData.Title,
		Caption:   photoReqData.Caption,
		PhotoURL:  photoReqData.PhotoURL,
		UserID:    photoReqData.UserID,
		UpdatedAt: photoReqData.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *PhotoRepositoryImpl) Delete(photoReqData model.Photo) error {
	err := pr.DB.Delete(&photoReqData).Error
	if err != nil {
		return err
	}

	return nil
}
