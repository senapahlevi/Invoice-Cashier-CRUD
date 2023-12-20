package itemdetail

import (
	"encoding/json"
	"inv-system-go/databases"
	"inv-system-go/invoice"
	"inv-system-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDatabaseCalculate(database *databases.DB) {
	db = database.DB
}

func CalculateDetailItem(c *gin.Context) {
	var invoiceInput invoice.InvoicesRequest

	var detailItems []models.DetailItem
	var invoice models.Invoices
	err := c.ShouldBindJSON(&invoiceInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := json.Unmarshal([]byte(invoiceInput.DetailItemJSON), &detailItems); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i := range detailItems {
		detailItems[i].Amount = float64(detailItems[i].Quantity) * detailItems[i].UnitPrice
	}

	invoice.SubTotal = 0
	invoice.Tax = 0
	invoice.GrandTotal = 0

	for _, item := range detailItems {
		item.Amount = float64(item.Quantity) * item.UnitPrice
		invoice.DetailItems = append(invoice.DetailItems, item)
	}

	for _, item := range detailItems {
		invoice.SubTotal += item.Amount
	}

	invoice.Tax = invoice.SubTotal * 0.1
	invoice.GrandTotal = invoice.SubTotal - invoice.Tax

	c.JSON(http.StatusOK, gin.H{"data": invoice})
}
