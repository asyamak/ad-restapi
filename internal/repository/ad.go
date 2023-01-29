package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"ad-api/internal/entity"
)

type CreateAds interface {
	CreateAd(ad *entity.Ad) error
	GetAdsByPrice(pricePreference string, offset int) ([]entity.DisplayAds, error)
	GetAdsByDate(datePreference string, offset int) ([]entity.DisplayAds, error)
	DeleteAdById(guid string) error
	GetAdByGuid(guid string) (entity.DisplayAd, error)
}

type CreateAdRepository struct {
	db *sql.DB
}

func NewCreateAdRepository(db *sql.DB) CreateAds {
	return &CreateAdRepository{
		db: db,
	}
}

func (r *CreateAdRepository) CreateAd(ad *entity.Ad) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("repository: create ad: transaction begin: %w", err)
	}

	stmt, err := tx.Prepare(`INSERT INTO ad (name, guid, description, price, timestamp) VALUES ($1, $2, $3, $4, $5) RETURNING id;`)
	if err != nil {
		return fmt.Errorf("repository: create ad: prepare1: %w", err)
	}
	defer stmt.Close()

	var ad_id int
	if err := stmt.QueryRow(ad.Name, ad.Guid, ad.Description, ad.Price, time.Now()).Scan(&ad_id); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("repository: create ad: rollback1: %w", err)
		}
		return fmt.Errorf("repository: create ad: query row1: %w", err)
	}

	////
	stmt, err = tx.Prepare(`INSERT INTO photo (guid, link) VALUES ($1, $2) RETURNING id;`)
	if err != nil {
		fmt.Printf("error add photos %v\n", err)
		return fmt.Errorf("repository: create ad: prepare2: %w", err)
	}

	var links_id []int
	for i := range ad.Photos {
		var id int
		if err := stmt.QueryRow(ad.Guid, ad.Photos[i].Link).Scan(&id); err != nil {
			if err := tx.Rollback(); err != nil {
				return fmt.Errorf("repository: create ad: rollback2: %w", err)
			}
			return fmt.Errorf("repository: create ad: query row2: %w", err)
		}
		links_id = append(links_id, id)
	}

	////
	stmt, err = tx.Prepare(`INSERT INTO ad_photos (ad_id, photo_id) VALUES ($1, $2);`)
	if err != nil {
		return fmt.Errorf("repository: create ad: prepare3: %w", err)
	}

	for i := range links_id {
		if _, err := stmt.Exec(ad_id, links_id[i]); err != nil {
			if err := tx.Rollback(); err != nil {
				return fmt.Errorf("repository: create ad: rollback3: %w", err)
			}
			return fmt.Errorf("repository: create ad: exec: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("repository: create ad: commit: %w", err)
	}

	return nil
}

// GetAdsByPrice function sorts all ads by price
func (r *CreateAdRepository) GetAdsByPrice(pricePreference string, offset int) ([]entity.DisplayAds, error) {
	query := "SELECT guid, name, price FROM ad ORDER BY price LIMIT 10 OFFSET $1;"

	if pricePreference == "DESC" {
		query = "SELECT guid, name, price FROM ad ORDER BY price DESC LIMIT 10 OFFSET $1;"
	}

	var ads []entity.Ad
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("repository: get ads by price: transaction begin: %w", err)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("repository: get ads by price: prepare1: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(offset)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, fmt.Errorf("repository: get ads by price: rollback: %w", err)
		}
		return nil, fmt.Errorf("repository: get ads by price: query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var ad entity.Ad
		if err := rows.Scan(&ad.Guid, &ad.Name, &ad.Price); err != nil {
			return nil, fmt.Errorf("repository: get ads by price: rows scan: %w", err)
		}
		ads = append(ads, ad)
	}

	query = `SELECT link FROM photo WHERE guid = $1 LIMIT 1;`
	stmt, err = tx.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("repository: get ads by price: prepare2: %w", err)
	}

	var display_ads []entity.DisplayAds

	for _, w := range ads {
		var photo_link entity.Photos
		var display_ad entity.DisplayAds

		err := stmt.QueryRow(w.Guid).Scan(&photo_link.Link)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return nil, fmt.Errorf("repository: get ads by price: rollback1: %w", err)
			}
			return nil, fmt.Errorf("repository: get ads by price: query row1: %w", err)
		}

		display_ad.Name = w.Name
		display_ad.Price = w.Price
		display_ad.Link = photo_link.Link
		display_ads = append(display_ads, display_ad)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repository: get ads by price: rows error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("repository: get ads by price: commit: %w", err)
	}

	return display_ads, nil
}

