
In Go projects, the internal/ folder is a special directory designed for **private application code**.
This code is meant to be used **only within the project** (and not be exposed to other Go modules or external dependencies). It provides a way to **encapsulate** certain parts of the codebase that you don't want to be directly accessed by other projects or packages.

It is a private package, while `pkg/` or `cmd/` are public packages.

### Why Does the `internal/` Folder Exist?

- **Encapsulation**: The primary goal of the `internal/` folder is to **restrict access** to certain parts of the codebase. Any package inside the `internal/` folder is **only accessible to other code within the same project**. It can't be imported by external projects or packages, which helps keep your application’s private code protected.
    
- **Prevent External Dependencies**: By using `internal/`, you ensure that only the public packages (those in the `pkg/` or `cmd/` folders) are visible outside of your module. Any implementation details (like services, models, or handlers) should stay private and internal to your project.
    
- **Modular Design**: It helps in organizing the codebase in a modular way, clearly separating **public** and **private** code.

### Key Characteristics of the `internal/` Folder

1. **Private Code**: Code within `internal/` cannot be imported by any package outside of the project (including other Go modules).
    
2. **Access Control**: The import restriction is enforced by the **Go toolchain** itself, meaning Go will throw an error if any external code attempts to import something from `internal/`.
    
3. **Used for App Logic**: This is where you put **core application logic**, data models, business services, and utility code that you want to keep internal to your project.

The `internal/` folder contains the **core logic** of your application, such as:

- Handlers
    
- Business logic services
    
- Models
    
- Repositories (for database access)
    
- Middlewares
    
- Utilities
    
- Configuration files

By clearly separating the public and private parts of the code, you make it easier to maintain and refactor the project. You can change internal code without worrying about breaking external consumers.


# Typical flow in a Go backend:

```cmd
Client Request
      ↓
Handler (internal/handlers)
      ↓
Service (internal/service)
      ↓
Repository (internal/repository)
      ↓
Database
```

# Sub Layers inside Internal

- [[config]]
- [[handlers]]
- [[Domain]]
- [[DTO]]
- [[repository]]
- [[service]]