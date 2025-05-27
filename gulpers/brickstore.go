package gulpers

import (
	"github.com/Choff3/gulper/gemini"
	"github.com/Choff3/gulper/utils"
)

func BrickStorePub() string {
	const website = "https://www.brickstorepub.com"
	const venue = "Brick Store Pub"
	const prompt = "The brewery name comes first followed by the beer name. When storing price, take the higher price." // If under the Cask section, put (cask) after the beer name."

	beerStr := gemini.GetMenuHTML("https://www.brickstorepub.com/drinks-menu#menu=draught-beer", prompt)

	utils.Gulp(beerStr, venue, website, true)

	// TODO: Add error checking
	return "Success!"
}
