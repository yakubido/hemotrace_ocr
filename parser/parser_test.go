package processor

import (
	"reflect"
	"testing"
)

func TestExtractMarkers(t *testing.T) {
	sample := `Гемоглобин 145 г/л
ALT: 35 U/L
Креатинин 95 мкмоль/л
Глюкоза 5.6 ммоль/л`

	expected := []MarkerResult{
		{"Гемоглобин", 145, "г/л"},
		{"ALT", 35, "U/L"},
		{"Креатинин", 95, "мкмоль/л"},
		{"Глюкоза", 5.6, "ммоль/л"},
	}

	results := ExtractMarkers(sample)
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %+v, got %+v", expected, results)
	}
}

func TestExtractMarkersWithInvalidLines(t *testing.T) {
	sample := `Some irrelevant text
Гемоглобин: 140 г/л
ALT: thirty five
Кальций: 2.4 ммоль/л`

	expected := []MarkerResult{
		{"Гемоглобин", 140, "г/л"},
		{"Кальций", 2.4, "ммоль/л"},
	}

	results := ExtractMarkers(sample)
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %+v, got %+v", expected, results)
	}
}
