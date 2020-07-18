package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// Boostrap global variable
type Boostrap struct {
	name    string
	href    string
	snippet string
}

func main() {
	// Initialize a slice to hold structs
	var arr []Boostrap
	// Collect all the Boostrap Components
	arr = getBootstrapCompenonts(arr)
	// Collet all the code snippets for each struct in slice
	arr = getCodeSnippets(arr)
	// write slice in to a text file
	writeFile(arr)
	// fmt.Printf("%+v\n", arr2)
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
			link := strings.TrimSpace(e.Attr("href"))

			text := strings.TrimSpace(e.Text)

			//creating struct
			b := Boostrap{name: text, href: link, snippet: ""}

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

	// return the slice to save it globally
	return data
}

// Grabs all the code snippets for all the Boostrap Components
func getCodeSnippets(data []Boostrap) []Boostrap {
	for index, element := range data {
		// Instantiate default collector
		c := colly.NewCollector()
		// fmt.Println(index)
		// fmt.Printf("%+v\n", element)

		// On every a element which has this selector
		c.OnHTML("body > div > div > main", func(e *colly.HTMLElement) {
			// for each code snippet, add it to the relative index
			e.ForEach("figure", func(_ int, e *colly.HTMLElement) {
				// fmt.Printf(e.Text)
				data[index].snippet = data[index].snippet + e.Text
				fmt.Printf(data[index].snippet)
			})
		})
		// Limit the number of threads started by colly to two
		// when visiting links which domains' matches "*httpbin.*" glob
		c.Limit(&colly.LimitRule{
			RandomDelay: 30 * time.Second,
		})

		// Before making a request print "Visiting ..."
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting for Code Snippets", r.URL.String())
		})
		// Start scraping on https://getboostrap.com
		c.Visit("https://getbootstrap.com" + element.href)
	}
	// return the slice to save globally
	return data
}

func writeFile(data []Boostrap) error {
	// set file path
	filePath := "./test.txt"
	// Create a file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// loop over slice and write each struct in file
	for _, element := range data {
		fmt.Fprintln(file, element)
	}
	// success
	return nil
}
