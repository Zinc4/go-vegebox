package user

import "time"

type User struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `json:"name" gorm:"type:varchar(255)"`
	Password   string    `json:"password" gorm:"type:varchar(255)"`
	Email      string    `json:"email" gorm:"type:varchar(255);unique"`
	Avatar     string    `json:"avatar" form:"avatar" gorm:"type:varchar(255)"`
	Role       string    `json:"role" gorm:"type:varchar(255) default:user"`
	IsVerified bool      `json:"is_verified" gorm:"default:false"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
	DeleteAt   time.Time `json:"delete_at"`
}

type OTP struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId     int       `json:"user_id" gorm:"index;unique"`
	User       User      `json:"user" gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OTP        string    `json:"otp" gorm:"type:varchar(255)"`
	ExpiredOTP int64     `json:"expired_otp" gorm:"type:bigint"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}
