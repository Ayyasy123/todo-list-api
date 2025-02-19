package usecase

import (
	"github.com/Ayyasy123/todo-list-api/models"
	"github.com/Ayyasy123/todo-list-api/repository"
)

type ChecklistUsecase interface {
	CreateChecklist(checklist *models.Checklist) error
	GetChecklists() ([]models.Checklist, error)
	GetChecklistByID(id int) (*models.Checklist, error)
	UpdateChecklist(checklist *models.Checklist) error
	DeleteChecklist(id int) error
}

type checklistUsecase struct {
	checklistRepo repository.ChecklistRepository
}

func NewChecklistUsecase(checklistRepo repository.ChecklistRepository) ChecklistUsecase {
	return &checklistUsecase{checklistRepo: checklistRepo}
}

func (c *checklistUsecase) CreateChecklist(checklist *models.Checklist) error {
	return c.checklistRepo.CreateChecklist(checklist)
}

func (c *checklistUsecase) GetChecklists() ([]models.Checklist, error) {
	return c.checklistRepo.GetChecklists()
}

func (c *checklistUsecase) GetChecklistByID(id int) (*models.Checklist, error) {
	return c.checklistRepo.GetChecklistByID(id)
}

func (c *checklistUsecase) UpdateChecklist(checklist *models.Checklist) error {
	return c.checklistRepo.UpdateChecklist(checklist)
}

func (c *checklistUsecase) DeleteChecklist(id int) error {
	return c.checklistRepo.DeleteChecklist(id)
}
