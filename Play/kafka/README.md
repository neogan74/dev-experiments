# Kafka Playground

A comprehensive Kafka lab/playground with Docker Compose setup and examples in Python, Java (Spring Boot), and Go.

## 📋 Prerequisites

- Docker & Docker Compose
- Python 3.8+ (for Python examples)
- Java 17+ & Gradle (for Java examples)
- Go 1.25+ (for Go examples)

## 🚀 Quick Start

### 1. Start Kafka Infrastructure

```bash
docker-compose up -d
```

This will start:
- **Kafka** on `localhost:9092`
- **Zookeeper** on `localhost:2181`
- **Schema Registry** on `localhost:8081`
- **Kafka UI** on `localhost:8080`

### 2. Verify Services

Access Kafka UI at: http://localhost:8080

Check running containers:
```bash
docker-compose ps
```

## 📚 Examples

### Python Examples

#### Setup
```bash
cd python-examples
pip install -r requirements.txt
```

#### Run Producer
```bash
python producer.py [topic-name] [num-messages]

# Examples:
python producer.py                    # Uses defaults: test-topic, 10 messages
python producer.py my-topic 20        # Send 20 messages to my-topic
```

#### Run Consumer
```bash
python consumer.py [topic-name]

# Examples:
python consumer.py                    # Consumes from test-topic
python consumer.py my-topic           # Consumes from my-topic
```

### Java (Spring Boot) Examples

#### Setup
```bash
cd java-examples
./gradlew build
```

#### Run Application
```bash
./gradlew bootRun
```

The application includes:
- `KafkaProducerService` - Service for producing messages
- `KafkaConsumerService` - Service for consuming messages (auto-listens on startup)
- `ProducerRunner` - CommandLineRunner to send messages on startup (commented out by default)

To enable automatic message production on startup, uncomment the code in `ProducerRunner.java:14-30`.

### Go Examples

#### Setup
```bash
cd go-examples
go mod download
```

#### Run Producer
```bash
go run producer.go [topic-name] [num-messages]

# Examples:
go run producer.go                    # Uses defaults: test-topic, 10 messages
go run producer.go my-topic 20        # Send 20 messages to my-topic
```

#### Run Consumer
```bash
go run consumer.go [topic-name]

# Examples:
go run consumer.go                    # Consumes from test-topic
go run consumer.go my-topic           # Consumes from my-topic
```

## 🧪 Testing the Setup

1. **Start Kafka**:
   ```bash
   docker-compose up -d
   ```

2. **Open two terminals**

3. **Terminal 1 - Start Consumer** (choose your language):
   ```bash
   # Python
   cd python-examples && python consumer.py

   # Java
   cd java-examples && ./gradlew bootRun

   # Go
   cd go-examples && go run consumer.go
   ```

4. **Terminal 2 - Send Messages** (choose your language):
   ```bash
   # Python
   cd python-examples && python producer.py

   # Go
   cd go-examples && go run producer.go
   ```

5. **View in Kafka UI**: Open http://localhost:8080 to see topics, messages, and consumer groups

## 🛠️ Useful Commands

### Docker Commands
```bash
# Start all services
docker-compose up -d

# Stop all services
docker-compose down

# View logs
docker-compose logs -f kafka

# Restart services
docker-compose restart
```

### Kafka CLI Commands (from within container)
```bash
# Create a topic
docker exec -it kafka kafka-topics --create --topic my-topic --bootstrap-server localhost:19092 --partitions 3 --replication-factor 1

# List topics
docker exec -it kafka kafka-topics --list --bootstrap-server localhost:19092

# Describe topic
docker exec -it kafka kafka-topics --describe --topic test-topic --bootstrap-server localhost:19092

# Console producer
docker exec -it kafka kafka-console-producer --topic test-topic --bootstrap-server localhost:19092

# Console consumer
docker exec -it kafka kafka-console-consumer --topic test-topic --from-beginning --bootstrap-server localhost:19092

# Delete topic
docker exec -it kafka kafka-topics --delete --topic test-topic --bootstrap-server localhost:19092
```

## 📁 Project Structure

```
kafka/
├── docker-compose.yml          # Kafka infrastructure setup
├── README.md                   # This file
├── python-examples/
│   ├── requirements.txt
│   ├── producer.py
│   └── consumer.py
├── java-examples/
│   ├── build.gradle
│   ├── settings.gradle
│   └── src/
│       └── main/
│           ├── java/com/example/kafka/
│           │   ├── KafkaExamplesApplication.java
│           │   ├── model/Message.java
│           │   ├── producer/KafkaProducerService.java
│           │   ├── consumer/KafkaConsumerService.java
│           │   └── runner/ProducerRunner.java
│           └── resources/
│               └── application.yml
└── go-examples/
    ├── go.mod
    ├── producer.go
    └── consumer.go
```

## 🎯 Use Cases

This playground is perfect for:
- Learning Kafka basics
- Testing Kafka producers and consumers
- Experimenting with different languages
- Understanding consumer groups
- Testing schema registry
- Prototyping Kafka-based applications

## 🔧 Configuration

### Kafka Configuration
- Bootstrap servers: `localhost:9092`
- Internal broker communication: `kafka:19092`

### Consumer Groups
- Python: `python-consumer-group`
- Java: `java-consumer-group`
- Go: `go-consumer-group`

### Topics
- Default topic: `test-topic`
- Auto-create enabled (topics are created automatically when produced to)

## 📊 Monitoring

Access Kafka UI at http://localhost:8080 to monitor:
- Topics and their configurations
- Messages in topics
- Consumer groups and lag
- Broker information

## 🧹 Cleanup

```bash
# Stop and remove all containers
docker-compose down

# Stop and remove all containers + volumes (clears all data)
docker-compose down -v
```

## 📝 Notes

- All examples use JSON serialization for messages
- Consumer groups are configured to read from the earliest offset
- The Java application uses Spring Boot with auto-configuration
- Go examples use the Sarama library (IBM/sarama)
- Python examples use confluent-kafka library

Happy Kafka learning! 🎉
