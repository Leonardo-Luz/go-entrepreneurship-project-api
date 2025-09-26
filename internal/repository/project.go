package repository

import (
	"github.com/google/uuid"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/model"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetAll() ([]model.Project, error)
	GetByID(id uuid.UUID) (*model.Project, error)
	Create(model *model.Project) error
	Update(model *model.Project) error
	Delete(id uuid.UUID) error
}

type projectRepository struct {
	database *gorm.DB
}

func NewProjectRepository(database *gorm.DB) ProjectRepository {
	return &projectRepository{database: database}
}

func (repo *projectRepository) GetAll() ([]model.Project, error) {
	var projects []model.Project

	if err := repo.database.Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

func (repo *projectRepository) GetByID(id uuid.UUID) (*model.Project, error) {
	var project *model.Project

	if err := repo.database.First(&project, id).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func (repo *projectRepository) Create(model *model.Project) error {
	return repo.database.Create(model).Error
}

func (repo *projectRepository) Update(model *model.Project) error {
	return repo.database.Save(model).Error
}

func (repo *projectRepository) Delete(id uuid.UUID) error {
	return repo.database.Delete(&model.Project{}, id).Error
}
