package catalog

// This Catalog is designed to hardcode Wiki Addresses and Games. More will be added over time.

// Define Game Elements
type Game struct {
	Name string
	Slug string
	CharacterPath string
}

// Define Wiki Elements. Nested Map to assist in mapping game URLs.
type Wiki struct {
	Name  string
	URL   string
	Games map[string]Game
}

// Literal to hold Wiki Data
var Wikis = map[string]Wiki{
	"dustloop": {
		Name: "Dustloop",
		URL:  DustloopURL,
		Games: map[string]Game{
			"ggst": {
					Name: "Guilty Gear Strive", 
					Slug: "GGST", 
					CharacterPath: "GGST/{character}",
					},
			"blazbluecf": {
					Name: "Blazblue Centralfiction", 
					Slug: "BBCF", 
					CharacterPath: "BBCF/{character}",
					},
			},
	},
	"supercombo": {
		Name: "SuperCombo",
		URL:  SuperComboURL,
		Games: map[string]Game{
			"sf6": {
					Name: "Street Fighter 6", 
					Slug: "SF6", 
					CharacterPath: "Street_Fighter_6/{character}",
			},
			"3s":  {
					Name: "Street Fighter III: 3rd Strike", 
					Slug: "3S",
					CharacterPath: "Street_Fighter_3/{character}",
			},
		},
	},
}
