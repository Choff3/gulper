package gulpers

import (
	"fmt"

	"github.com/Choff3/gulper/gemini"
	"github.com/Choff3/gulper/utils"
	"github.com/gocolly/colly"
)

const website = "https://www.theporterbeerbar.com"

const venue = "The Porter"

const prompt = "The brewery name comes first, then a dot is used to separate the brewery name from from the beer name."

func getMenuURL() string {

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

func Porter() string {
	// beerStr := gemini.GetMenuPDF(getMenuURL(), prompt)
	beerStr := gemini.GetMenuPDF("https://www.theporterbeerbar.com/wordpress/wp-content/uploads/2025/05/MAY-17-DRAFT-MENU.pdf", prompt)
	fmt.Println("Beer String:\n", beerStr)

	utils.Gulp(beerStr, venue, website, true)

	// TODO: Add error checking
	return "Success!"
}
