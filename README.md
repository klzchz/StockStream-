# StockStream

A high-performance trading system built in Go for order matching and asset management.

## Features

- **Order Management**: Buy/sell order creation and execution
- **Asset Trading**: Support for multiple financial instruments
- **Investor Profiles**: User management and portfolio tracking
- **Transaction Records**: Complete trade history and audit trail
- **Order Matching**: Efficient price-time priority matching algorithm

## Project Structure

```
internal/
└── market/
    └── entity/
        ├── asset.go          # Asset entity and operations
        ├── book.go           # Order book and matching logic
        ├── investor.go       # Investor management
        ├── order.go          # Order creation and handling
        ├── order_queue.go    # Priority queue for orders
        ├── transaction.go    # Transaction recording
        └── book_test.go      # Comprehensive test suite
```

## Getting Started

### Prerequisites
- Go 1.21 or higher

### Installation

```bash
git clone https://github.com/klzchz/StockStream-.git
cd StockStream-
go mod tidy
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./internal/market/entity

# Run specific test
go test -run TestBuyAsset ./internal/market/entity
```

## Usage Example

```go
// Create a new order book
book := entity.NewBook("AAPL", "Apple Inc.")

// Create buy and sell orders
buyOrder := entity.NewOrder("1", "investor1", "AAPL", 100, 150.0, "buy")
sellOrder := entity.NewOrder("2", "investor2", "AAPL", 50, 149.0, "sell")

// Add orders to book
book.AddBuyOrder(buyOrder)
book.AddSellOrder(sellOrder)
```

## Test Coverage

Current test coverage: **96.2%**

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request
