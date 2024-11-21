package storage

import (
	"context"
	"user/internal/domain/models"
)

func (s *Storage) Users(ctx context.Context) ([]models.User, error) {
	var users []models.User

	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Storage) User(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Storage) CreateUser(ctx context.Context, user *models.User) error {
	return s.db.Create(user).Error
}

func (s *Storage) UpdateUser(ctx context.Context, user *models.User, id string) error {
	return s.db.Model(user).Where("id = ?", id).Updates(user).Error
}

func (s *Storage) DeleteUser(ctx context.Context, id string) error {
	return s.db.Delete(&models.User{}, id).Error
}
