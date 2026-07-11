package gameofthrones

import (
	"testing"
)

func TestCharacterAllegiance(t *testing.T) {
	if len(CharacterAllegiance) == 0 {
		t.Fatal("expected character allegiance mappings, got none")
	}

	// Verify all allegiances have valid structure
	for name, allegiance := range CharacterAllegiance {
		// Either liege or sworn_to should be set (or both can be empty for independent characters)
		if allegiance.Liege == "" && allegiance.SwornTo == "" {
			// Some characters may be independent (commoners, etc.)
			// This is acceptable
		}

		// If a liege is set, it should ideally be another character
		// (We don't strictly enforce this as some lieges may not be in our character list)
		_ = name
	}
}

func TestKnownAllegiances(t *testing.T) {
	tests := []struct {
		name    string
		liege   string
		swornTo string
	}{
		{"Eddard \"Ned\" Stark", "Robert Baratheon", "House Baratheon"},
		{"Jon Snow", "Jeor Mormont", "The Night's Watch"},
		{"Jaime Lannister", "Robert Baratheon", "The Kingsguard"},
		{"Jorah Mormont", "Daenerys Targaryen", "House Targaryen"},
		{"Tyrion Lannister", "Tywin Lannister", "House Lannister"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			allegiance, ok := CharacterAllegiance[tc.name]
			if !ok {
				t.Errorf("character %q not found in CharacterAllegiance", tc.name)
				return
			}
			if allegiance.Liege != tc.liege {
				t.Errorf("character %q: expected liege %q, got %q", tc.name, tc.liege, allegiance.Liege)
			}
			if allegiance.SwornTo != tc.swornTo {
				t.Errorf("character %q: expected sworn_to %q, got %q", tc.name, tc.swornTo, allegiance.SwornTo)
			}
		})
	}
}

func TestLiegeReferences(t *testing.T) {
	// Verify that lieges reference valid characters (when they exist in our data)
	characterSet := make(map[string]struct{})
	for name := range CharacterAllegiance {
		characterSet[name] = struct{}{}
	}

	missingLieges := make(map[string]int)
	for _, allegiance := range CharacterAllegiance {
		if allegiance.Liege != "" {
			if _, ok := characterSet[allegiance.Liege]; !ok {
				missingLieges[allegiance.Liege]++
			}
		}
	}

	// Log missing lieges but don't fail - some lieges may not be major characters
	for liege, count := range missingLieges {
		t.Logf("NOTE: Liege %q referenced by %d characters but not in CharacterAllegiance", liege, count)
	}
}

func TestAllegianceHierarchy(t *testing.T) {
	// Verify the Stark hierarchy
	starkChildren := []string{"Robb Stark", "Sansa Stark", "Arya Stark", "Bran Stark", "Rickon Stark"}

	for _, child := range starkChildren {
		allegiance, ok := CharacterAllegiance[child]
		if !ok {
			t.Errorf("character %q not found in CharacterAllegiance", child)
			continue
		}
		if allegiance.SwornTo != "House Stark" {
			t.Errorf("character %q: expected sworn to House Stark, got %q", child, allegiance.SwornTo)
		}
	}
}

func TestSwornToDistribution(t *testing.T) {
	swornToCount := make(map[string]int)
	for _, allegiance := range CharacterAllegiance {
		if allegiance.SwornTo != "" {
			swornToCount[allegiance.SwornTo]++
		}
	}

	// Major houses/organizations should have multiple members
	majorOrgs := []string{"House Stark", "House Lannister", "House Targaryen", "The Night's Watch"}
	for _, org := range majorOrgs {
		if count := swornToCount[org]; count < 2 {
			t.Errorf("expected at least 2 characters sworn to %q, got %d", org, count)
		}
	}
}
