package usecase

import (
	"errors"
	"fmt"
	"strings"

	// "ad-api/internal/controllers"
	"ad-api/internal/entity"
	"ad-api/internal/repository"
)

var (
	ErrDiscriptionLength = errors.New("invalid discription length")
	ErrNameLength        = errors.New("invalid name legth")
	ErrLinkNumber        = errors.New("invalid link number")
)

type AdsUsecase interface {
	CreateAd(ad *entity.Ad) (string, error)
	GetAds(search *entity.Search) ([]entity.DisplayAds, error)
	DeleteById(guid string) error
	GetOneAdById(guid string) (entity.DisplayAd, error)
}

type AdUsecase struct {
	Repository repository.CreateAds
}

func NewAdUsecase(r repository.CreateAds) AdsUsecase {
	return &AdUsecase{
		Repository: r,
	}
}

func (u *AdUsecase) CreateAd(requestAd *entity.Ad) (string, error) {
	// requestAd.Guid = requestAd.Guid

	if err := validation(requestAd); err != nil {
		return "", fmt.Errorf("usecase: create ad: validation: %w", err)
	}

	if err := u.Repository.CreateAd(requestAd); err != nil {
		return "", err
	}

	return requestAd.Guid, nil
}

func validation(ad *entity.Ad) error {
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

var ErrWrongQuery = errors.New("invalid query request")

func (u *AdUsecase) GetAds(search *entity.Search) ([]entity.DisplayAds, error) {
	offset := (search.Page - 1) * 10

	var (
		ads []entity.DisplayAds
		err error
	)

	search.PricePreference = strings.ToUpper(search.PricePreference)
	search.DatePreference = strings.ToUpper(search.DatePreference)

	if search.DatePreference == "" && search.PricePreference != "" {
		ads, err = u.Repository.GetAdsByPrice(search.PricePreference, offset)
		if err != nil {
			return nil, fmt.Errorf("usecase: get ads: price preference: %w", err)
		}
	} else {
		ads, err = u.Repository.GetAdsByDate(search.DatePreference, offset)
		if err != nil {
			return nil, fmt.Errorf("usecase: get ads: date preference: %w", err)
		}
	}

	return ads, nil
}

var ErrUuidLength = errors.New("invalid length of uuid")

// DeleteById function receive guid of ad, checks it and deletes it accordingly
func (u *AdUsecase) DeleteById(guid string) error {
	if len(guid) != 36 {
		return fmt.Errorf("usecase: delete by id: %w", ErrUuidLength)
	}

	guid = strings.TrimSpace(guid)

	if err := u.Repository.DeleteAdById(guid); err != nil {
		return err
	}

	return nil
}

func (u *AdUsecase) GetOneAdById(guid string) (entity.DisplayAd, error) {
	if len(guid) != 36 {
		return entity.DisplayAd{}, fmt.Errorf("usecase: get one ad by id: %w", ErrUuidLength)
	}

	ad, err := u.Repository.GetAdByGuid(guid)
	if err != nil {
		return entity.DisplayAd{}, err
	}

	return ad, nil
}
