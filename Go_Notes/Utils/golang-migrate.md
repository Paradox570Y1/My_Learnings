
**golang-migrate** is an open-source database migration tool and Go (Golang) library for managing schema changes across multiple environments and databases. It provides both a command-line interface and a programmable API, enabling developers to apply, roll back, and version database schema updates reliably. It is one of the most widely adopted migration tools in the Go ecosystem.

### Key facts

- **Initial release:** Forked from _mattes/migrate_ (2016)
    
- **Latest stable version:** v4.x (2025)
    
- **Supported databases:** PostgreSQL, MySQL/MariaDB, SQLite, MongoDB, SQL Server, CockroachDB, and others
    
- **License:** MIT
    
- **Source repository:** github.com/golang-migrate/migrate
    

### Functionality

golang-migrate organizes database changes into sequential “up” and “down” migration files (for applying and reverting changes). These files can be sourced from local filesystems, Git repositories, or cloud storage such as AWS S3 or Google Cloud Storage. The tool applies migrations in order, tracks version history in the target database, and ensures changes are reproducible across environments.

### CLI and library usage

The CLI provides simple commands for creating, applying, and rolling back migrations:

```cmd
migrate -source file://migrations -database postgres://localhost:5432/dbname up
```

The Go library allows integration directly into applications:

```go
m, _ := migrate.New("file://migrations", "postgres://localhost:5432/dbname")  
m.Up()
```

It supports graceful shutdowns, stable APIs (v3–v4), and thread-safe operation. Docker images are also available for containerized workflows.

### Design principles

The project emphasizes simplicity and reliability through “dumb drivers” — each database driver performs minimal logic, while the core library orchestrates the process. Migrations fail fast on errors to prevent partial or inconsistent updates. This design prioritizes safety over automatic recovery.

### Ecosystem and integrations

golang-migrate integrates with CI/CD pipelines and complements schema management tools such as Atlas, which can automatically plan migrations for golang-migrate users. It remains framework-agnostic, making it suitable for Go and non-Go projects alike.