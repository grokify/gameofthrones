package gameofthrones

import (
	"testing"
)

func TestResourceCategories(t *testing.T) {
	categories := ResourceCategories()
	if len(categories) == 0 {
		t.Fatal("expected resource categories, got none")
	}

	expected := []ResourceCategory{
		ResourceCategoryCommunication,
		ResourceCategoryWeapons,
		ResourceCategoryFinance,
		ResourceCategoryKnowledge,
		ResourceCategorySacred,
		ResourceCategorySecurity,
		ResourceCategoryGovernance,
		ResourceCategoryMilitary,
	}

	if len(categories) != len(expected) {
		t.Errorf("expected %d categories, got %d", len(expected), len(categories))
	}
}

func TestOrganizationResources(t *testing.T) {
	if len(OrganizationResources) == 0 {
		t.Fatal("expected organization resources, got none")
	}

	// Verify resource structure
	validCategories := make(map[ResourceCategory]struct{})
	for _, c := range ResourceCategories() {
		validCategories[c] = struct{}{}
	}
	// Add Commerce which is defined separately
	validCategories[ResourceCategoryCommerce] = struct{}{}

	for org, resources := range OrganizationResources {
		if len(resources) == 0 {
			t.Errorf("organization %q has empty resources slice", org)
			continue
		}

		for i, resource := range resources {
			if resource.Name == "" {
				t.Errorf("organization %q resource %d: name is empty", org, i)
			}
			if _, ok := validCategories[resource.Category]; !ok {
				t.Errorf("organization %q resource %d (%s): invalid category %q", org, i, resource.Name, resource.Category)
			}
		}
	}
}

func TestKnownOrganizationResources(t *testing.T) {
	tests := []struct {
		org         string
		hasResource string
		category    ResourceCategory
	}{
		{"Winterfell", "Ravens", ResourceCategoryCommunication},
		{"Winterfell", "Godswood", ResourceCategorySacred},
		{"King's Landing", "Iron Throne", ResourceCategoryGovernance},
		{"King's Landing", "Wildfire Cache", ResourceCategoryWeapons},
		{"Casterly Rock", "Gold Mines", ResourceCategoryFinance},
		{"Dragonstone", "Dragonglass", ResourceCategoryWeapons},
	}

	for _, tc := range tests {
		t.Run(tc.org+"_"+tc.hasResource, func(t *testing.T) {
			resources, ok := OrganizationResources[tc.org]
			if !ok {
				t.Errorf("organization %q not found in OrganizationResources", tc.org)
				return
			}

			found := false
			for _, r := range resources {
				if r.Name == tc.hasResource {
					found = true
					if r.Category != tc.category {
						t.Errorf("resource %q: expected category %q, got %q", tc.hasResource, tc.category, r.Category)
					}
					break
				}
			}
			if !found {
				t.Errorf("organization %q: expected to have resource %q", tc.org, tc.hasResource)
			}
		})
	}
}

func TestResourceCategoryDistribution(t *testing.T) {
	distribution := make(map[ResourceCategory]int)
	for _, resources := range OrganizationResources {
		for _, r := range resources {
			distribution[r.Category]++
		}
	}

	// Communication should be common (ravens everywhere)
	if distribution[ResourceCategoryCommunication] < 5 {
		t.Errorf("expected at least 5 communication resources, got %d", distribution[ResourceCategoryCommunication])
	}

	// Governance should be common
	if distribution[ResourceCategoryGovernance] < 5 {
		t.Errorf("expected at least 5 governance resources, got %d", distribution[ResourceCategoryGovernance])
	}
}

func TestMajorLocationsHaveResources(t *testing.T) {
	majorLocations := []string{
		"Winterfell",
		"King's Landing",
		"Casterly Rock",
		"Dragonstone",
		"Highgarden",
	}

	for _, loc := range majorLocations {
		resources, ok := OrganizationResources[loc]
		if !ok {
			t.Errorf("major location %q not found in OrganizationResources", loc)
			continue
		}
		if len(resources) < 3 {
			t.Errorf("major location %q: expected at least 3 resources, got %d", loc, len(resources))
		}
	}
}
