package handler

import (
	"net/http"
	"strconv"

	"github.com/Ayyasy123/todo-list-api/models"
	"github.com/Ayyasy123/todo-list-api/usecase"
	"github.com/gin-gonic/gin"
)

type ChecklistHandler struct {
	checklistUsecase usecase.ChecklistUsecase
}

func NewChecklistHandler(checklistUsecase usecase.ChecklistUsecase) ChecklistHandler {
	return ChecklistHandler{checklistUsecase: checklistUsecase}
}

func (ch *ChecklistHandler) CreateChecklist(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	var checklist models.CreateChecklistReq
	if err := c.ShouldBindJSON(&checklist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checklist.UserID = userID.(int)

	err := ch.checklistUsecase.CreateChecklist(&models.Checklist{UserID: checklist.UserID, Title: checklist.Title})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checklist created successfully"})
}

func (ch *ChecklistHandler) GetChecklists(c *gin.Context) {
	checklists, err := ch.checklistUsecase.GetChecklists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checklists})
}

func (ch *ChecklistHandler) GetChecklistByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	checklist, err := ch.checklistUsecase.GetChecklistByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checklist})
}

func (ch *ChecklistHandler) DeleteChecklist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = ch.checklistUsecase.DeleteChecklist(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checklist deleted successfully"})
}