// GetAdsByDate function sorts ads by date and displays them according page number
func (r *CreateAdRepository) GetAdsByDate(datePreference string, offset int) ([]entity.DisplayAds, error) {
	query := "SELECT guid, name, price FROM ad ORDER BY timestamp LIMIT 10 OFFSET $1;"

	if datePreference == "DESC" {
		query = "SELECT guid, name, price FROM ad ORDER BY timestamp DESC LIMIT 10 OFFSET $1;"
	}

	var ads []entity.Ad

	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("repository: get ads by date: transaction begin: %w", err)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("repository: get ads by date: prepare: %w", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(offset)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, fmt.Errorf("repository: get ads by date: rollback: %w", err)
		}
		return nil, fmt.Errorf("repository: get ads by date: query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var ad entity.Ad
		if err := rows.Scan(&ad.Guid, &ad.Name, &ad.Price); err != nil {
			return nil, fmt.Errorf("repository: get ads by date: rows scan: %w", err)
		}
		ads = append(ads, ad)
	}

	query = `SELECT link FROM photo WHERE guid = $1 LIMIT 1;`
	stmt, err = tx.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("repository: get ads by date: prepare1: %w", err)
	}

	var display_ads []entity.DisplayAds

	for _, w := range ads {

		var display_ad entity.DisplayAds
		var photo_link string

		err := stmt.QueryRow(w.Guid).Scan(&photo_link)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return nil, fmt.Errorf("repository: get ads by date: rollback1: %w", err)
			}
			log.Printf("link doesn't found\n")
			continue
		}

		display_ad.Name = w.Name
		display_ad.Price = w.Price
		display_ad.Link = photo_link
		display_ads = append(display_ads, display_ad)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repository: get ads by date: rows error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("repository: get ads by date: commit: %w", err)
	}

	return display_ads, nil
}

func (r *CreateAdRepository) DeleteAdById(guid string) error {
	query := `DELETE FROM ad WHERE guid = $1;`
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("repository: delete ad by id: transaction begin: %w", err)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return fmt.Errorf("repository: delete ad by id: prepare: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(guid)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("repository: delete ad by id: rollback: %w", err)
		}
		return fmt.Errorf("repository: delete ad by id: exec: %w", err)
	}

	query = `DELETE FROM photo WHERE guid = $1;`
	stmt, err = tx.Prepare(query)
	if err != nil {
		return fmt.Errorf("repository: delete ad by id: prepare1: %w", err)
	}

	_, err = stmt.Exec(guid)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("repository: delete ad by id: rollback1: %w", err)
		}
		return fmt.Errorf("repository: delete ad by id: exec1: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("repository: delete ad by id: commit: %w", err)
	}

	return nil
}

func (r *CreateAdRepository) GetAdByGuid(guid string) (entity.DisplayAd, error) {
	query := "SELECT name, description, price FROM ad WHERE guid = $1;"
	tx, err := r.db.Begin()
	if err != nil {
		return entity.DisplayAd{}, fmt.Errorf("repository: get ad by guid: transaction begin: %w", err)
	}

	var ad entity.DisplayAd

	stmt, err := tx.Prepare(query)
	if err != nil {
		return entity.DisplayAd{}, fmt.Errorf("repository: get ad by guid: prepare: %w", err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(guid).Scan(
		&ad.Name,
		&ad.Description,
		&ad.Price,
	)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return entity.DisplayAd{}, fmt.Errorf("repository: get ad by guid: rollback: %w", err)
		}
		return entity.DisplayAd{}, fmt.Errorf("repository: get ad by guid: query row: %w", err)
	}

	query = `SELECT link FROM photo WHERE guid = $1;`
	stmt, err = tx.Prepare(query)
	if err != nil {
		return entity.DisplayAd{}, fmt.Errorf("repository: get ad by guid: prepare: %w", err)
	}

	rows, err := stmt.Query(guid)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return entity.DisplayAd{}, fmt.Errorf("repository: get ad by guid: rollback: %w", err)
		}
		log.Printf("link doesn't found\n")

	}

	for rows.Next() {
		var photo_link entity.Links

		if err := rows.Scan(&photo_link.Link); err != nil {
			if err := tx.Rollback(); err != nil {
				return entity.DisplayAd{}, fmt.Errorf("repository: get ad by guid: rollback1: %w", err)
			}
			return entity.DisplayAd{}, fmt.Errorf("repository: get ad by guid: rows scan: %w", err)
		}
		ad.Links = append(ad.Links, photo_link)
	}

	if err := tx.Commit(); err != nil {
		return entity.DisplayAd{}, fmt.Errorf("repository: get ad by guid: commit: %w", err)
	}

	return ad, nil
}
