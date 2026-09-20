bin := "bin/pinguin"

# Show recipes
default:
    @just --list --unsorted

# Build into ./bin
build:
    go build -o {{bin}} ./cmd/pinguin

# Build and run as root (ICMP needs raw sockets)
run *args: build
    sudo ./{{bin}} {{args}}

# Run tests
test:
    go test -race ./...

# Format and vet
lint:
    gofmt -s -w .
    go vet ./...

# Remove build output
clean:
    rm -rf bin
