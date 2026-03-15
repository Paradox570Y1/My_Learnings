```go
package database

import (
	"database/sql"
	"log"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB // single global database instance
func Connect() {
    //define dsn
	db, err := sql.Open("mysql", dsn)
	//handle err
	err = db.Ping()
	//handle err
	DB = db
	log.Println("database connected successfully")
}
```

`*sql.DB` is the **database handle (connection manager)** provided by the Go standard package `database/sql`.

It **does not represent a single connection** to the database.  
Instead, it represents a **pool of connections** that Go manages automatically.

So think of it as:

```txt
Application
     │
     ▼
 *sql.DB  (connection pool manager)
     │
 ┌───┼───────────┐
 ▼   ▼           ▼
conn1 conn2    conn3
     │
     ▼
   MySQL
```

Your code talks to `*sql.DB`, and it internally uses connections from the pool.


## Where It Comes From

You create it using:

```go
db, err := sql.Open("mysql", dsn)
```

- `"mysql"` → database driver
- `dsn` → connection string

Example DSN

```go
user:password@tcp(localhost:3306)/appdb?parseTime=true
```


## Important: `sql.Open` Does NOT Connect

This surprises many developers.

```go
db, err := sql.Open(...)
```

This only:

- validates arguments
- prepares the connection pool

It **does not establish a database connection immediately**.

To confirm the database is reachable, you use:

```go
db.Ping()
```

## What `*sql.DB` Manages

### 1️⃣ Connection Pool

Go automatically manages multiple database connections.

Example pool:

```txt
max open connections = 10
```

Your application may execute queries simultaneously using different connections.

## Connection Pool Settings (Very Important)

Production Go services **always configure the pool**.

Example:

```txt
db.SetMaxOpenConns(25)  
db.SetMaxIdleConns(10)  
db.SetConnMaxLifetime(time.Hour)
```

## Why Go Uses a Pool

Opening database connections is **expensive**.

Without pooling:

```txt
Request → open DB connection → query → close connection
```

This would be slow.

With pooling:

```txt
Request → reuse existing connection → query
```

Much faster.