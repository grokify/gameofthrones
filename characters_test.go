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

func TestGetDemoOrganizations(t *testing.T) {
	demoOrgs, err := GetDemoOrganizations()
	if err != nil {
		t.Fatalf("GetDemoOrganizations() error: %v", err)
	}

	if len(demoOrgs.OrganizationsMap) == 0 {
		t.Fatal("expected demo organizations, got none")
	}

	// Verify each organization has required fields
	for name, org := range demoOrgs.OrganizationsMap {
		if org.Name == "" {
			t.Errorf("organization %q: name is empty", name)
		}
		if org.AreaCode == 0 {
			t.Errorf("organization %q: area code is zero", name)
		}
		if org.Phone == 0 {
			t.Errorf("organization %q: phone is zero", name)
		}
		if org.Domain == "" {
			t.Errorf("organization %q: domain is empty", name)
		}

		// Verify E164 format
		e164 := org.E164()
		if e164 == "" {
			t.Errorf("organization %q: E164() returned empty string", name)
		}
		if e164[0] != '+' {
			t.Errorf("organization %q: E164() should start with '+', got %q", name, e164)
		}
	}

	// Verify area codes are spread across organizations (using stride)
	areaCodes := make(map[uint16]struct{})
	for _, org := range demoOrgs.OrganizationsMap {
		areaCodes[org.AreaCode] = struct{}{}
	}
	if len(areaCodes) < len(demoOrgs.OrganizationsMap)/2 {
		t.Error("expected area codes to be diverse across organizations")
	}
}

func TestGetDemoCharacters(t *testing.T) {
	demoChars, err := GetDemoCharacters()
	if err != nil {
		t.Fatalf("GetDemoCharacters() error: %v", err)
	}

	if len(demoChars.CharactersMap) == 0 {
		t.Fatal("expected demo characters, got none")
	}

	// Verify phone numbers are unique
	phoneNumbers := make(map[string]string)
	for name, char := range demoChars.CharactersMap {
		if len(char.Character.PhoneNumbers) > 0 {
			phone := char.Character.PhoneNumbers[0].Value
			if existingName, exists := phoneNumbers[phone]; exists {
				t.Errorf("duplicate phone number %s: %s and %s", phone, existingName, name)
			}
			phoneNumbers[phone] = name

			// Verify E164 format
			if phone[0] != '+' {
				t.Errorf("character %q: phone should be E164 format (start with '+'), got %q", name, phone)
			}
		}

		// Verify email is present
		if len(char.Character.Emails) == 0 {
			t.Errorf("character %q: no email address", name)
		}
	}
}

func TestDemoCharactersNamesSorted(t *testing.T) {
	demoChars, err := GetDemoCharacters()
	if err != nil {
		t.Fatalf("GetDemoCharacters() error: %v", err)
	}

	names := demoChars.NamesSorted()
	if len(names) == 0 {
		t.Fatal("expected sorted names, got none")
	}

	// Verify names are sorted
	for i := 1; i < len(names); i++ {
		if names[i-1] > names[i] {
			t.Errorf("names not sorted: %q > %q", names[i-1], names[i])
		}
	}
}
