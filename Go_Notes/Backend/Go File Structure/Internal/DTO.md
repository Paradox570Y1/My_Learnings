The **DTO (Data Transfer Object) pattern** is widely used in Go `internal/handlers` to **decouple HTTP request/response payloads from internal domain models**. This keeps your API layer clean, secure, and maintainable. Here’s a crisp yet detailed explanation.

>NOTE
>[[Domain]]
# 1. What a DTO Is

A **DTO** is a struct that defines **only the data your handler needs to receive or return**, separate from internal domain or database models.

- **Request DTO** – represents incoming request data
    
- **Response DTO** – represents outgoing response data
    

Example:

```go
// Request DTO  
type CreateUserRequest struct {  
    Name  string `json:"name"`  
    Email string `json:"email"`  
}  
```
  
// Response DTO  
```go
type UserResponse struct {  
    ID    string `json:"id"`  
    Name  string `json:"name"`  
    Email string `json:"email"`  
}
```

---

# 2. Why Use DTOs in Handlers

1. **Separation of concerns** – keeps *HTTP layer* independent of domain models
    
2. **Security** – avoid exposing sensitive fields (like password hashes)
    
3. **Validation** – DTOs are designed for *incoming request validation*
    
4. **Flexibility** – request/response structure can differ from internal models
    
5. **API versioning** – easy to change DTOs without touching services
    

---

# 3. How DTOs Are Used in Handlers

### 3.1 Receive a request

```go
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {  
    var req CreateUserRequest  
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {  
        http.Error(w, "invalid request", http.StatusBadRequest)  
        return  
    }  
      
    // validate  
    if req.Email == "" || req.Name == "" {  
        http.Error(w, "name and email required", http.StatusBadRequest)  
        return  
    }  
  
    // convert DTO → domain model  
    user := domain.User{  
        Name:  req.Name,  
        Email: req.Email,  
    }  
  
    createdUser, err := h.service.CreateUser(user)  
    if err != nil {  
        http.Error(w, err.Error(), http.StatusInternalServerError)  
        return  
    }  
  
    // convert domain model → response DTO  
    res := UserResponse{  
        ID:    createdUser.ID,  
        Name:  createdUser.Name,  
        Email: createdUser.Email,  
    }  
  
    json.NewEncoder(w).Encode(res)  
}
```

---

### 3.2 DTO Conversion

- **Incoming request → DTO → domain model**
    
- **Domain model → DTO → HTTP response**
    

This keeps **handlers thin** and **services unaware of transport layer**.

---

# 4. Advantages in Large Go Projects

- Consistent request/response handling
    
- Easier to add API validation (e.g., using `validator` package)
    
- DTOs can handle **different API versions** without changing services
    
- Protects internal models (prevents accidental exposure of fields)
    

---

# 5. Example with API Versioning

```go
// v1  
type CreateUserRequestV1 struct {  
    Name  string `json:"name"`  
    Email string `json:"email"`  
}  
  
// v2 adds Phone field  
type CreateUserRequestV2 struct {  
    Name  string `json:"name"`  
    Email string `json:"email"`  
    Phone string `json:"phone"`  
}
```

Handlers can switch between versions without touching `service` or `repository` layers.

## **Solution: Mapping & Defaults**

The general approach is:

1. **DTO → Domain Model Mapping**
    
    - Service layer takes the DTO and **creates a domain model**, filling in missing fields with defaults, generated values, or derived data.
        
2. **Do not make DTO contain everything**
    
    - DTO is intentionally minimal; service/domain layer is responsible for completing the domain object.

# Example in Go

```go
// DTO received from HTTP request
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Domain model
type User struct {
    ID        int
    Name      string
    Email     string
    Password  string // internal
    CreatedAt time.Time
    Status    string
}

// Service function: maps DTO to domain model
func CreateUser(req CreateUserRequest) (*User, error) {
    user := &User{
        Name:      req.Name,
        Email:     req.Email,
        Password:  generateRandomPassword(), // fill internal field
        CreatedAt: time.Now(),                // fill missing timestamp
        Status:    "active",                  // default value
    }
    // save user to DB
    return user, nil
}
```


✅ Points to note:

- DTO only contains **what comes from the client**.
    
- Service **completes the domain model** as needed.
    
- This ensures domain logic **doesn’t rely on HTTP layer**.

### Cleaner Way To convert DTO to Domain model


In bigger Go projects, people sometimes use small **mapping helpers**:

A **mapping helper** is usually a **method or function** that converts a DTO into a domain model (or vice versa). Instead of writing the mapping inline in the service every time, you centralize it.

```go
// Mapping helper attached to the DTO
func (req CreateUserRequest) ToDomain() *User {
    return &User{
        Name:      req.Name,
        Email:     req.Email,
        Password:  generateRandomPassword(), // fill internal field
        CreatedAt: time.Now(),                // default timestamp
        Status:    "active",                  // default status
    }
}
```

Then your **service function becomes simpler**:
```go
func CreateUser(req CreateUserRequest) (*User, error) {
    user := req.ToDomain() // mapping handled by helper
    // save user to DB
    return user, nil
}
```

This keeps mapping **organized**, but mapping logic **still belongs to service/domain**, not DTO.

- `&User{...}` is a **composite literal with `&`**.
    
- This **creates a new `User` struct in memory** and returns a **pointer to it**.
    
- Every time you call `ToDomain()`, a **brand new `User` instance** is created.
    

So it **does not point to the same struct every time**. Each call produces a **separate object in memory**.