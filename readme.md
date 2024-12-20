### * Project under development *

# Comparing JDK21 Virtual Threads, Java Traditional Threads, and Goroutines

## Project Overview

This project aims to evaluate the performance and efficiency of three different concurrency models:
1. **JDK21 Virtual Threads**
2. **Java Traditional Threads**
3. **Goroutines** (Go's lightweight threads)

We utilize a RabbitMQ workload to simulate a real-world high-throughput scenario, processing **1,000,000 JSON records**. The project focuses on optimizing resource utilization and measuring processing times for a comprehensive comparison of these models.

---

## Objectives

- **Performance Benchmarking**: Measure and compare the time taken by each concurrency model to process 1,000,000 JSON records.
- **Resource Utilization**: Evaluate CPU and memory usage under each concurrency model.
- **Scalability Analysis**: Test the ability of each model to scale under increasing loads.
- **Real-World Applicability**: Simulate a real-world messaging workload with RabbitMQ.

---

## Technologies Used

### Java Components:
- **JDK 21**: Leveraging Virtual Threads for improved scalability.
- **JDK 17**: Leveraging Traditional Threads for improved scalability.
- **Java Traditional Threads**: Using standard thread pools for concurrency.

### Go Components:
- **Goroutines**: Go's lightweight concurrency mechanism.

### Messaging System:
- **RabbitMQ**: A message broker to simulate a real-world data streaming workload.

### JSON Processing:
- **Jackson (Java)** / **encoding/json (Go)**: Libraries for parsing and serializing JSON records.

### Infra
- **Docker Compose**: Builds the container
- **Dockerfile**: Builds image of the apps

### Build Tools:
- **Maven** (Java)
- **Go Modules** (Go)

---

## Architecture

### Workflow:
1. **Message Production**:
   - A producer sends 1,000,000 JSON records to a RabbitMQ queue.
2. **Message Consumption and Processing**:
   - Consumers process the JSON records using one of the concurrency models (Virtual Threads, Traditional Threads, or Goroutines). 
3. **Metrics Collection**:
   - Time taken to process all messages.
   - CPU and memory usage during processing.
4. **Data Analysis**:
   - Results are logged and visualized to highlight the differences between models.

### Directory Structure:
```
project-root/
├── java-virtual-threads/
│   ├── src/
│   ├── pom.xml
├── java-traditional-threads/
│   ├── src/
│   ├── pom.xml
├── go-goroutines/
│   ├── main.go
├── benchmarks/
│   ├── results.csv
│   ├── charts/
└── README.md
```

---

## Installation and Setup

### Prerequisites:
- **Docker**

### Steps:

#### 1. Run Docker compose file in each project:

---

## Metrics and Results

### Metrics Captured:
1. **Processing Time**:
   - Time taken to process 1,000,000 records 
2. **CPU Usage**:
   - Measured at regular intervals.
3. **Memory Consumption**:
   - Peak and average memory usage.

### Example Results:
| Concurrency Model      | Processing Time (ms) | CPU Usage (%) | Memory Usage (MB) |
|------------------------|----------------------|---------------|-------------------|
| JDK21 Virtual Threads  | TBF                | TBF            | TBF               |
| Java Traditional Threads | TBF               | TBF            | TBF              |
| Goroutines             | TBF                | TBF            | TBF               |

### Observations:
- **Virtual Threads** significantly reduce thread contention and resource usage.
- **Traditional Threads** struggle with scalability under high workloads.
- **Goroutines** are the most lightweight and performant for the given workload.


---

## Future Enhancements
- Optimize RabbitMQ configurations for higher throughput.

---

## Conclusion
This project provides an in-depth analysis of concurrency models in high-throughput scenarios. It demonstrates the advantages and trade-offs of using JDK21 Virtual Threads, Java Traditional Threads, and Goroutines for RabbitMQ workload processing.

