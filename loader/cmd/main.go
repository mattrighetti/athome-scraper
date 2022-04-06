package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"athome_loader/database"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	db, _ := database.Open(dbPath)
	if err := db.RunMigrations(); err != nil {
		log.Fatal(err)
	}

	contentPath := os.Getenv("JSON_OUT")
	content, err := ioutil.ReadFile(contentPath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var apartments []database.Apartment
	err = json.Unmarshal(content, &apartments)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	for _, apartment := range apartments {
		if err = db.Save(&apartment); err != nil {
			log.Printf("%v", err)
		}
	}

	log.Printf("Done!")
}
