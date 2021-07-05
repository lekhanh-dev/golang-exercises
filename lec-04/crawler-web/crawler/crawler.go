package crawler

import (
	"crawler-web/models"
	repo "crawler-web/repository"
	"crawler-web/storage"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func CrawlerDataForMovie(url string, db *storage.DB) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	src := string(body)
	space := regexp.MustCompile(`\s+`)
	src = space.ReplaceAllString(src, " ")

	regexMovieName := regexp.MustCompile(`<td class="titleColumn">.*?<a href=.*?>(.*?)</a>.*?</td>`)
	movienames := regexMovieName.FindAllStringSubmatch(src, -1)

	regexImage := regexp.MustCompile(`<td class="posterColumn">.*?<img src="(.*?)" width.*?</td>`)
	movieImages := regexImage.FindAllStringSubmatch(src, -1)

	regexYearManufacture := regexp.MustCompile(`<span class="secondaryInfo">(.*?)</span>`)
	movieYearManufactures := regexYearManufacture.FindAllStringSubmatch(src, -1)

	regexRating := regexp.MustCompile(`<td class="ratingColumn imdbRating">.*?ratings">(.*?)</strong>`)
	movieRatings := regexRating.FindAllStringSubmatch(src, -1)

	for i := 0; i < len(movienames); i++ {
		movie := models.Movie{Name: movienames[i][1], Image: movieImages[i][1], YearManufacture: movieYearManufactures[i][1]}
		rating, _ := strconv.ParseFloat(movieRatings[i][1], 32)
		movie.Rating = float32(rating)
		repo.CreateMovie(db, movie)
	}

}
