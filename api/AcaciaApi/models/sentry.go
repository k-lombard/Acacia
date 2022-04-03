package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Sentry struct {
	ID            uuid.UUID `sql:",pk" json:"id" gorm:"type:uuid;primaryKey;default:null;"`
	GeolocationID uuid.UUID `json:"geolocation_id" gorm:"default:null"`
	Alias         string    `json:"alias" gorm:"default:null"`
}

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"access_uuid"`
	RefreshUuid  string `json:"refresh_uuid"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

type AccessDetails struct {
	AccessUuid string `json:"access_uuid"`
	SentryId   string `json:"sentry_id"`
}

type SentryList struct {
	Sentries []Sentry `json:"sentries"`
}

func (i *Sentry) Bind(r *http.Request) error {
	if i.Alias == "" {
		return fmt.Errorf("Alias is a required field.")
	}
	return nil
}

func (*SentryList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Sentry) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*TokenDetails) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
