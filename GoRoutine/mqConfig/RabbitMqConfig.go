package mqconfig

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/shirou/gopsutil/mem"
	"github.com/streadway/amqp"
)

const rabbitMQURL = "amqp://guest:guest@rabbitmq:5672/"
const queueName = "test_queue"

type RabbitMQConfig struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
	mu      sync.Mutex // For thread-safe QoS updates
}

// NewRabbitMQConfig initializes and returns a RabbitMQConfig instance.
func NewRabbitMQConfig() (*RabbitMQConfig, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQConfig{
		conn:    conn,
		channel: channel,
		queue:   queueName,
	}, nil
}

// Close gracefully closes the RabbitMQ connection and channel.
func (r *RabbitMQConfig) Close() {
	if r.channel != nil {
		_ = r.channel.Close()
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
}

// GetOptimalWorkerCount calculates the optimal number of goroutines based on system resources.
func GetOptimalWorkerCount() int {
	cores := runtime.NumCPU()

	// Fetch available memory using gopsutil.
	availableMemoryMB := getAvailableMemoryMB()
	averageMessageSizeMB := 0.001 // Assuming 1 KB per message.

	// Calculate maximum workers based on memory and cores.
	maxWorkersByMemory := availableMemoryMB / averageMessageSizeMB
	maxWorkersByCores := cores * 2 // A heuristic: 2 goroutines per core.

	// The final count is the minimum of the two constraints.
	return min(int(maxWorkersByMemory), maxWorkersByCores)
}

// getAvailableMemoryMB fetches available memory using gopsutil.
func getAvailableMemoryMB() float64 {
	vmStats, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("Failed to fetch memory stats: %v", err)
	}
	return float64(vmStats.Available) / (1024 * 1024) // Convert bytes to MB
}

// min returns the smaller of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// AdjustQos dynamically adjusts RabbitMQ QoS based on calculated workers.
func (r *RabbitMQConfig) AdjustQos(workerCount int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.channel.Qos(workerCount, 0, false)
	if err != nil {
		log.Printf("Failed to adjust QoS: %v", err)
		return err
	}
	log.Printf("Adjusted QoS to worker count: %d", workerCount)
	return nil
}

// StartDynamicQosAdjustment periodically recalculates and adjusts QoS.
func (r *RabbitMQConfig) StartDynamicQosAdjustment(interval time.Duration) {
	go func() {
		for {
			workerCount := GetOptimalWorkerCount()
			if err := r.AdjustQos(workerCount); err != nil {
				log.Printf("QoS adjustment failed: %v", err)
			}
			time.Sleep(interval)
		}
	}()
}
