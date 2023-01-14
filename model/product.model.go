package model

import "time"

type Product struct {
	Id          string
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BasePrice   int64
	SellPrice   int64
	Stock       int64
	Description string
	IsActive    bool
}

type ProductRequest struct {
	Name        string `json:"name"`
	BasePrice   int64  `json:"base_price"`
	SellPrice   int64  `json:"sell_price"`
	Stock       int64  `json:"stock"`
	Description string `json:"description"`
} // @name ProductRequest

type ProductResponse struct {
	Name        string `json:"name"`
	SellPrice   int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Description string `json:"description"`
} // @name ProductResponse
