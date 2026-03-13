The **`net/http` package** in Go is the standard library used to build **HTTP servers and clients**. If you want to work with web APIs, microservices, or web servers in Go, this package is the core tool.


**HTTP (HyperText Transfer Protocol)** is a **request–response protocol**.

1. Client (browser, mobile app, curl) sends a **request**
    
2. Server processes it
    
3. Server sends a **response**

Example:

```txt
GET /users HTTP/1.1  
Host: example.com
```

Server responds

```txt
HTTP/1.1 200 OK
Content-Type: application/json

{"name":"Rahul"}
```

Important parts:

|Term|Meaning|
|---|---|
|Request|Client message to server|
|Response|Server reply|
|Method|GET, POST, PUT, DELETE|
|Header|Metadata|
|Body|Actual data|
|Status Code|Result of request|

# 3. Core Concept: `http.Handler`

Everything in Go's HTTP server is built around **`http.Handler`**.

Definition:

```go
type Handler interface {  
    ServeHTTP(ResponseWriter, *Request)  
}
```

This means **any type implementing `ServeHTTP` is a Handler.**

### Handler Responsibility

A handler:

1. Receives the request
    
2. Processes it
    
3. Writes the response
    

---

# 4. Request Object (`http.Request`)

The **`Request`** struct represents the incoming HTTP request.

Example fields:

```go
type Request struct {  
    Method string  
    URL *url.URL  
    Header Header  
    Body io.ReadCloser  
    Host string  
    RemoteAddr string  
}

Example usage:

func handler(w http.ResponseWriter, r *http.Request) {  
    fmt.Println(r.Method)  
    fmt.Println(r.URL.Path)  
}
```

Common fields:

|Field|Meaning|
|---|---|
|`Method`|GET / POST|
|`URL.Path`|request path|
|`Header`|request headers|
|`Body`|request payload|
|`Query()`|URL query parameters|
|`Context()`|request context|

Example query:

/users?id=10

id := r.URL.Query().Get("id")

---

# 5. ResponseWriter

`http.ResponseWriter` is used to send the response back.

Interface:

```go
type ResponseWriter interface {  
    Header() Header  
    Write([]byte)  
    WriteHeader(statusCode int)  
}
```

Example:

```go
func handler(w http.ResponseWriter, r *http.Request) {  
    w.WriteHeader(200)  
    w.Write([]byte("Hello"))  
}
```

Example JSON response:

```go
func handler(w http.ResponseWriter, r *http.Request) {  
    w.Header().Set("Content-Type", "application/json")  
    json.NewEncoder(w).Encode(map[string]string{  
        "message": "hello",  
    })  
}
```

---

# 6. Simplest HTTP Server

```go
package main  
  
import (  
    "fmt"  
    "net/http"  
)  
  
func handler(w http.ResponseWriter, r *http.Request) {  
    fmt.Fprintf(w, "Hello World")  
}  
  
func main() {  
    http.HandleFunc("/", handler)  
    http.ListenAndServe(":8080", nil)  
}
```

Server runs on:

http://localhost:8080

---

# 7. `http.HandleFunc` vs `http.Handle`

### HandleFunc

`http.HandleFunc("/hello", handler)`

Used with **functions**

### Handle

`http.Handle("/hello", handlerObject)`

Used with **Handler interface**

---

# 8. `http.HandlerFunc`

Go provides an adapter:

```go
type HandlerFunc func(ResponseWriter, *Request)
```

It converts functions into `Handler`.

Example:

```go
func hello(w http.ResponseWriter, r *http.Request) {  
    fmt.Fprint(w, "hello")  
}  
```
  
```go
http.Handle("/hello", http.HandlerFunc(hello))
```

But usually we use:

```go
http.HandleFunc()
```

---

# 9. Creating Custom Handlers

You can implement your own handler struct.

Example:

```go
type UserHandler struct{}  
  
func (h UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {  
    fmt.Fprintln(w, "User handler")  
}
```
Register it:

```go
http.Handle("/users", UserHandler{})
```

---

# 10. `http.ServeMux` (Router)

Go's built-in router.

```txt
ServeMux  
   |  
   |--- /users  
   |--- /products
```

Example:

```txt
mux := http.NewServeMux()  
  
mux.HandleFunc("/users", usersHandler)  
mux.HandleFunc("/products", productsHandler)  
```
  
```go
http.ListenAndServe(":8080", mux)
```

Routing rules:

```txt
/      matches everything  
/users matches /users  
/users/ matches /users/*
```
---

# 11. Middleware Concept

Middleware wraps handlers.

Example:

```txt
Request  
   ↓  
Middleware  
   ↓  
Handler  
   ↓  
Response
```

