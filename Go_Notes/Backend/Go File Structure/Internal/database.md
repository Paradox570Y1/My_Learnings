# First Principle: What is the `database` Directory?

The `database` directory is responsible for **setting up the connection to the database** and making it available to other layers like repositories.

Typical location:

```txt
internal/database
```

Its job is **NOT to run queries directly for business logic**.

Instead, it:

- connects to MySQL
    
- configures connection pooling
    
- manages DB lifecycle
    
- provides a DB instance for repositories


# Where It Fits in Backend Architecture

Typical Go backend architecture:

```txt
cmd/server/main.go  
        │  
        ▼  
   routes layer  
        │  
        ▼  
   handlers layer  
        │  
        ▼  
   services layer  
        │  
        ▼  
   repository layer  
        │  
        ▼  
   database layer  
        │  
        ▼  
      MySQL
```


So:

```txt
database = infrastructure layer
```

It simply **gives access to the DB**.


# Typical Production Folder Structure

Example:

```txt
internal/  
│  
├── database/  
│   ├── mysql.go  
│   ├── migrations.go
│   └── connection.go  
│  
├── repository/  
├── services/  
├── handlers/  
├── routes/
```

The database directory usually contains:

```txt
connection setup  
DB config  
migration logic
```

`migrations.go` - go files can also migrate database it requires dependencies like `golang-migrate` or `goose` or `gorm` migrations but in my case i will be using `.sql` files.



# Why We Separate the Database Layer

Imagine you write DB code directly in handlers.

Bad example:

```go
func GetUsers(c *gin.Context) {  
    db, _ := sql.Open("mysql", "root:pass@/test")  
  
    rows, _ := db.Query("SELECT * FROM users")  
}
```

Problems:

- opens connection every request
    
- hard to test
    
- messy architecture
    
- cannot reuse connection
    

Instead we create **one DB connection for the whole app**.


# Required Go Packages

For MySQL in Go you usually use:

```txt
database/sql  
github.com/go-sql-driver/mysql
```

These are the standard.

Install MySQL driver:

```bash
go get github.com/go-sql-driver/mysql
```


# Basic MySQL Connection Code

Typical file:

```txt
internal/database/mysql.go
```

Example:

```go
package database  
  
import (  
	"database/sql"  
	"fmt"  
	"log"  
  
	_ "github.com/go-sql-driver/mysql"  
)  
  
func Connect() *sql.DB {  
  
	dsn := "user:password@tcp(localhost:3306)/mydb"  
  
	db, err := sql.Open("mysql", dsn)  
	if err != nil {  
		log.Fatal("Database connection failed:", err)  
	}  
  
	err = db.Ping()  
	if err != nil {  
		log.Fatal("Database unreachable:", err)  
	}  
  
	fmt.Println("Database connected")  
  
	return db  
}
```

Explanation:

|Part|Meaning|
|---|---|
|sql.Open|creates DB connection|
|mysql|driver|
|dsn|database connection string|
|db.Ping|verifies connection|

# What is DSN (Database Source Name)?

DSN tells Go **how to connect to MySQL**.

Example:

```txt
user:password@tcp(localhost:3306)/dbname
```

Breakdown:

```go
user        → MySQL username  
password    → MySQL password  
localhost   → DB host  
3306        → MySQL port  
dbname      → database name
```

Example:

```txt
root:root@tcp(127.0.0.1:3306)/ecommerce
```



# Connection Pooling

Go automatically maintains a **pool of connections**.

You should configure it.

Example:

```go
db.SetMaxOpenConns(25)  
db.SetMaxIdleConns(10)  
db.SetConnMaxLifetime(5 * time.Minute)
```

Meaning:

|Setting|Meaning|
|---|---|
|MaxOpenConns|max connections to DB|
|MaxIdleConns|idle connections|
|ConnMaxLifetime|reuse time|

Why needed?

Without this:

```txt
too many connections  
DB overload
```

# Production Example Database File

Example:

```txt
internal/database/mysql.go
```

