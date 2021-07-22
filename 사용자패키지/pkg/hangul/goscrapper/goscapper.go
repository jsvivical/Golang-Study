package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatal("could not create file, err : %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)
	c.OnHTML(".internship_meta", func(e *colly.HTMLElement)){
		writer.Writer([]string {
			e.ChildText("a"),
			e.ChildText("span"),
			})
		})

		for i := 0; i < 312; i++{
			fmt.Printf("Scrapping page : %d\n", i)
			c.Visit("https://internshala.com/internships/pages-" + strconv.Itoa())
		}

	}
}
