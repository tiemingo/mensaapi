package mensaapi

import "time"

var TipsMap = map[string]string{
	"veg": "Vegetarisch",
	"van": "Vegan",
	"bio": "Bio",
	"S":   "Schwein",
	"R":   "Rind",
	"L":   "Lamm",
	"W":   "Wildfleisch",
	"G":   "Geflügel",
	"F":   "Fisch",
	"T":   "Tintenfisch",
}

var TipsMapEn = map[string]string{
	"veg": "Vegetarian",
	"van": "Vegan",
	"bio": "Bio",
	"S":   "Pork",
	"R":   "Beef",
	"L":   "Lamb",
	"W":   "Game meat",
	"G":   "Poultry",
	"F":   "Fish",
	"T":   "Squid",
}

type Mensa struct {
	Locations map[string]MensaLocation
}

type MensaLocation struct {
	Name         string // Facility
	SubLocations map[string]Sublocation
}

type Sublocation struct {
	Name   string // Location
	Legend map[string]string
	Days   map[time.Time]Item
}

type Day struct {
	Date       time.Time
	Categories map[string][]Item
}

type Item struct {
	Category  string
	Name      string
	Info      string
	Prices    Prices
	Date      time.Time
	Additives Additives // Allergenes + Food restrictions
	Nutrition Nutrition
}

type Prices struct {
	PriceStudent float64
	PriceWorker  float64
	PriceGuest   float64
}

type Nutrition struct {
	Cal           float64
	Protein       float64
	Fat           float64
	PartSaturated float64
	Carbs         float64
	PartSuguar    float64
	Salt          float64
}

type Additives struct {
	Allergenes []string
	Tips       []string // Food restrictions
}
