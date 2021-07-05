package models

type Movie struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Image           string  `json:"image"`
	YearManufacture string  `json:"year_manufacture"`
	Rating          float32 `json:"rating"`
}
