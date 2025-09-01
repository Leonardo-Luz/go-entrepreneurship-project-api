package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	GitRepo     string    `json:"git_repo" gorm:"nullable"`
	Description string    `json:"description" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	UserID uuid.UUID `json:"user_id" gorm:"not null"`
	User   User      `json:"user" gorm:"foreignKey:UserID;references:ID"`

	Contributors []User `json:"contributors,omitempty" gorm:"many2many:project_contributors;"`
}

func (u *Project) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID, err = uuid.NewUUID()
		if err != nil {
			return err
		}
	}
	return nil
}
