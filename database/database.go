package database

import (
	"gorm.io/gorm"
	"log"
)

// Seed creates the tables in the DB
func Seed(db *gorm.DB) error {
	log.Println("seeding db...")

	// drop tables
	for _, table := range SeedTables {
		q := table.Down

		log.Printf("%s", q)

		err := db.Exec(q).Error
		if err != nil {
			return err
		}
	}

	// create tables
	for _, table := range SeedTables {
		q := table.Up

		log.Printf("%s", q)

		err := db.Exec(q).Error
		if err != nil {
			return err
		}
	}

	// (optional) add dummy data to db for testing
	//for _, table := range SeedTables {
	//	q := table.SeedDummy
	//
	//	log.Printf("%s", q)
	//
	//	err := db.Exec(q).Error
	//	if err != nil {
	//		return err
	//	}
	//}

	log.Println("seeding done!")

	return nil
}
