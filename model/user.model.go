package model

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID     `json:"id" gorm:"primaryKey;default:uuid_generate_v4(); not null; constraint:OnDelete:CASCADE;"`
	Name      string        `json:"name" gorm:"not null"`
	CreatedAt int64         `json:"created_at" gorm:"autoCreateTime:nano; not null"`
	UpdatedAt int64         `json:"updated_at" gorm:"autoUpdateTime:nano; not null"`
	IsActive  bool          `json:"is_active" gorm:"default:true; not null"`
	Username  string        `json:"username" gorm:"unique; not null"`
	Email     string        `json:"email" gorm:"unique; not null"`
	Password  string        `json:"password" gorm:"unique; not null"`
	Address   string        `json:"address" gorm:"not null"`
	IsAdmin   bool          `json:"is_admin" gorm:"default:false; not null"`
	Cart      []CartProduct `json:"cart" gorm:"foreignKey:UserId;references:Id"`
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

func (u User) EmptyUserStruct() error {
	if u.Id.String() == "00000000-0000-0000-0000-000000000000" {
		return errors.New("error not found")
	}
	return nil
}
