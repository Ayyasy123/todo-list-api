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

	r.POST("/register", userHandler.RegisterUser) //2. Api daftar baru
	r.POST("/login", userHandler.LoginUser)       // 1. Api login
}

func SetupChecklistRoutes(db *gorm.DB, r *gin.Engine) {
	checklistRepository := repository.NewChecklistRepository(db)
	checklistUsecase := usecase.NewChecklistUsecase(checklistRepository)
	checklistHandler := handler.NewChecklistHandler(checklistUsecase)

	checklistRoutes := r.Group("/checklists")
	checklistRoutes.Use(middleware.JWTAuth())
	{
		checklistRoutes.POST("/", checklistHandler.CreateChecklist)      //3. API untuk membuat checklist
		checklistRoutes.GET("/", checklistHandler.GetChecklists)         // 5. API untuk menampilkan checklist-checklist yang sudah dibuat
		checklistRoutes.GET("/:id", checklistHandler.GetChecklistByID)   // 6. API Detail Checklist (Berisi item-item to-do yang sudah dibuat)
		checklistRoutes.DELETE("/:id", checklistHandler.DeleteChecklist) // 4. API untuk menghapus checklist
	}
}

func SetupItemRoutes(db *gorm.DB, r *gin.Engine) {
	itemRepository := repository.NewItemRepository(db)
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	itemHandler := handler.NewItemHandler(itemUsecase)

	itemRoutes := r.Group("/items")
	itemRoutes.Use(middleware.JWTAuth())
	{
		itemRoutes.POST("/", itemHandler.CreateItem)           // 7. API untuk membuat item-item to-do di dalam checklist
		itemRoutes.GET("/", itemHandler.GetItemByID)           // 8. API detail item
		itemRoutes.PUT("/:id", itemHandler.UpdateItem)         // 9. API untuk mengubah item-item di dalam checklist
		itemRoutes.PATCH("/:id", itemHandler.UpdateItemStatus) // 10. API untuk mengubah status dari item di dalam checklist
		itemRoutes.DELETE("/:id", itemHandler.DeleteItem)      // 11. API untuk menghapus item dari checklist
	}
}
