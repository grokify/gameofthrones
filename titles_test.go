package gameofthrones

import (
	"testing"
)

func TestTitleCategories(t *testing.T) {
	categories := TitleCategories()
	if len(categories) == 0 {
		t.Fatal("expected title categories, got none")
	}

	expected := []TitleCategory{
		TitleCategoryGovernance,
		TitleCategoryMilitary,
		TitleCategoryReligious,
		TitleCategoryAcademic,
		TitleCategoryCommerce,
	}

	if len(categories) != len(expected) {
		t.Errorf("expected %d categories, got %d", len(expected), len(categories))
	}
}

func TestCharacterTitles(t *testing.T) {
	if len(CharacterTitles) == 0 {
		t.Fatal("expected character titles, got none")
	}

	// Verify title structure
	validStations := make(map[Station]struct{})
	for _, s := range Stations() {
		validStations[s] = struct{}{}
	}

	validCategories := make(map[TitleCategory]struct{})
	for _, c := range TitleCategories() {
		validCategories[c] = struct{}{}
	}

	for name, titles := range CharacterTitles {
		if len(titles) == 0 {
			t.Errorf("character %q has empty titles slice", name)
			continue
		}

		for i, title := range titles {
			if title.Name == "" {
				t.Errorf("character %q title %d: name is empty", name, i)
			}
			if _, ok := validStations[title.Station]; !ok {
				t.Errorf("character %q title %d (%s): invalid station %q", name, i, title.Name, title.Station)
			}
			if _, ok := validCategories[title.Category]; !ok {
				t.Errorf("character %q title %d (%s): invalid category %q", name, i, title.Name, title.Category)
			}
		}
	}
}

func TestKnownCharacterTitles(t *testing.T) {
	tests := []struct {
		name       string
		titleCount int
		hasTitle   string
	}{
		{"Eddard \"Ned\" Stark", 3, "Hand of the King"},
		{"Jon Snow", 3, "Lord Commander of the Night's Watch"},
		{"Daenerys Targaryen", 5, "Mother of Dragons"},
		{"Jaime Lannister", 3, "Kingslayer"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			titles, ok := CharacterTitles[tc.name]
			if !ok {
				t.Errorf("character %q not found in CharacterTitles", tc.name)
				return
			}
			if len(titles) < tc.titleCount {
				t.Errorf("character %q: expected at least %d titles, got %d", tc.name, tc.titleCount, len(titles))
			}

			found := false
			for _, title := range titles {
				if title.Name == tc.hasTitle {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("character %q: expected to have title %q", tc.name, tc.hasTitle)
			}
		})
	}
}

func TestTitleCategoryDistribution(t *testing.T) {
	distribution := make(map[TitleCategory]int)
	for _, titles := range CharacterTitles {
		for _, title := range titles {
			distribution[title.Category]++
		}
	}

	// Governance and Military should be most common
	if distribution[TitleCategoryGovernance] == 0 {
		t.Error("expected governance titles")
	}
	if distribution[TitleCategoryMilitary] == 0 {
		t.Error("expected military titles")
	}
}
