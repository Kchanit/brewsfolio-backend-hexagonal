package database

import (
	"fmt"
	"log"

	"github.com/Kchanit/brewsfolio-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
func SeedUserData() {
	usersData := []models.User{
		{Name: "John Doe", Email: "jdoe@gmail.com", Password: HashPassword("123123"), Role: "USER"},
		{Name: "Jane Doe", Email: "jadoe@gmail.com", Password: HashPassword("123123"), Role: "USER"},
		{Name: "Admin", Email: "admin@gmail.com", Password: HashPassword("123123"), Role: "ADMIN"},
	}

	for _, user := range usersData {
		fmt.Println("User password: ", user.Password)
		err := DBConn.Create(&user).Error
		if err != nil {
			log.Fatalf("Error seeding user data: %v", err)
		}
	}
	fmt.Println("Seed user data completed")
}

func SeedBeerData() {
	beersData := []models.Beer{
		{Name: "Buzz", Tagline: "A Real Bitter Experience.", Description: "A light, crisp and bitter IPA brewed with English and American hops. A small batch brewed only once.", ImageURL: "https://images.punkapi.com/v2/keg.png", Abv: 4.5, Ibu: 60, Ebc: 20, Srm: 10, Ph: 4.4, AttenuationLevel: 75},
		{Name: "Trashy Blonde", Tagline: "You Know You Shouldn't", Description: "A titillating, neurotic, peroxide punk of a Pale Ale. Combining attitude, style, substance, and a little bit of low self esteem for good measure; what would your mother say? The seductive lure of the sassy passion fruit hop proves too much to resist. All that is even before we get onto the fact that there are no additives, preservatives, pasteurization or strings attached. All wrapped up with the customary BrewDog bite and imaginative twist.", ImageURL: "https://images.punkapi.com/v2/2.png", Abv: 4.1, Ibu: 41.5, Ebc: 15, Srm: 15, Ph: 4.4, AttenuationLevel: 76},
		{Name: "Berliner Weisse With Yuzu - B-Sides", Tagline: "Japanese Citrus Berliner Weisse.", Description: "Japanese citrus fruit intensifies the sour nature of this German classic.", ImageURL: "https://images.punkapi.com/v2/keg.png", Abv: 4.2, Ibu: 8, Ebc: 8, Srm: 4, Ph: 3.2, AttenuationLevel: 83},
		{Name: "Pilsen Lager", Tagline: "Unleash the Yeast Series.", Description: "Our Unleash the Yeast series was an epic experiment into the differences in aroma and flavour provided by switching up your yeast. We brewed up a wort with a light caramel note and some toasty biscuit flavour, and hopped it with Amarillo and Centennial for a citrusy bitterness. Everything else is down to the yeast. Pilsner yeast ferments with no fruity esters or spicy phenols, although it can add a hint of butterscotch.", ImageURL: "https://images.punkapi.com/v2/4.png", Abv: 6.3, Ibu: 55, Ebc: 30, Srm: 15, Ph: 4.4, AttenuationLevel: 80},
		{Name: "Avery Brown Dredge", Tagline: "Bloggers' Imperial Pilsner.", Description: "An Imperial Pilsner in collaboration with beer writers. Tradition. Homage. Revolution. We wanted to showcase the awesome backbone of the Czech brewing tradition, the noble Saaz hop, and also tip our hats to the modern beers that rock our world, and the people who make them.", ImageURL: "https://images.punkapi.com/v2/5.png", Abv: 7.2, Ibu: 59, Ebc: 10, Srm: 5, Ph: 4.4, AttenuationLevel: 67},
		{Name: "Electric India", Tagline: "Vibrant Hoppy Saison.", Description: "Re-brewed as a spring seasonal, this beer – which appeared originally as an Equity Punk shareholder creation – retains its trademark spicy, fruity edge. A perfect blend of Belgian Saison and US IPA, crushed peppercorns and heather honey are also added to produce a genuinely unique beer.", ImageURL: "https://images.punkapi.com/v2/6.png", Abv: 5.2, Ibu: 38, Ebc: 15, Srm: 7.5, Ph: 4.4, AttenuationLevel: 88.9},
		{Name: "AB:12", Tagline: "Imperial Black Belgian Ale.", Description: "An Imperial Black Belgian Ale aged in old Invergordon Scotch whisky barrels with mountains of raspberries, tayberries and blackberries in each cask. Decadent but light and dry, this beer would make a fantastic base for ageing on pretty much any dark fruit - we used raspberries, tayberries and blackberries beause they were local.", ImageURL: "https://images.punkapi.com/v2/7.png", Abv: 11.2, Ibu: 35, Ebc: 80, Srm: 40, Ph: 5.3, AttenuationLevel: 84},
		{Name: "Fake Lager", Tagline: "Bohemian Pilsner.", Description: "Fake is the new black. Fake is where it is at. Fake Art, fake brands, fake breasts, and fake lager. We want to play our part in the ugly fallout from the Lager Dream. Say hello to Fake Lager – a zesty, floral 21st century faux masterpiece with added BrewDog bitterness.", ImageURL: "https://images.punkapi.com/v2/8.png", Abv: 4.7, Ibu: 40, Ebc: 12, Srm: 6, Ph: 4.4, AttenuationLevel: 78},
		{Name: "AB:07", Tagline: "Whisky Cask-Aged Scotch Ale.", Description: "Whisky cask-aged imperial scotch ale. Beer perfect for when the rain is coming sideways. Liquorice, plum and raisin temper the warming alcohol, producing a beer capable of holding back the Scottish chill.", ImageURL: "https://images.punkapi.com/v2/9.png", Abv: 12.5, Ibu: 30, Ebc: 84, Srm: 42, Ph: 5.6, AttenuationLevel: 83},
		{Name: "Bramling X", Tagline: "Single Hop IPA Series - 2011.", Description: "Good old Bramling Cross is elegant, refined, assured, (boring) and understated. Understated that is unless you hop the living daylights out of a beer with it. This is Bramling Cross re-invented and re-imagined, and shows just what can be done with English hops if you use enough of them. Poor Bramling Cross normally gets lost in a woeful stream of conformist brown ales made by sleepy cask ale brewers. But not anymore. This beer shows that British hops do have some soul, and is a fruity riot of blackberries, pears, and plums. Reminds me of the bramble, apple and ginger jam my grandmother used to make.", ImageURL: "https://images.punkapi.com/v2/10.png", Abv: 7.5, Ibu: 75, Ebc: 22, Srm: 11, Ph: 4.4, AttenuationLevel: 80.9},
		{Name: "Misspent Youth", Tagline: "Milk & Honey Scotch Ale.", Description: "The brainchild of our small batch brewer, George Woods. A dangerously drinkable milk sugar- infused Scotch Ale.", ImageURL: "https://images.punkapi.com/v2/keg.png", Abv: 7.3, Ibu: 30, Ebc: 120, Srm: 60, Ph: 4.4, AttenuationLevel: 74.7},
		{Name: "Arcade Nation", Tagline: "Seasonal Black IPA.", Description: "Running the knife-edge between an India Pale Ale and a Stout, this particular style is one we truly love. Black IPAs are a great showcase for the skill of our brew team, balancing so many complex and twisting flavours in the same moment. The citrus, mango and pine from the hops – three of our all-time favourites – play off against the roasty dryness from the malt bill at each and every turn.", ImageURL: "https://images.punkapi.com/v2/12.png", Abv: 5.3, Ibu: 60, Ebc: 200, Srm: 100, Ph: 4.2, AttenuationLevel: 77},
		{Name: "Movember", Tagline: "Moustache-Worthy Beer.", Description: "A deliciously robust, black malted beer with a decadent dark, dry cocoa flavour that provides an enticing backdrop to the Cascade hops.", ImageURL: "https://images.punkapi.com/v2/13.png", Abv: 4.5, Ibu: 50, Ebc: 140, Srm: 70, Ph: 5.2, AttenuationLevel: 74.5},
		{Name: "Alpha Dog", Tagline: "Existential Red Ale.", Description: "A fusion of caramel malt flavours and punchy New Zealand hops. A session beer you can get your teeth into.", ImageURL: "https://images.punkapi.com/v2/14.png", Abv: 4.5, Ibu: 42, Ebc: 62, Srm: 31, Ph: 4.4, AttenuationLevel: 72.8},
		{Name: "Mixtape 8", Tagline: "An Epic Fusion Of Old Belgian, American New Wave, And Scotch Whisky.", Description: "This recipe is for the Belgian Tripel base. A blend of two huge oak aged beers – half a hopped up Belgian Tripel, and half a Triple India Pale Ale. Both aged in single grain whisky barrels for two years and blended, each beer brings its own character to the mix. The Belgian Tripel comes loaded with complex spicy, fruity esters, and punchy citrus hop character.", ImageURL: "https://images.punkapi.com/v2/15.png", Abv: 14.5, Ibu: 50, Ebc: 40, Srm: 20, Ph: 4.4, AttenuationLevel: 85},
		{Name: "Libertine Porter", Tagline: "Dry-Hopped Aggressive Porter.", Description: "An avalanche of cross-continental hop varieties give this porter a complex spicy, resinous and citrusy aroma, with a huge malt bill providing a complex roasty counterpoint. Digging deeper into the flavour draws out cinder toffee, bitter chocolate and hints of woodsmoke.", ImageURL: "https://images.punkapi.com/v2/16.png", Abv: 6.1, Ibu: 45, Ebc: 219, Srm: 109.5, Ph: 4.4, AttenuationLevel: 70.1},
		{Name: "AB:06", Tagline: "Imperial Black IPA.", Description: "Our sixth Abstrakt, this imperial black IPA combined dark malts with a monumental triple dry-hop, using an all-star team of some of our favourite American hops. Roasty and resinous.", ImageURL: "https://images.punkapi.com/v2/17.png", Abv: 11.2, Ibu: 150, Ebc: 70, Srm: 35, AttenuationLevel: 87},
		{Name: "Russian Doll – India Pale Ale", Tagline: "Nesting Hop Bomb.", Description: "The levels of hops vary throughout the range. We love hops, so all four beers are big, bitter badasses, but by tweaking the amount of each hop used later in the boil and during dry- hopping, we can balance the malty backbone with some unexpected flavours. Simcoe is used in the whirlpool for all four beers, and yet still lends different characters to each", ImageURL: "https://images.punkapi.com/v2/18.png", Abv: 6, Ibu: 70, Ebc: 25, Srm: 12.5, Ph: 5.2, AttenuationLevel: 79.3},
		{Name: "Hello My Name Is Mette-Marit", Tagline: "Lingonberry Double IPA.", Description: "We sent this beer to Norway where it was known as 'Hello, my name is Censored’. You can make up your own mind as to why. This brew was a red berry explosion, with a reisnous bitter edge layered with dry berry tartness.", ImageURL: "https://images.punkapi.com/v2/19.png", Abv: 8.2, Ibu: 70, Ph: 4.4, AttenuationLevel: 83},
		{Name: "Rabiator", Tagline: "Imperial Wheat Beer", Description: "Imperial Wheat beer / Weizenbock brewed by a homesick German in leather trousers. Think banana bread, bubble gum and David Hasselhoff.", ImageURL: "https://images.punkapi.com/v2/keg.png", Abv: 10.27, Ibu: 26, Ebc: 24, Srm: 12, Ph: 4.4, AttenuationLevel: 85},
		{Name: "Vice Bier", Tagline: "Hoppy Wheat Bier.", Description: "Our take on the classic German Kristallweizen. A clear German wheat beer, layers of bubblegum and vanilla perfectly balanced with the American and New Zealand hops.", ImageURL: "https://images.punkapi.com/v2/keg.png", Abv: 4.3, Ibu: 25, Ebc: 30, Srm: 15, Ph: 4, AttenuationLevel: 81.8},
		{Name: "Devine Rebel (w/ Mikkeller)", Tagline: "Oak-aged Barley Wine.", Description: "Two of Europe's most experimental, boundary-pushing brewers, BrewDog and Mikkeller, combined forces to produce a rebellious beer that combined their respective talents and brewing skills. The 12.5% Barley Wine fermented well, and the champagne yeast drew it ever closer to 12.5%. The beer was brewed with a single hop variety and was going to be partially aged in oak casks.", ImageURL: "https://images.punkapi.com/v2/22.png", Abv: 12.5, Ibu: 100, Ebc: 36, Srm: 18, Ph: 4.4, AttenuationLevel: 68},
		{Name: "Storm", Tagline: "Islay Whisky Aged IPA.", Description: "Dark and powerful Islay magic infuses this tropical sensation of an IPA. Using the original Punk IPA as a base, we boosted the ABV to 8% giving it some extra backbone to stand up to the peated smoke imported directly from Islay.", ImageURL: "https://images.punkapi.com/v2/23.png", Abv: 8, Ibu: 60, Ebc: 12, Srm: 6, Ph: 4.4, AttenuationLevel: 86},
		{Name: "The End Of History", Tagline: "The World's Strongest Beer.", Description: "The End of History: The name derives from the famous work of philosopher Francis Fukuyama, this is to beer what democracy is to history. Complexity defined. Floral, grapefruit, caramel and cloves are intensified by boozy heat.", ImageURL: "https://images.punkapi.com/v2/24.png", Abv: 55, Ph: 4.4, AttenuationLevel: 100},
		{Name: "Bad Pixie", Tagline: "Spiced Wheat Beer.", Description: "2008 Prototype beer, a 4.7% wheat ale with crushed juniper berries and citrus peel.", ImageURL: "https://images.punkapi.com/v2/25.png", Abv: 4.7, Ibu: 45, Ebc: 8, Srm: 4, Ph: 4.4, AttenuationLevel: 79},
	}

	for _, beer := range beersData {
		if err := DBConn.Create(&beer).Error; err != nil {
			log.Fatalf("Error seeding beer data: %v", err)
		}
	}
}
