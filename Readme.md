# MiniRedis

A lightweight Redis clone written in Go with both TCP and HTTP interfaces.

## Features

### Core Commands
- **SET key value** – Store a key-value pair
- **SET key value EX seconds** – Store a key-value pair with expiration (TTL)
- **GET key** – Retrieve the value of a key
- **DEL key** – Delete a key
- **EXISTS key** – Check if a key exists

### Data Management
- **Automatic TTL cleanup** – Expired keys are automatically removed (runs every 1 second)
- **Persistence** – Data is saved to disk (`dump.json`) and restored on startup
- **Periodic snapshots** – Data is auto-saved every 5 seconds

### Interfaces
- **TCP Server** – Native Redis-like protocol interface
- **HTTP API** – RESTful endpoints for web integration

## Usage

### Starting the Server

```sh
go run main.go
```

### TCP Interface

Connect via TCP and use Redis-like commands:

```
> SET name Alice
OK
> GET name
Alice
> SET session abc123 EX 60
OK
> EXISTS session
1
> DEL name
1
> GET name
(nil)
```

### HTTP API

The HTTP API runs on port `8080` with the following endpoints:

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/set` | GET/POST | Set a key-value pair |
| `/get` | GET | Retrieve a value by key |
| `/del` | GET/POST | Delete a key |
| `/exists` | GET | Check if a key exists |
| `/stats` | GET | Get store statistics |

#### Examples

```sh
# Set a value
curl "http://localhost:8080/set?key=name&value=Alice"

# Get a value
curl "http://localhost:8080/get?key=name"

# Delete a key
curl "http://localhost:8080/del?key=name"

# Check if key exists
curl "http://localhost:8080/exists?key=name"

# Get stats
curl "http://localhost:8080/stats"
```

## Project Structure

```
redis-clone/
├── main.go          # Application entry point
├── store/           # Key-value store implementation
├── httpapi/         # HTTP API handlers
└── dump.json        # Persistence file (auto-generated)
```

## License

MIT