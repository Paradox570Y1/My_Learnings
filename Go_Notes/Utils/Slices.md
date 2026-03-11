## 1. What is a Slice?

A **slice** is a **dynamic, flexible view over an underlying array**.

Unlike arrays:

- **Arrays → fixed size**
    
- **Slices → dynamic size**
    

Slice type syntax:

```go
[]T
```

Example:

```go
numbers := []int{1, 2, 3}
```

Here:

- `[]int` → slice of integers
    

Key properties:

|Property|Meaning|
|---|---|
|Dynamic size|Can grow and shrink|
|Reference type|Points to an underlying array|
|Most used collection type in Go|Used everywhere in production|

---

# 2. Slice Internal Structure

A slice internally contains **three things**:

```cmd
Slice  
 ├── Pointer → underlying array  
 ├── Length (len)  
 └── Capacity (cap)
```

Example:

```go
arr := [5]int{10,20,30,40,50}  
s := arr[1:4]
```

Memory view:

```txt
Array:  [10 20 30 40 50]  
Index:   0  1  2  3  4  
```

```go
Slice s = arr[1:4]  
```
  
```txt
Pointer → index 1  
Length  = 3  (20,30,40)  
Capacity = 4 (20,30,40,50)
```

---

# 3. Creating Slices

## 3.1 Slice Literal

Most common in industry.

```go
nums := []int{1,2,3,4}
```

---

## 3.2 From an Array

```go
arr := [5]int{1,2,3,4,5}  
  
s1 := arr[1:4] // [2 3 4]  
s2 := arr[:]   // entire array
```

Syntax:

```cmd
slice[start:end]
```

Rules:

```txt
length = end - start  
capacity = len(array) - start
```

---

## 3.3 Using `make()` (Industry Standard)

Used when you want **control over length and capacity**.

```go
s := make([]int, 5)
```

Creates:

```txt
[0 0 0 0 0]  
len = 5  
cap = 5
```

With capacity:

```go
s := make([]int, 5, 10)
```

len = 5  
cap = 10

Industry use case:

- When building **large slices efficiently**
    

---

# 4. Length vs Capacity

## Length

Number of elements currently in slice.

len(slice)

## Capacity

Maximum size slice can grow **without reallocating memory**.

```txt
cap(slice)
```

Example:

```go
s := make([]int, 3, 5)  
fmt.Println(len(s)) // 3  
fmt.Println(cap(s)) // 5
```
  

---

# 5. Appending Elements

Slices grow using `append()`.

```go
s := []int{1,2,3}  
  
s = append(s, 4)
```

Result:

[1 2 3 4]

Append multiple values:

```go
s = append(s, 5,6,7)
```

Append another slice:

```go
a := []int{1,2}  
b := []int{3,4}  
a = append(a, b...)
```


Result:

[1 2 3 4]

---

# 6. What Happens When Capacity Exceeds?

If slice capacity is exceeded:

1. Go allocates a **new larger array**
    
2. Copies old elements
    
3. Returns new slice
    

Example:

```go
s := []int{1,2,3}  
s = append(s,4)
```
  

If capacity exceeded → **reallocation happens**.

Industry implication:

Large repeated appends without capacity planning can cause **performance overhead**.

---

# 7. Slices Share Memory

Slices **reference the same underlying array**.

Example:

```go
arr := [5]int{1,2,3,4,5}
a := arr[1:4]  
b := arr[2:5]  
  
b[0] = 100  
  
fmt.Println(a)
```

Output:

[2 100 4]

Why?

Both slices point to the **same array**.

---

# 8. Copying Slices

To avoid shared memory issues:

Use `copy()`.

```go
a := []int{1,2,3}  
  
b := make([]int, len(a))  
  
copy(b, a)
```

Now `a` and `b` are **independent**.

Industry rule:

> Use `copy()` when you need safe duplication.

---

# 9. Nil Slice vs Empty Slice

Important in production APIs.

### Nil slice

```go
var s []int
```

Properties:

```cmd
len = 0  
cap = 0  
s == nil
```

### Empty slice

```go
s := []int{}
```

Properties:

```cmd
len = 0  
cap = 0  
s != nil
```

Industry difference:

|Case|Use|
|---|---|
|nil slice|default state|
|empty slice|returning JSON responses|

Example:

```go
return []string{}
```

instead of:

```go
null
```

---

# 10. Slicing Operations

You can slice a slice.

```go
s := []int{1,2,3,4,5}  
  
a := s[:3]  
b := s[2:]
```

Result:

a = [1 2 3]  
b = [3 4 5]

---

# 11. Reslicing

Slices can grow **within capacity**.

Example:

```go
s := []int{1,2,3,4,5}  
  
a := s[:2]  
a = a[:4]
```

Result:

[1 2 3 4]

Because the capacity allowed it.

---

# 12. Passing Slices to Functions

Slices behave like **reference types**.

```go
func modify(s []int) {  
    s[0] = 100  
}  

a := []int{1,2,3}  
  
modify(a)
```


Result:

[100 2 3]

Because both reference same array.

---

# 13. Deleting Elements

Go has no built-in delete for slices.

Common pattern:

```go
s := []int{1,2,3,4}  
  
index := 1  
  
s = append(s[:index], s[index+1:]...)
```

Result:

[1 3 4]

---

# 14. Industry Best Practices

## 1. Prefer slices over arrays

Arrays are rarely used directly.

Use:
```go
[]T
```

---

## 2. Preallocate capacity

Bad:

```go
var s []int  
for i := 0; i < 10000; i++ {  
    s = append(s, i)  
}
```

Better:

```go
s := make([]int, 0, 10000)
```

---

## 3. Avoid unintended memory sharing

Copy slices if needed.

---

## 4. Return empty slices instead of nil in APIs

Better JSON output.

---

## 5. Be careful with large backing arrays

Example problem:

small := bigSlice[:10]

Even though slice is small, **entire large array stays in memory**.

Fix:

newSlice := append([]int(nil), small...)

---

# 15. When to Use What

|Situation|Use|
|---|---|
|Fixed memory block|Array|
|Dynamic collection|Slice|
|Unknown size input|Slice|
|Performance-sensitive bulk building|`make(..., capacity)`|
|Need independent copy|`copy()`|
|Concatenate collections|`append()`|

---

# 16. Real Industry Example

Processing API results:

func filterEven(nums []int) []int {  
  
    result := make([]int, 0, len(nums))  
  
    for _, n := range nums {  
        if n%2 == 0 {  
            result = append(result, n)  
        }  
    }  
  
    return result  
}

Best practices used:

- preallocated capacity
    
- append pattern
    
- slice return
    

---

# Final Mental Model

Think of a slice as:

Slice = (Pointer + Length + Capacity)

It is **a window over an array that can grow dynamically**.

---