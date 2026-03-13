The **`internal/middleware`** directory usually contains reusable components that **intercept and process HTTP requests before or after they reach handlers**.


>NOTE
>Router Libraries
>	-Gorilla Mux
      - Chi
      - Gin
     - Echo
# 1. What is Middleware in Go?

**Middleware** is a function that sits **between the HTTP request and the final handler**.

Flow:

```txt
Client Request  
      ↓  
Middleware (auth, logging, rate limit, etc.)  
      ↓  
Handler   
      ↓  
Response
```

In Go’s `net/http`, middleware usually wraps `http.Handler`.

>NOTE
>[[net_http]]

### Basic Idea

```go
func Middleware(next http.Handler) http.Handler
```

This means:

- Middleware receives the **next handler**
    
- It returns **another handler**
    
- It can execute **before and after** the next handler

# Why Middleware Exists

Middleware prevents repeating common logic in every handler.

Without middleware:

```go
func GetUser(w http.ResponseWriter, r *http.Request) {
    log.Println("Request received")

    if !isAuthenticated(r) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // actual logic
}
```

You would repeat this **in every endpoint**.

Middleware solves this.

---

# 3. Responsibilities of Middleware

Middleware usually handles **cross-cutting concerns**.

Common responsibilities:

### 1. Logging

Log every request.

Example:

```txt
GET /users 200 12ms  
POST /login 401 5ms
```

---

### 2. Authentication

Verify:

- JWT tokens
    
- API keys
    
- sessions
    

---

### 3. Authorization

Check permissions:

```txt
admin  
user  
guest
```

---

### 4. Request Validation

Examples:

- Check headers
    
- Content-Type validation
    
- Request size limit
    

---

### 5. Rate Limiting

Prevent abuse.

Example:

```txt
100 requests/minute
```

---

### 6. Panic Recovery

Prevent server crash.

```txt
panic → recover → return 500
```

---

### 7. CORS Handling

Allow browser cross-origin requests.

---

### 8. Request Context Injection

Add data into request context:

```txt
userID  
requestID  
traceID
```

---

# 4. Why Middleware Lives in `/internal/middleware`

### Benefits

1. Prevent external dependency
    
2. Keeps implementation private
    
3. Encourages modular design
    

Example:

```txt
/internal/middleware/auth.go  
/internal/middleware/logging.go  
/internal/middleware/recovery.go
```

---

# 5. Middleware Signature in Go

Standard middleware pattern:

```go
func MyMiddleware(next http.Handler) http.Handler
```

Example:

```go
func Logging(next http.Handler) http.Handler {  
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {  
  
        log.Println("Request started")  
  
        next.ServeHTTP(w, r)  
  
        log.Println("Request finished")  
    })  
}
```

Important line:

```go
next.ServeHTTP(w, r)
```

This calls the **next handler in the chain**.

---

# 6. Middleware Execution Flow

Example chain:

Recovery → Logging → Auth → Handler

Execution order:

```txt
Recovery start  
   Logging start  
      Auth start  
         Handler  
      Auth end  
   Logging end  
Recovery end
```

This nesting is important.

---

# 7. Example Project Structure

```go
project/  
│  
├── cmd/  
│    └── server/main.go  
│  
├── internal/  
│    ├── handlers/  
│    │     └── user_handler.go  
│    │  
│    ├── middleware/  
│    │     ├── auth.go  
│    │     ├── logging.go  
│    │     └── recovery.go  
│    │  
│    ├── service/  
│    └── repository/
```

---

# 8. Example: Logging Middleware

`internal/middleware/logging.go`

```go
package middleware  
  
import (  
    "log"  
    "net/http"  
    "time"  
)  
  
func Logging(next http.Handler) http.Handler {  
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {  
  
        start := time.Now()  
  
        log.Printf("Started %s %s", r.Method, r.URL.Path)  
  
        next.ServeHTTP(w, r)  
  
        log.Printf("Completed in %v", time.Since(start))  
    })  
}
```

---

# 9. Example: Auth Middleware

`internal/middleware/auth.go`

