package main

import (
	"context"
	"errors"
	"log"

	"github.com/yakubido/hemotrace_ocr/config"
	"github.com/yakubido/hemotrace_ocr/lib/sigtrap"
	"github.com/yakubido/hemotrace_ocr/queue"
	"github.com/yakubido/hemotrace_ocr/repo"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
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
	consumer, err := queue.New(logger.Named("Consumer"), cfg, tasks)
	if err != nil {
		logger.Fatal("queu consumer init", zap.Error(err))
	}

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error { return consumer.Run(ctx) })

	if err := group.Wait(); err != nil && !errors.Is(err, sigtrap.ErrSignalReceived) {
		logger.Error("service error", zap.Error(err))
	}

	logger.Info("shutting down")

	// 	mqURL := os.Getenv("MQ_URL")
	// 	if mqURL == "" {
	// 		mqURL = "amqp://guest:guest@localhost:5672/"
	// 	}

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
