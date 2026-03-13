- **Purpose**: This folder holds code related to **configuration and environment loading**.
    
- **What it does**: It typically reads configuration from files (like `.env`) or environment variables to set up things like DB connection strings, API keys, or other configuration parameters.
    
- **Why use it**: It centralizes configuration management, ensuring that all parts of the application use the same settings. Also, it helps in managing different environments (development, production, testing).
    
- **Example file (`config.go`)**: This file might have logic to parse configuration files and load them into application variables.

# About config directory

The **`internal/config` directory** is specifically used to **manage application configuration** in a centralized and controlled way.

`internal/config` is a package that **loads, stores, and provides configuration values** for the application.

Typical configs include:

- Environment variables
    
- Database credentials
    
- API keys
    
- Ports
    
- Feature flags
    
- Service URLs

Example structure
```cmd
internal/  
└── config/  
	├── config.go  
	├── env.go  
	└── config.yaml
```

Due to it's location inside internal package, external packages can't access it.


# Why a Dedicated Config Package Is Needed
Without a config package, configuration logic becomes **scattered across the codebase**.

Example of bad practice
```cmd
db.Connect(os.Getenv("DB_URL"))
server.Start(os.Getenv("PORT"))
```


Problems:

- Hard to maintain
    
- Repeated env calls
    
- Difficult to test
    
- No validation

A config package solves this by:

- Centralizing configuration
    
- Adding validation
    
- Supporting multiple sources (env, yaml, json)


## Responsibilities of `internal/config`

A config package usually handles:

### 1. Loading configuration

From:

- `.env`
    
- environment variables
    
- YAML/JSON files
    
- flags
    

Example libraries:

- `viper`
    
- `envconfig`
    
- `godotenv`

---
### 2. Defining config structure

Example:

```Go
type Config struct {
    Server ServerConfig
    DB     DBConfig
}

type ServerConfig struct {
    Port string
}

type DBConfig struct {
    Host     string
    User     string
    Password string
    Name     string
}
```

### 3. Validation

Ensuring required fields exist.

Example:

```cmd
DB_HOST must not be empty  
PORT must be valid
```


### 4. Providing config globally

The package exposes a **single config instance**.

Example:

```go
func Load() (*Config, error)
```

Used in `main.go`.

In case u want to make it singleton

```go
Typical Go singleton example:

var instance *Config
var once sync.Once

func Load() *Config {
    once.Do(func() {
        instance = &Config{
            // load config
        }
    })
    return instance
}
```

Here:

- `instance` is stored globally
    
- `sync.Once` ensures it is created **only once**
    
- Every call returns the **same instance**

# 5. Typical Implementation

Example `config.go`

```go
package config

import "os"

type Config struct {
    Port string
    DBUrl string
}

func Load() *Config {
    return &Config{
        Port:  getEnv("PORT", "8080"),
        DBUrl: getEnv("DB_URL", ""),
    }
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
```

# 6. How It Is Used in the Application

Example `main.go`

```go
package main  
  
import (  
    "myapp/internal/config"  
)  
  
func main() {  
    cfg := config.Load()  
  
    server := NewServer(cfg.Port)  
    server.Start()  
}
```

All services receive configuration via **dependency injection**.


- In advanced usage like in big projects it can support multiple environments
```cmd
config.dev.yaml
config.staging.yaml
config.prod.yaml
```

- Some systems reloads config without restart