```go
package middleware  
  
import (  
    "net/http"  
)  
  
func Auth(next http.Handler) http.Handler {  
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {  
  
        token := r.Header.Get("Authorization")  
  
        if token == "" {  
            http.Error(w, "Unauthorized", http.StatusUnauthorized)  
            return  
        }  
  
        next.ServeHTTP(w, r)  
    })  
}
```

---

# 10. Example: Recovery Middleware

Prevents server crashes.

```go
func Recovery(next http.Handler) http.Handler {  
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {  
  
        defer func() {  
            if err := recover(); err != nil {  
                http.Error(w, "Internal Server Error", 500)  
            }  
        }()  
  
        next.ServeHTTP(w, r)  
    })  
}
```

---

# 11. How Middleware is Used in `main.go`

Example:

```go
mux := http.NewServeMux()  
  
mux.Handle("/users", userHandler)  
  
handler := middleware.Logging(  
            middleware.Auth(  
                middleware.Recovery(mux),  
            ),  
          )  
http.ListenAndServe(":8080", handler)
```


Chain:

```txt
Logging  
   ↓  
Auth  
   ↓  
Recovery  
   ↓  
Mux Router
```

---

# 12. Middleware Chaining Pattern (Better)

Instead of nesting, many teams use a **chain function**.

Example:

```go
func Chain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {  
    for i := len(m) - 1; i >= 0; i-- {  
        h = m[i](h)  
    }  
    return h  
}

Usage:

handler := Chain(  
    mux,  
    middleware.Recovery,  
    middleware.Logging,  
    middleware.Auth,  
)
```

Much cleaner.

---

# 13. Context Middleware (Important)

Middleware often injects data into **request context**.

Example:

```go
type contextKey string  
  
const UserIDKey contextKey = "userID"

Middleware:

func UserContext(next http.Handler) http.Handler {  
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {  
  
        ctx := context.WithValue(r.Context(), UserIDKey, "123")  
  
        next.ServeHTTP(w, r.WithContext(ctx))  
    })  
}
```

Handler usage:

```go
userID := r.Context().Value(middleware.UserIDKey)
```

---

# 14. Must-Know Middleware Best Practices

### 1. Always call `next.ServeHTTP`

Unless stopping request intentionally.

---

### 2. Do not store request data in global variables

Use **context instead**.

---

### 3. Middleware must be fast

They run **on every request**.

---

### 4. Avoid heavy DB queries inside middleware

Except authentication if needed.

---

### 5. Order matters

Example:

```go
Recovery  
Logging  
Auth  
Handler
```

If recovery is inside logging, panic may crash logging.

---

### 6. Middleware should be reusable

Avoid coupling to specific handlers.

Bad:

```txt
middleware for only /users
```

Good:

```txt
generic auth middleware
```

---

# 15. Popular Go Middleware Libraries

Many frameworks provide middleware support.

Examples:

### Router Libraries

- Gorilla Mux
    
- Chi
    
- Gin
    
- Echo
    

Example in **Chi**:

```go
r := chi.NewRouter()  
  
r.Use(middleware.Logger)  
r.Use(middleware.Recoverer)
```

---

# 16. Real Production Middleware Stack

Typical production API:

```txt
Request  
 ↓  
RequestID  
 ↓  
RealIP  
 ↓  
Logger  
 ↓  
Recovery  
 ↓  
RateLimit  
 ↓  
Auth  
 ↓  
Handler
```

---

# 17. What Senior Go Developers Expect

A good middleware implementation should:

✔ be **stateless**  
✔ be **reusable**  
✔ use **context** properly  
✔ avoid **business logic**  
✔ follow **clean chain architecture**

---

# 18. Most Common Middleware Files

Typical directory:

```txt
internal/middleware/  
    auth.go  
    cors.go  
    logging.go  
    recovery.go  
    ratelimit.go  
    requestid.go
```

---

# 19. Middleware vs Interceptors

Conceptually similar.

|Technology|Name|
|---|---|
|Go HTTP|Middleware|
|Express.js|Middleware|
|gRPC|Interceptor|
|Spring Boot|Filter|

---

# 20. One-Line Mental Model

Middleware is basically:

decorators for HTTP handlers

They **wrap handlers with additional behavior**.