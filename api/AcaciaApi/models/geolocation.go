package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type GeolocationPosition struct {
	ID        uuid.UUID `sql:",pk" json:"id" gorm:"type:uuid;default:null;primaryKey;"`
	SentryID  uuid.UUID `sql:",fk" json:"sentry_id" gorm:"type:uuid;default:null;"`
	Accuracy  float64   `json:"accuracy" gorm:"default:null"`
	Latitude  float64   `json:"latitude" gorm:"default:null"`
	Longitude float64   `json:"longitude" gorm:"default:null"`
	Timestamp string    `json:"timestamp" gorm:"default:null"`
	Sentry    Sentry    `json:"sentry" gorm:"-"`
}

func (GeolocationPosition) TableName() string {
	return "geolocation"
}

type GeolocationPositionList struct {
	GeolocationPositions []GeolocationPosition `json:"geolocation_positions"`
}

func (i *GeolocationPosition) Bind(r *http.Request) error {
	if i.Latitude == 0 || i.Longitude == 0 || (i.SentryID).String() == "" {
		return fmt.Errorf("Accuracy, latitude, longitude, and id are required fields.")
	}
	return nil
}

func (*GeolocationPositionList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*GeolocationPosition) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
