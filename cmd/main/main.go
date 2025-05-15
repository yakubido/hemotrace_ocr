package main

import (
	"log"

	"github.com/yakubido/hemotrace_ocr/config"
	"github.com/yakubido/hemotrace_ocr/repo"
	"go.uber.org/zap"
)

// var version string

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logConfig := zap.NewProductionConfig()
	if cfg.Debug {
		logConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, _ := logConfig.Build()
	defer logger.Sync()

	tasks := repo.NewTask(logger.Named("tasks"), cfg)

	_ = tasks

	// 	mqURL := os.Getenv("MQ_URL")
	// 	if mqURL == "" {
	// 		mqURL = "amqp://guest:guest@localhost:5672/"
	// 	}

	// 	err := queue.Consume(mqURL, handleTask)
	// 	if err != nil {
	// 		log.Fatalf("[FATAL] Failed to start consumer: %v", err)
	// 	}
	// }

	// func handleTask(body []byte) {
	// 	var task OCRTask
	// 	if err := json.Unmarshal(body, &task); err != nil {
	// 		log.Printf("[ERROR] Invalid OCR task: %v", err)
	// 		return
	// 	}

	// 	absPath, _ := filepath.Abs(task.FilePath)
	// 	text, err := processor.ExtractText(absPath)
	// 	if err != nil {
	// 		log.Printf("[ERROR] OCR failed for %s: %v", task.FilePath, err)
	// 		return
	// 	}

	// 	markers := processor.ExtractMarkers(text)
	// 	result := OCRResult{
	// 		AnalysisID: task.AnalysisID,
	// 		Markers:    markers,
	// 	}

	// 	data, err := json.Marshal(result)
	// 	if err != nil {
	// 		log.Printf("[ERROR] JSON marshal failed: %v", err)
	// 		return
	// 	}

	// queue.PublishResult(data)
}
