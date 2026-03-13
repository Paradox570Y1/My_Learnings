Contains **entry points** of the application.

Example:
`cmd/server/main.go` - this is where server starts


Allows multiple binaries
	- binaries  mean separate executable programs built from the same Go project
	- In Go, every folder inside `cmd/` usually contains a `main.go` file, and each `main.go` becomes its own executable binary when compiled.
	- 
```txt
cmd/server
cmd/worker
cmd/cron
```