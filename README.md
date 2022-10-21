# Seesaw Engine

The Seesaw Engine is the binary application that supports the Seesaw Agent and Server.

## Building

```bash
go build -o engine ./cmd/engine
```

## Running

```bash
./engine -config=./config/server-mode.yml

./engine -config=./config/agent-mode.yml

./engine -config=./config/hybrid-mode.yml
```