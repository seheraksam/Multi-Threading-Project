package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/seheraksam/Multi-Threading-Project/database"
	"github.com/seheraksam/Multi-Threading-Project/models"
)

var C *colly.Collector

func init() {
	database.ConnecttoMongo()
}

func main() {
	err := database.Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDB'ye ping atılamadı: %v", err)
	}

	// Create a new Colly collector
	C = colly.NewCollector()

	// Define the URL you want to scrape
	url := "http://quotes.toscrape.com/page/1/"

	// Set up callbacks to handle scraping events
	C.OnHTML(".quote", func(e *colly.HTMLElement) {
		file, err := os.Create("productsare.csv")
		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		headers := []string{
			"Quote",
			"Author",
			"Tags",
		}
		writer.Write(headers)
		// Extract data from HTML elements
		quote := e.ChildText("span.text")
		author := e.ChildText("small.author")
		tags := e.ChildText("div.tags")

		// Clean up the extracted data
		quote = strings.TrimSpace(quote)
		author = strings.TrimSpace(author)
		tags = strings.TrimSpace(tags)
		for i := 0; i < 10; i++ {
			record := []string{
				quote,
				author,
				tags,
			}
			secord := models.Product{
				Quote:  quote,
				Author: author,
				Tags:   tags,
			}
			_, err = database.Client.Database("product").Collection("products").InsertOne(context.TODO(), secord)
			if err != nil {
				fmt.Println("Error inserting news:", err)
			}
			if err != nil {
				fmt.Println("Error inserting news:", err)
			}
			writer.Write(record)
		}

		defer writer.Flush()
		// Print the scraped data
		fmt.Printf("Quote: %s\nAuthor: %s\nTags: %s\n\n", quote, author, tags)
	})

	// Visit the URL and start scraping
	err = C.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}

func Scraper() {

	C.OnScraped(func(r *colly.Response) {

		// open the CSV file
		file, err := os.Create("productsare.csv")
		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		headers := []string{
			"Url",
			"Image",
			"Name",
			"Price",
		}
		writer.Write(headers)

	})
}
