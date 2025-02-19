package models

type Item struct {
	ID          int    `json:"id" gorm:"primary_key"`
	ChecklistID int    `json:"checklist_id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateItemReq struct {
	ChecklistID int    `json:"checklist_id"`
	Description string `json:"description"`
}

type UpdateItemReq struct {
	ChecklistID int    `json:"checklist_id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type ItemRes struct {
	ID          int    `json:"id"`
	ChecklistID int    `json:"checklist_id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
