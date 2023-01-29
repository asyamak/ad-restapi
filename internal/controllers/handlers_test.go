package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

// TestGetAds tests "/ads" handler by price preference
func TestGetAdsByPrice(t *testing.T) {
	ad := SearchInputRequest{
		Page:            1,
		PricePreference: "asc",
		DatePreference:  "",
	}
	obj, err := json.Marshal(&ad)
	if err != nil {
		t.Errorf("error test get ads by price: marshalling: %v\n", err)
	}

	resp, err := http.Post("http://localhost:9090/ads", "application/json", bytes.NewBuffer(obj))
	if err != nil {
		t.Errorf("error test get ads by price: http post: %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error test gets ads by price: status code difference: %v, %d\n", err, resp.StatusCode)
	}
}

// TestGetAdsByDate function testing "/ads" handler by date preference
func TestGetAdsByDate(t *testing.T) {
	ad := SearchInputRequest{
		Page:            1,
		PricePreference: "",
		DatePreference:  "desc",
	}
	obj, err := json.Marshal(&ad)
	if err != nil {
		t.Errorf("error test get ads: marshalling: %v\n", err)
	}

	resp, err := http.Post("http://localhost:9090/ads", "application/json", bytes.NewBuffer(obj))
	if err != nil {
		t.Errorf("error test get ads by date: http post: %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error test gets ads by date: status code difference: %v, %d\n", err, resp.StatusCode)
	}
}

// TestGetAd function testing "/ad" handler
func TestGetAd(t *testing.T) {
	id := GetId{
		Id: "8a6c0c14-03f1-429a-b434-97765822e0ff",
	}

	obj, err := json.Marshal(&id)
	if err != nil {
		t.Errorf("error: test get ad: marshalling: %v\n", err)
	}

	resp, err := http.Post("http://localhost:9090/ad", "application/json", bytes.NewBuffer(obj))
	if err != nil {
		t.Errorf("error: test get ad: http post : %v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("error: test get ad: status difference: %v, %d\n", err, resp.StatusCode)
	}
}

// TestCreateAd function testing "/ad/create"
func TestCreateAd(t *testing.T) {
	ad := AdRequest{
		Name:        "new ad",
		Description: "creating new advertisement",
		Price:       576.8,
		PhotoLinks: []string{
			"https://play-lh.googleusercontent.com/5HEkLqjWS6vJibQ9KIMRGHLwnCyTrcg2mNBH_i-VuOy6bc98EtR091teMYO8HKGcRz47",
			"https://play-lh.googleusercontent.com/3sdafaxcvsdasdLqjWS6vJidfdsfsOy6bc98EtR091teMYO8HKGcRz47",
			"https://play-lh.googleusercontent.com/9fdasfsfkLqjWS6vJiadfdfasdfadsfasfastgrtg35grecRz47",
		},
	}
	obj, err := json.Marshal(&ad)
	if err != nil {
		t.Errorf("error: test create ad: marshalling: %v\n", err)
	}
	resp, err := http.Post("http://localhost:9090/ad/create", "application/json", bytes.NewBuffer(obj))
	if err != nil {
		log.Printf("error: test create ad: http post: %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error: test create ad: status code diffrence: %v, %d\n", err, resp.StatusCode)
	}
}

// TestDeleteAd function testing "/ad/delete" handler
func TestDeleteAd(t *testing.T) {
	ad := DeleteAd{
		Id: "8a6c0c14-03f1-429a-b434-97765822e0ff",
	}

	obj, err := json.Marshal(&ad)
	if err != nil {
		t.Errorf("error: test delete ad: marshalling: %v\n", err)
	}

	resp, err := http.Post("http://localhost:9090/ad/create", "application/json", bytes.NewBuffer(obj))
	if err != nil {
		t.Errorf("error: test delete ad: http post: %v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("error: test delete ad: status code difference: %v, %d\n", err, resp.StatusCode)
	}
}
