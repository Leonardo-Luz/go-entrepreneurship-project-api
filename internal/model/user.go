package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	Role      Role      `json:"role" gorm:"type:user_role;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Projects []Project `json:"projects,omitempty" gorm:"foreignKey:UserID"`

	Followers []User `json:"followers,omitempty" gorm:"many2many:user_followers;joinForeignKey:UserID;joinReferences:FollowerID"`
	Following []User `json:"following,omitempty" gorm:"many2many:user_followers;joinForeignKey:FollowerID;joinReferences:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID, err = uuid.NewUUID()
		if err != nil {
			return err
		}
	}

	u.Role = RoleUser

	return u.hashPassword()
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// Check if password field was changed
	return u.hashPassword()
}

func (u *User) hashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
