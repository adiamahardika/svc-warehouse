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
	masterController := controller.MasterCategoryController(masterCategoryService)

	root := router.Group("/")
	{
		masterCategory := root.Group("/master-category")
		{
			masterCategory.POST("/", masterController.CreateMasterCategory)
			masterCategory.GET("/", masterController.ReadMasterCategory)
		}
	}

	error := server.ListenAndServe()
	if error != nil {
		fmt.Println(error)
	}
}
