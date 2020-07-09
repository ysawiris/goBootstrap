package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#bd-docs-nav > div.bd-toc-item.active > ul > li", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			link := e.Attr("href")
			// Print link
			fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			//Visit each link
			c.Visit("https://getbootstrap.com/docs/4.5/components/" + link)
		})

		// link := e.Attr("href")

		// // Print link
		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://getboostrap.com
	c.Visit("https://getbootstrap.com/docs/4.5/components/alerts/")
}
