package controllers

import "net/http"

func SetUpRouter(h *Handler) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/ads", h.GetAds)
	router.HandleFunc("/ad", h.GetOneAd)
	router.HandleFunc("/ad/create", h.CreateAd)
	router.HandleFunc("/ad/delete", h.DeleteAd)

	return router
}
