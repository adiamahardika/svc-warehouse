package router

import (
	"fmt"
	"net/http"
	"svc-warehouse/controller"
	"svc-warehouse/repository"
	"svc-warehouse/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AllRouter(db *gorm.DB) {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "token", "request-by", "signature-key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	repository := repository.Repository(db)

	masterCategoryService := service.MasterCategoryService(repository)
	masterCategoryController := controller.MasterCategoryController(masterCategoryService)

	masterProductService := service.MasterProductService(repository)
	masterProductController := controller.MasterProductController(masterProductService)

	productStatusService := service.ProductStatusService(repository)
	productStatusController := controller.ProductStatusController(productStatusService)

	reservationStatusService := service.ReservationStatusService(repository)
	reservationStatusController := controller.ReservationStatusController(reservationStatusService)

	root := router.Group("/")
	{
		masterCategory := root.Group("/master-category")
		{
			masterCategory.POST("/", masterCategoryController.CreateMasterCategory)
			masterCategory.GET("/", masterCategoryController.ReadMasterCategory)
			masterCategory.PUT("/:id", masterCategoryController.UpdateMasterCategory)
			masterCategory.DELETE("/:id", masterCategoryController.DeleteMasterCategory)
		}

		masterProduct := root.Group("/master-product")
		{
			masterProduct.POST("/", masterProductController.CreateMasterProduct)
			masterProduct.GET("/", masterProductController.ReadMasterProduct)
			masterProduct.PUT("/:id", masterProductController.UpdateMasterProduct)
			masterProduct.DELETE("/:id", masterProductController.DeleteMasterProduct)
		}

		productStatus := root.Group("/product-status")
		{
			productStatus.POST("/", productStatusController.CreateProductStatus)
			productStatus.GET("/", productStatusController.ReadProductStatus)
			productStatus.PUT("/:id", productStatusController.UpdateProductStatus)
			productStatus.DELETE("/:id", productStatusController.DeleteProductStatus)
		}

		reservationStatus := root.Group("/reservation-status")
		{
			reservationStatus.POST("/", reservationStatusController.CreateReservationStatus)
			reservationStatus.GET("/", reservationStatusController.ReadReservationStatus)
			reservationStatus.PUT("/:id", reservationStatusController.UpdateReservationStatus)
			reservationStatus.DELETE("/:id", reservationStatusController.DeleteReservationStatus)
		}
	}

	error := server.ListenAndServe()
	if error != nil {
		fmt.Println(error)
	}
}
