package gameofthrones

// Station represents a character's social rank in the feudal hierarchy.
type Station string

const (
	StationSmallfolk Station = "Smallfolk"
	StationKnight    Station = "Knight"
	StationLord      Station = "Lord"
	StationGreatLord Station = "Great Lord"
	StationRoyal     Station = "Royal"
)

// Stations returns all valid station values.
func Stations() []Station {
	return []Station{
		StationSmallfolk,
		StationKnight,
		StationLord,
		StationGreatLord,
		StationRoyal,
	}
}

// CharacterStation maps character display names to their social station.
// Station is derived from the highest title held by the character.
var CharacterStation = map[string]Station{
	// Royal - Kings, Queens, Princes, Princesses
	"Robert Baratheon":   StationRoyal,
	"Joffrey Baratheon":  StationRoyal,
	"Tommen Baratheon":   StationRoyal,
	"Myrcella Baratheon": StationRoyal,
	"Cersei Lannister":   StationRoyal, // Queen Regent, later Queen
	"Daenerys Targaryen": StationRoyal,
	"Viserys Targaryen":  StationRoyal, // Claimant to throne
	"Shireen Baratheon":  StationRoyal, // Princess
	"Stannis Baratheon":  StationRoyal, // King claimant

	// Great Lord - Lords Paramount, Wardens, Hand of the King
	"Eddard \"Ned\" Stark": StationGreatLord, // Lord of Winterfell, Warden of the North, Hand
	"Tywin Lannister":      StationGreatLord, // Lord of Casterly Rock, Warden of the West, Hand
	"Robb Stark":           StationGreatLord, // King in the North
	"Roose Bolton":         StationGreatLord, // Warden of the North (briefly)
	"Mace Tyrell":          StationGreatLord, // Lord of Highgarden, Warden of the South
	"Olenna Tyrell":        StationGreatLord, // Lady of Highgarden
	"Balon Greyjoy":        StationGreatLord, // Lord of the Iron Islands (implied)
	"Doran Martell":        StationGreatLord, // Prince of Dorne (implied)

	// Lord - Minor lords, Ladies, Lord Commanders, Masters
	"Jon Snow":                       StationLord, // Lord Commander of Night's Watch, King in the North
	"Tyrion Lannister":               StationLord, // Master of Coin, Hand of the Queen
	"Catelyn Stark":                  StationLord, // Lady of Winterfell
	"Sansa Stark":                    StationLord, // Lady of Winterfell, Queen in the North
	"Bran Stark":                     StationLord, // Lord of Winterfell, King of the Six Kingdoms
	"Ramsay Bolton":                  StationLord, // Lord of Winterfell (briefly)
	"Petyr \"Littlefinger\" Baelish": StationLord, // Lord of Harrenhal, Lord Protector of the Vale
	"Varys":                          StationLord, // Master of Whisperers
	"Jeor Mormont":                   StationLord, // Lord Commander of Night's Watch
	"Jorah Mormont":                  StationLord, // Exiled Lord of Bear Island
	"Theon Greyjoy":                  StationLord, // Prince of the Iron Islands
	"Yara Greyjoy":                   StationLord, // Lady of the Iron Islands
	"Ellaria Sand":                   StationLord, // Paramour, de facto ruler of Dorne
	"Margaery Tyrell":                StationLord, // Queen
	"Loras Tyrell":                   StationLord, // Lord Commander of the Kingsguard
	"Kevan Lannister":                StationLord, // Lord Regent, Hand of the King
	"Selyse Florent":                 StationLord, // Lady, Queen claimant
	"Beric Dondarrion":               StationLord, // Lord of Blackhaven
	"Daario Naharis":                 StationLord, // Commander of Second Sons
	"Grey Worm":                      StationLord, // Commander of the Unsullied
	"Qyburn":                         StationLord, // Master of Whisperers, Hand of the Queen

	// Knight - Ser, Kingsguard, warriors of note
	"Jaime Lannister":              StationKnight,
	"Barristan Selmy":              StationKnight,
	"Gregor Clegane":               StationKnight,
	"Sandor \"The Hound\" Clegane": StationKnight,
	"Brienne of Tarth":             StationKnight,
	"Bronn":                        StationKnight, // Knighted
	"Davos Seaworth":               StationKnight, // Ser Davos, Hand of the King
	"Lancel Lannister":             StationKnight,
	"Meryn Trant":                  StationKnight,
	"Rodrik Cassel":                StationKnight, // Ser Rodrik
	"Janos Slynt":                  StationKnight, // Commander of the Gold Cloaks
	"Podrick Payne":                StationKnight, // Squire, later knighted
	"Alliser Thorne":               StationKnight, // Ser Alliser
	"Tormund Giantsbane":           StationKnight, // Warrior leader (Free Folk equivalent)
	"Khal Drogo":                   StationKnight, // Khal (Dothraki equivalent)

	// Smallfolk and others - Common people, servants, bastards without title
	"Arya Stark":            StationLord,      // Lady, though unconventional
	"Rickon Stark":          StationLord,      // Lord
	"Samwell Tarly":         StationKnight,    // Maester in training, from noble house
	"Gendry":                StationSmallfolk, // Bastard blacksmith (later legitimized)
	"Gilly":                 StationSmallfolk,
	"Hot Pie":               StationSmallfolk,
	"Shae":                  StationSmallfolk,
	"Ros":                   StationSmallfolk,
	"Missandei":             StationSmallfolk, // Former slave, advisor
	"Ygritte":               StationSmallfolk, // Free Folk
	"Talisa Maegyr":         StationLord,      // Noblewoman of Volantis
	"Hodor":                 StationSmallfolk,
	"Osha":                  StationSmallfolk,
	"Irri":                  StationSmallfolk,
	"Doreah":                StationSmallfolk,
	"Melisandre":            StationKnight,    // Red Priestess, advisor
	"The High Sparrow":      StationLord,      // High Septon
	"Grand Maester Pycelle": StationLord,      // Grand Maester
	"Maester Luwin":         StationKnight,    // Maester
	"Maester Aemon":         StationLord,      // Maester, secret Targaryen prince
	"Thoros of Myr":         StationKnight,    // Red Priest, warrior
	"Jaqen H'ghar":          StationKnight,    // Faceless Man
	"The Waif":              StationSmallfolk, // Faceless Man acolyte
	"Grenn":                 StationSmallfolk,
	"Pypar":                 StationSmallfolk,
	"Rast":                  StationSmallfolk,
	"Eddison Tollett":       StationKnight, // Acting Lord Commander
	"Othell Yarwyck":        StationKnight, // First Builder
	"Bowen Marsh":           StationKnight, // First Steward
	"Olly":                  StationSmallfolk,
	"Little Sam":            StationSmallfolk,
	"Jojen Reed":            StationLord, // Noble house
	"Meera Reed":            StationLord, // Noble house
}
