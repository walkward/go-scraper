package scraper

import (
	"fmt"
	"testing"
)

func TestCrawl(t *testing.T) {
	result := Crawl("https://www.google.com/")
	if result != "/images/branding/googlelogo/1x/googlelogo_white_background_color_272x92dp.png" {
		t.Errorf("Bad result")
	}
	fmt.Println(result)
}
