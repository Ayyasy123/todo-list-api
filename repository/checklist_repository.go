package repository

import (
	"github.com/Ayyasy123/todo-list-api/models"
	"gorm.io/gorm"
)

type ChecklistRepository interface {
	CreateChecklist(checklist *models.Checklist) error
	GetChecklists() ([]models.Checklist, error)
	GetChecklistByID(id int) (*models.Checklist, error)
	UpdateChecklist(checklist *models.Checklist) error
	DeleteChecklist(id int) error
}

type checklistRepository struct {
	db *gorm.DB
}

func NewChecklistRepository(db *gorm.DB) ChecklistRepository {
	return &checklistRepository{db: db}
}

func (r *checklistRepository) CreateChecklist(checklist *models.Checklist) error {
	return r.db.Create(checklist).Error
}

func (r *checklistRepository) GetChecklists() ([]models.Checklist, error) {
	var checklists []models.Checklist
	err := r.db.Find(&checklists).Error
	return checklists, err
}

func (r *checklistRepository) GetChecklistByID(id int) (*models.Checklist, error) {
	var checklist models.Checklist
	err := r.db.First(&checklist, id).Preload("Items").Error
	return &checklist, err
}

func (r *checklistRepository) UpdateChecklist(checklist *models.Checklist) error {
	return r.db.Save(checklist).Error
}

func (r *checklistRepository) DeleteChecklist(id int) error {
	return r.db.Delete(&models.Checklist{}, id).Error
}
