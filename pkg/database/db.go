package database

import (
	"ad-api/config"
	"database/sql"
	"fmt"
)

func New(cfg *config.Config) (*sql.DB,error){
	dbConfig := fmt.Sprintf("user=%s dbname=%s host=%s port=%s password=%s sslmode=%s", cfg.User, cfg.DBname, cfg.Hostname,cfg.Port, cfg.Password,cfg.Ssl)

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil{
		return nil , err
	}

	if err = createTable(db); err != nil{
		return nil, err
	}

	return db, nil
}

const adTable = `CREATE TABLE IF NOT EXISTS ad (
	id SERIAL PRIMARY KEY,
	guid uuid,
	name VARCHAR(200),
	description varchar(2000),
	price FLOAT
);`

const photos = `CREATE TABLE IF NOT EXISTS photos(
	id SERIAL PRIMARY KEY,
	guid uuid,
	link TEXT
);`

const ad_photos = `CREATE TABLE IF NOT EXISTS ad_photos(
	id SERIAL PRIMARY KEY,
	ad_id INTEGER REGERENCES ad(id) DELETE ON CASCADE,
	photos_id INTEGER REFERENCES photos(id) DELETE ON CASCADE
);`


func createTable(db *sql.DB) error{
	_, err := db.Exec(adTable,photos,ad_photos)
	if err != nil {
		return err
	}
	return nil
}