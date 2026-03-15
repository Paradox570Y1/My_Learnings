# Understanding Go Interfaces and Dependency Injection

---

# 1. The Problem Interfaces Solve

Before learning interfaces, you need to understand the problem they solve.

Imagine you write a simple service.

```go
type MySQLDB struct {}

func (db MySQLDB) SaveUser(name string) {
    fmt.Println("Saving user to MySQL:", name)
}

type UserService struct {
    db MySQLDB
}

func (s UserService) Register(name string) {
    s.db.SaveUser(name)
}
```

**Usage**

```go
db := MySQLDB{}
service := UserService{db}

service.Register("Alice")
```

### What is the problem here?

`UserService` is **tightly coupled** to `MySQLDB`.

Meaning:

```
UserService ---> MySQLDB
```

Now imagine:

- Tomorrow you switch to PostgreSQL
- Or you want a **Mock DB for testing**
- Or MongoDB

You would need to **rewrite UserService**.

That is **bad design**.

---

# 2. The Key Idea of Interfaces

Instead of depending on a **concrete type**, we depend on **behavior**.

In Go:

> Behavior = Interface

Example:

```go
type UserRepository interface {
    SaveUser(name string)
}
```

This means:

Any type that has

```
SaveUser(string)
```

**Automatically satisfies this interface.**

There is **no `implements` keyword in Go.**

---

# 3. Implementing the Interface

Now our database struct implements it.

```go
type MySQLDB struct {}

func (db MySQLDB) SaveUser(name string) {
    fmt.Println("Saving user to MySQL:", name)
}
```

Because it has the method:

```
SaveUser(string)
```

Go automatically treats `MySQLDB` as `UserRepository`.

---

# 4. Using the Interface

Now the service depends on the **interface**, not the database.

```go
type UserService struct {
    repo UserRepository
}

func (s UserService) Register(name string) {
    s.repo.SaveUser(name)
}
```

Now this works:

```go
db := MySQLDB{}

service := UserService{
    repo: db,
}

service.Register("Alice")
```

Architecture now:

```
UserService ---> UserRepository (interface)
                      ↑
                      |
                   MySQLDB
```

This is called **decoupling**.

---

# 5. Why This Is Powerful

Now we can easily **swap implementations**.

### MySQL

```go
type MySQLDB struct{}

func (db MySQLDB) SaveUser(name string) {
    fmt.Println("Saving user in MySQL")
}
```

### MongoDB

```go
type MongoDB struct{}

func (db MongoDB) SaveUser(name string) {
    fmt.Println("Saving user in MongoDB")
}
```

Now this works without changing `UserService`.

```go
repo := MongoDB{}

service := UserService{repo}

service.Register("Alice")
```

Same service, **different database**.

---

# 6. Real Backend Architecture

In production Go apps the flow looks like:

```
Handler (HTTP)
     ↓
Service (Business Logic)
     ↓
Repository Interface
     ↓
Database Implementation
```

Example:

```
handler -> service -> repository(interface) -> mysql implementation
```

---

# 7. Dependency Injection (DI)

Dependency Injection means:

> Instead of creating dependencies inside a struct, we **inject them from outside**.

### ❌ BAD (No DI)

```go
type UserService struct {}

func (s UserService) Register(name string) {
    db := MySQLDB{}
    db.SaveUser(name)
}
```

Problems:

- Hard to test
- Hard to replace database
- Tight coupling

---

### ✅ GOOD (Dependency Injection)

```go
type UserService struct {
    repo UserRepository
}
```

Now dependency comes from outside.

```go
repo := MySQLDB{}

service := UserService{
    repo: repo,
}
```

This is called **manual dependency injection**.

---

# 8. Real Go Backend Example

This resembles **production code**.

### Repository Interface

```go
type UserRepository interface {
    CreateUser(name string) error
}
```

### MySQL Implementation

```go
type MySQLRepo struct {}

func (r MySQLRepo) CreateUser(name string) error {
    fmt.Println("Insert into MySQL:", name)
    return nil
}
```

### Service

```go
type UserService struct {
    repo UserRepository
}

func (s UserService) RegisterUser(name string) error {
    return s.repo.CreateUser(name)
}
```

### Main (Injection happens here)

```go
func main() {

    repo := MySQLRepo{}

    service := UserService{
        repo: repo,
    }

    service.RegisterUser("Alice")

}
```

The **main function wires everything together.**

---

# 9. Testing Becomes Extremely Easy

Now we can create a **mock repository**.

```go
type MockRepo struct{}

func (m MockRepo) CreateUser(name string) error {
    fmt.Println("Mock user saved")
    return nil
}
```

Test:

```go
repo := MockRepo{}

service := UserService{
    repo: repo,
}

service.RegisterUser("Alice")
```

Now you can test **without a real database**.

---

# 10. Golden Rule of Interfaces in Go

A famous Go principle:

> **Accept interfaces, return structs**

Meaning:

- Services **accept interfaces**
- Repositories **return concrete types**

Example:

```go
NewUserService(repo UserRepository)
```

---

# 11. Where Interfaces Should Live (Very Important)

In Go:

> Interfaces usually live **where they are used**, not where they are implemented.

Example:

```
internal/
  user/
    service.go
    repository.go   <-- interface here
    mysql_repo.go   <-- implementation
```

---

# 12. Visual Mental Model

Think of interfaces like **power sockets**.

```
Service (socket)
        ↓
Interface
        ↓
MySQL plug
Mongo plug
Mock plug
```

Anything can plug in **as long as it satisfies the interface**.

---

# 13. What You Should Practice

Build this small project.

```
go-di-example
```

Structure:

```
main.go

user/
  service.go
  repository.go
  mysql.go
  mock.go
```

Goal:

```
UserService -> UserRepository interface -> MySQL implementation
```

Then swap with **mock**.

---

# 14. Beginner Mistakes (Avoid These)

### ❌ Huge Interfaces

```go
type Repository interface {
  CreateUser()
  DeleteUser()
  UpdateUser()
  GetUser()
}
```

### ✅ Better

```go
type UserCreator interface {
   CreateUser()
}
```

Small interfaces are **more powerful**.

---

# 15. The 3 Ideas That Make Go DI Click

### 1️⃣ Interfaces describe behavior

```go
type Writer interface {
    Write([]byte)
}
```

### 2️⃣ Types automatically implement interfaces

No keyword needed.

### 3️⃣ Dependencies are injected

```go
service := UserService{repo}
```

---

If you'd like, I can also show:

- **How Uber, Docker, and Kubernetes structure interfaces**
- **The most important Go interface (`io.Reader`)**
- **A real production Go project architecture**

Those make Go interfaces **crystal clear**.