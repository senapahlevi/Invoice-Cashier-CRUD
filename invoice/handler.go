package invoice

import (
	"encoding/json"
	"fmt"
	"inv-system-go/databases"
	"inv-system-go/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDatabase(database *databases.DB) {
	db = database.DB
}

func CreateInvoice(c *gin.Context) {
	var invoiceInput InvoicesRequest
	var invoice models.Invoices
	var detailItems []models.DetailItem
	err := c.ShouldBindJSON(&invoiceInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	invoice.Subject = invoiceInput.Subject
	invoice.Status = invoiceInput.Status
	invoice.Address = invoiceInput.Address
	invoice.Customer = invoiceInput.Customer
	invoice.DetailItemJSON = invoiceInput.DetailItemJSON
	issuedDate, err := time.Parse("2006-01-02", invoiceInput.IssuedDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issued date format"})
		return
	}
	dueDate, err := time.Parse("2006-01-02", invoiceInput.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due date format"})
		return
	}
	invoice.IssuedDate = issuedDate
	invoice.DueDate = dueDate
	// save detail item
	if err := json.Unmarshal([]byte(invoiceInput.DetailItemJSON), &detailItems); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var detailItemsAmount float64
	for _, item := range detailItems {

		detailItemsAmount = float64(item.Quantity) * item.UnitPrice

	}

	var SubTotal float64
	for _, calculateItem := range detailItems {
		// SubTotal += float64(calculateItem.Quantity) * calculateItem.Amount
		SubTotal += float64(calculateItem.Quantity) * detailItemsAmount
	}
	tax := SubTotal * 0.1
	grandTotal := SubTotal - tax
	invoice.SubTotal = SubTotal
	invoice.Tax = tax
	invoice.GrandTotal = grandTotal
	result := db.Create(&invoice)

	for _, item := range detailItems {
		// biar bisa save id invoice id saat
		invoiceItem := models.DetailItem{
			InvoiceID: invoice.InvoiceID,
			ItemName:  item.ItemName,
			ItemType:  item.ItemType,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
			Amount:    float64(item.Quantity) * item.UnitPrice,
		}
		detailItemsAmount = float64(item.Quantity) * item.UnitPrice

		db.Create(&invoiceItem)
	}
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": invoice})
}

func GetInvoiceID(c *gin.Context) {
	var invoice models.Invoices
	id := c.Param("id")
	result := db.Preload("DetailItems").Where("invoice_id = ?", id).First(&invoice)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": invoice})
}

func GetInvoiceAll(c *gin.Context) {
	var invoice []models.Invoices
	result := db.Find(&invoice)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": invoice})
}

func DeleteDetailItems(c *gin.Context) {
	var detailItems models.DetailItem
	id := c.Param("id")
	result := db.Where("item_id = ?", id).First(&detailItems)
	fmt.Println("hello id", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "detail items not found"})
		return
	}

	// result = db.Delete(&detailItems)
	result = db.Where("item_id = ?", id).Delete(&detailItems)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "success deletes data"})

}

// indexing frontend

func GetIndexingInvoice(c *gin.Context) {
	var invoice []models.Invoices
	var search *gorm.DB

	status := c.Query("status")
	if status != "" {
		search = db.Where("status = ?", status)
	}

	issuedDateStart := c.Query("issued_date")
	issuedDateEnd := c.Query("due_date")

	if issuedDateStart != "" && issuedDateEnd != "" {
		issuedDateStartParsed, err := time.ParseInLocation("2006-01-02", issuedDateStart, time.Local)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		issuedDateEndParsed, err := time.ParseInLocation("2006-01-02", issuedDateEnd, time.Local)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		search = db.Where("issued_date BETWEEN ? AND ?", issuedDateStartParsed, issuedDateEndParsed)
	}

	customer := c.Query("customer")
	if customer != "" {
		search = db.Where("customer LIKE ?", "%"+customer+"%")
	}
	subject := c.Query("subject")
	if subject != "" {
		search = db.Where("subject LIKE ?", "%"+subject+"%")
	}
	totalItem := c.Query("total_item")
	if totalItem != "" {
		search = db.Where("total_item = ?", totalItem)
	}
	description := c.Query("description")
	if description != "" {
		search = db.Where("description LIKE ?", "%"+description+"%")
	}

	result := search.Preload("DetailItems").Find(&invoice)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoice})
}

// update invoice
func UpdateInvoice(c *gin.Context) {
	var invoiceInput InvoicesRequest
	var invoice models.Invoices
	var detailItems []models.DetailItem

	err := c.ShouldBindJSON(&invoiceInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	invoiceID := c.Param("id")
	result := db.First(&invoice, invoiceID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}
	invoice.Subject = invoiceInput.Subject
	invoice.Status = invoiceInput.Status
	invoice.Address = invoiceInput.Address
	invoice.Customer = invoiceInput.Customer
	invoice.DetailItemJSON = invoiceInput.DetailItemJSON
	issuedDate, err := time.Parse("2006-01-02", invoiceInput.IssuedDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issued date format"})
		return
	}
	dueDate, err := time.Parse("2006-01-02", invoiceInput.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due date format"})
		return
	}
	invoice.IssuedDate = issuedDate
	invoice.DueDate = dueDate
	// delete dahulu
	db.Where("invoice_id = ?", invoiceID).Delete(&models.DetailItem{})
	// Simpan baru
	if err := json.Unmarshal([]byte(invoiceInput.DetailItemJSON), &detailItems); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var detailItemsAmount float64
	for _, item := range detailItems {
		detailItemsAmount = float64(item.Quantity) * item.UnitPrice

		db.Create(&models.DetailItem{
			InvoiceID: invoice.InvoiceID,
			ItemName:  item.ItemName,
			ItemType:  item.ItemType,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
			Amount:    float64(item.Quantity) * item.UnitPrice,
		})
	}

	// Update total invoice
	var SubTotal float64
	for _, calculateItem := range detailItems {
		SubTotal += float64(calculateItem.Quantity) * detailItemsAmount
	}
	tax := SubTotal * 0.1
	grandTotal := SubTotal - tax
	invoice.SubTotal = SubTotal
	invoice.Tax = tax
	invoice.GrandTotal = grandTotal
	db.Save(&invoice)

	c.JSON(http.StatusOK, gin.H{"data": invoice})

}
