- **Purpose**: This folder contains the **HTTP handlers** or **controllers** that manage incoming HTTP requests and respond accordingly.
    
- **What it does**: It takes the incoming requests (from APIs, user interfaces, etc.), processes them (maybe by calling services or querying the database), and sends a response back to the client.
    
- **Why use it**: Separating the handlers makes the HTTP interface of the application clean and modular. Handlers only need to focus on handling requests, and they delegate business logic to services.
    
- **Example file (`user_handler.go`)**: It could define RESTful API routes for managing users, like creating a user, updating details, etc.

>NOTE
>[[DTO]]

In Go projects using **layered / clean architecture**, the `handlers` directory inside `internal` contains the **HTTP request handling logic** of the application.

- `internal/handlers` contains **functions or structs that handle incoming requests** from clients (HTTP, gRPC, etc.) and return responses.
- Handlers act as the **entry point to the application logic**.

Example responsibilities:

- Receive HTTP requests
    
- Parse request body / parameters
    
- Validate input
    
- Call service layer
    
- Format and send response

# Why a `handlers` Directory Is Needed

Separates Transport logic from Business logic
Without handlers, request logic often gets mixed with business logic.

Bad example:

```go
func CreateUser(w http.ResponseWriter, r *http.Request) {  
    body, _ := io.ReadAll(r.Body)  
    db.InsertUser(body)  
}
```

Problems:

- HTTP logic mixed with business logic
    
- Hard to test
    
- Hard to reuse services
    
- Difficult to maintain
    
---

# Typical Responsibilities of Handlers

Handlers usually perform:

### 1. Parsing requests

Example:

```go
json.NewDecoder(r.Body).Decode(&req)
```

They extract:

- JSON body
    
- Query parameters
    
- Path parameters
    
- Headers
    

---

### 2. Input validation

Example:

```go
if req.Email == "" {  
    http.Error(w, "email required", 400)  
}
```

---

### 3. Calling business services

Example:

```go
user, err := userService.CreateUser(req)
```

---

### 4. Formatting responses

Example:

```go
json.NewEncoder(w).Encode(user)
```

---

### 5. Handling HTTP status codes

Examples:

200 OK  
201 Created  
400 Bad Request  
500 Internal Server Error

---

# 6. Typical Handler Implementation

Example:

```go
package handlers  
  
import (  
    "encoding/json"  
    "net/http"  
)  
  
type UserHandler struct {  
    service UserService  
}  
  
func NewUserHandler(s UserService) *UserHandler {  
    return &UserHandler{service: s}  
}  
  
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {  
    var req CreateUserRequest  
  
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {  
        http.Error(w, "invalid request", http.StatusBadRequest)  
        return  
    }  
  
    user, err := h.service.CreateUser(req)  
    if err != nil {  
        http.Error(w, err.Error(), http.StatusInternalServerError)  
        return  
    }  
  
    json.NewEncoder(w).Encode(user)  
}
```

---

# 7. How Handlers Are Used in `main.go`

Handlers are connected to routes.

Example:

```go
cfg := config.Load()  
  
userService := service.NewUserService()  
userHandler := handlers.NewUserHandler(userService)  
  
http.HandleFunc("/users", userHandler.CreateUser)  
http.ListenAndServe(":8080", nil)
```

---

# 8. Common Files in `internal/handlers`

Example structure:

```txt
internal/  
 └── handlers/  
      ├── user_handler.go  
      ├── auth_handler.go  
      ├── health_handler.go  
      └── middleware.go
```
---

# 9. Handlers vs Services

|Layer|Responsibility|
|---|---|
|**Handlers**|HTTP request/response handling|
|**Services**|Business logic|
|**Repositories**|Database access|

Example:

Handler → "Create user request"  
Service → "Business rules for user creation"  
Repository → "Insert into database"


# 11. Best Practices for Handlers

✔ Keep handlers **thin**  
✔ Do not write business logic in handlers  
✔ Use **DTO structs** for requests/responses  
✔ Return proper HTTP status codes  
✔ Use middleware for cross-cutting concerns


Example responsibilities **not for handlers**:

- Complex business logic
    
- Database queries
    
- Long processing tasks