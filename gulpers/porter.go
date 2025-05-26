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
		colly.UserAgent(utils.GetUserAgent()),
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

	url := "https://www.theporterbeerbar.com/wordpress/wp-content/uploads/2025/05/MAY-17-DRAFT-MENU.pdf" //getMenu()

	beers := utils.GetBeersPDF(url)

	return beers
}
