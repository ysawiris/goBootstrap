package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// Boostrap global variable
type Boostrap struct {
	name string
	href string
	// snippet string
}

func main() {
	var arr []Boostrap

	arr = getBootstrapCompenonts(arr)
	getCodeSnippets(arr)

	// fmt.Printf("%+v\n", arr)

}

// Grab all the Boostrap Components
func getBootstrapCompenonts(data []Boostrap) []Boostrap {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has this selector
	c.OnHTML("#bd-docs-nav > div.bd-toc-item.active > ul > li", func(e *colly.HTMLElement) {
		// fmt.Printf(e.Text)
		e.ForEach("a[href]", func(_ int, e *colly.HTMLElement) {
			// saving the href to a variable
			link := e.Attr("href")

			string := strings.TrimSpace(e.Text)

			//creating struct
			b := Boostrap{name: string, href: link}

			// adding struct to slice
			data = append(data, b)

			fmt.Print("Data", data)

			// Print link
			fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		// Printing which url its visiting
		fmt.Println("Visiting for Compenonts", r.URL.String())
	})

	// Start scraping on https://getboostrap.com
	c.Visit("https://getbootstrap.com/docs/4.5/components/alerts/")

	// return the slice to send it to other functions
	return data
}

// Grabs all the code snippets for all the Boostrap Components
func getCodeSnippets(data []Boostrap) {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has this selector
	c.OnHTML("body > div > div > main > figure", func(e *colly.HTMLElement) {
		fmt.Printf(e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting for Code Snippets", r.URL.String())
	})

	for index, element := range data {
		fmt.Println(index)
		fmt.Printf("%+v\n", element)

		c.Visit("https://getbootstrap.com" + element.href)

		// On every a element which has this selector
		c.OnHTML("body > div > div > main > figure", func(e *colly.HTMLElement) {
			fmt.Printf(e.Text)
		})

	}
	// Start scraping on https://getboostrap.com
	c.Visit("https://getbootstrap.com/docs/4.5/components/alerts/")
}
