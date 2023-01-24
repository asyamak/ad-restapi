package usecase

import "ad-api/internal/repository"

type AdsUsecase interface{}


type AdUsecase struct{
	Repository repository.CreateAds
}


func NewAdUsecase(r repository.CreateAds)*AdUsecase{
	return &AdUsecase{
		
		Repository: r,
	}
}