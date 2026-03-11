## 1. What is a Map?

A **map** in Go is a **hash table** that stores **key–value pairs**.

Conceptually:

key → value

Example:

```go
ages := map[string]int{  
    "Alice": 25,  
    "Bob":   30,  
}
```

Meaning:

"Alice" → 25  
"Bob"   → 30

Key characteristics:

| Property            | Description                |
| ------------------- | -------------------------- |
| Unordered           | Maps do not preserve order |
| Key-value structure | Each key maps to a value   |
| Fast lookups        | Average O(1) lookup time   |
| Reference type      | Internally uses pointers   |
| Dynamic size        | Can grow automatically     |

---

# 2. Map Type Syntax

General form:

```go
map[KeyType]ValueType
```

Examples:

```go
map[string]int  
map[int]string  
map[string]bool  
map[string][]int
```

Example variable:

```go
scores := map[string]int{}
```

---

# 3. Creating Maps

## 3.1 Using Map Literal

Most common way.

```go
m := map[string]int{  
    "apple":  10,  
    "banana": 20,  
}
```

Creates a map with initial values.

---

## 3.2 Using `make()` (Industry Standard)

```go
m := make(map[string]int)
```

Creates an **empty map ready for use**.

You can optionally provide capacity:

```go
m := make(map[string]int, 100)
```

Capacity hint helps reduce **rehashing** in large maps.

Industry use case:

- When expecting **large datasets**
    

---

## 3.3 Nil Map

```go
var m map[string]int
```

This creates a **nil map**.

Properties:

```go
m == nil
```

Reading works:

```go
fmt.Println(m["a"]) // 0
```

But writing fails:

```go
m["a"] = 10 // panic
```

So you must initialize with `make()`.

---

# 4. Adding or Updating Values

Insert or update using assignment:

```go
m := make(map[string]int)  
  
m["apple"] = 10  
m["banana"] = 20
```

Update:

```go
m["apple"] = 50
```

Maps automatically handle insertion and updates.

---

# 5. Accessing Values

Example:

```go
value := m["apple"]
```

If key exists:

value = stored value

If key does NOT exist:

value = zero value of type

Example:

```go
m := map[string]int{}  
  
fmt.Println(m["missing"]) // 0
```

---

# 6. Checking if Key Exists

Use the **comma ok idiom**.

```go
value, ok := m["apple"]
```

Meaning:

value → stored value  
ok    → true if key exists

Example:

v, ok := m["apple"]  
  
```go
if ok {  
    fmt.Println("found:", v)  
} else {  
    fmt.Println("not found")  
}
```

This pattern is used **extensively in industry code**.

---

# 7. Deleting Elements

Use `delete()`.

```go
delete(m, "apple")
```

If the key does not exist → **no error**.

---

# 8. Map Length

Use `len()`.

len(m)

Example:

```go
m := map[string]int{  
    "a": 1,  
    "b": 2,  
}  
fmt.Println(len(m)) // 2
```
  
---

# 9. Iterating Over Maps

Use `range`.

```go
for key, value := range m {  
    fmt.Println(key, value)  
}
```

Example:

```go
m := map[string]int{  
    "a": 1,  
    "b": 2,  
}
```

Possible output:

b 2  
a 1

⚠️ Important:

**Map iteration order is random**.

Go intentionally randomizes iteration to prevent reliance on ordering.

---

# 10. Map of Complex Types

Maps can store **any type as values**.

Example:

### Map of slices

```go
students := map[string][]int{  
    "Alice": {90, 95},  
    "Bob":   {80, 85},  
}
```

### Map of structs

```go
type User struct {  
    Name string  
    Age  int  
}  
  
users := map[int]User{  
    1: {"Alice", 25},  
}
```

---

# 11. Valid Map Key Types

Keys must be **comparable types**.

Allowed:

```txt
string  
int  
float  
bool  
struct (if comparable)  
arrays
```

Not allowed:

```go
slice  
map  
function
```

Example invalid:

```go
map[[]int]string // ❌ compile error
```

Because slices are **not comparable**.

---

# 12. Maps Are Reference Types

Maps behave like references.

Example:

```go
a := map[string]int{"x": 1} 
b := a  
  
b["x"] = 100
``` 
  


Result:

```go
a["x"] = 100
```

Both refer to the **same map**.

---

# 13. Passing Maps to Functions

Maps are passed **by reference-like behavior**.

```go
func update(m map[string]int) {  
    m["x"] = 100  
}
```

Calling:

```go
m := map[string]int{"x": 1}
update(m)
```
  

Result:

```go
m["x"] = 100
```

---

# 14. Map Initialization Pattern (Industry)

Common pattern:
```go
counts := make(map[string]int)  
  
for _, word := range words {  
    counts[word]++  
}
```

This works because:

missing keys return zero value

Used heavily in:

- log processing
    
- analytics
    
- counters
    
- frequency maps
    

---

# 15. Nested Maps

Example:

```go
m := map[string]map[string]int{}
```

Usage:

```go
m["user1"] = map[string]int{  
    "score": 100,  
}
```

Common in:

- configuration
    
- hierarchical data
    

---

# 16. Map Concurrency Warning

Maps are **NOT safe for concurrent writes**.

This will crash:

```go
go func() {  
    m["a"] = 1  
}()  
```
  
```go
go func() {  
    m["b"] = 2  
}()
```

Solution:

Use:

### Mutex

```go
var mu sync.Mutex
```

or

### `sync.Map`

Used in highly concurrent environments.

---

# 17. Performance Characteristics

Maps use **hash tables**.

Average complexity:

|Operation|Complexity|
|---|---|
|Insert|O(1)|
|Lookup|O(1)|
|Delete|O(1)|

Worst case:

O(n)

But extremely rare due to hashing.

---

# 18. When to Use Maps

Use maps when you need:

|Use Case|Example|
|---|---|
|Fast lookups|userID → user|
|Counting|word frequency|
|Indexing|id → object|
|Deduplication|set implementation|
|Caching|key → cached value|

Example: deduplication

seen := map[string]bool{}  
  
```go
for _, v := range list {  
    if !seen[v] {  
        seen[v] = true  
    }  
}
```

```go
// Check if a key is present in the map like this:
	if val, ok := m["one"]; ok {
		// Do something
	}
```

---

# 19. Map vs Slice

|Situation|Use|
|---|---|
|Ordered data|Slice|
|Lookup by key|Map|
|Index-based access|Slice|
|Fast membership check|Map|
|Counting/grouping|Map|

Example:

Slice search:

O(n)

Map lookup:

O(1)

---

# 20. Real Industry Example

Counting API requests:

```go
func countRequests(paths []string) map[string]int {  
  
    counts := make(map[string]int)  
  
    for _, p := range paths {  
        counts[p]++  
    }  
  
    return counts  
}
```

Input:

```cmd
["/home","/login","/home"]
```

Output:

```go
{  
 "/home": 2,  
 "/login": 1  
}
```



# Final Mental Model

Think of a map as:

Map = Hash Table  
Key → Hash → Bucket → Value

Features:

fast lookups  
dynamic size  
unordered  
reference type