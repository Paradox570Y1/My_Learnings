package routes

import (
	"MyProject/internal/repository"
	"MyProject/internal/services"
	"MyProject/internal/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	facilityRepo := repository.NewFacilityRepo(db)
	orderRepo := repository.NewOrderRepo(db)

	facilityService := services.NewFacilityService(facilityRepo)
	orderService := services.NewOrderService(orderRepo, facilityRepo)

	facilityHandler := handlers.NewFacilityHandler(facilityService)
	orderHandler := handlers.NewOrderHandler(orderService)

	facilities := r.Group("/facilities")
	{
		facilities.GET("", facilityHandler.GetAll)
		facilities.GET("/:code", facilityHandler.GetByCode)
		facilities.POST("", facilityHandler.Create)
	}

	orders := r.Group("/orders")
	{
		orders.GET("", orderHandler.GetAll)
		orders.GET("/:id", orderHandler.GetByID)
		orders.POST("", orderHandler.Create)
	}
}