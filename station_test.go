package gameofthrones

import (
	"testing"
)

func TestStations(t *testing.T) {
	stations := Stations()
	if len(stations) == 0 {
		t.Fatal("expected stations, got none")
	}

	// Verify we have all expected stations
	expected := []Station{StationSmallfolk, StationKnight, StationLord, StationGreatLord, StationRoyal}
	if len(stations) != len(expected) {
		t.Errorf("expected %d stations, got %d", len(expected), len(stations))
	}

	stationSet := make(map[Station]struct{})
	for _, s := range stations {
		stationSet[s] = struct{}{}
	}

	for _, exp := range expected {
		if _, ok := stationSet[exp]; !ok {
			t.Errorf("expected station %q not found", exp)
		}
	}
}

func TestCharacterStation(t *testing.T) {
	if len(CharacterStation) == 0 {
		t.Fatal("expected character station mappings, got none")
	}

	// Verify all stations used are valid
	validStations := make(map[Station]struct{})
	for _, s := range Stations() {
		validStations[s] = struct{}{}
	}

	for name, station := range CharacterStation {
		if _, ok := validStations[station]; !ok {
			t.Errorf("character %q has invalid station %q", name, station)
		}
	}

	// Verify some known characters have correct stations
	tests := []struct {
		name    string
		station Station
	}{
		{"Robert Baratheon", StationRoyal},
		{"Eddard \"Ned\" Stark", StationGreatLord},
		{"Jon Snow", StationLord},
		{"Jaime Lannister", StationKnight},
		{"Gilly", StationSmallfolk},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			station, ok := CharacterStation[tc.name]
			if !ok {
				t.Errorf("character %q not found in CharacterStation", tc.name)
				return
			}
			if station != tc.station {
				t.Errorf("character %q: expected station %q, got %q", tc.name, tc.station, station)
			}
		})
	}
}

func TestStationDistribution(t *testing.T) {
	// Verify we have a reasonable distribution of stations
	distribution := make(map[Station]int)
	for _, station := range CharacterStation {
		distribution[station]++
	}

	// Each station should have at least one character
	for _, station := range Stations() {
		if count := distribution[station]; count == 0 {
			t.Errorf("station %q has no characters", station)
		}
	}
}
