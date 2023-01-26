package repository

import (
	"database/sql"
	"fmt"
	"time"

	"ad-api/internal/entity"
)

type CreateAds interface {
	CreateAd(ad entity.Ad) error
	// AddPhotos(photos []entity.Photos, guid string) ([]int, error)
	// InsertAdPhotos(adId int, photoId []int) error
	GetAdsAsc(search entity.Search, offset int) ([]entity.Ad, error)
	DeleteAdById(guid string) error
}

type CreateAdRepository struct {
	db *sql.DB
}

func NewCreateAdRepository(db *sql.DB) *CreateAdRepository {
	return &CreateAdRepository{
		db: db,
	}
}

func (r *CreateAdRepository) CreateAd(ad entity.Ad) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO ad (name, guid, description, price, timestamp) VALUES ($1, $2, $3, $4, $5) RETURNING id;`)
	if err != nil {
		return fmt.Errorf("create ad: prepare: %w", err)
	}
	defer stmt.Close()

	var ad_id int
	if err := stmt.QueryRow(ad.Name, ad.Guid, ad.Description, ad.Price, time.Now()).Scan(&ad_id); err != nil {
		if err = tx.Rollback(); err != nil {
			return fmt.Errorf("create ad: rollback: %w", err)
		}
		return fmt.Errorf("create ad: query row: %w", err)
	}

	////
	stmt, err = tx.Prepare(`INSERT INTO photo (guid, link) VALUES ($1, $2) RETURNING id;`)
	if err != nil {
		fmt.Printf("error add photos %v\n", err)
		return fmt.Errorf("add photos: prepare: %w", err)
	}

	var links_id []int
	for i := range ad.Photos {
		var id int
		if err := stmt.QueryRow(ad.Guid, ad.Photos[i].Link).Scan(&id); err != nil {
			if err = tx.Rollback(); err != nil {
				return fmt.Errorf("create ad: rollback: %w", err)
			}
			return fmt.Errorf("add photo: query row: %w", err)
		}
		links_id = append(links_id, id)
	}

	////
	stmt, err = tx.Prepare(`INSERT INTO ad_photos (ad_id, photo_id) VALUES ($1, $2);`)
	if err != nil {
		return fmt.Errorf("insert ad photos: prepare: %w", err)
	}

	for i := range links_id {
		if _, err = stmt.Exec(ad_id, links_id[i]); err != nil {
			if err = tx.Rollback(); err != nil {
				return fmt.Errorf("create ad: rollback: %w", err)
			}
			return fmt.Errorf("insert ad photos: exec: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repository: create ad: commit: %w", err)
	}

	return nil
}

// GetAdsAsc function sorts all ads by price asc and date of creation asc
func (r *CreateAdRepository) GetAdsAsc(search entity.Search, offset int) ([]entity.Ad, error) {
	query := fmt.Sprintf("SELECT id, guid, name, description, price FROM ad ORDER BY price %s, timestamp %s LIMIT 10 OFFSET $1;", search.PricePreference, search.DatePreference)

	var ads []entity.Ad

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ad entity.Ad
		if err = rows.Scan(&ad.Id, &ad.Guid, &ad.Name, &ad.Description, &ad.Price); err != nil {
			return nil, err
		}

		ads = append(ads, ad)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ads, nil
}

func (r *CreateAdRepository) DeleteAdById(guid string) error {
	query := `DELETE FROM ad WHERE guid = $1;`
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(guid)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	query = `DELETE FROM photo WHERE guid = $1;`
	stmt, err = tx.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(guid)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
