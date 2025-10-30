#!/usr/bin/env python3
"""
Simple Kafka Producer Example
Sends messages to a Kafka topic
"""

from confluent_kafka import Producer
import json
import time
import sys

# Kafka configuration
config = {
    'bootstrap.servers': 'localhost:9092',
    'client.id': 'python-producer'
}

# Create Producer instance
producer = Producer(config)


def delivery_callback(err, msg):
    """Callback called once message is delivered or failed"""
    if err:
        print(f'❌ Message delivery failed: {err}')
    else:
        print(f'✅ Message delivered to {msg.topic()} [{msg.partition()}] at offset {msg.offset()}')


def produce_messages(topic='test-topic', num_messages=10):
    """Produce messages to Kafka topic"""
    print(f"📤 Starting to produce {num_messages} messages to topic '{topic}'...\n")

    try:
        for i in range(num_messages):
            # Create message
            message = {
                'id': i,
                'message': f'Hello Kafka message #{i}',
                'timestamp': time.time()
            }

            # Produce message
            producer.produce(
                topic=topic,
                key=str(i),
                value=json.dumps(message),
                callback=delivery_callback
            )

            # Trigger delivery reports
            producer.poll(0)

            time.sleep(0.5)

        # Wait for all messages to be delivered
        print("\n⏳ Flushing remaining messages...")
        producer.flush()
        print("✨ All messages sent successfully!")

    except KeyboardInterrupt:
        print("\n⚠️  Interrupted by user")
    except Exception as e:
        print(f"❌ Error: {e}")
    finally:
        producer.flush()


if __name__ == '__main__':
    topic = sys.argv[1] if len(sys.argv) > 1 else 'test-topic'
    num_messages = int(sys.argv[2]) if len(sys.argv) > 2 else 10

    produce_messages(topic, num_messages)
