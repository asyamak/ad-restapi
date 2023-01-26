package controllers

import (
	"ad-api/internal/entity"
	"ad-api/internal/usecase"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Handler struct{
	usecase *usecase.Usecase

}

func NewHandler(u *usecase.Usecase) *Handler{
	return &Handler{
		usecase: u,
	}

}




func(h *Handler) GetAds(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	request := entity.Search{}
	
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Printf("page request %v\n",request)
	h.usecase.GetAds(request)

	fmt.Println(request)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

}


func(h *Handler) GetOneAd(w http.ResponseWriter, r *http.Request){}


func(h *Handler) CreateAd(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"incorrect method", http.StatusMethodNotAllowed)
		return
	}

	request := entity.Ad{}
	
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println(r.Body)
	fmt.Printf("request body %v\n",request)

	id, err := h.usecase.CreateAd(request)
	if err != nil {
		if errors.Is(err, usecase.ErrDiscriptionLength)||
			errors.Is(err, usecase.ErrLinkNumber)||
			errors.Is(err,usecase.ErrNameLength){
				http.Error(w, "incorrect input", http.StatusBadRequest)
				return
			}
		fmt.Printf("error created ad %v",err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
		
	}

	fmt.Printf("id from created ad %v\n",id)
	w.WriteHeader(http.StatusCreated)

}


func(h *Handler) DeleteAd(w http.ResponseWriter, r *http.Request){
	if r.Method != "DELETE"{
		return
	}
	id := r.URL.Query().Get("id")
	fmt.Printf("id in delete handler %v", id)
}