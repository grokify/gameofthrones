package gameofthrones

import (
	"testing"
)

func TestOrders(t *testing.T) {
	if len(Orders) == 0 {
		t.Fatal("expected orders, got none")
	}

	// Check for expected orders
	expected := []string{
		"The Kingsguard",
		"The Night's Watch",
		"The Citadel",
		"The Small Council",
	}

	orderSet := make(map[string]struct{})
	for _, order := range Orders {
		orderSet[order] = struct{}{}
	}

	for _, exp := range expected {
		if _, ok := orderSet[exp]; !ok {
			t.Errorf("expected order %q not found", exp)
		}
	}
}

func TestCharacterOrders(t *testing.T) {
	if len(CharacterOrders) == 0 {
		t.Fatal("expected character order mappings, got none")
	}

	// Create set of valid orders
	validOrders := make(map[string]struct{})
	for _, order := range Orders {
		validOrders[order] = struct{}{}
	}

	// Verify all orders assigned to characters are valid
	for name, orders := range CharacterOrders {
		if len(orders) == 0 {
			t.Errorf("character %q has empty orders slice", name)
			continue
		}

		for _, order := range orders {
			if _, ok := validOrders[order]; !ok {
				t.Errorf("character %q has invalid order %q", name, order)
			}
		}
	}
}

func TestKnownCharacterOrders(t *testing.T) {
	tests := []struct {
		name     string
		hasOrder string
	}{
		{"Jon Snow", "The Night's Watch"},
		{"Samwell Tarly", "The Night's Watch"},
		{"Jaime Lannister", "The Kingsguard"},
		{"Barristan Selmy", "The Kingsguard"},
		{"Tyrion Lannister", "The Small Council"},
		{"Grey Worm", "The Unsullied"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			orders, ok := CharacterOrders[tc.name]
			if !ok {
				t.Errorf("character %q not found in CharacterOrders", tc.name)
				return
			}

			found := false
			for _, order := range orders {
				if order == tc.hasOrder {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("character %q: expected to belong to order %q", tc.name, tc.hasOrder)
			}
		})
	}
}

func TestOrderMembership(t *testing.T) {
	// Count members per order
	orderMembers := make(map[string]int)
	for _, orders := range CharacterOrders {
		for _, order := range orders {
			orderMembers[order]++
		}
	}

	// Night's Watch should have multiple members
	if orderMembers["The Night's Watch"] < 5 {
		t.Errorf("expected at least 5 Night's Watch members, got %d", orderMembers["The Night's Watch"])
	}
}
