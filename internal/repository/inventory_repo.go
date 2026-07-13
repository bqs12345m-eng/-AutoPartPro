package repository

import (
	"AutoPartPro/internal/models"
)

// InventoryRepository يحدد العمليات المتاحة للمخزون
type InventoryRepository struct {
	data map[string]models.Part
}

// NewInventoryRepository ينشئ نسخة جديدة من المخزن
func NewInventoryRepository() *InventoryRepository {
	return &InventoryRepository{
		data: make(map[string]models.Part),
	}
}

// AddPart يقوم بإضافة قطعة جديدة للمخزون
func (r *InventoryRepository) AddPart(p models.Part) {
	r.data[p.ID] = p
}

// GetPart يقوم بجلب قطعة بناءً على الـ ID
func (r *InventoryRepository) GetPart(id string) (models.Part, bool) {
	p, exists := r.data[id]
	return p, exists
}
