package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Image struct {
	ID        uuid.UUID `sql:",pk" json:"id" gorm:"type:uuid;primaryKey;default:null;"`
	SentryID  uuid.UUID `sql:",fk" json:"sentry_id" gorm:"type:uuid;default:null"`
	Content   []byte    `json:"content" gorm:"default:null"`
	Timestamp string    `json:"timestamp" gorm:"default:null"`
}

type ImageList struct {
	Images []Image `json:"images"`
}

func (i *Image) Bind(r *http.Request) error {
	if i.Content == nil {
		return fmt.Errorf("Email and password are required fields.")
	}
	return nil
}

func (*ImageList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Image) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
