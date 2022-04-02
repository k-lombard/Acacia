package database

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/k-lombard/Acacia/AcaciaApi/models"
)

func (db Database) GetAllLocations() (*models.GeolocationPositionList, error) {
	list := &models.GeolocationPositionList{}
	err := db.Conn.Find(&list.GeolocationPositions).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) AddGeolocationPosition(loc *models.GeolocationPosition) (models.GeolocationPosition, error) {
	geolocationPositionTemp := models.GeolocationPosition{}
	err := db.Conn.First(&geolocationPositionTemp, "sentry_id = ?", loc.SentryID).Error
	if err != nil {
		geolocationPositionOut := models.GeolocationPosition{}
		if errors.Is(err, gorm.ErrRecordNotFound) {

			errNew := db.Conn.Create(&loc).Error
			if errNew != nil {
				return geolocationPositionOut, errNew
			}
			err2 := db.Conn.First(&geolocationPositionOut, "sentry_id = ?", loc.SentryID).Error
			if err2 != nil {
				return geolocationPositionOut, err2
			}
			fmt.Println("New geolocation_position record created with sentryID and timestamp: ", geolocationPositionOut.SentryID, geolocationPositionOut.Timestamp)
			return geolocationPositionOut, nil
		}
		return geolocationPositionOut, err
	} else {
		geolocationPositionOut := models.GeolocationPosition{}

		err := db.Conn.Model(&geolocationPositionOut).Where("sentry_id = ?", loc.SentryID).Updates(map[string]interface{}{
			"Accuracy":  loc.Accuracy,
			"Latitude":  loc.Latitude,
			"Longitude": loc.Longitude,
		}).Error
		if err != nil {
			return geolocationPositionOut, err
		}
		fmt.Println("Geolocation_position record updated with locationID and timestamp: ", geolocationPositionOut.ID, geolocationPositionOut.Timestamp)
		return geolocationPositionOut, nil
	}
}

func (db Database) GetGeolocationPositionBySentryId(sentryId uuid.UUID) (models.GeolocationPosition, error) {
	loc := models.GeolocationPosition{}
	err := db.Conn.First(&loc, "sentry_id = ?", sentryId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return loc, ErrNoMatch
		}
		return loc, err
	} else {
		return loc, nil
	}
}

func (db Database) DeleteGeolocationPosition(sentryId uuid.UUID) error {
	out := models.GeolocationPosition{}
	err := db.Conn.Where("sentry_id = ?", sentryId).Delete(&out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNoMatch
		}
		return err
	}
	fmt.Println("GeolocationPosition deleted with sentryID: ", out.ID)
	return nil
}

func (db Database) UpdateGeolocationPosition(sentryId uuid.UUID, locData models.GeolocationPosition) (models.GeolocationPosition, error) {
	loc := models.GeolocationPosition{}
	loc2 := models.GeolocationPosition{}
	errTwo := db.Conn.First(&loc2, "sentry_id = ?", sentryId).Error
	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return loc, ErrNoMatch
		}
		return loc, errTwo
	}
	err := db.Conn.Model(&loc).Where("sentry_id = ?", sentryId).Updates(map[string]interface{}{
		"Accuracy":  locData.Accuracy,
		"Latitude":  locData.Latitude,
		"Longitude": locData.Longitude,
	}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return loc, ErrNoMatch
		}
		return loc, err
	}
	errOut := db.Conn.First(&loc, "sentry_id = ?", sentryId).Error
	if errOut != nil {
		if errors.Is(errOut, gorm.ErrRecordNotFound) {
			return loc, ErrNoMatch
		}
		return loc, errOut
	}
	return loc, nil
}
