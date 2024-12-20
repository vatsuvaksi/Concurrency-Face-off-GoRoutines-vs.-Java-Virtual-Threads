import pika
import json
import time

# RabbitMQ connection parameters
rabbitmq_host = "rabbitmq"
queue_name = "test_queue"

# Retry mechanism
for attempt in range(1000):  # Retry up to 10 times
    try:
        connection = pika.BlockingConnection(pika.ConnectionParameters(host=rabbitmq_host))
        channel = connection.channel()
        break
    except pika.exceptions.AMQPConnectionError:
        print(f"Connection attempt {attempt + 1} failed. Retrying in 5 seconds...")
        time.sleep(5)
else:
    raise Exception("Failed to connect to RabbitMQ after 10 attempts.")

# Declare the queue
channel.queue_declare(queue=queue_name)

# Generate and publish 1,000,000 JSON messages
for i in range(1, 1000001):
    message = json.dumps({"id": i, "name": f"Name{i}", "email": f"user{i}@example.com", "age": i % 100})
    channel.basic_publish(exchange='', routing_key=queue_name, body=message)

print("Loaded 1,000,000 messages into the queue.")
connection.close()
