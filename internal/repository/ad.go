package repository

import (
	"database/sql"
	"fmt"
)

type CreateAds interface{}

type CreateAdRepository struct{
	db *sql.DB
}

func NewCreateAdRepository(db *sql.DB) *CreateAdRepository{
	return &CreateAdRepository{
		db: db,
	}
}

func(r *CreateAdRepository)CreateAd(links[]string,name,description string, price float32)error{
	query := `INSERT INTO ad (name, description,price)VALUES($1,$2,$3) RETURNING id;`

	
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("create ad: transaction begin: %w",err)
	}

	_, err = tx.Exec(query,name,description,price)
	if err != nil {
		if err = tx.Rollback(); err != nil{
		return fmt.Errorf("create ad: rollback: %w",err)
		}
		return fmt.Errorf("create ad: exec query 1: %w",err)
	}
	// var linksIds []int
	// query = `INSERT INTO photos (link) VALUES($1) RETURNING id;`
	// for _, link := range links{
	// 	linkId, err := tx.Exec(query,link)
	// 	if err != nil {
	// 		if err = tx.Rollback(); err != nil{
	// 			return fmt.Errorf("create ad: rollback: %w",err)
	// 			}
	// 			return fmt.Errorf("create ad: exec query 1: %w",err)
	// 	}
	// 	linksIds = append(linksIds, linkId)
	// }

	if err = tx.Commit(); err != nil{
		return fmt.Errorf("create ad: commit: %w",err)
	}
	
	return nil	
}
