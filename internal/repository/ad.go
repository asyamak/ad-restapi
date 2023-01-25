package repository

import (
	"ad-api/internal/entity"
	"database/sql"
	"fmt"
)

type CreateAds interface{
	CreateAd(ad entity.Ad)(int,error)
	AddPhotos(photos entity.Photos)(int,error)
	InsertAdPhotos(adId,photoId int)error
	GetAdsAsc(page,offset int) ([]entity.Ad, error)
}

type CreateAdRepository struct{
	db *sql.DB
}

func NewCreateAdRepository(db *sql.DB) *CreateAdRepository{
	return &CreateAdRepository{
		db: db,
	}
}

func(r *CreateAdRepository)CreateAd(ad entity.Ad)(int,error){
	query := `INSERT INTO ad (name, description, price) VALUES ($1, $2, $3) RETURNING id;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("create ad: prepare: %w",err)
	}
	// fmt.Println("in repository")
	var id int
	err = stmt.QueryRow(ad.Name,ad.Description,ad.Price).Scan(&id)
	if err != nil {
		return 0,fmt.Errorf("create ad: query row: %w",err)
	}	

	return id, nil
}

func(r *CreateAdRepository)AddPhotos(photos entity.Photos)(int,error){
	query := `INSERT INTO photo (guid,link) VALUES ($1, $2) RETURNING id;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		fmt.Printf("error add photos %v\n",err)
		return 0, fmt.Errorf("add photos: prepare: %w",err)
	}
	var id int

 	err = stmt.QueryRow(photos.Guid,photos.Link).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("add photo: query row: %w",err)
	}
	
	return id, nil
}

//InsertAdPhotos function insert ids of ad and photo link into ad_photos table
func(r *CreateAdRepository)InsertAdPhotos(adId,photoId int)error{
	query := `INSERT INTO ad_photos (ad_id, photo_id) VALUES ($1, $2);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("insert ad photos: prepare: %w",err)
	}

	_, err = stmt.Exec(adId,photoId)
	if err != nil {
		return fmt.Errorf("insert ad photos: exec: %w",err)
	}

	return nil
}

func(r *CreateAdRepository)GetAdsAsc(page,offset int) ([]entity.Ad, error){
	query := `SELECT id, guid, name, description, price FROM ads ORDER BY ASC LIMIT $1 OFFSET $2;`

	var ads []entity.Ad

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(page,offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		var ad entity.Ad
		err = rows.Scan(&ad.Id,&ad.Guid,&ad.Name,&ad.Description,&ad.Price)
		if err != nil {
			return nil, err
		}

		ads = append(ads, ad)

	}

	if err = rows.Err(); err != nil{
		return nil, err
	}

	return ads, nil
}
