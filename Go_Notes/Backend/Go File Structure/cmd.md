Contains **entry points** of the application.

Example:
`cmd/server/main.go` - this is where server starts


Allows multiple binaries
	- binaries  mean separate executable programs built from the same Go project
	- In a Go project, the `cmd/` folder typically contains different sub directories for various executable programs (like `cmd/server`, `cmd/cli`,  `cmd/cron`etc.), each with its own `main.go` file.

### Why Use This Structure?

The `cmd/` folder structure is often used in Go projects to separate different components or tools within a larger system. For example:

- `cmd/server` might contain the main application logic for a backend service.
    
- `cmd/cli` might contain a command-line interface for interacting with the service.

## About `main.go` file

Each `main.go` file contains a `main` function, which is the entry point of a Go program. When Go compiles this file, it creates an executable binary for that specific program.

For example:

- `cmd/server/main.go` will be compiled into a binary called `server` (on Linux/macOS) or `server.exe` (on Windows).
    
- `cmd/cli/main.go` will be compiled into a binary called `cli` or `cli.exe`.