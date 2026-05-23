package service

import (
	"context"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/repository"
)

type SettingService interface {
	GetAllSettings(ctx context.Context) ([]domain.Setting, error)
	GetSettingByKey(ctx context.Context, key string) (*domain.Setting, error)
	UpdateSetting(ctx context.Context, s *domain.Setting) error
}

type settingService struct {
	repo repository.SettingRepository
}

func NewSettingService(repo repository.SettingRepository) SettingService {
	return &settingService{repo}
}

func (s *settingService) GetAllSettings(ctx context.Context) ([]domain.Setting, error) {
	return s.repo.FindAll(ctx)
}

func (s *settingService) GetSettingByKey(ctx context.Context, key string) (*domain.Setting, error) {
	return s.repo.FindByKey(ctx, key)
}

func (s *settingService) UpdateSetting(ctx context.Context, setting *domain.Setting) error {
	return s.repo.Update(ctx, setting)
}
