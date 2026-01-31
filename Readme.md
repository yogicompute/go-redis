# MiniRedis

A lightweight Redis clone written in Go.

## Features

- **SET key value** – Store a key-value pair
- **SET key value EX seconds** – Store a key-value pair with expiration (TTL)
- **GET key** – Retrieve the value of a key
- **DEL key** – Delete a key
- **EXISTS key** – Check if a key exists
- **Automatic TTL cleanup** – Expired keys are automatically removed
- **Persistence** – Data is saved to disk (`dump.json`) and restored on startup
- **Periodic snapshots** – Data is auto-saved every 5 seconds

## Usage

```sh
go run main.go
```

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

## License

MIT