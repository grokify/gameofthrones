package gameofthrones

// Orders is a list of all major orders, brotherhoods, and formal organizations
// in Westeros and Essos. This differs from the Organizations list which includes
// geographic locations and houses.
var Orders = []string{
	"The Small Council",
	"The Kingsguard",
	"The Night's Watch",
	"The Citadel",
	"The Faith of the Seven",
	"The Gold Cloaks",
	"The Iron Fleet",
	"House Guard",
	"Brotherhood Without Banners",
	"The Faceless Men",
	"The Unsullied",
	"The Second Sons",
	"The Dothraki",
	"The Free Folk",
}

// CharacterOrders maps character display names to the orders they belong to.
// A character may belong to multiple orders over the course of the series.
var CharacterOrders = map[string][]string{
	// Night's Watch
	"Jon Snow":        {"The Night's Watch"},
	"Samwell Tarly":   {"The Night's Watch", "The Citadel", "The Small Council"},
	"Jeor Mormont":    {"The Night's Watch"},
	"Maester Aemon":   {"The Night's Watch", "The Citadel"},
	"Alliser Thorne":  {"The Night's Watch"},
	"Grenn":           {"The Night's Watch"},
	"Pypar":           {"The Night's Watch"},
	"Eddison Tollett": {"The Night's Watch"},
	"Othell Yarwyck":  {"The Night's Watch"},
	"Bowen Marsh":     {"The Night's Watch"},
	"Rast":            {"The Night's Watch"},
	"Olly":            {"The Night's Watch"},

	// Kingsguard
	"Jaime Lannister":              {"The Kingsguard"},
	"Barristan Selmy":              {"The Kingsguard"},
	"Meryn Trant":                  {"The Kingsguard"},
	"Gregor Clegane":               {"The Kingsguard"},
	"Brienne of Tarth":             {"The Kingsguard"},
	"Sandor \"The Hound\" Clegane": {"House Guard"}, // Not Kingsguard, personal guard

	// Small Council
	"Eddard \"Ned\" Stark":           {"The Small Council"},
	"Tyrion Lannister":               {"The Small Council"},
	"Tywin Lannister":                {"The Small Council"},
	"Petyr \"Littlefinger\" Baelish": {"The Small Council"},
	"Varys":                          {"The Small Council"},
	"Grand Maester Pycelle":          {"The Small Council", "The Citadel"},
	"Mace Tyrell":                    {"The Small Council"},
	"Kevan Lannister":                {"The Small Council"},
	"Qyburn":                         {"The Small Council"},
	"Davos Seaworth":                 {"The Small Council"},
	"Bronn":                          {"The Small Council"},

	// Gold Cloaks (City Watch)
	"Janos Slynt": {"The Gold Cloaks"},

	// Citadel / Maesters
	"Maester Luwin": {"The Citadel"},

	// Faith of the Seven
	"The High Sparrow": {"The Faith of the Seven"},
	"Lancel Lannister": {"The Faith of the Seven"},

	// Religious - Red Priests
	"Melisandre":    {"The Faith of the Seven"}, // Different faith, but similar role
	"Thoros of Myr": {"Brotherhood Without Banners"},

	// Brotherhood Without Banners
	"Beric Dondarrion": {"Brotherhood Without Banners"},

	// Faceless Men
	"Jaqen H'ghar": {"The Faceless Men"},
	"The Waif":     {"The Faceless Men"},
	"Arya Stark":   {"The Faceless Men"}, // Temporarily

	// Unsullied
	"Grey Worm": {"The Unsullied"},

	// Second Sons
	"Daario Naharis": {"The Second Sons"},

	// Dothraki
	"Khal Drogo":         {"The Dothraki"},
	"Daenerys Targaryen": {"The Dothraki"}, // Khaleesi

	// Free Folk
	"Tormund Giantsbane": {"The Free Folk"},
	"Ygritte":            {"The Free Folk"},
	"Osha":               {"The Free Folk"},
	"Gilly":              {"The Free Folk"},

	// Iron Fleet
	"Yara Greyjoy":  {"The Iron Fleet"},
	"Theon Greyjoy": {"The Iron Fleet"},
}
