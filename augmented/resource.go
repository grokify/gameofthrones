package augmented

import (
	got "github.com/grokify/gameofthrones"
)

// Resource embeds canonical GoT resource data and adds risk level for demos.
type Resource struct {
	got.Resource
	RiskLevel int `json:"risk_level"` // 1-5 scale
}

// NewResource creates an augmented Resource from a base Resource.
func NewResource(base got.Resource, riskLevel int) Resource {
	return Resource{
		Resource:  base,
		RiskLevel: riskLevel,
	}
}

// ResourceRiskLevels maps resource names to their risk levels (1-5 scale).
// Higher values indicate higher risk/sensitivity.
//
//	5 - Critical: Catastrophic impact if compromised (weapons of mass destruction, royal power)
//	4 - High: Significant impact (military assets, financial control)
//	3 - Medium: Moderate impact (sensitive locations, security)
//	2 - Low: Minor impact (communication, general knowledge)
//	1 - Minimal: Negligible impact (common areas, public spaces)
var ResourceRiskLevels = map[string]int{
	// Critical (5)
	"Treasury":       5,
	"Wildfire Cache": 5,
	"Iron Throne":    5,
	"Dragonglass":    5,
	"Gold Mines":     5,
	"Iron Bank":      5,

	// High (4)
	"Armory":                4,
	"Small Council Chamber": 4,
	"Dungeons":              4,
	"Red Keep":              4,
	"Iron Fleet":            4,
	"Painted Table":         4,
	"The Wall":              4,
	"Heartsbane":            4,
	"Faith Militant":        4,
	"Bloody Gate":           4,
	"Moon Door":             4,
	"Sky Cells":             4,
	"Mercenary Camp":        4,
	"War Chest":             4,
	"Khalasar":              4,
	"Giants":                4,
	"Arsenal":               4,

	// Medium (3)
	"Crypts":               3,
	"Great Hall":           3,
	"Castle Black":         3,
	"Eastwatch-by-the-Sea": 3,
	"The Shadow Tower":     3,
	"Pyke":                 3,
	"Salt Throne":          3,
	"Shipyards":            3,
	"Sunspear":             3,
	"The Eyrie":            3,
	"Riverrun":             3,
	"The Twins":            3,
	"Harrenhal":            3,
	"Mormont Keep":         3,
	"Evenfall Hall":        3,
	"Harbor":               3,
	"Throne Room":          3,
	"Gold Cloaks Barracks": 3,
	"Black Walls":          3,
	"Titan of Braavos":     3,
	"Crossbow Workshops":   3,
	"Hardhome":             3,
	"Slave Markets":        3,

	// Low (2)
	"Ravens":        2,
	"Library":       2,
	"Godswood":      2,
	"Gardens":       2,
	"Granaries":     2,
	"Winter Town":   2,
	"Mess Hall":     2,
	"The Citadel":   2,
	"Healing Halls": 2,
	"Archives":      2,
	"Myrish Lenses": 2,
	"Warg Network":  2,
	"Horse Herds":   2,
	"Sapphire Isle": 2,

	// Minimal (1)
	"Great Sept of Baelor":     1,
	"House of Black and White": 1,
	"Vaes Dothrak":             1,
	"Water Gardens":            1,
	"Red Temple":               1,
}

// GetRiskLevel returns the risk level for a resource name.
// Returns 2 (low) as default if the resource is not found.
func GetRiskLevel(resourceName string) int {
	if level, ok := ResourceRiskLevels[resourceName]; ok {
		return level
	}
	return 2 // Default to low risk
}

// GetAugmentedResources returns all resources for an organization with risk levels.
func GetAugmentedResources(org string) []Resource {
	baseResources, ok := got.OrganizationResources[org]
	if !ok {
		return nil
	}

	augmented := make([]Resource, 0, len(baseResources))
	for _, base := range baseResources {
		riskLevel := GetRiskLevel(base.Name)
		augmented = append(augmented, NewResource(base, riskLevel))
	}
	return augmented
}

// GetAllAugmentedResources returns all resources across all organizations with risk levels.
func GetAllAugmentedResources() map[string][]Resource {
	result := make(map[string][]Resource)
	for org := range got.OrganizationResources {
		result[org] = GetAugmentedResources(org)
	}
	return result
}

// FilterResourcesByRiskLevel returns resources with risk level >= minRisk.
func FilterResourcesByRiskLevel(resources []Resource, minRisk int) []Resource {
	var filtered []Resource
	for _, r := range resources {
		if r.RiskLevel >= minRisk {
			filtered = append(filtered, r)
		}
	}
	return filtered
}

// HighRiskResources returns all resources with risk level >= 4.
func HighRiskResources(org string) []Resource {
	return FilterResourcesByRiskLevel(GetAugmentedResources(org), 4)
}

// CriticalResources returns all resources with risk level == 5.
func CriticalResources(org string) []Resource {
	return FilterResourcesByRiskLevel(GetAugmentedResources(org), 5)
}
