package gulpers

import (
	"fmt"

	"github.com/Choff3/gulper/utils"
	"github.com/gocolly/colly"
)

func getMenu() string {

	var menu string

	// Instantiate default collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:138.0) Gecko/20100101 Firefox/138.0"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if e.Text == "Draft Menu" {
			menu = link
			return
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping
	c.Visit("https://www.theporterbeerbar.com")

	return menu
}

func GetPorterBeers() string {

	url := getMenu()
	prompt := "Get a list of all the beers on this menu and return each of them in an array of JSONs with each column as a key. The keys are: name, brewery, style, abv, description, and price."
	// Construct the full URL with the API key
	fullPrompt := fmt.Sprintf("%s %s", prompt, url)

	// beers := utils.GetBeers(url, prompt)
	beers := utils.GetBeers(fullPrompt)

	return beers
}
