package service

import (
	"cms/internal/models"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Login(username, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	if user.Status != 1 {
		return nil, errors.New("account is disabled")
	}

	if user.Password != password {
		return nil, errors.New("invalid username or password")
	}

	return &user, nil
}

func (s *UserService) GetList() ([]models.User, error) {
	var users []models.User
	err := s.db.Order("id ASC").Find(&users).Error
	return users, err
}

func (s *UserService) Create(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UserService) Update(user *models.User) error {
	updates := map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
		"status":   user.Status,
	}

	if user.Password != "" {
		updates["password"] = user.Password
	}

	return s.db.Model(user).Updates(updates).Error
}

func (s *UserService) Delete(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}

func (s *UserService) ChangePassword(id uint, currentPassword, newPassword string) error {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}

	if user.Password != currentPassword {
		return errors.New("current password is incorrect")
	}

	return s.db.Model(&user).Update("password", newPassword).Error
}

func (s *UserService) InitDefaultUser() error {
	var count int64
	s.db.Model(&models.User{}).Count(&count)
	if count > 0 {
		return nil
	}

	defaultUser := &models.User{
		Username: "admin",
		Password: "123456",
		Email:    "admin@example.com",
		Role:     "admin",
		Status:   1,
	}
	return s.db.Create(defaultUser).Error
}
