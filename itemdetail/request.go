package itemdetail

type DetailItemRequest struct {
	ItemID    int     `gorm:"item_id" json:"item_id"`
	InvoiceID int     `gorm:"invoice_id" json:"invoice_id"`
	ItemName  string  `gorm:"item_name" json:"item_name"`
	ItemType  string  `gorm:"item_type" json:"item_type"`
	Quantity  int     `gorm:"quantity" json:"quantity"`
	UnitPrice float64 `gorm:"unit_price" json:"unit_price"`
	Amount    float64 `gorm:"amount" json:"amount"`
}
