In a Go project, database **migrations** are usually stored as **versioned `.sql` files** in a `migrations/` folder and executed by tools like:

- [[golang-migrate]] (`migrate`)
    
- [[Goose]]
    
- [[Atlas]]
    

The most common in Go projects is **golang-migrate**. It expects **two files per migration**:  
`<version>_name.up.sql` and `<version>_name.down.sql`.

---

# 1️⃣ Folder Structure

Example Go module layout:

your-project/  
 ├── cmd/  
 ├── internal/  
 ├── migrations/  
 │    ├── 0001_init_schema.up.sql  
 │    └── 0001_init_schema.down.sql  
 ├── go.mod

---