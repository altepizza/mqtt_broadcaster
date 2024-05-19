import paho.mqtt.client as mqtt
from config import settings, logger

logger.info("Starting MQTT client")

broker_host = settings.MQTT_BROKER
broker_port = settings.MQTT_PORT

source_prefix = settings.MQTT_DISTRIBUTION_PREFIX
target_prefixes = settings.MQTT_TARGET_PREFIXES


# Define the on_connect callback function
def on_connect(client, userdata, flags, reason_code, properties):
    logger.info(f"Connected with result code {reason_code}")
    client.subscribe(source_prefix + "/#")


# Define the on_message callback function
def on_message(client, userdata, msg):
    logger.debug(f"Received message: {msg.payload.decode()} on topic {msg.topic}")
    for prefix in target_prefixes:
        new_topic = msg.topic.replace(source_prefix, prefix)
        client.publish(new_topic, msg.payload)


client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)
client.username_pw_set(username=settings.MQTT_USER, password=settings.MQTT_PASSWORD)

# Set the on_connect and on_message functions as the respective callbacks for the client
client.on_connect = on_connect
client.on_message = on_message

# Connect the client to the broker
client.connect(broker_host, broker_port, 60)

# Start the client loop
client.loop_start()

while True:
    pass
