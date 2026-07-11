package gameofthrones

import (
	"testing"
)

func TestLifeStatuses(t *testing.T) {
	statuses := LifeStatuses()
	if len(statuses) == 0 {
		t.Fatal("expected life statuses, got none")
	}

	expected := []LifeStatus{StatusLiving, StatusDeceased}
	if len(statuses) != len(expected) {
		t.Errorf("expected %d statuses, got %d", len(expected), len(statuses))
	}
}

func TestCharacterStatus(t *testing.T) {
	if len(CharacterStatus) == 0 {
		t.Fatal("expected character status mappings, got none")
	}

	// Verify all statuses used are valid
	validStatuses := make(map[LifeStatus]struct{})
	for _, s := range LifeStatuses() {
		validStatuses[s] = struct{}{}
	}

	for name, status := range CharacterStatus {
		if _, ok := validStatuses[status]; !ok {
			t.Errorf("character %q has invalid status %q", name, status)
		}
	}

	// Verify some known characters have correct statuses
	tests := []struct {
		name   string
		status LifeStatus
	}{
		{"Jon Snow", StatusLiving},
		{"Sansa Stark", StatusLiving},
		{"Arya Stark", StatusLiving},
		{"Bran Stark", StatusLiving},
		{"Tyrion Lannister", StatusLiving},
		{"Eddard \"Ned\" Stark", StatusDeceased},
		{"Robert Baratheon", StatusDeceased},
		{"Cersei Lannister", StatusDeceased},
		{"Daenerys Targaryen", StatusDeceased},
		{"Joffrey Baratheon", StatusDeceased},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			status, ok := CharacterStatus[tc.name]
			if !ok {
				t.Errorf("character %q not found in CharacterStatus", tc.name)
				return
			}
			if status != tc.status {
				t.Errorf("character %q: expected status %q, got %q", tc.name, tc.status, status)
			}
		})
	}
}

func TestStatusDistribution(t *testing.T) {
	living := 0
	deceased := 0
	for _, status := range CharacterStatus {
		switch status {
		case StatusLiving:
			living++
		case StatusDeceased:
			deceased++
		}
	}

	// Verify we have both living and deceased characters
	if living == 0 {
		t.Error("expected some living characters")
	}
	if deceased == 0 {
		t.Error("expected some deceased characters")
	}

	// Game of Thrones is known for killing characters
	if deceased < living {
		t.Logf("NOTE: More living (%d) than deceased (%d) characters - this is unusual for GoT", living, deceased)
	}
}
