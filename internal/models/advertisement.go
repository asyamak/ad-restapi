package models

type Ad struct{
	Id int `json:"id"`
	Guid []byte `json:"guid"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
}

type Photos struct{
	Id int `json:"id"`
	Guid []byte `json:"guid"`
	Link string `json:"link"`
}

type as_photos struct{
	Id int `json:"id"`
	Ad_id int `json:"ad_id"`
	Photos_id int `json:"photos_id"`
}