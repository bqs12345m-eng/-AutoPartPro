package models

// Part يمثل هيكل بيانات قطعة الغيار في نظامنا
type Part struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	SKU   string  `json:"sku"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
