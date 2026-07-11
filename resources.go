package gameofthrones

// ResourceCategory represents the type of resource.
type ResourceCategory string

const (
	ResourceCategoryCommunication ResourceCategory = "Communication"
	ResourceCategoryWeapons       ResourceCategory = "Weapons"
	ResourceCategoryFinance       ResourceCategory = "Finance"
	ResourceCategoryKnowledge     ResourceCategory = "Knowledge"
	ResourceCategorySacred        ResourceCategory = "Sacred"
	ResourceCategorySecurity      ResourceCategory = "Security"
	ResourceCategoryGovernance    ResourceCategory = "Governance"
	ResourceCategoryMilitary      ResourceCategory = "Military"
)

// ResourceCategories returns all valid resource category values.
func ResourceCategories() []ResourceCategory {
	return []ResourceCategory{
		ResourceCategoryCommunication,
		ResourceCategoryWeapons,
		ResourceCategoryFinance,
		ResourceCategoryKnowledge,
		ResourceCategorySacred,
		ResourceCategorySecurity,
		ResourceCategoryGovernance,
		ResourceCategoryMilitary,
	}
}

// Resource represents a valuable asset or location within an organization.
type Resource struct {
	Name     string           `json:"name"`
	Category ResourceCategory `json:"category"`
}

