package usecase

import (
	"ad-api/internal/controllers"
	"ad-api/internal/entity"
	"ad-api/internal/repository"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type AdsUsecase interface{
	CreateAd(ad entity.Ad)(string,error)
	GetAds(search entity.Search)
}


type AdUsecase struct{
	Repository repository.CreateAds
}


func NewAdUsecase(r repository.CreateAds)*AdUsecase{
	return &AdUsecase{
		Repository: r,
	}
}

func (u *AdUsecase)CreateAd(requestAd controllers.AdRequest)(string,error){
	var ad entity.Ad
	ad.Guid = uuid.NewString()

	err := validation(ad)
	if err != nil {
		return "",err
	}
	
	adId, err := u.Repository.CreateAd(ad)
	if err != nil {
		return  "",err
	}
	 
	photoId, err := u.Repository.AddPhotos(ad.Photos)
	if err != nil {
		return "",err
	}

	err = u.Repository.InsertAdPhotos(adId,photoId)
	if err != nil {
		return "",err
	}

	return ad.Guid, nil
}

var (
	ErrDiscriptionLength = errors.New("invalid discription length")
	ErrNameLength = errors.New("invalid name legth")
	ErrLinkNumber = errors.New("invalid link number")
)

func validation(ad entity.Ad) error{
	if len(ad.Description) > 1000{
		return ErrDiscriptionLength
	}
	if len(ad.Name) > 200{
		return ErrNameLength
	}

	// if len(ad.Photos.Link) > 3{
	// 	return ErrLinkNumber
	// }
	return nil
}

func (u *AdUsecase)GetAds(search entity.Search){
	
	offset := (search.Page - 1) * 10
	
	if search.PricePreference == "asc"{
		ads,err := u.Repository.GetAdsAsc(search,offset)
		if err != nil {
			fmt.Printf("error getads: %v\n",err)
		}
		fmt.Println(ads)
	}



}