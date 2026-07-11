package gameofthrones

// LifeStatus represents whether a character is living or deceased.
type LifeStatus string

const (
	StatusLiving   LifeStatus = "Living"
	StatusDeceased LifeStatus = "Deceased"
)

// LifeStatuses returns all valid life status values.
func LifeStatuses() []LifeStatus {
	return []LifeStatus{
		StatusLiving,
		StatusDeceased,
	}
}

// CharacterStatus maps character display names to their life status at the end of the series.
// Status is based on the character's state at the conclusion of Season 8.
var CharacterStatus = map[string]LifeStatus{
	// Living characters at end of series
	"Jon Snow":           StatusLiving,
	"Sansa Stark":        StatusLiving,
	"Arya Stark":         StatusLiving,
	"Bran Stark":         StatusLiving,
	"Tyrion Lannister":   StatusLiving,
	"Samwell Tarly":      StatusLiving,
	"Brienne of Tarth":   StatusLiving,
	"Davos Seaworth":     StatusLiving,
	"Bronn":              StatusLiving,
	"Podrick Payne":      StatusLiving,
	"Gendry":             StatusLiving,
	"Gilly":              StatusLiving,
	"Yara Greyjoy":       StatusLiving,
	"Grey Worm":          StatusLiving,
	"Tormund Giantsbane": StatusLiving,
	"Hot Pie":            StatusLiving,
	"Little Sam":         StatusLiving,
	"Meera Reed":         StatusLiving,

	// Deceased characters
	"Eddard \"Ned\" Stark":           StatusDeceased,
	"Robert Baratheon":               StatusDeceased,
	"Jaime Lannister":                StatusDeceased,
	"Catelyn Stark":                  StatusDeceased,
	"Cersei Lannister":               StatusDeceased,
	"Daenerys Targaryen":             StatusDeceased,
	"Jorah Mormont":                  StatusDeceased,
	"Petyr \"Littlefinger\" Baelish": StatusDeceased,
	"Viserys Targaryen":              StatusDeceased,
	"Robb Stark":                     StatusDeceased,
	"Theon Greyjoy":                  StatusDeceased,
	"Joffrey Baratheon":              StatusDeceased,
	"Sandor \"The Hound\" Clegane":   StatusDeceased,
	"Khal Drogo":                     StatusDeceased,
	"Tywin Lannister":                StatusDeceased,
	"Margaery Tyrell":                StatusDeceased,
	"Stannis Baratheon":              StatusDeceased,
	"Melisandre":                     StatusDeceased,
	"Jeor Mormont":                   StatusDeceased,
	"Varys":                          StatusDeceased,
	"Shae":                           StatusDeceased,
	"Ygritte":                        StatusDeceased,
	"Talisa Maegyr":                  StatusDeceased,
	"Missandei":                      StatusDeceased,
	"Jaqen H'ghar":                   StatusLiving, // Faceless Men continue
	"Tommen Baratheon":               StatusDeceased,
	"Roose Bolton":                   StatusDeceased,
	"The High Sparrow":               StatusDeceased,
	"Grand Maester Pycelle":          StatusDeceased,
	"Meryn Trant":                    StatusDeceased,
	"Hodor":                          StatusDeceased,
	"Grenn":                          StatusDeceased,
	"Osha":                           StatusDeceased,
	"Rickon Stark":                   StatusDeceased,
	"Ros":                            StatusDeceased,
	"Gregor Clegane":                 StatusDeceased,
	"Janos Slynt":                    StatusDeceased,
	"Lancel Lannister":               StatusDeceased,
	"Rodrik Cassel":                  StatusDeceased,
	"Maester Luwin":                  StatusDeceased,
	"Irri":                           StatusDeceased,
	"Doreah":                         StatusDeceased,
	"Kevan Lannister":                StatusDeceased,
	"Barristan Selmy":                StatusDeceased,
	"Rast":                           StatusDeceased,
	"Maester Aemon":                  StatusDeceased,
	"Pypar":                          StatusDeceased,
	"Alliser Thorne":                 StatusDeceased,
	"Othell Yarwyck":                 StatusDeceased,
	"Loras Tyrell":                   StatusDeceased,
	"Beric Dondarrion":               StatusDeceased,
	"Eddison Tollett":                StatusDeceased,
	"Selyse Florent":                 StatusDeceased,
	"Shireen Baratheon":              StatusDeceased,
	"Jojen Reed":                     StatusDeceased,
	"Thoros of Myr":                  StatusDeceased,
	"Olly":                           StatusDeceased,
	"Mace Tyrell":                    StatusDeceased,
	"The Waif":                       StatusDeceased,
	"Bowen Marsh":                    StatusDeceased,
	"Ramsay Bolton":                  StatusDeceased,
	"Ellaria Sand":                   StatusDeceased, // Presumed dead in dungeon
	"Daario Naharis":                 StatusLiving,
	"Olenna Tyrell":                  StatusDeceased,
	"Myrcella Baratheon":             StatusDeceased,
	"Qyburn":                         StatusDeceased,
}
