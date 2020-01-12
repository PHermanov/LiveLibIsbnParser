package main

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("parameter required, example: https://www.livelib.ru/reader/user_name/read/listview/biglist~")
	}

	var myUrl = os.Args[1]
	const fileName = "out.txt"

	fileContent := ""

	for i := 1; ; i++ {
		var currentPageUrl = myUrl + strconv.Itoa(i)

		response, err := http.Get(currentPageUrl)
		if err != nil {
			log.Fatal(err)
		}

		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal("error loading HTTP response body. ", err)
		}

		if len(document.Find("span[itemprop='isbn']").Nodes) == 0 {
			break
		}

		document.Find("span[itemprop='isbn']").Each(func(index int, element *goquery.Selection) {
			html, _ := element.Html()
			if html != "" {
				fileContent += html + "\n"
			}
		})

		response.Body.Close()
	}

	ioutil.WriteFile(fileName, []byte(fileContent), 0644)

	log.Println("file successfully created")
}
