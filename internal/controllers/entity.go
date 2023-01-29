package controllers

type AdRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	PhotoLinks  []string `json:"photo_links"`
}

type GetId struct {
	Id string `json:"id"`
}

type SearchInputRequest struct {
	Page            int    `json:"page"`
	PricePreference string `json:"price"`
	DatePreference  string `json:"date"`
}

type DeleteAd struct {
	Id string `json:"id"`
}