```go
package database  
  
import (  
	"database/sql"  
	"log"  
	"time"  
  
	_ "github.com/go-sql-driver/mysql"  
)  
  
func Connect() *sql.DB {  
  
	dsn := "root:password@tcp(localhost:3306)/mydb"  
  
	db, err := sql.Open("mysql", dsn)  
	if err != nil {  
		log.Fatal(err)  
	}  
  
	if err := db.Ping(); err != nil {  
		log.Fatal(err)  
	}  
  
	db.SetMaxOpenConns(25)  
	db.SetMaxIdleConns(10)  
	db.SetConnMaxLifetime(time.Hour)  
  
	return db  
}
```


# How `main.go` Uses Database

Example:

```txt
cmd/server/main.go
```

```go
package main  
  
import (  
	"myapp/internal/database"  
	"myapp/internal/routes"  
)  
  
func main() {  
  
	db := database.Connect()  
  
	router := routes.SetupRouter(db)  
  
	router.Run(":8080")  
}
```

Flow:

```txt
main  
 ↓  
database.Connect()  
 ↓  
db instance  
 ↓  
routes  
 ↓  
handlers  
 ↓  
repositories
```


# How Repositories Use the Database

Example:

```txt
internal/repository/user_repository.go
```

```go
type UserRepository struct {  
	db *sql.DB  
}  
  
func NewUserRepository(db *sql.DB) *UserRepository {  
	return &UserRepository{db: db}  
}
```

Query example:

```go
func (r *UserRepository) GetUsers() ([]User, error) {  
  
	rows, err := r.db.Query("SELECT id, name FROM users")  
	if err != nil {  
		return nil, err  
	}  
  
	defer rows.Close()  
  
	var users []User  
  
	for rows.Next() {  
		var u User  
		rows.Scan(&u.ID, &u.Name)  
		users = append(users, u)  
	}  
  
	return users, nil  
}
```

Notice:

```txt
routes → handlers → services → repository → db
```


# Database Directory Responsibilities (Summary)

The database directory should only handle:

```txt
DB connection  
DB configuration  
DB health check  
Providing DB instance
```

Not:

```txt
business logic  
queries  
validation  
routing
```

Queries belong in **repository layer**.


# Advanced Production Setup

In larger systems the database directory may contain:

```txt
internal/database  
│  
├── mysql.go  
├── migrations.go  
├── seed.go  
├── config.go  
└── transaction.go
```

Possible responsibilities:

|File|Purpose|
|---|---|
|mysql.go|DB connection|
|migrations.go|schema migrations|
|seed.go|initial data|
|config.go|DB configuration|
|transaction.go|transaction helpers|



# Health Check Endpoint

Sometimes DB health is exposed.

Example:

```txt
GET /health
```

Handler checks:

```go
err := db.Ping()
```

If DB fails:

health = unhealthy

Used by:

- Kubernetes
    
- load balancers
    
- monitoring


# A Real Production Flow

Request:

```txt
GET /users
```

Flow:

```txt
Client  
  ↓  
Router  
  ↓  
Handler  
  ↓  
Service  
  ↓  
Repository  
  ↓  
Database  
  ↓  
MySQL
```

Database layer only **supplies the DB connection**.


# Common Beginner Mistakes

### Opening DB per request

Bad:

```go
sql.Open inside handler
```

Always open **once at startup**.


### Running queries in database package

Bad:

```txt
internal/database/user_queries.go
```

Queries should live in **repositories**.


### Hardcoding credentials

Bad:

```txt
dsn := "root:password@..."
```

Use environment variables.

Example:

```go
os.Getenv("DB_USER")
```


# Production DSN Example

Example:

```go
user:password@tcp(localhost:3306)/dbname?parseTime=true
```

Important flag:

```go
parseTime=true
```

Allows MySQL `DATETIME` → Go `time.Time`.


# One Important Insight

In production systems, **database performance determines system scalability**.

So this layer is critical for:

```go
connection pooling  
transactions  
timeouts  
migrations
```