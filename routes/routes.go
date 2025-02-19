package routes

import (
	"github.com/Ayyasy123/todo-list-api/handler"
	"github.com/Ayyasy123/todo-list-api/middleware"
	"github.com/Ayyasy123/todo-list-api/repository"
	"github.com/Ayyasy123/todo-list-api/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(db *gorm.DB, r *gin.Engine) {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	r.POST("/register", userHandler.RegisterUser)
	r.POST("/login", userHandler.LoginUser)
}

func SetupChecklistRoutes(db *gorm.DB, r *gin.Engine) {
	checklistRepository := repository.NewChecklistRepository(db)
	checklistUsecase := usecase.NewChecklistUsecase(checklistRepository)
	checklistHandler := handler.NewChecklistHandler(checklistUsecase)

	checklistRoutes := r.Group("/checklists")
	checklistRoutes.Use(middleware.JWTAuth())
	{
		checklistRoutes.POST("/", checklistHandler.CreateChecklist)
		checklistRoutes.GET("/", checklistHandler.GetChecklists)
		checklistRoutes.GET("/:id", checklistHandler.GetChecklistByID)
		checklistRoutes.DELETE("/:id", checklistHandler.DeleteChecklist)
	}

}

func SetupItemRoutes(db *gorm.DB, r *gin.Engine) {
	itemRepository := repository.NewItemRepository(db)
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	itemHandler := handler.NewItemHandler(itemUsecase)

	itemRoutes := r.Group("/items")
	itemRoutes.Use(middleware.JWTAuth())
	{
		itemRoutes.POST("/", itemHandler.CreateItem)
		itemRoutes.GET("/", itemHandler.GetItemsByChecklistID)
		itemRoutes.PUT("/:id", itemHandler.UpdateItem)
		itemRoutes.DELETE("/:id", itemHandler.DeleteItem)
	}
}
