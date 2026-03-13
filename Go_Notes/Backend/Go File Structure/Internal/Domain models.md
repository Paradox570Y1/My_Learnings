## **Domain Models**

A **domain model** represents the **core business logic** of your application. It defines the entities and behaviors that are relevant to your problem space, independent of how the data comes in or goes out.

In Go, a domain model is typically represented by a `struct` with fields that match the real-world concepts in your application.

**Example:**

```go
type User struct {  
    ID       int  
    Name     string  
    Email    string  
    Password string // internal detail, not exposed via API  
}
```

Key points about domain models:

- They contain **business rules** and **state**.
    
- They may have **methods** that operate on the data.
    
- They are **decoupled from the transport layer** (HTTP, gRPC, etc.).