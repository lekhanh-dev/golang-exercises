package repository

import (
	"crawler-web/models"
	"crawler-web/storage"
	"log"
)

func GetMovie(db *storage.DB) {
	results, err := db.SQL.Query("SELECT * FROM movies")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var movie models.Movie
		err = results.Scan(&movie.ID, &movie.Image, &movie.Name, &movie.YearManufacture, &movie.YearManufacture)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(movie.Name)
	}
}

func CreateMovie(db *storage.DB, movie models.Movie) {
	insert, err := db.SQL.Query("INSERT INTO movies (image, name, year_manufacture, rating ) VALUES ( ?, ?, ?, ? )", movie.Image, movie.Name, movie.YearManufacture, movie.Rating)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
