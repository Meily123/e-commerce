package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	Id          uuid.UUID            `json:"id" gorm:"primaryKey; default:uuid_generate_v4(); not null;"`
	UserId      uuid.UUID            `json:"user_id" gorm:"not null"`
	Products    []TransactionProduct `json:"products" gorm:"foreignKey:TransactionId; references:Id"`
	TotalSum    int                  `json:"total_sum"`
	TotalMargin int                  `json:"total_margin"`
	TotalItem   int                  `json:"total_item"`
	IsPaid      bool                 `json:"is_paid" gorm:"default:false;"`
	CreatedAt   int64                `json:"created_at" gorm:"autoCreatedTime:nano;"`
	PaidAt      int64                `json:"paid_at"`
	IsActive    bool                 `json:"is_active" gorm:"default:false"`
}

type TransactionProduct struct {
	Id            uuid.UUID `json:"id"  gorm:"primaryKey; default:uuid_generate_v4(); not null;"`
	TransactionId uuid.UUID `json:"transaction_id"  gorm:"not null"`
	Item          Product   `json:"item"  gorm:"not null; constraint:OnDelete:CASCADE"`
	ItemId        uuid.UUID `json:"item_id"  gorm:"not null"`
	Quantity      int       `json:"quantity"  gorm:"not null"`
	Sum           int       `json:"sum"  gorm:"not null"`
	Margin        int       `json:"margin" gorm:"not null"`
}

func (t *Transaction) BeforeCreate(*gorm.DB) (err error) {
	t.Id = uuid.New()
	return
}

func (tp *TransactionProduct) BeforeCreate(*gorm.DB) (err error) {
	tp.Id = uuid.New()
	return
}
