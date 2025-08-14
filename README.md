# GBroker - Trading System

Trading system that processes buy and sell orders through Kafka, automatically executing transactions when orders match.

## Prerequisites

- Docker and Docker Compose
- Go 1.19+
- WSL2 (if using Windows)
- GCC (for Kafka C libraries)

## Environment Setup

### 1. WSL2 Configuration

If using WSL2, add to `/etc/hosts`:

```bash
sudo nano /etc/hosts
```

Add these lines:
```
127.0.0.1 host.docker.internal
127.0.0.1 kafka
```

### 2. Go + CGO Configuration

Enable C library in Go environment variables:

```bash
go env -w CGO_ENABLED=1
```

If you get the error:
```
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in $PATH
```

Install build tools:

```bash
sudo apt update
sudo apt install build-essential
```

## Installation

1. Clone the repository:
```bash
git clone <repo-url>
cd gbroker
```

2. Start Kafka services:
```bash
docker-compose up -d
```

3. Wait for containers to start (about 30 seconds):
```bash
docker ps
```

4. Install Go dependencies:
```bash
go mod tidy
```

## Running the System

1. Start the trading service:
```bash
go run cmd/trade/main.go
```

2. Access Kafka Control Center:
```
http://localhost:9021
```

## Testing Transactions

### Method 1: Via Control Center (Web Interface)

1. Access `http://localhost:9021`
2. Go to **Topics** → **input**
3. Click **Produce a new message to this topic**
4. Paste a JSON order in the **Value** field

### Method 2: Via Terminal

```bash
docker exec -it gbroker-kafka-1 kafka-console-producer --bootstrap-server localhost:9092 --topic input
```

## Order Examples

### Sell Order:
```json
{"order_id": "1", "investor_id": "Mari", "asset_id": "asset1", "current_shares": 10, "shares": 5, "price": 5.0, "order_type": "SELL"}
```

### Buy Order:
```json
{"order_id": "2", "investor_id": "João", "asset_id": "asset1", "current_shares": 0, "shares": 5, "price": 5.0, "order_type": "BUY"}
```

### Complete Transaction Example

1. Send a sell order:
```json
{"order_id": "1", "investor_id": "Mari", "asset_id": "asset1", "current_shares": 10, "shares": 5, "price": 5.0, "order_type": "SELL"}
```

2. Send a buy order with same price:
```json
{"order_id": "2", "investor_id": "Celia", "asset_id": "asset1", "current_shares": 8, "shares": 5, "price": 5.0, "order_type": "BUY"}
```

3. Watch the automatic transaction execution in the Go terminal.

## Order Fields

- `order_id`: Unique order ID (string)
- `investor_id`: Investor ID (string)
- `asset_id`: Asset ID (string)
- `current_shares`: Current investor's shares (int)
- `shares`: Number of shares to trade (int)
- `price`: Price per share (float)
- `order_type`: Order type ("BUY" or "SELL")

## Monitoring

### View processed messages:
- **Input topic**: `http://localhost:9021` → Topics → input
- **Output topic**: `http://localhost:9021` → Topics → output

### System logs:
```bash
# View Go program logs
go run cmd/trade/main.go

# View Kafka logs
docker logs gbroker-kafka-1

# View Control Center logs
docker logs gbroker-control-center-1
```

## Troubleshooting

### Error "Failed to resolve 'kafka:9092'"
Check if you added the entries to `/etc/hosts` as instructed above.

### CGO/GCC Error
```bash
go env -w CGO_ENABLED=1
sudo apt update
sudo apt install build-essential
```

### Containers won't start
```bash
docker-compose down
docker-compose up -d
```

### Consumer not receiving messages
1. Stop the Go program (Ctrl+C)
2. Restart: `go run cmd/trade/main.go`
3. Send a new message

### Clear Kafka data
```bash
docker-compose down -v
docker-compose up -d
```

## Architecture

```
[Kafka Producer] → [Topic: input] → [Go Consumer] → [Trading Engine] → [Topic: output]
```

The system:
1. Receives orders via `input` topic
2. Processes and matches orders
3. Automatically executes transactions
4. Publishes results to `output` topic

## Sample Output

```
Received message: {"order_id":"1","investor_id":"Mari",...}
Received message: {"order_id":"2","investor_id":"Celia",...}
Processed order: {
  "order_id": "1",
  "investor_id": "Mari",