Example logging middleware:

```go
func loggingMiddleware(next http.Handler) http.Handler {  
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {  
  
        fmt.Println("Request:", r.URL.Path)  
  
        next.ServeHTTP(w, r)  
  
        fmt.Println("Response sent")  
    })  
}
```

Use it:

```go
mux := http.NewServeMux()  
mux.HandleFunc("/hello", hello)  
  
wrapped := loggingMiddleware(mux)  
  
http.ListenAndServe(":8080", wrapped)
```

---

# 12. Parsing Request Body

Example JSON request:

```go
POST /users  
{  
 "name":"rahul"  
}
```

Code:

```go
type User struct {  
    Name string `json:"name"`  
}  

func handler(w http.ResponseWriter, r *http.Request) {  
  
    var user User  
  
    json.NewDecoder(r.Body).Decode(&user)  
  
    fmt.Println(user.Name)  
}
```
  


---

# 13. Sending JSON Response

```go
func handler(w http.ResponseWriter, r *http.Request) {  
  
    response := map[string]string{  
        "status": "ok",  
    }  
  
    w.Header().Set("Content-Type", "application/json")  
  
    json.NewEncoder(w).Encode(response)  
}
```

---

# 14. HTTP Methods

Check method:

```go
if r.Method == http.MethodPost {  
}

Constants:

http.MethodGet  
http.MethodPost  
http.MethodPut  
http.MethodDelete  
http.MethodPatch

Example REST style:

switch r.Method {  
  
case http.MethodGet:  
    getUser()  
  
case http.MethodPost:  
    createUser()  
  
}
```

---

# 15. URL Parameters

Example:

```cmd
/users?id=10
```

Code:

```go
id := r.URL.Query().Get("id")
```

---

# 16. Headers

Read header:

```go
auth := r.Header.Get("Authorization")
```

Write header:

```go
w.Header().Set("Content-Type", "application/json")
```

---

# 17. Status Codes

Examples:

```go
w.WriteHeader(http.StatusOK)  
w.WriteHeader(http.StatusBadRequest)  
w.WriteHeader(http.StatusNotFound)
```

Common ones:

|Code|Meaning|
|---|---|
|200|OK|
|201|Created|
|400|Bad Request|
|401|Unauthorized|
|404|Not Found|
|500|Internal Server Error|

---

# 18. Serving Static Files

Example:

```go
fs := http.FileServer(http.Dir("./static"))  
  
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

Access:

```go
/static/image.png
```

---

# 19. Context (Important for Production)

Each request has a **context**.

Used for:

- cancellation
    
- deadlines
    
- tracing
    
- request scoped data
    

Example:

```go
ctx := r.Context()  
  
select {  
case <-ctx.Done():  
    fmt.Println("request cancelled")  
}
```

---

# 20. Graceful Server Setup

Instead of `ListenAndServe`, production servers use:

```go
server := &http.Server{  
    Addr: ":8080",  
    Handler: mux,  
}  
  
server.ListenAndServe()
```

You can configure:

- timeouts
    
- TLS
    
- shutdown
    

---

# 21. HTTP Client (net/http)

Example request:

```go
resp, err := http.Get("https://api.github.com")
```

Custom client:

```go
client := &http.Client{}  
  
req, _ := http.NewRequest("GET", url, nil)  
  
resp, err := client.Do(req)
```

---

# 22. Full Example API

```go
package main  
  
import (  
    "encoding/json"  
    "net/http"  
)  
  
type User struct {  
    Name string `json:"name"`  
}  
  
func createUser(w http.ResponseWriter, r *http.Request) {  
  
    if r.Method != http.MethodPost {  
        w.WriteHeader(http.StatusMethodNotAllowed)  
        return  
    }  
  
    var user User  
  
    json.NewDecoder(r.Body).Decode(&user)  
  
    json.NewEncoder(w).Encode(user)  
}  
  
func main() {  
  
    mux := http.NewServeMux()  
  
    mux.HandleFunc("/users", createUser)  
  
    http.ListenAndServe(":8080", mux)  
}
```

---

# 23. Real Production Stack

Typical Go web stack:

```txt
net/http  
   ↓  
Router (gorilla/mux, chi, gin)  
   ↓  
Middleware  
   ↓  
Handlers  
   ↓  
Service Layer  
   ↓  
Repository  
   ↓  
Database
```

---

# 24. Key Things to Master

To truly understand `net/http`, focus on:

1. `http.Handler`
    
2. `http.HandlerFunc`
    
3. `http.ResponseWriter`
    
4. `http.Request`
    
5. `ServeMux`
    
6. Middleware pattern
    
7. JSON encoding/decoding
    
8. Context
    
9. HTTP client