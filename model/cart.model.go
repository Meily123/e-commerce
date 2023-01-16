package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartProduct struct {
	Id       uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4(); not null"`
	Item     Product
	Quantity int
	UserId   uuid.UUID
	ItemId   uuid.UUID
}

func (c *CartProduct) BeforeCreate(*gorm.DB) (err error) {
	c.Id = uuid.New()
	return
}
