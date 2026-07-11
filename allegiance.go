package gameofthrones

// Allegiance represents a character's feudal relationships.
type Allegiance struct {
	Liege   string `json:"liege,omitempty"`    // Direct superior (feudal lord)
	SwornTo string `json:"sworn_to,omitempty"` // House/Order allegiance
}

// CharacterAllegiance maps character display names to their feudal allegiances.
// This represents the primary allegiance during the main events of the series.
var CharacterAllegiance = map[string]Allegiance{
	// House Stark
	"Eddard \"Ned\" Stark": {Liege: "Robert Baratheon", SwornTo: "House Baratheon"},
	"Catelyn Stark":        {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Robb Stark":           {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Sansa Stark":          {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Arya Stark":           {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Bran Stark":           {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Rickon Stark":         {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Rodrik Cassel":        {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Maester Luwin":        {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Hodor":                {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Osha":                 {Liege: "Bran Stark", SwornTo: "House Stark"},

	// House Baratheon
	"Robert Baratheon":  {Liege: "", SwornTo: "The Iron Throne"},
	"Stannis Baratheon": {Liege: "Robert Baratheon", SwornTo: "House Baratheon"},
	"Selyse Florent":    {Liege: "Stannis Baratheon", SwornTo: "House Baratheon"},
	"Shireen Baratheon": {Liege: "Stannis Baratheon", SwornTo: "House Baratheon"},
	"Davos Seaworth":    {Liege: "Stannis Baratheon", SwornTo: "House Baratheon"},
	"Melisandre":        {Liege: "Stannis Baratheon", SwornTo: "The Lord of Light"},

	// House Baratheon of King's Landing
	"Joffrey Baratheon":  {Liege: "Robert Baratheon", SwornTo: "House Baratheon"},
	"Tommen Baratheon":   {Liege: "Cersei Lannister", SwornTo: "House Baratheon"},
	"Myrcella Baratheon": {Liege: "Cersei Lannister", SwornTo: "House Baratheon"},

	// House Lannister
	"Tywin Lannister":  {Liege: "Robert Baratheon", SwornTo: "The Iron Throne"},
	"Cersei Lannister": {Liege: "Tywin Lannister", SwornTo: "House Lannister"},
	"Jaime Lannister":  {Liege: "Robert Baratheon", SwornTo: "The Kingsguard"},
	"Tyrion Lannister": {Liege: "Tywin Lannister", SwornTo: "House Lannister"},
	"Kevan Lannister":  {Liege: "Tywin Lannister", SwornTo: "House Lannister"},
	"Lancel Lannister": {Liege: "Tywin Lannister", SwornTo: "House Lannister"},

	// House Targaryen
	"Daenerys Targaryen": {Liege: "", SwornTo: "House Targaryen"},
	"Viserys Targaryen":  {Liege: "", SwornTo: "House Targaryen"},
	"Jorah Mormont":      {Liege: "Daenerys Targaryen", SwornTo: "House Targaryen"},
	"Missandei":          {Liege: "Daenerys Targaryen", SwornTo: "House Targaryen"},
	"Grey Worm":          {Liege: "Daenerys Targaryen", SwornTo: "House Targaryen"},
	"Daario Naharis":     {Liege: "Daenerys Targaryen", SwornTo: "House Targaryen"},
	"Barristan Selmy":    {Liege: "Daenerys Targaryen", SwornTo: "House Targaryen"},

	// Night's Watch (sworn to the realm, not a house)
	"Jon Snow":        {Liege: "Jeor Mormont", SwornTo: "The Night's Watch"},
	"Jeor Mormont":    {Liege: "", SwornTo: "The Night's Watch"},
	"Samwell Tarly":   {Liege: "Jon Snow", SwornTo: "The Night's Watch"},
	"Maester Aemon":   {Liege: "Jeor Mormont", SwornTo: "The Night's Watch"},
	"Alliser Thorne":  {Liege: "Jon Snow", SwornTo: "The Night's Watch"},
	"Grenn":           {Liege: "Jon Snow", SwornTo: "The Night's Watch"},
	"Pypar":           {Liege: "Jon Snow", SwornTo: "The Night's Watch"},
	"Eddison Tollett": {Liege: "Jon Snow", SwornTo: "The Night's Watch"},
	"Othell Yarwyck":  {Liege: "Jon Snow", SwornTo: "The Night's Watch"},
	"Bowen Marsh":     {Liege: "Jon Snow", SwornTo: "The Night's Watch"},
	"Rast":            {Liege: "Jon Snow", SwornTo: "The Night's Watch"},
	"Olly":            {Liege: "Jon Snow", SwornTo: "The Night's Watch"},

	// House Tyrell
	"Mace Tyrell":     {Liege: "Robert Baratheon", SwornTo: "The Iron Throne"},
	"Olenna Tyrell":   {Liege: "Mace Tyrell", SwornTo: "House Tyrell"},
	"Margaery Tyrell": {Liege: "Mace Tyrell", SwornTo: "House Tyrell"},
	"Loras Tyrell":    {Liege: "Mace Tyrell", SwornTo: "House Tyrell"},

	// House Greyjoy
	"Theon Greyjoy": {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"}, // Ward
	"Yara Greyjoy":  {Liege: "", SwornTo: "House Greyjoy"},

	// House Bolton
	"Roose Bolton":  {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Ramsay Bolton": {Liege: "Roose Bolton", SwornTo: "House Bolton"},

	// Free Folk
	"Tormund Giantsbane": {Liege: "", SwornTo: "The Free Folk"},
	"Ygritte":            {Liege: "", SwornTo: "The Free Folk"},
	"Gilly":              {Liege: "", SwornTo: "The Free Folk"},

	// Dothraki
	"Khal Drogo": {Liege: "", SwornTo: "Dothraki"},

	// King's Landing / Crown servants
	"Petyr \"Littlefinger\" Baelish": {Liege: "Robert Baratheon", SwornTo: "The Iron Throne"},
	"Varys":                          {Liege: "Robert Baratheon", SwornTo: "The Iron Throne"},
	"Grand Maester Pycelle":          {Liege: "Robert Baratheon", SwornTo: "The Iron Throne"},
	"Janos Slynt":                    {Liege: "Robert Baratheon", SwornTo: "The Iron Throne"},
	"Meryn Trant":                    {Liege: "Robert Baratheon", SwornTo: "The Kingsguard"},
	"Sandor \"The Hound\" Clegane":   {Liege: "Joffrey Baratheon", SwornTo: "House Baratheon"},
	"Gregor Clegane":                 {Liege: "Tywin Lannister", SwornTo: "House Lannister"},
	"The High Sparrow":               {Liege: "", SwornTo: "The Faith of the Seven"},
	"Qyburn":                         {Liege: "Cersei Lannister", SwornTo: "House Lannister"},
	"Bronn":                          {Liege: "Tyrion Lannister", SwornTo: "House Lannister"},
	"Podrick Payne":                  {Liege: "Tyrion Lannister", SwornTo: "House Lannister"},
	"Shae":                           {Liege: "Tyrion Lannister", SwornTo: "House Lannister"},

	// Dorne
	"Ellaria Sand": {Liege: "", SwornTo: "House Martell"},

	// Braavos
	"Jaqen H'ghar": {Liege: "", SwornTo: "The Faceless Men"},
	"The Waif":     {Liege: "", SwornTo: "The Faceless Men"},

	// Brotherhood Without Banners
	"Beric Dondarrion": {Liege: "", SwornTo: "Brotherhood Without Banners"},
	"Thoros of Myr":    {Liege: "Beric Dondarrion", SwornTo: "Brotherhood Without Banners"},

	// Volantis
	"Talisa Maegyr": {Liege: "Robb Stark", SwornTo: "House Stark"},

	// Tarth
	"Brienne of Tarth": {Liege: "Catelyn Stark", SwornTo: "House Stark"},

	// House Reed
	"Jojen Reed": {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},
	"Meera Reed": {Liege: "Eddard \"Ned\" Stark", SwornTo: "House Stark"},

	// House Tarly
	// Samwell is under Night's Watch

	// Dothraki handmaidens
	"Irri":   {Liege: "Daenerys Targaryen", SwornTo: "House Targaryen"},
	"Doreah": {Liege: "Daenerys Targaryen", SwornTo: "House Targaryen"},

	// Commoners
	"Gendry":  {Liege: "", SwornTo: ""},
	"Hot Pie": {Liege: "", SwornTo: ""},
	"Ros":     {Liege: "", SwornTo: ""},

	// Night's Watch temporary
	"Little Sam": {Liege: "Samwell Tarly", SwornTo: ""},
}
