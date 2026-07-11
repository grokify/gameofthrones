package gameofthrones

import (
	"testing"
)

func TestCharacters(t *testing.T) {
	chars := Characters()
	if len(chars) == 0 {
		t.Fatal("expected characters, got none")
	}

	// Verify we have a reasonable number of characters
	if len(chars) < 50 {
		t.Errorf("expected at least 50 characters, got %d", len(chars))
	}

	// Verify character structure is populated
	for i, char := range chars {
		if char.Actor.DisplayName == "" {
			t.Errorf("character %d: actor display name is empty", i)
		}
		if char.Character.DisplayName == "" {
			t.Errorf("character %d: character display name is empty", i)
		}
	}
}

func TestNewCharacterSimple(t *testing.T) {
	opts := NewCharacterSimpleOpts{
		ActorName:       "Sean Bean",
		GivenName:       "Eddard",
		FamilyName:      "Stark",
		NickName:        "Ned",
		AddOrganization: true,
	}
	char := NewCharacterSimple(opts)

	if char.Actor.DisplayName != "Sean Bean" {
		t.Errorf("expected actor 'Sean Bean', got %q", char.Actor.DisplayName)
	}
	if char.Character.Name.GivenName != "Eddard" {
		t.Errorf("expected given name 'Eddard', got %q", char.Character.Name.GivenName)
	}
	if char.Character.Name.FamilyName != "Stark" {
		t.Errorf("expected family name 'Stark', got %q", char.Character.Name.FamilyName)
	}
	if char.Character.NickName != "Ned" {
		t.Errorf("expected nickname 'Ned', got %q", char.Character.NickName)
	}
	if char.Organization.Name != "Winterfell" {
		t.Errorf("expected organization 'Winterfell', got %q", char.Organization.Name)
	}
}

func TestGetOrganizationForUser(t *testing.T) {
	tests := []struct {
		familyName string
		wantOrg    string
	}{
		{"Stark", "Winterfell"},
		{"Lannister", "Casterly Rock"},
		{"Targaryen", "Dragonstone"},
		{"Snow", "Winterfell"},
		{"Unknown", ""},
	}

	for _, tc := range tests {
		t.Run(tc.familyName, func(t *testing.T) {
			user := NewCharacterSimple(NewCharacterSimpleOpts{
				FamilyName:      tc.familyName,
				AddOrganization: false,
			})
			org := GetOrganizationForUser(user.Character)
			if org.Name != tc.wantOrg {
				t.Errorf("GetOrganizationForUser(%q) = %q, want %q", tc.familyName, org.Name, tc.wantOrg)
			}
		})
	}
}

func TestOrganizations(t *testing.T) {
	if len(Organizations) == 0 {
		t.Fatal("expected organizations, got none")
	}

	// Check for expected organizations
	expected := []string{"Winterfell", "King's Landing", "Casterly Rock", "Dragonstone"}
	orgSet := make(map[string]struct{}, len(Organizations))
	for _, org := range Organizations {
		orgSet[org] = struct{}{}
	}

	for _, exp := range expected {
		if _, ok := orgSet[exp]; !ok {
			t.Errorf("expected organization %q not found", exp)
		}
	}
}
