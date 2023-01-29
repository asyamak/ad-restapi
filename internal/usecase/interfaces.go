package usecase

import "ad-api/internal/repository"

type AdService struct {
	AdsUsecase
}

func NewUsecase(r *repository.Repository) *AdService {
	return &AdService{
		AdsUsecase: NewAdUsecase(r.CreateAds),
	}
}
