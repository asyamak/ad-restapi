package controllers

import "net/http"

func SetUpRouter(h *Handler) *http.ServeMux{
	router := http.NewServeMux()
	router.HandleFunc("/ad",h.GetAds)
	router.HandleFunc("/ad/{id}",h.GetOneAd)
	router.HandleFunc("/ad/create/",h.CreateAd)
	router.HandleFunc("/ad/delete/{id}",h.DeleteAd)
	return router
}