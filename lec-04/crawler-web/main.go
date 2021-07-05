package main

import (
	"crawler-web/crawler"
	"crawler-web/storage"
	"log"
	"os"
)

func main() {
	dbName := "crawler_web"
	dbPass := "123456"
	dbHost := "127.0.0.1"
	dbPort := "3306"

	connection, err := storage.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	url := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	crawler.CrawlerDataForMovie(url, connection)
}

// func main() {
// 	url := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
// 	res, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := io.ReadAll(res.Body)
// 	res.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	src := string(body)
// 	space := regexp.MustCompile(`\s+`)
// 	src = space.ReplaceAllString(src, " ")
// 	// regexpn := regexp.MustCompile(`(?s)<tdclass="titleColumn">(.*?)</td>`)
// 	regexpn := regexp.MustCompile(`<td class="titleColumn">.*?<a href=.*?>(.*?)</a>.*?</td>`)
// 	movienames := regexpn.FindAllStringSubmatch(src, -1)

// 	for _, v := range movienames {
// 		fmt.Println(v[1])
// 	}
// 	// sub := movienames[0][0]
// 	// regexnew := regexp.MustCompile(`<a href="*?>(.*?)</a>`)
// 	// sub1 := regexnew.FindAllStringSubmatch(sub, -1)
// 	// fmt.Println(sub1)
// 	// 	// fmt.Println(moviename[1])
// 	fmt.Println(len(movienames))
// 	// 	// WriteFile("output.txt", src)
// }

// func main() {
// 	str1 := `<html><body>
// 			<form name="query" action="http://www.example.net/action.php" method="post">
// 				<textarea type="text" name="nameiknow">The text I want</textarea>
// 				<div id="button">
// 					<input type="submit" value="Submit" />
// 				</div>
// 				<p><textarea type="text" name="nameiknow">Ok baby</textarea></p>
// 			</form>
// 			</body></html>`

// 	re := regexp.MustCompile(`<textarea.*?>(.*)</textarea>`)

// 	submatchall := re.FindAllStringSubmatch(str1, -1)
// 	for _, element := range submatchall {
// 		fmt.Println(element)
// 	}
// }
