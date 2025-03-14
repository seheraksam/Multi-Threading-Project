package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gocolly/colly"
	"github.com/seheraksam/Multi-Threading-Project/models"
)

var C *colly.Collector
var Products []models.Product

func main() {
	C = colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)
	C.Visit("https://www.scrapingcourse.com/ecommerce")

	C.OnHTML("li.product", func(e *colly.HTMLElement) {

		// initialize a new Product instance
		product := models.Product{}

		// scrape the target data
		product.Url = e.ChildAttr("a", "href")
		product.Image = e.ChildAttr("img", "src")
		product.Name = e.ChildText(".product-name")
		product.Price = e.ChildText(".price")
		spew.Dump(Products)
		// add the product instance with scraped data to the list of products
		Products = append(Products, product)

	})
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

		for _, product := range Products {

			record := []string{
				product.Url,
				product.Image,
				product.Name,
				product.Price,
			}
			writer.Write(record)
		}

		defer writer.Flush()
	})
}
