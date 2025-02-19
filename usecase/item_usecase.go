package usecase

import (
	"github.com/Ayyasy123/todo-list-api/models"
	"github.com/Ayyasy123/todo-list-api/repository"
)

type ItemUsecase interface {
	CreateItem(req *models.CreateItemReq) (*models.ItemRes, error)
	GetItemsByChecklistID(checklistID int) ([]models.Item, error)
	UpdateItem(id int, req *models.UpdateItemReq) (*models.ItemRes, error)
	DeleteItem(id int) error
}

type itemUsecase struct {
	itemRepo repository.ItemRepository
}

func NewItemUsecase(itemRepo repository.ItemRepository) ItemUsecase {
	return &itemUsecase{itemRepo: itemRepo}
}

func (r *itemUsecase) CreateItem(req *models.CreateItemReq) (*models.ItemRes, error) {
	item := &models.Item{
		ChecklistID: req.ChecklistID,
		Description: req.Description,
		Completed:   false,
	}

	err := r.itemRepo.CreateItem(item)
	if err != nil {
		return nil, err
	}

	return &models.ItemRes{
		ID:          item.ID,
		ChecklistID: item.ChecklistID,
		Description: item.Description,
		Completed:   item.Completed,
	}, nil
}

func (r *itemUsecase) GetItemsByChecklistID(checklistID int) ([]models.Item, error) {
	return r.itemRepo.GetItemsByChecklistID(checklistID)
}

func (r *itemUsecase) UpdateItem(id int, req *models.UpdateItemReq) (*models.ItemRes, error) {
	item := &models.Item{
		ID:          id,
		ChecklistID: req.ChecklistID,
		Description: req.Description,
		Completed:   req.Completed,
	}

	err := r.itemRepo.UpdateItem(item)
	if err != nil {
		return nil, err
	}

	return &models.ItemRes{
		ID:          item.ID,
		ChecklistID: item.ChecklistID,
		Description: item.Description,
		Completed:   item.Completed,
	}, nil
}

func (r *itemUsecase) DeleteItem(id int) error {
	return r.itemRepo.DeleteItem(id)
}