// OrganizationResources maps organization names to their available resources.
var OrganizationResources = map[string][]Resource{
	"Winterfell": {
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Armory", Category: ResourceCategoryWeapons},
		{Name: "Crypts", Category: ResourceCategorySacred},
		{Name: "Godswood", Category: ResourceCategorySacred},
		{Name: "Library", Category: ResourceCategoryKnowledge},
		{Name: "Treasury", Category: ResourceCategoryFinance},
		{Name: "Great Hall", Category: ResourceCategoryGovernance},
		{Name: "Winter Town", Category: ResourceCategoryGovernance},
	},
	"King's Landing": {
		{Name: "Iron Throne", Category: ResourceCategoryGovernance},
		{Name: "Wildfire Cache", Category: ResourceCategoryWeapons},
		{Name: "Treasury", Category: ResourceCategoryFinance},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Dungeons", Category: ResourceCategorySecurity},
		{Name: "Small Council Chamber", Category: ResourceCategoryGovernance},
		{Name: "Red Keep", Category: ResourceCategoryGovernance},
		{Name: "Great Sept of Baelor", Category: ResourceCategorySacred},
		{Name: "Gold Cloaks Barracks", Category: ResourceCategoryMilitary},
	},
	"Casterly Rock": {
		{Name: "Gold Mines", Category: ResourceCategoryFinance},
		{Name: "Treasury", Category: ResourceCategoryFinance},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Armory", Category: ResourceCategoryWeapons},
		{Name: "Great Hall", Category: ResourceCategoryGovernance},
		{Name: "Dungeons", Category: ResourceCategorySecurity},
	},
	"Dragonstone": {
		{Name: "Dragonglass", Category: ResourceCategoryWeapons},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Painted Table", Category: ResourceCategoryGovernance},
		{Name: "Harbor", Category: ResourceCategoryMilitary},
		{Name: "Throne Room", Category: ResourceCategoryGovernance},
	},
	"The Wall": {
		{Name: "The Wall", Category: ResourceCategorySecurity},
		{Name: "Castle Black", Category: ResourceCategoryMilitary},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Armory", Category: ResourceCategoryWeapons},
		{Name: "Library", Category: ResourceCategoryKnowledge},
		{Name: "Mess Hall", Category: ResourceCategoryGovernance},
	},
	"Night's Watch": {
		{Name: "Castle Black", Category: ResourceCategoryMilitary},
		{Name: "Eastwatch-by-the-Sea", Category: ResourceCategoryMilitary},
		{Name: "The Shadow Tower", Category: ResourceCategoryMilitary},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Armory", Category: ResourceCategoryWeapons},
		{Name: "The Wall", Category: ResourceCategorySecurity},
	},
	"Highgarden": {
		{Name: "Treasury", Category: ResourceCategoryFinance},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Gardens", Category: ResourceCategorySacred},
		{Name: "Great Hall", Category: ResourceCategoryGovernance},
		{Name: "Granaries", Category: ResourceCategoryGovernance},
	},
	"Iron Islands": {
		{Name: "Iron Fleet", Category: ResourceCategoryMilitary},
		{Name: "Pyke", Category: ResourceCategoryGovernance},
		{Name: "Salt Throne", Category: ResourceCategoryGovernance},
		{Name: "Shipyards", Category: ResourceCategoryMilitary},
	},
	"Dorne": {
		{Name: "Water Gardens", Category: ResourceCategorySacred},
		{Name: "Sunspear", Category: ResourceCategoryGovernance},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Treasury", Category: ResourceCategoryFinance},
	},
	"The Vale": {
		{Name: "The Eyrie", Category: ResourceCategoryGovernance},
		{Name: "Moon Door", Category: ResourceCategorySecurity},
		{Name: "Sky Cells", Category: ResourceCategorySecurity},
		{Name: "Bloody Gate", Category: ResourceCategoryMilitary},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
	},
	"The Riverlands": {
		{Name: "Riverrun", Category: ResourceCategoryGovernance},
		{Name: "The Twins", Category: ResourceCategoryGovernance},
		{Name: "Harrenhal", Category: ResourceCategoryGovernance},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
	},
	"Bear Island": {
		{Name: "Mormont Keep", Category: ResourceCategoryGovernance},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Armory", Category: ResourceCategoryWeapons},
	},
	"The Dreadfort": {
		{Name: "Dungeons", Category: ResourceCategorySecurity},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Great Hall", Category: ResourceCategoryGovernance},
	},
	"Horn Hill": {
		{Name: "Heartsbane", Category: ResourceCategoryWeapons},
		{Name: "Library", Category: ResourceCategoryKnowledge},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Treasury", Category: ResourceCategoryFinance},
	},
	"Order of Maesters": {
		{Name: "The Citadel", Category: ResourceCategoryKnowledge},
		{Name: "Library", Category: ResourceCategoryKnowledge},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
		{Name: "Healing Halls", Category: ResourceCategoryKnowledge},
		{Name: "Archives", Category: ResourceCategoryKnowledge},
	},
	"The Sparrows": {
		{Name: "Great Sept of Baelor", Category: ResourceCategorySacred},
		{Name: "Faith Militant", Category: ResourceCategoryMilitary},
		{Name: "Dungeons", Category: ResourceCategorySecurity},
	},
	"Braavos": {
		{Name: "Iron Bank", Category: ResourceCategoryFinance},
		{Name: "House of Black and White", Category: ResourceCategorySacred},
		{Name: "Titan of Braavos", Category: ResourceCategorySecurity},
		{Name: "Arsenal", Category: ResourceCategoryMilitary},
	},
	"Dothraki": {
		{Name: "Vaes Dothrak", Category: ResourceCategorySacred},
		{Name: "Khalasar", Category: ResourceCategoryMilitary},
		{Name: "Horse Herds", Category: ResourceCategoryMilitary},
	},
	"Second Sons": {
		{Name: "Mercenary Camp", Category: ResourceCategoryMilitary},
		{Name: "War Chest", Category: ResourceCategoryFinance},
	},
	"Free Folk": {
		{Name: "Hardhome", Category: ResourceCategoryGovernance},
		{Name: "Warg Network", Category: ResourceCategoryCommunication},
		{Name: "Giants", Category: ResourceCategoryMilitary},
	},
	"Myr": {
		{Name: "Myrish Lenses", Category: ResourceCategoryKnowledge},
		{Name: "Crossbow Workshops", Category: ResourceCategoryWeapons},
	},
	"Volantis": {
		{Name: "Black Walls", Category: ResourceCategorySecurity},
		{Name: "Slave Markets", Category: ResourceCategoryCommerce},
		{Name: "Red Temple", Category: ResourceCategorySacred},
	},
	"Tarth": {
		{Name: "Evenfall Hall", Category: ResourceCategoryGovernance},
		{Name: "Sapphire Isle", Category: ResourceCategoryFinance},
		{Name: "Ravens", Category: ResourceCategoryCommunication},
	},
}

// ResourceCategoryCommerce is defined for completeness but not in the main constants
// as it appears only in Volantis resources.
const ResourceCategoryCommerce ResourceCategory = "Commerce"
