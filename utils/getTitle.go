package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/seheraksam/Multi-Threading-Project/initializers"
	"golang.org/x/net/html"
)

var ctx = context.Background()

// HTML içinden başlık çeken fonksiyon
func getTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	var title string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return title, nil
}

// Cache kontrolü yapıp başlık çeken fonksiyon
func FetchTitle(url string) (string, error) {
	// Önce Redis Cache'i kontrol et
	cachedTitle, err := initializers.RedisClient.Get(ctx, url).Result()
	if err == nil {
		fmt.Println("🔵 Cache'den alındı:", url)
		return cachedTitle, nil
	}

	// Cache'te yoksa Web Scraper kullan
	title, err := getTitle(url)
	if err != nil {
		return "", err
	}

	// Cache'e ekle (5 dakika boyunca)
	initializers.RedisClient.Set(ctx, url, title, 5*time.Minute)

	fmt.Println(" Web Scraper ile alındı:", url)
	return title, nil
}
