# Purpose of models Directory

The `models` directory is typically where you define **data structures** that represent the core entities of your application. These are sometimes called “[[Domain models]]” or “business models”.

### Key Roles:

1. **Centralizing Core Data Structures**
    
    - Example: `User`, `Product`, `Order`.
        
    - Having them in one place makes it easy to maintain and share across your internal services.
        
2. **Type Safety**
    
    - Go is strongly typed. Having defined models helps avoid passing unstructured data (like `map[string]interface{}`) around.
        
3. **Abstraction**
    
    - Models decouple your **domain entities** from storage, API layer, or business logic. You can use the same model in your services, database layer, or REST API layer.
        
4. **Consistency**
    
    - Shared definitions across packages prevent mismatch issues.
        
    - Example: Ensures that all parts of the app use the same definition of a `User`.

## 3. Typical Content of `models` Directory

You usually find:

### a) Structs for Entities

```go
package models  
  
type User struct {  
    ID       int64  
    Name     string  
    Email    string  
    IsActive bool  
}  
  
type Product struct {  
    ID    int64  
    Name  string  
    Price float64  
}
```

### b) Validation / Helper Methods

Some methods on models for validation or convenience:

```go
func (u *User) IsValidEmail() bool {  
    // simple validation example  
    return strings.Contains(u.Email, "@")  
}
```

### c) Constants / Enums

For status codes or type definitions:

```go
const (  
    UserStatusActive   = "active"  
    UserStatusInactive = "inactive"  
)
```

### d) Database Tags (optional)

If using an ORM like GORM:

```go
type User struct {  
    ID    int64  `gorm:"primaryKey"`  
    Name  string `gorm:"size:100"`  
    Email string `gorm:"uniqueIndex"`  
}
```



## 5. Significance in Large Applications

In a large Go application:

- `internal/models` acts as the **single source of truth** for your domain entities.
    
- Helps maintain **clean architecture**, separating:
    
    - `models` → domain/data
        
    - `service` → business logic
        
    - `handlers` → HTTP/API endpoints
        
- Makes testing easier because you can create mock instances of your models without touching other layers.


## 6. Best Practices for `models` Directory

1. **Keep it Focused**  
    Only include core domain models, not DTOs or API-specific structures.
    
2. **Avoid Business Logic Here**  
    Models should represent data. Validation can exist, but complex logic belongs in services.
    
3. **Use Small, Composable Structs**
    
    - Easier to test and maintain.
        
    - For example, separate `UserProfile` from `UserAuth`.
        
4. **Tag Carefully**  
    If using JSON, database, or protobuf, tags should be clear and consistent.
    
5. **Documentation**
    
    - Briefly document each struct and its fields.
        
    - Makes it clear what each model represents.

## 7. Example Directory Layout

```txt
internal/  
└── models/  
    ├── user.go       // User entity and helpers  
    ├── product.go    // Product entity  
    ├── order.go      // Order entity  
    └── enums.go      // Constants / enums
```

Usage in service:

```go
package service  
  
import "myapp/internal/models"  
  
func CreateUser(name, email string) *models.User {  
    return &models.User{  
        Name:  name,  
        Email: email,  
        IsActive: true,  
    }  
}
```


✅ **Summary:**

- `internal/models` = **centralized place for domain entities**.
    
- Makes your code **type-safe, maintainable, and encapsulated**.
    
- Promotes **clean architecture** by separating core data from logic or external APIs.
    
- Best practices: keep it focused, avoid business logic, and document clearly.