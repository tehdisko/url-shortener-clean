package usecase

import (
	"url-shortener-clean/internal/entity"
)

type repository interface {
	Save(*entity.URL) (*entity.URL, error)
	FindByID(*entity.URL) (*entity.URL, error)
}

type UrlUseCase struct {
	repo repository
}

func NewURLUseCase(r repository) *UrlUseCase {
	return &UrlUseCase{repo: r}
}

func (uc *UrlUseCase) Shorten(url *entity.URL) (*entity.URL, error) {
	return uc.repo.Save(url)
}

func (uc *UrlUseCase) Expand(url *entity.URL) (*entity.URL, error) {
	return uc.repo.FindByID(url)
}
