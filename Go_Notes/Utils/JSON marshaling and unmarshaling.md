JSON marshaling/unmarshaling in Go is something you'll use **constantly as a backend developer**, especially when building APIs.

Since you're learning **Go for backend development**, let's go step-by-step from basics to real API usage.

---

# 1. First: What JSON Is

JSON (JavaScript Object Notation) is a **text format for data**.

Example JSON:

```json
{
  "name": "Alice",
  "age": 25,
  "email": "alice@example.com"
}
```

Backend servers usually:

- receive JSON from clients *(request body)*
- send JSON in responses

Example API response:

```json
{
  "status": "success",
  "user_id": 42
}
```

Go needs a way to convert between **Go structs and JSON**.

That’s where **marshaling** and **unmarshaling** come in.

---

# 2. Marshaling vs Unmarshaling

## Marshaling

Convert:

```
Go struct → JSON
```

## Unmarshaling

Convert:

```
JSON → Go struct
```

Think of it like **translation between two languages**.

---

# 3. The Package Used

Go provides JSON handling in the standard library.

```go
import "encoding/json"
```

This package handles **JSON encoding and decoding**.

---

# 4. Creating a Go Struct

Suppose we have this JSON:

```json
{
  "name": "Alice",
  "age": 25
}
```

We represent it in Go as:

```go
type User struct {
    Name string
    Age  int
}
```

---

# 5. Marshaling (Go → JSON)

Example program:

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string
    Age  int
}

func main() {

    user := User{
        Name: "Alice",
        Age:  25,
    }

    jsonData, err := json.Marshal(user)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(jsonData))
}
```

Output:

```
{"Name":"Alice","Age":25}
```

---

# 6. What `json.Marshal()` Does

```go
jsonData, err := json.Marshal(user)
```

This function:

1️⃣ Reads the struct  
2️⃣ Converts it into JSON  
3️⃣ Returns bytes

Return type:

```
[]byte
```

That’s why we convert it to string:

```go
string(jsonData)
```

---

# 7. Why the Output Has Capitalized Keys

Output:

```json
{"Name":"Alice","Age":25}
```

Because Go uses the **struct field names as JSON keys**.

To control this, we use **JSON tags**.

---

# 8. JSON Tags (Very Important)

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

Now output becomes:

```json
{"name":"Alice","age":25}
```

These are called **struct tags**.

Format:

```
`json:"fieldname"`
```

---

# 9. Unmarshaling (JSON → Go Struct)

Example JSON:

```json
{"name":"Alice","age":25}
```

Code:

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {

    jsonData := []byte(`{"name":"Alice","age":25}`)

    var user User

    err := json.Unmarshal(jsonData, &user)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(user.Name)
    fmt.Println(user.Age)
}
```

Output:

```
Alice
25
```

---

# 10. Why We Use `&user`

```go
json.Unmarshal(jsonData, &user)
```

We pass a **pointer**.

Reason:

`Unmarshal` needs to **modify the struct**.

If we passed:

```go
json.Unmarshal(jsonData, user)
```

Go would not be able to **write values into it**.

---

# 11. Real Backend Example (API Request)

Example HTTP request body:

```json
{
  "email": "alice@example.com",
  "password": "123456"
}
```

Struct:

```go
type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
```

Handler example:

```go
var req LoginRequest

json.NewDecoder(r.Body).Decode(&req)
```

Now you can access:

```
req.Email
req.Password
```

This is **unmarshaling request JSON**.

---

# 12. Sending JSON Response

Example response struct:

```go
type LoginResponse struct {
    Token string `json:"token"`
}
```

Handler:

```go
resp := LoginResponse{
    Token: "abc123",
}

json.NewEncoder(w).Encode(resp)
```

This automatically does:

```
Go struct → JSON → HTTP response
```

---

# 13. Nested JSON Example

JSON:

```json
{
  "name": "Alice",
  "address": {
    "city": "Delhi",
    "zip": "110001"
  }
}
```

Struct:

```go
type Address struct {
    City string `json:"city"`
    Zip  string `json:"zip"`
}

type User struct {
    Name    string  `json:"name"`
    Address Address `json:"address"`
}
```

Go handles **nested JSON automatically**.

---

# 14. Ignoring Fields

If you don't want a field in JSON:

```go
type User struct {
    Name     string `json:"name"`
    Password string `json:"-"`
}
```

`-` means **ignore this field**.

---

# 15. Optional Fields (`omitempty`)

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
}
```

If Email is empty:

```json
{"name":"Alice"}
```

The field is **omitted**.

---

# 16. The Two Functions You'll Use Most

## Convert struct → JSON

```go
json.Marshal()
```

or

```go
json.NewEncoder()
```

## Convert JSON → struct

```go
json.Unmarshal()
```

or

```go
json.NewDecoder()
```

---

# 17. What Happens Internally

The JSON package uses **reflection**.

It:

1️⃣ Inspects struct fields  
2️⃣ Reads tags  
3️⃣ Converts types to JSON

Example:

```
Age int
```

Becomes:

```json
"age": 25
```

---

# 18. Common Beginner Mistakes

## ❌ Using lowercase struct fields

```go
type User struct {
    name string
}
```

This will **not work**.

Why?

Go only exports **capitalized fields**.

Correct:

```go
Name string
```

---

## ❌ Forgetting pointer in Unmarshal

Wrong:

```go
json.Unmarshal(data, user)
```

Correct:

```go
json.Unmarshal(data, &user)
```

---

# 19. Typical Backend Flow Using JSON

```
Client
  ↓
HTTP Request (JSON)
  ↓
Unmarshal → Go struct
  ↓
Business logic
  ↓
Marshal → JSON
  ↓
HTTP Response
```

This happens in **every API endpoint**.

---

# What You Should Learn Next

Since you're learning **Go backend architecture**, the next topics that connect perfectly are:

- DTOs in Go APIs *(very important for clean architecture)*
- Struct tags deep dive
- How Go JSON decoding works with interfaces
- Handling unknown JSON fields
- Performance tricks used in production APIs