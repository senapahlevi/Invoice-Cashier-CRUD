package models

type DetailItem struct {
	ItemID int `gorm:"item_id" json:"item_id"`
	// InvoiceID int `gorm:"invoice_id" json:"invoice_id"`
	InvoiceID int `gorm:"invoice_id;foreignKey:InvoiceID"` // Tambahkan foreign key

	ItemName  string  `gorm:"item_name" json:"item_name"`
	ItemType  string  `gorm:"item_type" json:"item_type"`
	Quantity  int     `gorm:"quantity" json:"quantity"`
	UnitPrice float64 `gorm:"unit_price" json:"unit_price"`
	Amount    float64 `gorm:"amount" json:"amount"`
}

func (DetailItem) TableName() string {
	return "detail_items"
}
