package queue

import (
	"context"
	"os"

	"github.com/streadway/amqp"
	"github.com/yakubido/hemotrace_ocr/config"
	"github.com/yakubido/hemotrace_ocr/repo"
	"go.uber.org/zap"
)

var resultQueue = "ocr_results"

type Service struct {
	logger      *zap.Logger
	tasks       repo.Tasks
	amqpChannel *amqp.Channel
}

func New(logger *zap.Logger, cfg *config.Config, tasks repo.Tasks) (*Service, error) {
	svc := Service{
		logger: logger,
		tasks:  tasks,
	}
	// topic := ""

	conn, err := amqp.Dial(cfg.QueueConsumer.DSN) //"amqp://user:user@20.241.196.70:5672/my_vhost"
	if err != nil {
		logger.Fatal("Can't connect to AMQP", zap.Error(err))
	}
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	if err != nil {
		svc.logger.Fatal("Can't create a amqpChannel", zap.Error(err))
	}
	defer amqpChannel.Close()

	svc.amqpChannel = amqpChannel

	return &svc, nil
}

func (svc *Service) Run(ctx context.Context) error {
	queue, err := svc.amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	if err != nil {
		svc.logger.Fatal("Could not declare `add` queue", zap.Error(err))
	}

	err = svc.amqpChannel.Qos(1, 0, false)
	if err != nil {
		svc.logger.Fatal("Could not configure QoS", zap.Error(err))
	}

	autoAck, exclusive, noLocal, noWait := false, false, false, false
	messageChannel, err := svc.amqpChannel.Consume(
		queue.Name,
		"",
		autoAck,
		exclusive,
		noLocal,
		noWait,
		nil,
	)
	if err != nil {
		svc.logger.Fatal("Could not register consumer", zap.Error(err))
	}

	stopChan := make(chan bool)
	go func() {
		svc.logger.Info("Consumer ready", zap.Int("PID", os.Getpid()))
		for d := range messageChannel {
			svc.logger.Info("Received a message:", zap.String("Body", string(d.Body)))

			if err := d.Ack(false); err != nil {
				svc.logger.Error("Error acknowledging message", zap.Error(err))
			} else {
				svc.logger.Info("Acknowledged message")
			}
		}
	}()
	<-stopChan

	return nil
}
