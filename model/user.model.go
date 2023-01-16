package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	CreatedAt int64     `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt int64     `json:"updated_at" gorm:"autoUpdateTime:nano"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password" gorm:"unique"`
	Address   string    `json:"address"`
	IsAdmin   bool      `json:"is_admin" gorm:"default:false"`
}

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
} // @name UserRequest

type UserEditRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
} // @name UserEditRequest

type UserResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	IsAdmin  bool   `json:"is_admin"`
} // @name UserResponse

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
} //@name LoginRequest

func (u *User) BeforeCreate(*gorm.DB) (err error) {
	u.Id = uuid.New()
	return
}
