package augmented

import (
	"strings"
	"testing"

	got "github.com/grokify/gameofthrones"
)

func TestGetCharacters(t *testing.T) {
	chars, err := GetCharacters()
	if err != nil {
		t.Fatalf("GetCharacters() error: %v", err)
	}

	if chars.Count() == 0 {
		t.Fatal("expected characters, got none")
	}

	// Verify phone numbers are unique
	phoneNumbers := make(map[string]string)
	for _, char := range chars.CharactersMap {
		phone := char.GetPhone()
		if phone != "" {
			if existingName, exists := phoneNumbers[phone]; exists {
				t.Errorf("duplicate phone number %s: %s and %s", phone, existingName, char.DisplayName())
			}
			phoneNumbers[phone] = char.DisplayName()

			// Verify E164 format
			if !strings.HasPrefix(phone, "+") {
				t.Errorf("character %q: phone should be E164 format (start with '+'), got %q", char.DisplayName(), phone)
			}
		}

		// Verify email is present
		email := char.GetEmail()
		if email == "" {
			t.Errorf("character %q: no email address", char.DisplayName())
		}

		// Verify email format
		if !strings.Contains(email, "@") {
			t.Errorf("character %q: invalid email format %q", char.DisplayName(), email)
		}
	}
}

func TestCharactersSorted(t *testing.T) {
	chars, err := GetCharacters()
	if err != nil {
		t.Fatalf("GetCharacters() error: %v", err)
	}

	names := chars.NamesSorted()
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

func TestGetOrganizations(t *testing.T) {
	orgs, err := GetOrganizations()
	if err != nil {
		t.Fatalf("GetOrganizations() error: %v", err)
	}

	if len(orgs.OrganizationsMap) == 0 {
		t.Fatal("expected organizations, got none")
	}

	// Verify each organization has required fields
	for name, org := range orgs.OrganizationsMap {
		if org.Organization.Name == "" {
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
		if !strings.HasPrefix(e164, "+") {
			t.Errorf("organization %q: E164() should start with '+', got %q", name, e164)
		}
	}
}

func TestGetAugmentedResources(t *testing.T) {
	resources := GetAugmentedResources("Winterfell")
	if len(resources) == 0 {
		t.Fatal("expected resources for Winterfell, got none")
	}

	for _, r := range resources {
		if r.Name == "" {
			t.Error("resource name is empty")
		}
		if r.RiskLevel < 1 || r.RiskLevel > 5 {
			t.Errorf("resource %q: risk level %d out of range (1-5)", r.Name, r.RiskLevel)
		}
	}
}

func TestResourceRiskLevels(t *testing.T) {
	// Verify critical resources have high risk levels
	criticalResources := []string{"Treasury", "Wildfire Cache", "Iron Throne"}
	for _, name := range criticalResources {
		level := GetRiskLevel(name)
		if level < 5 {
			t.Errorf("critical resource %q: expected risk level 5, got %d", name, level)
		}
	}

	// Verify unknown resources get default level
	unknownLevel := GetRiskLevel("Unknown Resource")
	if unknownLevel != 2 {
		t.Errorf("unknown resource: expected default risk level 2, got %d", unknownLevel)
	}
}

func TestHighRiskResources(t *testing.T) {
	highRisk := HighRiskResources("King's Landing")
	if len(highRisk) == 0 {
		t.Fatal("expected high risk resources for King's Landing, got none")
	}

	for _, r := range highRisk {
		if r.RiskLevel < 4 {
			t.Errorf("high risk resource %q: expected risk level >= 4, got %d", r.Name, r.RiskLevel)
		}
	}
}

func TestCriticalResources(t *testing.T) {
	critical := CriticalResources("King's Landing")
	if len(critical) == 0 {
		t.Fatal("expected critical resources for King's Landing, got none")
	}

	for _, r := range critical {
		if r.RiskLevel != 5 {
			t.Errorf("critical resource %q: expected risk level 5, got %d", r.Name, r.RiskLevel)
		}
	}
}

func TestGetCharactersByOrganization(t *testing.T) {
	byOrg, err := GetCharactersByOrganization()
	if err != nil {
		t.Fatalf("GetCharactersByOrganization() error: %v", err)
	}

	if len(byOrg) == 0 {
		t.Fatal("expected organization groups, got none")
	}

	// Winterfell should have Stark family members
	winterfell, ok := byOrg["Winterfell"]
	if !ok {
		t.Error("expected Winterfell in organization groups")
	} else if len(winterfell) < 3 {
		t.Errorf("expected at least 3 Winterfell characters, got %d", len(winterfell))
	}
}

func TestGetCharactersByStation(t *testing.T) {
	byStation, err := GetCharactersByStation()
	if err != nil {
		t.Fatalf("GetCharactersByStation() error: %v", err)
	}

	if len(byStation) == 0 {
		t.Fatal("expected station groups, got none")
	}

	// Verify we have characters in multiple stations
	stations := []got.Station{got.StationRoyal, got.StationLord, got.StationKnight, got.StationSmallfolk}
	for _, station := range stations {
		chars, ok := byStation[station]
		if !ok || len(chars) == 0 {
			t.Errorf("expected characters in station %q", station)
		}
	}
}

func TestCharacterSCIMUser(t *testing.T) {
	chars, err := GetCharacters()
	if err != nil {
		t.Fatalf("GetCharacters() error: %v", err)
	}

	// Test first character
	for _, char := range chars.CharactersMap {
		user := char.SCIMUser()
		if user.DisplayName == "" {
			t.Error("SCIM user display name is empty")
		}
		if len(user.Emails) == 0 {
			t.Error("SCIM user has no emails")
		}
		break // Just test one
	}
}

func TestDomainForOrganization(t *testing.T) {
	tests := []struct {
		org    string
		suffix string
	}{
		{"Night's Watch", ".org"},
		{"Free Folk", ".org"},
		{"Winterfell", ".com"},
		{"King's Landing", ".com"},
	}

	for _, tc := range tests {
		t.Run(tc.org, func(t *testing.T) {
			domain := DomainForOrganization(tc.org)
			if !strings.HasSuffix(domain, tc.suffix) {
				t.Errorf("organization %q: expected domain ending with %q, got %q", tc.org, tc.suffix, domain)
			}
		})
	}
}

func TestFilterResourcesByRiskLevel(t *testing.T) {
	resources := GetAugmentedResources("King's Landing")

	// Filter for high risk (>= 4)
	filtered := FilterResourcesByRiskLevel(resources, 4)

	for _, r := range filtered {
		if r.RiskLevel < 4 {
			t.Errorf("filtered resource %q has risk level %d, expected >= 4", r.Name, r.RiskLevel)
		}
	}
}

func TestGetAllAugmentedResources(t *testing.T) {
	all := GetAllAugmentedResources()
	if len(all) == 0 {
		t.Fatal("expected resources for all organizations, got none")
	}

	// Should have resources for major locations
	for _, loc := range []string{"Winterfell", "King's Landing", "Casterly Rock"} {
		if _, ok := all[loc]; !ok {
			t.Errorf("expected resources for %q", loc)
		}
	}
}
