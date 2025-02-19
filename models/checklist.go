package models

type Checklist struct {
	ID        int    `json:"id" gorm:"primary_key"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Items     []Item `json:"items"`
}

type CreateChecklistReq struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
}

type UpdateChecklistReq struct {
	Title string `json:"title"`
}

type ChecklistRes struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Items     []Item `json:"items,omitempty"`
}
