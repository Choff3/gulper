package gulpers

import (
	"io"
	"log"
	"net/http"

	"github.com/Choff3/gulper/gemini"
	"github.com/Choff3/gulper/utils"
)

func BrickStorePub() string {
	const website = "https://www.brickstorepub.com"
	const venue = "Brick Store Pub"
	const prompt = "The brewery name comes first followed by the beer name. When storing price, take the higher price. If under the Cask section, put (cask) after the beer name."

	beerStr := gemini.GetMenuHTML(getPageContent("https://www.brickstorepub.com/drinks-menu#menu=draught-beer"), prompt)

	utils.Gulp(beerStr, venue, website, true)

	// TODO: Add error checking
	return "Success!"
}

func getPageContent(url string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", utils.GetUserAgent())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)
	return sb
}
