package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.goodreads.com"),
	)

	var quotes []Quote
	pageNum := 1
	baseURL := "https://www.goodreads.com/author/quotes/22782.George_Carlin"

	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		quoteText := strings.TrimSpace(e.ChildText(".quoteText"))
		author := strings.TrimSpace(e.ChildText(".authorOrTitle"))

		quote := Quote{
			Text:   quoteText,
			Author: author,
		}
		quotes = append(quotes, quote)
	})

	c.OnHTML(".next_page", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		if nextPage != "" {
			pageNum++
			nextURL := fmt.Sprintf("%s?page=%d", baseURL, pageNum)
			c.Visit(nextURL)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(baseURL)

	jsonData, err := json.MarshalIndent(quotes, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("quotes.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("Scraped %d quotes from %d pages. Saved to quotes.json\n", len(quotes), pageNum)
}

