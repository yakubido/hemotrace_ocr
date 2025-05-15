package processor

import (
	"fmt"
	"log"

	"github.com/otiai10/gosseract/v2"
)

// OCRTask представляет входящее сообщение из очереди
type OCRTask struct {
	AnalysisID int    `json:"analysis_id"`
	FilePath   string `json:"file_path"`
}

// OCRResult представляет структуру ответа после обработки
type OCRResult struct {
	AnalysisID int            `json:"analysis_id"`
	Markers    []MarkerResult `json:"markers"`
}

type MarkerResult struct {
}

// ExtractText извлекает текст из изображения или PDF с помощью Tesseract OCR
func ExtractText(filePath string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetImage(filePath)
	if err != nil {
		log.Printf("[ERROR] Failed to set image: %v", err)
		return "", err
	}

	text, err := client.Text()
	if err != nil {
		log.Printf("[ERROR] OCR extraction failed: %v", err)
		return "", err
	}

	fmt.Printf("[INFO] OCR extracted text (%d chars)\n", len(text))
	return text, nil
}

func ExtractMarkers(data string) []MarkerResult {
	return nil
}
