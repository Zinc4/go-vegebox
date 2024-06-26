package user

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	SaveOTP(otp OTP) (OTP, error)
	FindOTP(userID int, otp string) (OTP, error)
	UpdateUser(user User) (User, error)
	DeleteOTP(otp OTP) error
	DeleteUserOTP(userID int) error
	ExistingUser(email string) error
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) SaveOTP(otp OTP) (OTP, error) {
	err := r.db.Create(&otp).Error
	if err != nil {
		return otp, err
	}
	return otp, nil
}

func (r *repository) FindOTP(userID int, otp string) (OTP, error) {
	var otpData OTP
	err := r.db.Where("user_id = ? AND otp = ? AND expired_otp > ?", userID, otp, time.Now().Unix()).Find(&otpData).Error
	if err != nil {
		return otpData, err
	}
	return otpData, nil
}

func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.Model(&user).Updates(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteOTP(otp OTP) error {
	err := r.db.Delete(&otp).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteUserOTP(userID int) error {
	err := r.db.Where("user_id = ?", userID).Delete(&OTP{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ExistingUser(email string) error {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	result := r.db.Where("id = ?", user.ID).Updates(&user)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, gorm.ErrRecordNotFound
	}

	return user, nil
}
