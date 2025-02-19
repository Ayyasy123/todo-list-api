package usecase

import (
	"github.com/Ayyasy123/todo-list-api/models"
	"github.com/Ayyasy123/todo-list-api/repository"
)

type ChecklistUsecase interface {
	CreateChecklist(checklist *models.Checklist) error
	GetChecklists() ([]models.ChecklistRes, error)
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

func (c *checklistUsecase) GetChecklists() ([]models.ChecklistRes, error) {
	checklists, err := c.checklistRepo.GetChecklists()
	if err != nil {
		return nil, err
	}

	var res []models.ChecklistRes
	for _, checklist := range checklists {
		res = append(res, models.ChecklistRes{
			ID:        checklist.ID,
			UserID:    checklist.UserID,
			Title:     checklist.Title,
			CreatedAt: checklist.CreatedAt,
			UpdatedAt: checklist.UpdatedAt,
			Items:     checklist.Items,
		})
	}

	return res, nil
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
