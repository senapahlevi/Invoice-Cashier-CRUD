package invoice

import (
	"inv-system-go/models"
)

type InvoicesRequest struct {
	InvoiceID      int                 `gorm:"primaryKey;invoice_id" json:"invoice_id"`
	Subject        string              `gorm:"subject" json:"subject"`
	Status         string              `gorm:"status" json:"status"`
	IssuedDate     string              `gorm:"issued_date" json:"issued_date"`
	DueDate        string              `gorm:"due_date" json:"due_date"`
	SubTotal       float64             `gorm:"sub_total" json:"sub_total"`
	Tax            float64             `gorm:"tax" json:"tax"`
	GrandTotal     float64             `gorm:"grand_total" json:"grand_total"`
	DetailItemJSON string              `gorm:"detail_item_json" json:"detail_item_json"`
	Customer       string              `gorm:"customer" json:"customer"`
	Address        string              `gorm:"address" json:"address"`
	DetailItems    []models.DetailItem `gorm:"foreignKey:InvoiceID" json:"detail_items"`
}
