package gameofthrones

// TitleCategory represents the domain of a title.
type TitleCategory string

const (
	TitleCategoryGovernance TitleCategory = "Governance"
	TitleCategoryMilitary   TitleCategory = "Military"
	TitleCategoryReligious  TitleCategory = "Religious"
	TitleCategoryAcademic   TitleCategory = "Academic"
	TitleCategoryCommerce   TitleCategory = "Commerce"
)

// TitleCategories returns all valid title category values.
func TitleCategories() []TitleCategory {
	return []TitleCategory{
		TitleCategoryGovernance,
		TitleCategoryMilitary,
		TitleCategoryReligious,
		TitleCategoryAcademic,
		TitleCategoryCommerce,
	}
}

// Title represents a formal title held by a character.
type Title struct {
	Name     string        `json:"name"`
	Station  Station       `json:"station"`
	Category TitleCategory `json:"category"`
}

// CharacterTitles maps character display names to their titles.
// Characters may hold multiple titles over the course of the series.
var CharacterTitles = map[string][]Title{
	"Eddard \"Ned\" Stark": {
		{Name: "Lord of Winterfell", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Warden of the North", Station: StationGreatLord, Category: TitleCategoryGovernance},
		{Name: "Hand of the King", Station: StationGreatLord, Category: TitleCategoryGovernance},
	},
	"Robert Baratheon": {
		{Name: "King of the Andals and the First Men", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Lord of the Seven Kingdoms", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Protector of the Realm", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Jaime Lannister": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "Lord Commander of the Kingsguard", Station: StationLord, Category: TitleCategoryMilitary},
		{Name: "Kingslayer", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Catelyn Stark": {
		{Name: "Lady of Winterfell", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Cersei Lannister": {
		{Name: "Queen of the Seven Kingdoms", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Queen Regent", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Lady of Casterly Rock", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Daenerys Targaryen": {
		{Name: "Queen of the Andals and the First Men", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Mother of Dragons", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Khaleesi of the Great Grass Sea", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Breaker of Chains", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "The Unburnt", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Jorah Mormont": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "Lord of Bear Island", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Petyr \"Littlefinger\" Baelish": {
		{Name: "Lord of Harrenhal", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Lord Protector of the Vale", Station: StationGreatLord, Category: TitleCategoryGovernance},
		{Name: "Master of Coin", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Viserys Targaryen": {
		{Name: "King of the Andals and the First Men (claimant)", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Jon Snow": {
		{Name: "Lord Commander of the Night's Watch", Station: StationLord, Category: TitleCategoryMilitary},
		{Name: "King in the North", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Warden of the North", Station: StationGreatLord, Category: TitleCategoryGovernance},
	},
	"Sansa Stark": {
		{Name: "Lady of Winterfell", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Queen in the North", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Arya Stark": {
		{Name: "Lady of Winterfell", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Robb Stark": {
		{Name: "King in the North", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Lord of Winterfell", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Theon Greyjoy": {
		{Name: "Prince of the Iron Islands", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Bran Stark": {
		{Name: "Lord of Winterfell", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Three-Eyed Raven", Station: StationLord, Category: TitleCategoryReligious},
		{Name: "King of the Andals and the First Men", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Joffrey Baratheon": {
		{Name: "King of the Andals and the First Men", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Lord of the Seven Kingdoms", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Sandor \"The Hound\" Clegane": {
		{Name: "The Hound", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "Sworn Shield", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Tyrion Lannister": {
		{Name: "Hand of the King", Station: StationGreatLord, Category: TitleCategoryGovernance},
		{Name: "Hand of the Queen", Station: StationGreatLord, Category: TitleCategoryGovernance},
		{Name: "Master of Coin", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Khal Drogo": {
		{Name: "Khal", Station: StationGreatLord, Category: TitleCategoryGovernance},
	},
	"Tywin Lannister": {
		{Name: "Lord of Casterly Rock", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Warden of the West", Station: StationGreatLord, Category: TitleCategoryGovernance},
		{Name: "Hand of the King", Station: StationGreatLord, Category: TitleCategoryGovernance},
	},
	"Davos Seaworth": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "Hand of the King", Station: StationGreatLord, Category: TitleCategoryGovernance},
		{Name: "Onion Knight", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Samwell Tarly": {
		{Name: "Maester", Station: StationKnight, Category: TitleCategoryAcademic},
		{Name: "Grand Maester", Station: StationLord, Category: TitleCategoryAcademic},
	},
	"Margaery Tyrell": {
		{Name: "Queen of the Seven Kingdoms", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Stannis Baratheon": {
		{Name: "King of the Andals and the First Men (claimant)", Station: StationRoyal, Category: TitleCategoryGovernance},
		{Name: "Lord of Dragonstone", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Melisandre": {
		{Name: "Red Priestess", Station: StationKnight, Category: TitleCategoryReligious},
		{Name: "The Red Woman", Station: StationKnight, Category: TitleCategoryReligious},
	},
	"Jeor Mormont": {
		{Name: "Lord Commander of the Night's Watch", Station: StationLord, Category: TitleCategoryMilitary},
	},
	"Bronn": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "Lord of Highgarden", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Master of Coin", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Varys": {
		{Name: "Master of Whisperers", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "The Spider", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Tormund Giantsbane": {
		{Name: "Chieftain of the Free Folk", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Brienne of Tarth": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "Lord Commander of the Kingsguard", Station: StationLord, Category: TitleCategoryMilitary},
	},
	"Ramsay Bolton": {
		{Name: "Lord of Winterfell", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Warden of the North", Station: StationGreatLord, Category: TitleCategoryGovernance},
	},
	"Ellaria Sand": {
		{Name: "Paramour of Prince Oberyn", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Daario Naharis": {
		{Name: "Commander of the Second Sons", Station: StationLord, Category: TitleCategoryMilitary},
	},
	"Missandei": {
		{Name: "Advisor to Queen Daenerys", Station: StationKnight, Category: TitleCategoryGovernance},
	},
	"Tommen Baratheon": {
		{Name: "King of the Andals and the First Men", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Roose Bolton": {
		{Name: "Lord of the Dreadfort", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Warden of the North", Station: StationGreatLord, Category: TitleCategoryGovernance},
	},
	"The High Sparrow": {
		{Name: "High Septon", Station: StationLord, Category: TitleCategoryReligious},
	},
	"Grand Maester Pycelle": {
		{Name: "Grand Maester", Station: StationLord, Category: TitleCategoryAcademic},
	},
	"Gregor Clegane": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "The Mountain", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Lancel Lannister": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Rodrik Cassel": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "Master-at-Arms", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Maester Luwin": {
		{Name: "Maester", Station: StationKnight, Category: TitleCategoryAcademic},
	},
	"Kevan Lannister": {
		{Name: "Lord Regent", Station: StationGreatLord, Category: TitleCategoryGovernance},
		{Name: "Hand of the King", Station: StationGreatLord, Category: TitleCategoryGovernance},
	},
	"Barristan Selmy": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "Lord Commander of the Kingsguard", Station: StationLord, Category: TitleCategoryMilitary},
	},
	"Maester Aemon": {
		{Name: "Maester", Station: StationKnight, Category: TitleCategoryAcademic},
	},
	"Alliser Thorne": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "First Ranger", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Loras Tyrell": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "Knight of Flowers", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Beric Dondarrion": {
		{Name: "Lord of Blackhaven", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "The Lightning Lord", Station: StationLord, Category: TitleCategoryMilitary},
	},
	"Podrick Payne": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Eddison Tollett": {
		{Name: "Lord Commander of the Night's Watch", Station: StationLord, Category: TitleCategoryMilitary},
	},
	"Yara Greyjoy": {
		{Name: "Lady of the Iron Islands", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Grey Worm": {
		{Name: "Commander of the Unsullied", Station: StationLord, Category: TitleCategoryMilitary},
	},
	"Olenna Tyrell": {
		{Name: "Lady of Highgarden", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Queen of Thorns", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Shireen Baratheon": {
		{Name: "Princess", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Jojen Reed": {
		{Name: "Heir to Greywater Watch", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Meera Reed": {
		{Name: "Lady of Greywater Watch", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Thoros of Myr": {
		{Name: "Red Priest", Station: StationKnight, Category: TitleCategoryReligious},
	},
	"Mace Tyrell": {
		{Name: "Lord of Highgarden", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Warden of the South", Station: StationGreatLord, Category: TitleCategoryGovernance},
		{Name: "Master of Coin", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Qyburn": {
		{Name: "Master of Whisperers", Station: StationLord, Category: TitleCategoryGovernance},
		{Name: "Hand of the Queen", Station: StationGreatLord, Category: TitleCategoryGovernance},
	},
	"Myrcella Baratheon": {
		{Name: "Princess", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Gendry": {
		{Name: "Lord of Storm's End", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Rickon Stark": {
		{Name: "Lord of Winterfell", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Selyse Florent": {
		{Name: "Queen (claimant)", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Meryn Trant": {
		{Name: "Ser", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Janos Slynt": {
		{Name: "Commander of the City Watch", Station: StationKnight, Category: TitleCategoryMilitary},
		{Name: "Lord of Harrenhal", Station: StationLord, Category: TitleCategoryGovernance},
	},
	"Othell Yarwyck": {
		{Name: "First Builder", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Bowen Marsh": {
		{Name: "First Steward", Station: StationKnight, Category: TitleCategoryMilitary},
	},
	"Talisa Maegyr": {
		{Name: "Queen in the North", Station: StationRoyal, Category: TitleCategoryGovernance},
	},
	"Jaqen H'ghar": {
		{Name: "Faceless Man", Station: StationKnight, Category: TitleCategoryReligious},
	},
}
