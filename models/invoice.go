package models

import "time"

type Invoices struct {
	InvoiceID      int          `gorm:"primaryKey;invoice_id" json:"invoice_id"`
	Subject        string       `gorm:"subject" json:"subject"`
	Status         string       `gorm:"status" json:"status"`
	IssuedDate     time.Time    `gorm:"issued_date" json:"issued_date"`
	DueDate        time.Time    `gorm:"due_date" json:"due_date"`
	SubTotal       float64      `gorm:"sub_total" json:"sub_total"`
	Tax            float64      `gorm:"tax" json:"tax"`
	GrandTotal     float64      `gorm:"grand_total" json:"grand_total"`
	DetailItemJSON string       `gorm:"detail_item_json" json:"detail_item_json"`
	Customer       string       `gorm:"customer" json:"customer"`
	Address        string       `gorm:"address" json:"address"`
	DetailItems    []DetailItem `gorm:"foreignKey:InvoiceID" json:"detail_items"`
}

func (Invoices) TableName() string {
	return "invoices"
}
