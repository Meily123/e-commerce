package model

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartProduct struct {
	Id       uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4(); not null;constraint:OnDelete:CASCADE;"`
	Item     Product   `gorm:"constraint:OnDelete:CASCADE"`
	Quantity int
	UserId   uuid.UUID
	ItemId   uuid.UUID
}

type CartProductRequest struct {
	ItemId   string `json:"product_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
} // @name CartRequest

type CartProductEditRequest struct {
	Quantity int `json:"quantity" binding:"required"`
} // @name CartEditRequest

type CartProductResponse struct {
	ItemId   string `json:"product_id"`
	Quantity int    `json:"quantity"`
} // @name CartResponse

func (c *CartProduct) BeforeCreate(*gorm.DB) (err error) {
	c.Id = uuid.New()
	return
}

func (c CartProduct) EmptyProductStruct() error {
	empty := CartProduct{}
	if c == empty {
		return errors.New("error not found")
	}
	return nil
}

func (c CartProduct) IsThisUserCart(UserId string) error {
	if c.UserId.String() != UserId {
		return errors.New("error you don't have access")
	}
	return nil
}
