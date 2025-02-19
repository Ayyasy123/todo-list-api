package repository

import (
	"github.com/Ayyasy123/todo-list-api/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item *models.Item) error
	GetItemsByChecklistID(checklistID int) ([]models.Item, error)
	UpdateItem(item *models.Item) error
	DeleteItem(id int) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) CreateItem(item *models.Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) GetItemsByChecklistID(checklistID int) ([]models.Item, error) {
	var items []models.Item
	err := r.db.Where("checklist_id = ?", checklistID).Find(&items).Error
	return items, err
}

func (r *itemRepository) UpdateItem(item *models.Item) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) DeleteItem(id int) error {
	return r.db.Delete(&models.Item{}, id).Error
}
