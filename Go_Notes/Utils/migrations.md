In a Go project, database **migrations** are usually stored as **versioned `.sql` files** in a `migrations/` folder and executed by tools like:

- [[golang-migrate]] (`migrate`)
    
- [[Goose]]
    
- [[Atlas]]
    

The most common in Go projects is **golang-migrate**. It expects **two files per migration**:  
`<version>_name.up.sql` and `<version>_name.down.sql`.

---

# 1️⃣ Folder Structure

Example Go module layout:

```txt
your-project/  
 ├── cmd/  
 ├── internal/  
 ├── migrations/  
 │    ├── 0001_init_schema.up.sql  
 │    └── 0001_init_schema.down.sql  
 ├── go.mod
```

---

# Dependencies Required
```txt
migration library
    go get -u github.com/golang-migrate/migrate/v4
mysql driver
	go get -u github.com/go-sql-driver/mysql
file source driver
	go get -u github.com/golang-migrate/migrate/v4/source/file
mysql migration driver
go get -u github.com/golang-migrate/migrate/v4/database/mysql
```

To install above dependencies just need to install these two
```txt
go get -u github.com/golang-migrate/migrate/v4
go get -u github.com/go-sql-driver/mysql
```

> NOTE
> [[database-sql dependency]]

# Architecture
```txt
Migration files
(.sql)
      │
      ▼
File Source Driver
(reads SQL files)
      │
      ▼
Migration Library
(manages versions and execution order)
      │
      ▼
MySQL Migration Driver
(executes migrations safely)
      │
      ▼
MySQL Go Driver
(communicates with database)
      │
      ▼
MySQL Server
(executes SQL)
```

### 1. Migration Library

`github.com/golang-migrate/migrate/v4`

The **main migration engine**.  
It finds migration files, runs them in order, tracks which ones already ran, and supports rollback.  
It stores migration status in the `schema_migrations` table.

---

### 2. MySQL Driver

`github.com/go-sql-driver/mysql`

Allows **Go to communicate with MySQL**.  
It sends SQL queries from Go to the MySQL server and returns results.  
Without it, Go cannot execute queries like `SELECT`, `INSERT`, or `CREATE TABLE`.

---

### 3. File Source Driver

`github.com/golang-migrate/migrate/v4/source/file`

Lets the migration tool **read migration `.sql` files from your local folder**.  
It loads and sorts the migration files (like `0001.up.sql`).  
These files are then passed to the migration engine to execute.

---

### 4. MySQL Migration Driver

`github.com/golang-migrate/migrate/v4/database/mysql`

Connects the migration engine **specifically to MySQL**.  
It runs the migration SQL on the database and manages the `schema_migrations` tracking table.  
It also handles MySQL-specific behavior during migrations.
# Real Example Flow

Suppose you run:

```txt
m.Up()
```

What happens internally:

### Step 1

Migration library starts.

### Step 2

File source driver reads:

```txt
000001_create_users.up.sql
```

### Step 3

Migration library checks:

```txt
schema_migrations
```

### Step 4

MySQL migration driver prepares query.

### Step 5

MySQL Go driver sends query.

### Step 6

MySQL server executes:

```sql
CREATE TABLE users
```

### Step 7

Migration version saved.

```txt
schema_migrations  
version = 1
```