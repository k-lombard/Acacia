package database

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/k-lombard/Acacia/AcaciaApi/models"
)

func (db Database) GetAllSentries() (*models.SentryList, error) {
	list := &models.SentryList{}
	if err := db.Conn.Find(&list.Sentries).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) AddSentry(sentry *models.Sentry) (models.Sentry, error) {
	sentryOut := models.Sentry{}
	if err := db.Conn.Create(&sentry).Error; err != nil {
		return sentryOut, err
	}
	if err2 := db.Conn.Where("alias = ?", sentry.Alias).First(&sentryOut).Error; err2 != nil {
		return sentryOut, err2
	}

	fmt.Println("New sentry record created with ID ", sentryOut.ID)
	return sentryOut, nil
}

func (db Database) GetSentryById(sentryId uuid.UUID) (models.Sentry, error) {
	sentry := models.Sentry{}
	err := db.Conn.First(&sentry, "id = ?", sentryId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sentry, ErrNoMatch
		} else {
			return sentry, err
		}
	}
	return sentry, nil
}

func (db Database) DeleteSentry(sentryId uuid.UUID) error {
	sentry := &models.Sentry{}
	err := db.Conn.Delete(&sentry, "sentry_id = ?", sentryId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNoMatch
		} else {
			return err
		}
	}
	fmt.Println("Sentry deleted with ID: ", sentry.ID)
	return nil
}

func (db Database) UpdateSentry(sentryId uuid.UUID, sentryData models.Sentry) (models.Sentry, error) {
	sentry := models.Sentry{}
	err := db.Conn.Model(&sentry).Where("id = ?", sentryId).Updates(map[string]interface{}{
		"Alias":         sentryData.Alias,
		"GeolocationID": sentryData.GeolocationID,
	}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sentry, ErrNoMatch
		}
		return sentry, err
	}
	errFinal := db.Conn.First(&sentry, "id = ?", sentryId).Error
	if errFinal != nil {
		if errors.Is(errFinal, gorm.ErrRecordNotFound) {
			return sentry, ErrNoMatch
		}
		return sentry, errFinal
	}
	return sentry, nil
}
