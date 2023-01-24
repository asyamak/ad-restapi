package repository

import "database/sql"

type Repository struct{
	CreateAds

}

func NewAdRepository(db *sql.DB) *Repository{
	return &Repository{
		
		CreateAds: NewCreateAdRepository(db),
	}
}