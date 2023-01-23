package usecase

import "ad-api/internal/repository"

type Usecase struct{
	AdsUsecaser 
}

func NewUsecase(r *repository.Repository) *Usecase{
	return &Usecase{
		AdsUsecaser: NewAdUsecase(r.CreateAds),
	}
}