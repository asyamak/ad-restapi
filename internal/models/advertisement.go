package models

type Ad struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Photos []string `json:"photos"`
	Price float32 `json:"price"`
}