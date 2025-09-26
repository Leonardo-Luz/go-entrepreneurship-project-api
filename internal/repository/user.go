package repository

import (
	"github.com/google/uuid"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id uuid.UUID) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(model *model.User) error
	Update(model *model.User) error
	Delete(id uuid.UUID) error
}

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{database: database}
}

func (repo *userRepository) GetAll() ([]model.User, error) {
	var users []model.User

	if err := repo.database.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *userRepository) GetByID(id uuid.UUID) (*model.User, error) {
	var user *model.User

	if err := repo.database.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepository) GetByEmail(email string) (*model.User, error) {
	var user *model.User

	if err := repo.database.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepository) Create(model *model.User) error {
	return repo.database.Create(model).Error
}

func (repo *userRepository) Update(model *model.User) error {
	return repo.database.Save(model).Error
}

func (repo *userRepository) Delete(id uuid.UUID) error {
	return repo.database.Delete(&model.User{}, id).Error
}
