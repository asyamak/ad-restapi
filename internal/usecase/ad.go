package usecase

import (
	"errors"
	"fmt"
	"strings"

	// "ad-api/internal/controllers"
	"ad-api/internal/entity"
	"ad-api/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrDiscriptionLength = errors.New("invalid discription length")
	ErrNameLength        = errors.New("invalid name legth")
	ErrLinkNumber        = errors.New("invalid link number")
)

type AdsUsecase interface {
	CreateAd(ad entity.Ad) (string, error)
	GetAds(search entity.Search) ([]entity.Ad, error)
	DeleteById(guid string) error
}

type AdUsecase struct {
	Repository repository.CreateAds
}

func NewAdUsecase(r repository.CreateAds) *AdUsecase {
	return &AdUsecase{
		Repository: r,
	}
}

func (u *AdUsecase) CreateAd(requestAd entity.Ad) (string, error) {
	requestAd.Guid = uuid.NewString()

	err := validation(requestAd)
	if err != nil {
		return "", err
	}

	err = u.Repository.CreateAd(requestAd)
	if err != nil {
		return "", err
	}

	return requestAd.Guid, nil
}

func validation(ad entity.Ad) error {
	if len(ad.Description) > 1000 {
		return ErrDiscriptionLength
	}
	if len(ad.Name) > 200 {
		return ErrNameLength
	}

	if len(ad.Photos) > 3 {
		return ErrLinkNumber
	}
	return nil
}

func (u *AdUsecase) GetAds(search entity.Search) ([]entity.Ad, error) {
	offset := (search.Page - 1) * 10
	var (
		ads []entity.Ad
		err error
	)

	search.PricePreference = strings.ToUpper(search.PricePreference)
	search.DatePreference = strings.ToUpper(search.DatePreference)

	ads, err = u.Repository.GetAdsAsc(search, offset)
	if err != nil {
		return nil, fmt.Errorf("error get ads: %w", err)
	}

	return ads, nil
}

var ErrUuidLength = errors.New("invalid length of uuid")

func (u *AdUsecase) DeleteById(guid string) error {
	if len(guid) > 16 || len(guid) < 16 {
		return ErrUuidLength
	}
	guid = strings.TrimSpace(guid)
	fmt.Println(guid)
	err := u.Repository.DeleteAdById(guid)
	if err != nil {
		return err
	}
	return nil
}
