# MQTT Message Broadcaster

## Description

This application subscribes to a specified MQTT prefix and bradcasts any received
messages to a list of other prefixes.

### Example

Imagine you have multiple MQTT clients that can only subscribe to one topic-prefix
each (like [awtrix3](https://github.com/Blueforcer/awtrix3/issues/540)).

* Client_A subscribed to `client_a/#`
* Client_B subscribed to `client_b/#`
* Client_C ...
* ...

In order to broadcast a message to all clients you now need to send the message
to every client while remembering all your topic-prefixes.

With this tool you can simply send the message to your own topic-prefix and it will
then be broadcasted to all configured clients.

Message to `broadcast/custom/weather` will be translated into several messages:

* `client_a/custom/weather`
* `client_b/custom/weather`
* `client_c/custom/weather`
* `...`

## Usage

```bash
cp .env.example .env
```

Alter the settings in our `.env` file.

```bash
docker compose build
docker compose up
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss
what you would like to change.
