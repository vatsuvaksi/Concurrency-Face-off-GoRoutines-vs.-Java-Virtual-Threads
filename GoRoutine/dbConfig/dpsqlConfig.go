package dbconfig

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/streadway/amqp"
)

// DatabaseConfig holds the DB connection.
type DatabaseConfig struct {
	db *sql.DB
}

// InitDB initializes the database connection.
func InitDB(user, password, host, dbname string, port int) *DatabaseConfig {
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", user, password, host, dbname, port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return &DatabaseConfig{db: db}
}

// Close closes the database connection.
func (dc *DatabaseConfig) Close() {
	if dc.db != nil {
		_ = dc.db.Close()
	}
}

// SaveMessages saves multiple messages to the database concurrently.
func (dc *DatabaseConfig) SaveMessages(messages []amqp.Delivery) error {
	for _, msg := range messages {
		var data map[string]interface{}
		if err := json.Unmarshal(msg.Body, &data); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			msg.Nack(false, false)
			continue
		}

		id := int(data["id"].(float64))
		name := data["name"].(string)
		email := data["email"].(string)
		age := int(data["age"].(float64))

		_, err := dc.db.Exec("INSERT INTO users (id, name, email, age) VALUES ($1, $2, $3, $4)", id, name, email, age)
		if err != nil {
			log.Printf("Failed to save to database: %v", err)
			msg.Nack(false, true)
			continue
		}

		log.Printf("Message processed and saved: %v", data)
		msg.Ack(false)
	}
	return nil
}
