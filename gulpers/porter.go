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

// func GetPorterBeers() string {
// 	prompt := "The brewery name comes first, then a dot is used to separate the brewery name from from the beer name."

// 	url := getMenu()

// 	beers := utils.GetBeersPDF(url, prompt)

// 	return beers
// }

func GetPorterBeers() string {
	beerString := `
	[
	{
		"name": "Blueberry Spaceship Box",
		"brewery": "Superstition Meadery",
		"style": "Cider, Mead, and Fruit Wine",
		"abv": "5.5%",
		"price": "9.50",
		"other": "Fruited Cider: The top rated hard cider in the world. A spectacular blend of sweet & tart, of apple & blueberry"
	},
	{
		"name": "LESS is More Pet Nat",
		"brewery": "Botanist & Barrel",
		"style": "Dry Cider",
		"abv": "8.5%",
		"price": "11.00",
		"other": "Lightly effervescent, dry w/ subtle minerality, notes of honey, tropical & stone fruit, and a racy acidity"
	},
	{
		"name": "Beemosa",
		"brewery": "Superstition Meadery",
		"style": "Sparkling Mead",
		"abv": "6.0%",
		"price": "9.00",
		"other": "Bright bubbles, candied orange & a lively minerality w/ vinous notes of pinot grigio & a splash of honey in the finish"
	},
	{
		"name": "Rancio",
		"brewery": "Frederiksdal",
		"style": "Cherry Wine",
		"abv": "15.0%",
		"price": "17.00",
		"other": "Made w/ Danish Stevns cherries. Exceptional, full-bodied & complex taste experience. Pure wine meditation"
	},
	{
		"name": "BEER BEER",
		"brewery": "Pipeworks Brewing Co.",
		"style": "American Lager",
		"abv": "4.4%",
		"price": "7.00",
		"other": "Just a crisp, clean, American lager. The kind of beer your dad would drink while mowing the lawn. Simple, delicious BEER"
	},
	{
		"name": "The Silence of Apocalypse",
		"brewery": "Sceptre Brewing Arts",
		"style": "Mexican Lager",
		"abv": "4.4%",
		"price": "7.50",
		"other": "Easy drinking & extremely refreshing. Brewed w/ regional malts, corn, a kiss of noble hops, & a Mexican yeast strain"
	},
	{
		"name": "Kölsch",
		"brewery": "Privat-Brauerei Heinrich Reissdorf",
		"style": "Kölsch",
		"abv": "4.8%",
		"price": "7.50",
		"other": "The Cologne beer specialty w/ a brewing tradition since 1894; pleasant, full-bodied & unmistakably drinkable"
	},
	{
		"name": "Dent",
		"brewery": "Barrique Brewing & Blending",
		"style": "American Lager",
		"abv": "4.7%",
		"price": "8.00",
		"other": "Light, subtly sweet malt body w/ a smooth, crisp finish. Delicate floral & citrus notes from Audacia hops"
	},
	{
		"name": "Super Super Turbo Turbo",
		"brewery": "Halfway Crooks",
		"style": "Intercontinental-style Rice Lager",
		"abv": "5.0%",
		"price": "7.50",
		"other": "Lean & clean. Delicate aromas of lager yeast and bread dough, sweet, floral, refreshing"
	},
	{
		"name": "Domácí Pivo",
		"brewery": "Barrique Brewing & Blending",
		"style": "Italian-Style Pilsner",
		"abv": "5.0%",
		"price": "8.50",
		"other": "Slight malt body, balanced w/ fresh hoppy notes of Bermudagrass clippings & lime zest"
	},
	{
		"name": "Jas",
		"brewery": "Green Bench Brewing Co.",
		"style": "Czech-style Pale Lager",
		"abv": "5.2%",
		"price": "8.00",
		"other": "Golden, clear, delicate, drinkable. Brewed w/ imported Czech malts & hops for a perfectly balanced profile"
	},
	{
		"name": "Pil-Zen",
		"brewery": "Fonta Flora Brewery",
		"style": "Pilsner",
		"abv": "5.2%",
		"price": "8.00",
		"other": "A pilsner-style lagerbier aged on jasmine green tea"
	},
	{
		"name": "Champagne Velvet",
		"brewery": "Upland Brewing Co.",
		"style": "Pilsener",
		"abv": "5.5%",
		"price": "7.50",
		"other": "\"The Beer with the Million Dollar Flavor.\" Literally. Insured for $1 million dollars in 1935, this Midwestern Lager legacy was Resur- rected post-Prohibition when the recipe was discovered on a hand-written scrap of paper. A refreshing, smooth, light, everyday craft beer"
	},
	{
		"name": "Tampa Timeshare",
		"brewery": "Wrecking Bar",
		"style": "Fruited Gose",
		"abv": "4.5%",
		"price": "7.50",
		"other": "Soured w/ Atlanta Fresh Creamery yogurt, salted w/ a blend of sea salts from Beautiful Briny Sea in Grant Park. Then gen- erously dosed w/ Thai & sweet basil, fermented on grapefruit juice & pulp, aged on grapefruit zest"
	},
	{
		"name": "Sorry Not Sour",
		"brewery": "Prairie Artisan Ales",
		"style": "Fruit Ale",
		"abv": "4.5%",
		"price": "8.00",
		"other": "Blonde Ale with blood orange, mango, and apple"
	},
	{
		"name": "Lord Fog",
		"brewery": "Three Taverns Brewery",
		"style": "Sour Ale",
		"abv": "5.0%",
		"price": "8.00",
		"other": "This classic London Fog inspired variant of Lord Grey, our earl grey tea sour ale, has additions of lactose, vanilla, & lavender"
	},
	{
		"name": "Death Rattle",
		"brewery": "Green Bench Brewing Co.",
		"style": "Fruited Sour",
		"abv": "5.9%",
		"price": "9.50",
		"other": "Complex aromas & flavors of blackberry preserves, oak & fruit tannin, w/ a prickly, medium-high acidity & carbonation"
	},
	{
		"name": "Doragon Nectar",
		"brewery": "Tripping Animals Brewing Co. (Collaboration w/ Mikkeller)",
		"style": "Fruited Sour",
		"abv": "6.0%",
		"price": "9.00",
		"other": "An electrifying fusion of fruit & acidity, conditioned w/ strawberry, sweet cherry, tart cherry, peach, melon, raspberry, & a touch of lemon salt"
	},
	{
		"name": "Out of Order Blue Milk",
		"brewery": "RaR Brewing",
		"style": "Smoothie Sour",
		"abv": "6.0%",
		"price": "8.00",
		"other": "\"Blue Milk\" is made with milk from a Bantha, but also tastes like a blue raspberry float"
	},
	{
		"name": "Sun Kissed",
		"brewery": "Edmund's Oast Brewing Co.",
		"style": "Fruited Wheat Beer",
		"abv": "6.3%",
		"price": "8.00",
		"other": "Blend of sour & crisp wheat beers w/tangerine puree. Mildly tart & bursting w/ flavors of fresh-squeezed citrus juice!"
	},
	{
		"name": "Coolship Resurgam",
		"brewery": "Allagash Brewing Co.",
		"style": "American Wild Ale",
		"abv": "6.3%",
		"price": "10.00",
		"other": "Blend of wild ales; aromas of apricot & lemon zest. Notes of tropical fruit and a light funk lead to a clean, tart, dry finish"
	},
	{
		"name": "This Deli Needs a Bigger Double Lime Coconut Muffin Lassi Gose",
		"brewery": "Evil Twin",
		"style": "Smoothie/Pastry Sour",
		"abv": "7.0%",
		"price": "9.50",
		"other": "Brewed with kiwi, coconut, lime, sea salt, & marshmallow. Contains milk sugar. Collab with Omnipollo"
	},
	{
		"name": "Thrilla in Manila",
		"brewery": "Cøntrast Artisan Ales",
		"style": "Fruited Berliner Weisse",
		"abv": "7.4%",
		"price": "8.50",
		"other": "With coconut, lime, & agave, this sour drinks like a tropical cocktail"
	},
	{
		"name": "Mama Luna",
		"brewery": "Jester King",
		"style": "American Blonde Ale",
		"abv": "5.0%",
		"price": "8.00",
		"other": "Fermented crisp & clean w/ California Ale yeast and filtered to a brilliant clarity"
	},
	{
		"name": "Climate Promise",
		"brewery": "Perennial Artisan Ales",
		"style": "Saison",
		"abv": "5.0%",
		"price": "8.00",
		"other": "French-style Saison with honeysuckle, rose hips, dandelion, orange peel, & lemon peel"
	},
	{
		"name": "Bräuweisse",
		"brewery": "Ayinger Privatbrauerei",
		"style": "Hefeweizen",
		"abv": "5.1%",
		"price": "8.50",
		"other": "A light, quintessential expression of the style; refined, flowery character, unmistakable banana aroma, scarcely detected bitter tone. Full-bodied, soft & mild w/ a lively, champagne-like sparkle"
	},
	{
		"name": "Town & Country",
		"brewery": "Wild East Brewing Co.",
		"style": "Saison",
		"abv": "5.2%",
		"price": "8.00",
		"other": "Brewed w/ pilsner malt, a Blaugies/Dupont yeast strain, & Alora hops. Medium-light body w/ notes of Asian pear, pepper, & tangy apricot. Clean & approachable w/ a dry, smooth finish."
	},
	{
		"name": "Grisette",
		"brewery": "Moody Tongue Brewing Co.",
		"style": "Farmhouse Ale",
		"abv": "5.5%",
		"price": "8.00",
		"other": "Refreshing, light body w/ notes of wheat, straw, and chamomile. Lightly tart finish. Hopped w/ French Strisselspalt"
	},
	{
		"name": "La Chouffe",
		"brewery": "Brasserie d'Achouffe",
		"style": "Belgian Blonde Ale",
		"abv": "8.0%",
		"price": "10.00",
		"other": "Citrus notes followed by a refreshing touch, a light hop taste, fresh coriander notes, & fruity accents"
	},
	{
		"name": "Curieux",
		"brewery": "Allagash Brewing Co.",
		"style": "Belgian Tripel",
		"abv": "10.2%",
		"price": "9.50",
		"other": "Aged in Jim Beam barrels & blended w/ fresh Tripel. A rich, golden ale w/ notes of coconut, vanilla, & a hint of bourbon"
	},
		{
		"name": "Big Spin",
		"brewery": "DSSOLVR",
		"style": "Pale Ale",
		"abv": "5.0%",
		"price": "8.00",
		"other": "Bright, crisp, citrus-meets-tropical fruit, slightly floral, slightly dank, super balanced, incredibly crushable, and VERY clean on the backend. Like a whisper of hops across a citrus grove. It just keeps you wanting more"
	},
	{
		"name": "Pip Pip ESB",
		"brewery": "Cøntrast Artisan Ales",
		"style": "Extra Special Bitter",
		"abv": "5.8%",
		"price": "8.00",
		"other": "Dreamy British ESB: Toasty breadcrust, delicate chocolate covered raspberry, light caramel & toffee, toasted pecan"
	},
	{
		"name": "The Substance",
		"brewery": "Bissell Brothers",
		"style": "Hazy IPA",
		"abv": "6.6%",
		"price": "9.00",
		"other": "Our precious flagship. Brightly dank & just mysterious enough. Falconer's Flight, Centennial, Apollo, Chinook, & Simcoe hops"
	},
	{
		"name": "Cowboy Coast",
		"brewery": "Fonta Flora (Collaboration w/ Firestone Walker)",
		"style": "West Coast IPA",
		"abv": "7.0%",
		"price": "9.50",
		"other": "Brewed with heritage rice, dry-hopped with hand-selected mosaic, nelson sauvin, riwaka & el dorado"
	},
	{
		"name": "Millenial Falcon",
		"brewery": "Mason Ale Works",
		"style": "West Coast IPA",
		"abv": "7.0%",
		"price": "9.00",
		"other": "A galactic double-dry hopped West Coast IPA with Citra and Mosaic hops"
	},
	{
		"name": "Gouda",
		"brewery": "Other Half Brewing Co.",
		"style": "Double Hazy IPA",
		"abv": "8.0%",
		"price": "7.50",
		"other": "Hopped w/ Citra, Strata, Comet, and Eclipse: Comet's a '60s baby, but the era's light lagers weren't a fit. It's big on zesty citrus & pineapple. Eclipse is a new Australian hop w/ loads of sweet orange & fruit flavors"
	},
	{
		"name": "Lord Octomoss",
		"brewery": "Hop Butcher For The World",
		"style": "Double Hazy IPA",
		"abv": "8.0%",
		"price": "10.00",
		"other": "Mosaic, Citra & Simcoe-hopped Double India Pale Ale"
	},
	{
		"name": "XI-DDH Cosmik Debris",
		"brewery": "Creature Comforts Brewing Co.",
		"style": "Double IPA",
		"abv": "8.0%",
		"price": "8.50",
		"other": "Celebrating 11 yrs w/ a special release of Cosmik Debris -the first ever Creature Comforts DIPA- this time double-dry hopped"
	},
	{
		"name": "Extra Extra Bayou Juice",
		"brewery": "Weldwerks Brewing Co.",
		"style": "Double Hazy IPA",
		"abv": "8.2%",
		"price": "8.00",
		"other": "Brewed w/ Cascade, Citra, El Dorado, Mosaic, & Simcoe hops. Collaboration with Parish Brewing Co."
	},
	{
		"name": "Oboe Day",
		"brewery": "Creature Comforts Brewing Co.",
		"style": "Double Hazy IPA",
		"abv": "8.2%",
		"price": "8.50",
		"other": "Showcasing the almighty Citra hop, this double hazy is inspired by Andre 3000 as a reminder to always stay curious"
	},
	{
		"name": "Convivial IPA",
		"brewery": "Weldwerks Brewing Co.",
		"style": "Double Hazy IPA",
		"abv": "8.5%",
		"price": "8.50",
		"other": "Brewed with Russian River Brewing Co., using Brewer's Gold, Citra, Nectaron, & Simcoe Hops"
	},
	{
		"name": "Cuddle Buddies",
		"brewery": "Tripping Animals Brewing Co.",
		"style": "Triple Hazy Oat Cream IPA",
		"abv": "10.0%",
		"price": "9.50",
		"other": "Collab w/ Magnanimous. Epic hop lineup: Citra Lupomax, Topaz, & Nelson Sauvin. Contains lactose."
	},
	{
		"name": "Digital Dissonance",
		"brewery": "Wild East Brewing Co.",
		"style": "English-Style Brown Ale",
		"abv": "4.2%",
		"price": "8.00",
		"other": "Moderately malty w/ nutty toffee notes, light caramel & chocolate, balanced w/ earthy hops. Semi-dry finish. Collaboration w/ Good Word Brewing in Duluth, GA"
	},
	{
		"name": "Hidden Providence",
		"brewery": "Sceptre",
		"style": "Munich-Style Dunkel",
		"abv": "4.7%",
		"price": "7.50",
		"other": "Deceptively drinkable - herbaceous German hops & a clean lager yeast meld with caramel & toffee Munich malts"
	},
		{
		"name": "El Tecolote",
		"brewery": "Round Trip Brewing Co.",
		"style": "Vienna Mexican Lager",
		"abv": "4.9%",
		"price": "7.50",
		"other": "Malty body & rich, amber hue; refreshing & smooth w/ a light, crisp finish. Lemondrop hops add a citrusy punch"
	},
	{
		"name": "Bone Church",
		"brewery": "Creature Comforts Brewing Co.",
		"style": "Czech-Style Amber Lager",
		"abv": "5.2%",
		"price": "7.50",
		"other": "A touch of malty toast, firm bitterness, and amazing drinkability"
	},
	{
		"name": "Amberish Energy",
		"brewery": "Barrique Brewing & Blending",
		"style": "Kellerbier",
		"abv": "5.3%",
		"price": "8.50",
		"other": "Smooth & balanced malt body w/ notes of toasted grain, biscuit, & a hint of honey, w/ gentle bitterness & earthy hop character"
	},
	{
		"name": "Black Lager",
		"brewery": "Moody Tongue Brewing Co.",
		"style": "Schwarzbier",
		"abv": "5.3%",
		"price": "7.50",
		"other": "Aromas of espresso & roasted pecan, flavors of black plum & Russian black bread, light-bodied but volcanic black in color"
	},
	{
		"name": "Another, Please",
		"brewery": "Creature Comforts Brewing Co.",
		"style": "Altbier",
		"abv": "5.5%",
		"price": "7.50",
		"other": "A classic altbier, not made with lager yeasts, but made in the lager tradition, altbier simple means 'old beer' as a reference to the deep history of this style that predates lager yeasts! Dark amber, dry, and bitter; this beer style is designed to be quite quaffable"
	},
	{
		"name": "Dunkel",
		"brewery": "Bierstadt Lagerhaus",
		"style": "Munich Style Dunkel",
		"abv": "6.0%",
		"price": "10.00",
		"other": "Dark, flavors of bread crusts and chocolate. Franconian classic"
	},
	{
		"name": "Augur Porter",
		"brewery": "Dutchess Ales",
		"style": "English-Style Porter",
		"abv": "4.6%",
		"price": "7.50",
		"other": "Classic English Porter w/ a touch of East Coast nuance. Dry, crisp, satisfying proper ale w/ a dash of toasty biscuit & cacao. Maris Otter, barley, wheat, & crystal malts. Willamette, Bramling Cross, & East Kent Goldings hops give a woodsy, fresh presence"
	},
	{
		"name": "Slay the Psychonut",
		"brewery": "Three Taverns Brewery",
		"style": "Double Pastry Stout",
		"abv": "8.0%",
		"price": "8.50",
		"other": "Saturated with a sweet confection of fresh & toasted coconut, almonds, & rich chocolate, this coffee milk stout will take your taste buds on a twisted island adventure"
	},
	{
		"name": "Dark Chocolate Truffle",
		"brewery": "Untitled Art",
		"style": "Double Pastry Stout",
		"abv": "10.0%",
		"price": "8.00",
		"other": "Decadent, like a dark chocolate truffle. Velvety smooth mouthfeel, rich cocoa notes, & subtle roasted undertones"
	},
	{
		"name": "Coffee & Cakes",
		"brewery": "Hubbard's Cave Brewing",
		"style": "Imperial Stout",
		"abv": "12.0%",
		"price": "8.50",
		"other": "Imperial Stout with coffee & maple syrup"
	},
	{
		"name": "Peak Bloom Harvest Cider",
		"brewery": "Eden Specialty Ciders",
		"style": "Semi-Dry Cider",
		"abv": "6.2%",
		"price": "10.00",
		"other": "Apple sunshine in a can. Lush apple balanced by light tannin & soft tartness. Notes of white grape, applesauce, & lime"
	},
	{
		"name": "Pseudo Sue",
		"brewery": "Toppling Goliath Brewing Co.",
		"style": "Pale Ale",
		"abv": "5.8%",
		"price": "7.50",
		"other": "Bright w/ a bite. Citra-hopped, medium body w/ a mild bitterness. Ferocious aromas of citrus & mango; hints of passionfruit"
	},
	{
		"name": "Dark Lager",
		"brewery": "Sacred Profane Brewing",
		"style": "Czech-style Tmavé",
		"abv": "4.0%",
		"price": "8.00",
		"other": "Rich & complex w/ a robust character, boasting hints of caramel & toasted malt, balanced by a slight hop bitterness"
	},
	{
		"name": "Death & Taxes",
		"brewery": "Moonlight Brewing Co.",
		"style": "San Francisco-Style Black Lager",
		"abv": "5.3%",
		"price": "9.50",
		"other": "Deceptively light-bodied and highly drinkable. Drinks like iced coffee with a different effect"
	},
	{
		"name": "Carlow",
		"brewery": "The Drowned Lands Brewery",
		"style": "Irish Dry Stout",
		"abv": "5.5%",
		"price": "10.00",
		"other": "We're tasting baker's chocolate, a perfect pour-over coffee w/ Irish cream, and oven-fresh dark molasses cookies"
	},
	{
		"name": "Double Darwin's Forehead",
		"brewery": "Fonta Flora Brewery",
		"style": "Imperial Salted Brown Porter",
		"abv": "8.0%",
		"price": "11.00",
		"other": "Brewed to comically celebrate Charles Darwin, the complex & contemplative father of evolution"
	}
	]
	`

	return beerString
}
