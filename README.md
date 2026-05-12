# cache-management

A lightweight, interface-driven Go library for managing cache clients. Currently provides a Redis backend with a clean abstraction layer, making it easy to swap or extend cache implementations without changing consumer code.

## Features

- Simple `Client` interface for `Set`, `Get`, and `Close` operations
- Redis backend powered by [`go-redis/v9`](https://github.com/redis/go-redis)
- `Bootstrap` helper for quick setup with environment-based configuration
- `Eval` support for running atomic Lua scripts (e.g. token bucket rate limiting)
- Designed to be imported as a Go module by other services

## Project Structure

```
cache-management/
├── Bootstrap/      # Bootstrap() helper — initializes and returns a configured client
├── Internal/       # Redis client implementation (RedisClient struct)
├── Types/          # Client interface definition
├── go.mod
└── main.go
```

## Installation

```bash
go get github.com/dev-AdiR/cache-management
```

## Usage

### Bootstrap a Redis client

```go
import cacheBootstrap "github.com/dev-AdiR/cache-management/Bootstrap"

app := cacheBootstrap.Bootstrap(0) // 0 = Redis DB index
defer app.Client.Close()
```

`Bootstrap` reads the Redis URL from your environment (via `godotenv`) and returns an `App` struct with a `Client` field that satisfies the `types.Client` interface.

### The Client interface

```go
type Client interface {
    Set(ctx context.Context, key string, value []byte, expiration time.Duration) error
    Get(ctx context.Context, key string) ([]byte, error)
    Close() error
}
```

### Using Eval (Lua scripts)

`Eval` is available on the underlying `*RedisClient` struct. If your consumer (e.g. a rate limiter) needs it, define a broader interface locally:

```go
// In your consumer package
type ScriptableClient interface {
    types.Client
    Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
}
```

Since `*RedisClient` already implements all these methods, it satisfies the interface automatically with no changes to this library.

## Dependencies

| Package | Purpose |
|---|---|
| `github.com/redis/go-redis/v9` | Redis client |
| `github.com/joho/godotenv` | Load Redis URL from `.env` |

## Environment Variables

Create a `.env` file in your project root:

```env
REDIS_URL=localhost:6379
```

## Example: Rate Limiter Integration

This library was built to back a token bucket rate limiter. The rate limiter calls `Eval` with a Lua script to atomically check and consume tokens on each request:

```go
allowed, err := client.Eval(
    ctx,
    luaScript,
    []string{key},
    10,                // capacity
    5,                 // refill rate (tokens/sec)
    time.Now().Unix(), // current timestamp
).(*redis.Cmd).Int()
```

The Lua script runs entirely inside Redis, ensuring the read-modify-write is race-condition free.

## License

MIT