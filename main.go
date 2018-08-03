package main

import (
	"fmt"

	"./scraper"
)

// Crawl page to find images
func main() {
	result := scraper.Crawl("https://www.google.com/")
	fmt.Println(result)
}
