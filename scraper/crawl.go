package scraper

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Crawl page to find images
func Crawl(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Errors:", err)
	}

	defer resp.Body.Close() // Close the request after function is complete
	tokenizer := html.NewTokenizer(resp.Body)

	for {
		token := tokenizer.Next()

		switch {
		case token == html.ErrorToken:
			// End of the document, we're done
			return "None found"
		case token == html.StartTagToken:
			tokenType := tokenizer.Token()

			isAnchor := tokenType.Data == "img"
			if isAnchor {
				// Looping through through token attributes to check for src
				for i := 0; i < len(tokenType.Attr); i++ {
					srcTag := tokenType.Attr[i]
					if srcTag.Key == "src" {
						return srcTag.Val
					}
				}
			}
		}
	}
}
