package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"ad-api/internal/entity"
	"ad-api/internal/usecase"

	"github.com/google/uuid"
)

type Handler struct {
	adService *usecase.AdService
}

func NewHandler(uc *usecase.AdService) *Handler {
	return &Handler{
		adService: uc,
	}
}

func (h *Handler) GetAds(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request SearchInputRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ads, err := h.adService.GetAds(&entity.Search{
		Page:            request.Page,
		PricePreference: request.PricePreference,
		DatePreference:  request.DatePreference,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(ads); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetOneAd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect method", http.StatusMethodNotAllowed)
		return
	}
	var id GetId

	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	request, err := h.adService.GetOneAdById(id.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(request); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateAd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect method", http.StatusMethodNotAllowed)
		return
	}

	var request AdRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var photosTemp []entity.Photos
	for _, link := range request.PhotoLinks {
		photosTemp = append(photosTemp, entity.Photos{
			Link: link,
		})
	}

	id, err := h.adService.CreateAd(&entity.Ad{
		Name:        request.Name,
		Guid:        uuid.New().String(),
		Description: request.Description,
		Price:       request.Price,
		Photos:      photosTemp,
	})
	if err != nil {
		fmt.Printf("error in handler create ad: %v\n", err)
		if errors.Is(err, usecase.ErrDiscriptionLength) ||
			errors.Is(err, usecase.ErrLinkNumber) ||
			errors.Is(err, usecase.ErrNameLength) {
			http.Error(w, "incorrect input", http.StatusBadRequest)
			return
		}
		fmt.Printf("error created ad %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(id); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) DeleteAd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "incorrect method", http.StatusMethodNotAllowed)
		return
	}

	var request DeleteAd

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.adService.DeleteById(request.Id)
	if err != nil {
		if errors.Is(err, usecase.ErrUuidLength) {
			log.Printf("error incorrect guid")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(request); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
