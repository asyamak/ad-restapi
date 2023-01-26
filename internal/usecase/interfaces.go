package usecase

import "ad-api/internal/repository"

type Usecase struct {
	AdsUsecase
}

func NewUsecase(r *repository.Repository) *Usecase {
	return &Usecase{
		AdsUsecase: NewAdUsecase(r.CreateAds),
	}
}
