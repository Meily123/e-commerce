package model

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id           uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4();not nul; constraint:OnDelete:CASCADE;"`
	Name         string    `json:"name" gorm:"not null"`
	CreatedAt    int64     `json:"created_at" gorm:"autoCreateTime:nano;not null"`
	UpdatedAt    int64     `json:"updated_at" gorm:"autoUpdateTime:nano;not null"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	BasePrice    int       `json:"base_price" gorm:"not null"`
	SellPrice    int       `json:"sell_price" gorm:"not null"`
	Stock        int       `json:"stock" gorm:"not null"`
	Descriptions string    `json:"descriptions"`
}

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	BasePrice   int    `json:"base_price" binding:"required"`
	SellPrice   int    `json:"sell_price" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Description string `json:"description"`
} // @name ProductRequest

type ProductEditRequest struct {
	Name        string `json:"name" binding:"required"`
	BasePrice   int    `json:"base_price" binding:"required"`
	SellPrice   int    `json:"sell_price" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Description string `json:"description"`
} // @name ProductEditRequest

type ProductResponse struct {
	Name        string `json:"name"`
	SellPrice   int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
} // @name ProductResponse

func (p *Product) BeforeCreate(*gorm.DB) (err error) {
	p.Id = uuid.New()
	return
}

func (s Product) EmptyProductStruct() error {
	empty := Product{}
	if s == empty {
		return errors.New("error not found")
	}
	return nil
}
