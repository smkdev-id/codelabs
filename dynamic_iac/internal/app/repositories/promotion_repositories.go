package repositories

import (
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
	"gorm.io/gorm"
)

type PromotionRepository interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}

type PromotionRepositoryImpl struct {
	db *gorm.DB
}

// NewPromotionRepository creates a new instance of PromotionRepository
func NewPromotionRepository(db *gorm.DB) PromotionRepository {
	return &PromotionRepositoryImpl{
		db: db,
	}
}

// CreatePromotion creates a new promotion in the database
func (r *PromotionRepositoryImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	// Implementasi kamu taruh disini
}

// GetAllPromotions throw all data that recorded in the database
func (r *PromotionRepositoryImpl) GetAllPromotions() ([]models.Promotion, error) {
	// Implementasi kamu taruh disini
}

// GetPromotionByPromotionID will throw data based on promotionID request
func (r *PromotionRepositoryImpl) GetPromotionbyPromotionID(PromotionID string) (models.Promotion, error) {
	// Implementasi kamu taruh disini
}

// UpdatePromotion will update data based on promotionID request
func (r *PromotionRepositoryImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	// Implementasi kamu taruh disini
}

// DeletePromotionByPromotionID will delete data based on promotionID request
func (r *PromotionRepositoryImpl) DeletePromotionbyPromotionID(promotionID string) error {
	// Implementasi kamu taruh disini
}
