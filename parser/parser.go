package processor

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// MarkerResult представляет результат одного маркера после парсинга текста
type MarkerResult struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// ExtractMarkers разбирает текст OCR и извлекает маркеры с их значениями и единицами измерения
func ExtractMarkers(text string) []MarkerResult {
	lines := strings.Split(text, "\n")
	var results []MarkerResult

	// Пример шаблона: "Гемоглобин  145 г/л"
	re := regexp.MustCompile(`(?i)([a-zа-яё\-\s]+)[\s:]+([0-9]+(?:\.[0-9]*)?)\s*([a-zа-яё/µ%]+)?`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		matches := re.FindStringSubmatch(line)
		if len(matches) >= 3 {
			name := strings.TrimSpace(matches[1])
			valueStr := matches[2]
			unit := ""
			if len(matches) >= 4 {
				unit = strings.TrimSpace(matches[3])
			}

			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				log.Printf("[WARN] Failed to parse value for '%s': %v", name, err)
				continue
			}

			results = append(results, MarkerResult{
				Name:  name,
				Value: value,
				Unit:  unit,
			})
		}
	}

	log.Printf("[INFO] Extracted %d marker(s) from OCR text", len(results))
	return results
}