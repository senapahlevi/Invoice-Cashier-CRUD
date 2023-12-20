package main

import (
	"inv-system-go/databases"
	"inv-system-go/invoice"
	"inv-system-go/itemdetail"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := databases.SetDB()
	if err != nil {
		log.Fatal(err)
	}

	invoice.SetDatabase(db)
	itemdetail.SetDatabaseCalculate(db)

	router := gin.Default()
	api := router.Group("/api")

	api.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*", "http://localhost:3000"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
		// ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))

	api.POST("/create-invoice", invoice.CreateInvoice)
	api.DELETE("/detail-items/:id", invoice.DeleteDetailItems)
	api.GET("/invoice/:id", invoice.GetInvoiceID)
	api.GET("/invoice", invoice.GetInvoiceAll)
	api.GET("/invoice-indexing", invoice.GetIndexingInvoice)
	api.POST("/calculate", itemdetail.CalculateDetailItem)
	api.PUT("/update-invoice/:id", invoice.UpdateInvoice)

	router.Run(":8080")

}
