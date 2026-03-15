When you build a **production Go backend**, the `internal/routes` directory is usually where **all HTTP routes of your application are registered and organized**.

Think of it like the **traffic controller of your API**.  
It decides:

- Which **URL path** calls which **handler**
    
- Which **middlewares** apply
    
- How **different modules expose endpoints**
    

# First Principle: What is a Route?

A **route** maps a **URL + HTTP method → handler function**.

Example request:

```txt
GET /users/12
```

This must map to something like:

```txt
GetUserHandler()
```

Example mapping:

```txt
GET     /users       -> ListUsers  
GET     /users/:id   -> GetUser  
POST    /users       -> CreateUser
```

This mapping is exactly what the **routes directory manages**.


# Why a `routes` Directory Exists

In a **small Go app**, routing is often inside `main.go`.

Example:

```go
r := gin.Default()  
  
r.GET("/users", handler.GetUsers)  
r.POST("/users", handler.CreateUser)  
  
r.Run(":8080")
```

This becomes **unmanageable in large systems** because:

- 50+ endpoints
    
- multiple teams
    
- multiple services
    

So we move routing logic to:

```txt
internal/routes
```

Typical purpose:

```txt
main.go → start server  
routes → define endpoints  
handlers → implement logic  
services → business logic  
repository → database
```


# Typical Production Structure

A common structure looks like this:

```go
internal/  
│  
├── routes/  
│   ├── router.go  
│   ├── user_routes.go  
│   ├── auth_routes.go  
│   └── health_routes.go  
│  
├── handlers/  
├── services/  
├── repository/
```

Each file registers routes **for one domain**.
This keeps routing **modular and scalable**.

# Responsibilities of the Routes Layer

The routes layer has **very specific responsibilities**.

It should:

✔ Define API endpoints  
✔ Connect routes → handlers  
✔ Apply middleware  
✔ Group endpoints  
✔ Version APIs

It should **NOT**:

❌ contain business logic  
❌ query database  
❌ validate deeply  
❌ perform calculations

Routes layer = **wiring layer**


# Typical Router Setup

Most Go backends use frameworks like:

- **Gin**
    
- **Chi**
    
- **Echo**
    
- **Fiber**
    

Example with **Gin**.

### router.go

```go
internal/routes/router.go
```

Example:

```go
package routes  
  
import (  
	"github.com/gin-gonic/gin"  
)  
  
func SetupRouter() *gin.Engine {  
  
	r := gin.Default()  
  
	// register routes  
	RegisterUserRoutes(r)  
	RegisterAuthRoutes(r)  
  
	return r  
}
```

This file acts as the **central router builder**.


Explanation:

```txt
/users        -> GetUsers  
/users/:id    -> GetUser
```

Routes grouped using:

```go
r.Group("/users")
```


# Using Route Groups

Groups help with:

- prefixes
    
- middlewares
    
- versioning
    

Example:

```go
/api/v1/users  
/api/v1/auth

Example code:

api := r.Group("/api")  
  
v1 := api.Group("/v1")  
  
RegisterUserRoutes(v1)  
RegisterAuthRoutes(v1)
```

This creates:

```txt
/api/v1/users  
/api/v1/auth
```

Very common in production APIs.


# Routes With Middleware

Routes layer is responsible for attaching middleware.

Example:

```go
auth := r.Group("/auth")  
{  
	auth.POST("/login", handlers.Login)  
	auth.POST("/signup", handlers.Signup)  
}
```

Protected routes:

```go
users := r.Group("/users")  
users.Use(middleware.AuthMiddleware())  
  
{  
	users.GET("/", handlers.GetUsers)  
}
```

Now:

```go
GET /users
```

requires authentication.


# How `main.go` Uses Routes

Example:

```txt
cmd/server/main.go
```

```go
package main  
  
import (  
	"myapp/internal/routes"  
)  
  
func main() {  
  
	router := routes.SetupRouter()  
  
	router.Run(":8080")  
}
```

Flow:

```txt
main.go  
   ↓  
routes.SetupRouter()  
   ↓  
RegisterUserRoutes  
RegisterAuthRoutes  
   ↓  
handlers
```


# Full Request Flow (Production)

When a request arrives:

```go
Client Request  
     ↓  
Router  
     ↓  
Middleware  
     ↓  
Handler  
     ↓  
Service  
     ↓  
Repository  
     ↓  
Database
```

Routes layer sits **between server and handlers**.


# A Real Production Routes Folder

Example:

internal/routes

```txt
routes/  
│  
├── router.go  
├── health_routes.go  
├── auth_routes.go  
├── user_routes.go  
├── product_routes.go  
└── order_routes.go
```

Example endpoints:

```txt
GET    /health  
POST   /auth/login  
POST   /auth/register  
GET    /users  
GET    /products  
POST   /orders
```

Each module controls its own endpoints.



# Health Route Example

Example:

internal/routes/health_routes.go

```go
package routes  
  
import (  
	"github.com/gin-gonic/gin"  
)  
  
func RegisterHealthRoutes(r *gin.Engine) {  
  
	r.GET("/health", func(c *gin.Context) {  
		c.JSON(200, gin.H{  
			"status": "ok",  
		})  
	})  
}
```

Used by:

- Kubernetes
    
- load balancers
    
- monitoring tools



# Advanced Production Router

Sometimes routes receive dependencies.

Example:

```go
func RegisterUserRoutes(r *gin.RouterGroup, h *handlers.UserHandler) {  
  
	users := r.Group("/users")  
  
	users.GET("/", h.GetUsers)  
	users.POST("/", h.CreateUser)  
}
```

Router injects handler.

Example:

handlers -> services -> repositories

This allows **dependency injection**.


# Example Full Setup

Production style.

### router.go

```go
func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {  
  
	r := gin.Default()  
  
	api := r.Group("/api/v1")  
  
	RegisterUserRoutes(api, userHandler)  
  
	return r  
}
```


# Common Beginner Mistakes

### Putting logic in routes

Bad:

```go
routes/user_routes.go

users.GET("/", func(c *gin.Context){  
    users := db.FindAll()  
})
```

Routes should **never access DB**.


### One giant router file

Bad:

```txt
router.go (2000 lines)
```

Better:

```go
user_routes.go  
product_routes.go  
auth_routes.go
```


# Real Industry Example (Simplified)

Example endpoints:

```go
GET /api/v1/users  
GET /api/v1/users/:id  
POST /api/v1/auth/login  
POST /api/v1/orders
```

Routes connect them to:

```go
handlers/user_handler.go  
handlers/auth_handler.go  
handlers/order_handler.go
```


# Summary

The **routes directory**:

Purpose:

```txt
Define API endpoints  
Connect routes → handlers  
Group endpoints  
Apply middleware  
Version APIs
```

Responsibilities:

```txt
URL structure  
HTTP methods  
Route grouping  
Middleware binding
```

Not responsible for:

```txt
Business logic  
Database  
Validation
```


✅ **One important insight**

The routes layer is where you **design your public API structure**.

A clean routes layer = **clean API design**.