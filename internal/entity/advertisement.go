package entity

type Ad struct {
	Id          int
	Guid        string   `json:"guid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	Photos      []Photos `json:"photo"`
}

type Photos struct {
	Id   int
	Guid string
	Link string `json:"link"`
}

type Ad_photos struct {
	Id        int `json:"id"`
	Ad_id     int `json:"ad_id"`
	Photos_id int `json:"photo_id"`
}

type Search struct {
	Page            int    `json:"page"`
	PricePreference string `json:"price"`
	DatePreference  string `json:"date"`
}

type DisplayAds struct {
	Name  string  `json:"name"`
	Link  string  `json:"link"`
	Price float32 `json:"price"`
}
