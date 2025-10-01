#!/usr/bin/env python3
"""
Simple Kafka Consumer Example
Consumes messages from a Kafka topic
"""

from confluent_kafka import Consumer, KafkaError
import json
import sys

# Kafka configuration
config = {
    'bootstrap.servers': 'localhost:9092',
    'group.id': 'python-consumer-group',
    'auto.offset.reset': 'earliest',
    'enable.auto.commit': True
}

# Create Consumer instance
consumer = Consumer(config)


def consume_messages(topic='test-topic'):
    """Consume messages from Kafka topic"""
    print(f"📥 Starting to consume messages from topic '{topic}'...")
    print("Press Ctrl+C to stop\n")

    # Subscribe to topic
    consumer.subscribe([topic])

    try:
        message_count = 0
        while True:
            # Poll for messages
            msg = consumer.poll(timeout=1.0)

            if msg is None:
                continue

            if msg.error():
                if msg.error().code() == KafkaError._PARTITION_EOF:
                    print(f'📌 Reached end of partition {msg.partition()}')
                else:
                    print(f'❌ Error: {msg.error()}')
                continue

            # Process message
            message_count += 1
            key = msg.key().decode('utf-8') if msg.key() else None
            value = json.loads(msg.value().decode('utf-8'))

            print(f"📨 Message #{message_count}")
            print(f"   Topic: {msg.topic()}")
            print(f"   Partition: {msg.partition()}")
            print(f"   Offset: {msg.offset()}")
            print(f"   Key: {key}")
            print(f"   Value: {value}")
            print()

    except KeyboardInterrupt:
        print("\n⚠️  Consumer interrupted by user")
    except Exception as e:
        print(f"❌ Error: {e}")
    finally:
        # Close consumer
        print("🔒 Closing consumer...")
        consumer.close()


if __name__ == '__main__':
    topic = sys.argv[1] if len(sys.argv) > 1 else 'test-topic'
    consume_messages(topic)
