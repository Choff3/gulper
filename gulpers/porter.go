package gulpers

import (
	"fmt"

	"github.com/Choff3/gulper/gemini"
	"github.com/Choff3/gulper/utils"
	"github.com/gocolly/colly"
)

func Porter() string {
	const website = "https://www.theporterbeerbar.com"
	const venue = "The Porter"
	const prompt = "The brewery name comes first, then a • is used to separate the brewery name from from the beer name. When storing style, only store the style not the description that comes after."

	beerStr := gemini.GetMenuPDF(getMenuURL(website), prompt)

	utils.Gulp(beerStr, venue, website, true)

	// TODO: Add error checking
	return "Success!"
}

func getMenuURL(website string) string {

	var menu string

	// Instantiate default collector
	c := colly.NewCollector(
		colly.UserAgent(utils.GetUserAgent()),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if e.Text == "Draft Menu" {
			menu = link
			return
		}
	})

	// Start scraping
	c.Visit(website)

	return menu
}
