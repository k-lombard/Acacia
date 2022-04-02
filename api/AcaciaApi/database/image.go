package database

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/k-lombard/Acacia/AcaciaApi/models"
)

func (db Database) GetAllImages() (*models.ImageList, error) {
	list := &models.ImageList{}
	if err := db.Conn.Find(&list.Images).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) AddImage(image *models.Image) (models.Image, error) {
	imageOut := models.Image{}
	if err := db.Conn.Create(&image).Error; err != nil {
		return imageOut, err
	}
	if err2 := db.Conn.Where("sentry_id = ?", image.SentryID).First(&imageOut).Error; err2 != nil {
		return imageOut, err2
	}

	fmt.Println("New sentry record created with ID ", imageOut.ID)
	return imageOut, nil
}

func (db Database) GetImagesBySentryId(sentryId uuid.UUID) (*models.ImageList, error) {
	list := &models.ImageList{}
	err := db.Conn.Where("sentry_id = ?", sentryId).Find(&list.Images).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return list, ErrNoMatch
		} else {
			return list, err
		}
	}
	return list, nil
}

func (db Database) GetImageById(imageId uuid.UUID) (models.Image, error) {
	image := models.Image{}
	err := db.Conn.First(&image, "id = ?", imageId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return image, ErrNoMatch
		} else {
			return image, err
		}
	}
	return image, nil
}

func (db Database) DeleteImage(imageId uuid.UUID) error {
	image := &models.Image{}
	err := db.Conn.Delete(&image, "id = ?", imageId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNoMatch
		} else {
			return err
		}
	}
	fmt.Println("Image deleted with ID: ", image.ID)
	return nil
}
