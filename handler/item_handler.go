package handler

import (
	"net/http"
	"strconv"

	"github.com/Ayyasy123/todo-list-api/models"
	"github.com/Ayyasy123/todo-list-api/usecase"
	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemHandler(itemUsecase usecase.ItemUsecase) ItemHandler {
	return ItemHandler{itemUsecase: itemUsecase}
}

func (i *ItemHandler) CreateItem(c *gin.Context) {
	var req models.CreateItemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := i.itemUsecase.CreateItem(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item created successfully", "data": res})
}

func (i *ItemHandler) GetItemsByChecklistID(c *gin.Context) {
	checklistID, err := strconv.Atoi(c.Param("checklist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist ID"})
		return
	}

	items, err := i.itemUsecase.GetItemsByChecklistID(checklistID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (i *ItemHandler) UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var req models.UpdateItemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := i.itemUsecase.UpdateItem(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully", "data": res})
}

func (i *ItemHandler) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	err = i.itemUsecase.DeleteItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
