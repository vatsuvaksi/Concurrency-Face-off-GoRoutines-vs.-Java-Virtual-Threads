package main

import (
	dbconfig "goRoutineWorker/dbConfig"
	mqconfig "goRoutineWorker/mqConfig"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Initialize RabbitMQ
	rabbitMQ := initializeRabbitMQ()
	defer rabbitMQ.Close()

	// Initialize Database
	dbConfig := initializeDatabase()
	defer dbConfig.Close()

	// Start Consumer
	startConsumer(rabbitMQ, dbConfig)
}

func initializeRabbitMQ() *mqconfig.RabbitMQConfig {
	rabbitMQ, err := mqconfig.NewRabbitMQConfig()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ config: %v", err)
	}
	return rabbitMQ
}

func initializeDatabase() *dbconfig.DatabaseConfig {
	return dbconfig.InitDB("user", "password", "psql", "mydatabase", 5432)
}

func startConsumer(rabbitMQ *mqconfig.RabbitMQConfig, dbConfig *dbconfig.DatabaseConfig) {
	handler := func(messages []amqp.Delivery) {
		if err := dbConfig.SaveMessages(messages); err != nil {
			log.Printf("Error saving messages to DB: %v", err)
		}
	}

	qos := mqconfig.GetOptimalWorkerCount()
	consumer := mqconfig.NewConsumer(rabbitMQ, handler)
	consumer.Init(qos)
}
