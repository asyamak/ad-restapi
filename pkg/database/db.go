package database

import (
	"ad-api/config"
	"database/sql"

	_ "github.com/lib/pq"
)

func New(cfg *config.Config) (*sql.DB,error){
	
	// dbConfig := fmt.Sprintf("user=%s dbname=%s host=%s port=%s password=%s sslmode=%s", cfg.User, cfg.DBname, cfg.Hostname,cfg.Port, cfg.Password,cfg.Ssl)
	
	db, err := sql.Open("postgres", cfg.DatabaseUrl)
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
	guid TEXT,
	name VARCHAR(200),
	description varchar(2000),
	price FLOAT
);`

const photos = `CREATE TABLE IF NOT EXISTS photo(
	id SERIAL PRIMARY KEY,
	guid TEXT,
	link TEXT
);`

const ad_photos = `CREATE TABLE IF NOT EXISTS ad_photos(
	id SERIAL PRIMARY KEY,
	ad_id INTEGER REFERENCES ad(id) ON DELETE CASCADE,
	photo_id INTEGER REFERENCES photo(id) ON DELETE CASCADE 
);`


func createTable(db *sql.DB) error{
	tables := []string{adTable,photos,ad_photos}
	for _, table := range tables{
	_, err := db.Exec(table)
	if err != nil {
		return err
	}
}
	return nil
}