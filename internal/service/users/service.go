package sevice

import (
	"context"
	"fmt"
	"log/slog"
	"user/internal/domain/models"
)

type UserStorage interface {
	User(ctx context.Context, id string) (*models.User, error)
	Users(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User, id string) error
	DeleteUser(ctx context.Context, id string) error
}

type UserService struct {
	storage UserStorage
	log     *slog.Logger
}

func New(storage UserStorage, log *slog.Logger) *UserService {
	log = log.With(slog.String("service", "user"))
	return &UserService{storage: storage, log: log}
}

func (s *UserService) User(ctx context.Context, id string) (*models.User, error) {
	const op = "User"
	log := s.log.With(slog.String("op", op))
	log.Info("start get user")

	user, err := s.storage.User(ctx, id)
	if err != nil {
		log.Error("failed get user", err.Error())
		return nil, fmt.Errorf("failed get user: %w", err)
	}

	log.Info("success get user")
	return user, nil
}

func (s *UserService) Users(ctx context.Context) ([]models.User, error) {
	const op = "Users"
	log := s.log.With(slog.String("op", op))
	log.Info("start get users")

	users, err := s.storage.Users(ctx)
	if err != nil {
		log.Error("failed get users", err.Error())
		return nil, fmt.Errorf("failed get users: %w", err)
	}

	log.Info("success get users")
	return users, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	const op = "CreateUser"
	log := s.log.With(slog.String("op", op))
	log.Info("start create user")

	if err := s.storage.CreateUser(ctx, user); err != nil {
		log.Error("failed create user", err.Error())
		return fmt.Errorf("failed create user: %w", err)
	}

	log.Info("success create user")
	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User, id string) error {
	const op = "UpdateUser"
	log := s.log.With(slog.String("op", op))
	log.Info("start update user")

	if err := s.storage.UpdateUser(ctx, user, id); err != nil {
		log.Error("failed update user", err.Error())
		return fmt.Errorf("failed update user: %w", err)
	}

	log.Info("success update user")
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	const op = "DeleteUser"
	log := s.log.With(slog.String("op", op))
	log.Info("start delete user")

	if err := s.storage.DeleteUser(ctx, id); err != nil {
		log.Error("failed delete user", err.Error())
		return fmt.Errorf("failed delete user: %w", err)
	}

	log.Info("success delete user")
	return nil
}
