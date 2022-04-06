package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type DB struct {
	*sqlx.DB
	*Configuration
}

func Open(filepath string) (*DB, error) {
	db, err := sqlx.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	configPath := os.Getenv("CONFIG_PATH")
	config := ParseConfig(configPath)

	return &DB{
		DB:            db,
		Configuration: &config,
	}, nil
}

func (db *DB) RunMigrations() error {
	_, err := db.Exec(db.Migrations)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func (db *DB) Get(id uint32) (bool, error) {
	var ids []int
	const q = `SELECT listingId FROM apartment WHERE listingId = ?`

	if err := db.Select(&ids, q, id); err != nil {
		log.Fatal(err)
	}

	return len(ids) > 0, nil
}

func (db *DB) Update(a *Apartment) error {
	_, err := db.Exec(db.UpdateQuery, a.Found, a.ListingId, a.ListingAgencyReference, a.IsSoldProperty, a.Region, a.CityName, a.Lon, a.Lat, a.Price, a.ChargesPrice, a.Caution, a.AgencyFee, a.PropertySubType, a.PublisherId, a.PublisherRemoteVisit, a.PublisherPhone, a.PublisherName, a.PublisherAthomeId, a.PropertySurface, a.BuildingYear, a.FloorNumber, a.BathroomsCount, a.BedroomsCount, a.BalconiesCount, a.GaragesCount, a.CarparkCount, a.HasLivingRoom, a.HasKitchen, a.Availability, a.MediaConcatString(), a.Description, a.Link, a.ListingId)
	return err
}

func (db *DB) UpdateToFalse(a *Apartment) error {
	_, err := db.Exec(`UPDATE apartment SET found = false WHERE listingId = ?`, a.ListingId)
	return err
}

func (db *DB) Insert(a *Apartment) error {
	_, err := db.Exec(db.InsertQuery, a.Found, a.ListingId, a.ListingAgencyReference, a.IsSoldProperty, a.Region, a.CityName, a.Lon, a.Lat, a.Price, a.ChargesPrice, a.Caution, a.AgencyFee, a.PropertySubType, a.PublisherId, a.PublisherRemoteVisit, a.PublisherPhone, a.PublisherName, a.PublisherAthomeId, a.PropertySurface, a.BuildingYear, a.FloorNumber, a.BathroomsCount, a.BedroomsCount, a.BalconiesCount, a.GaragesCount, a.CarparkCount, a.HasLivingRoom, a.HasKitchen, a.Availability, a.MediaConcatString(), a.Description, a.Link)
	return err
}

func (db *DB) Save(apartment *Apartment) error {
	exists, err := db.Get(apartment.ListingId)
	if err != nil {
		return err
	}

	if exists {
		if !apartment.Found {
			err = db.UpdateToFalse(apartment)
		} else {
			err = db.Update(apartment)
		}
	} else {
		err = db.Insert(apartment)
	}

	if err != nil {
		return err
	}

	return nil
}
