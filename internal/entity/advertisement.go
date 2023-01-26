package entity

type Ad struct{
	Id int `json:"id"`
	Guid string `json:"guid"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	Photos Photos `json:"photo"`
	Date string `json:"date_creation"`
	Link1 string `json:"link1"`
	Link2 string `json:"link2"`
	Link3 string `json:"link3"`
}

type Photos struct{
	Id int `json:"id"`
	Guid string `json:"guid"`
	Link string `json:"link"`
}

type Ad_photos struct{
	Id int `json:"id"`
	Ad_id int `json:"ad_id"`
	Photos_id int `json:"photo_id"`
}

type Search struct{
	Page int `json:"page"`
	PricePreference string `json:"price"`
}
