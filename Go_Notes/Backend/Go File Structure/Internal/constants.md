The `constants` directory is about **centralizing fixed values that should never change during runtime**. It helps prevent **magic strings, duplicated values, and inconsistencies** across the codebase.

Let’s build the concept from scratch and then move toward how it’s used in a **real Go backend (like the one you’re building with MySQL)**.

# What the `constants` Directory Is

The `constants` directory contains **global constant values used throughout the application**.

These values:

- **do not change during runtime**
    
- are used **in multiple places**
    
- should have **one single source of truth**
    

Example structure:

```txt
internal/  
├── constants/  
│   ├── roles.go  
│   ├── errors.go  
│   ├── status.go  
│   └── headers.go
```

Each file defines constants related to a **specific concept


# Why We Need a Constants Directory

Without constants, developers often write **hardcoded values everywhere**.

Example (bad practice):

```go
if user.Role == "admin" {  
    // allow access  
}
```

Another developer might write:

```go
if user.Role == "ADMIN" {  
}
```

Or:

```go
if user.Role == "administrator" {  
}
```

This creates **bugs and inconsistencies**.

Using constants solves this.

Example:

```go
if user.Role == constants.RoleAdmin {  
}
```

Now the value exists **in one place only**.

The constants package is **shared by all layers**.

```txt
routes     → HTTP constants
handlers   → error constants
services   → role constants
repository → status constants
```


# What Should Go Inside Constants

Constants are usually things like:

### Application roles

```txt
admin  
user  
moderator
```

### HTTP statuses

```txt
active  
inactive  
pending  
deleted
```

### Error codes

```txt
invalid_credentials  
user_not_found  
unauthorized
```
### Headers

```txt
Authorization  
Content-Type
```

### Token types

```txt
access_token  
refresh_token
```

These values are **fixed definitions** used across the app.


# Basic Example: Role Constants

File:

```txt
internal/constants/roles.go
```

Example:

```go
package constants  
  
const (  
    RoleAdmin = "admin"  
    RoleUser  = "user"  
    RoleGuest = "guest"  
)
```

Usage in service:

```go
if user.Role == constants.RoleAdmin {  
    // allow admin operation  
}
```

# Example: Error Constants

Instead of writing error strings everywhere.

Bad:

```go
return errors.New("user not found")
```

Better:

```go
internal/constants/errors.go

package constants  
  
const (  
    ErrUserNotFound       = "user not found"  
    ErrInvalidCredentials = "invalid credentials"  
    ErrUnauthorized       = "unauthorized"  
)
```

Usage:

```go
return errors.New(constants.ErrUserNotFound)
```

This ensures **consistent error responses**.


# Example: API Response Status

File:

```txt
internal/constants/status.go
```

Example:

```go
package constants  
  
const (  
    StatusActive   = "active"  
    StatusInactive = "inactive"  
    StatusPending  = "pending"  
)
```

Used in repository or service:

```go
user.Status = constants.StatusActive
```

This avoids accidental typos.


# Example: HTTP Header Constants

File:

```txt
internal/constants/headers.go
```

Example:

```go
package constants  
  
const (  
    HeaderAuthorization = "Authorization"  
    HeaderContentType   = "Content-Type"  
)
```

Usage in middleware:

```go
token := c.GetHeader(constants.HeaderAuthorization)
```

# Example: Pagination Constants

Sometimes apps define default pagination values.

File:

```txt
internal/constants/pagination.go
```

Example:

```go
package constants  
  
const (  
    DefaultPage  = 1  
    DefaultLimit = 20  
    MaxLimit     = 100  
)
```

Usage:

```go
if limit == 0 {  
    limit = constants.DefaultLimit  
}
```


# 10. Constants vs Variables

Important distinction.

### Constants

Cannot change.

```go
const AppName = "MyBackend"
```

### Variables

Can change.

```go
var port = 8080
In the constants directory you should use:
```

const

almost always.


# Grouping Constants

Go allows grouped constants:

```txt
const (  
    RoleAdmin = "admin"  
    RoleUser  = "user"  
)
```

This is preferred over:

```txt
const RoleAdmin = "admin"  
const RoleUser = "user"
```

because it keeps related values organized.


# Typed Constants (Advanced)

Sometimes constants are typed.

Example:

```go
type UserStatus string  
  
const (  
    StatusActive   UserStatus = "active"  
    StatusInactive UserStatus = "inactive"  
)
```

This improves **type safety**.

Example:

```txt
var status UserStatus
```

Now only allowed values are from this type.

# Typical Constants Directory in Production

A production backend might look like:

```txt
internal/constants/  
│  
├── roles.go  
├── errors.go  
├── headers.go  
├── status.go  
├── pagination.go  
├── token.go  
└── permissions.go
```

Each file holds **constants related to one concept**.


# Common Mistakes with Constants

### 1. Putting everything in one file

Bad:

```txt
constants/constants.go (1000 lines)
```

Better:

```txt
roles.go  
errors.go  
status.go
```

---

### 2. Storing configuration in constants

Bad:

```go
const DB_PASSWORD = "123"
```

Secrets belong in **environment variables**, not constants.

---

### 3. Using constants for values that might change

Bad example:

```go
const Discount = 10
```

If this could change, it should be in **config or database**.


# Constants vs Config (Important Difference)

Developers often confuse these.

### Constants

Compile-time values.

Example:

const RoleAdmin = "admin"

Never changes.

### Config

Environment-dependent values.

Example:

```txt
DB_HOST  
JWT_SECRET  
PORT
```

Loaded from **environment variables or config files**.


# Mental Model

Think of constants as **the dictionary of your application**.

It defines the **official words the system understands**.

Example:

```txt
admin  
user  
active  
inactive  
authorization
```

If everyone uses the same dictionary, the system stays consistent.