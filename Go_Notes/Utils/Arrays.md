# **Arrays in Go – Complete Guide**

## **1. What is an Array in Go?**

An **array** in Go is a **fixed-size, homogeneous sequence of elements**. All elements have the same type, and the **size is part of the array’s type**.

- **Fixed size** → cannot grow or shrink
    
- **Homogeneous** → all elements are the same type
    
- **Value type** → arrays are **copied on assignment**
    

**Syntax:**

```go
var a [5]int // array of 5 integers
```

Key points:

- `[5]int` → array of 5 elements of type `int`
    
- The length `5` is part of the type. `[5]int` and `[6]int` are **different types**
    

---

## **2. Declaring Arrays**

### a) Explicit declaration

```go
var numbers [3]int        // [0,0,0] default values  
var names [2]string       // ["",""] default strings
```

### b) Declaration with initialization

```go
var primes = [5]int{2, 3, 5, 7, 11}
```

### c) Implicit size (using `...`)

```go
arr := [...]int{1, 2, 3, 4} // compiler counts size automatically
```

### d) Partially initialized array

```go
arr := [5]int{1, 2} // equivalent to [1, 2, 0, 0, 0]
```

---

## **3. Accessing and Modifying Array Elements**

Arrays are **indexed starting at 0**.

```go
arr := [3]int{10, 20, 30}  
fmt.Println(arr[0]) // 10  
  
arr[1] = 50  
fmt.Println(arr)    // [10, 50, 30]
```

- Accessing out-of-range index → **compile-time error** in static size or runtime panic if dynamic
    

---

## **4. Array Length**

- Use the built-in `len()` function:
    

```go
arr := [4]string{"a", "b", "c", "d"}  
fmt.Println(len(arr)) // 4
```

- Length is **part of the type**, so `[4]int` ≠ `[5]int`
    

---

## **5. Iterating Over Arrays**

### a) Using traditional `for` loop

```go
arr := [3]int{1, 2, 3}  
for i := 0; i < len(arr); i++ {  
    fmt.Println(i, arr[i])  
}
```

### b) Using `range`

```go
for index, value := range arr {  
    fmt.Println(index, value)  
}  
  
// ignore index: _  
for _, value := range arr {  
    fmt.Println(value)  
}
```

---

## **6. Arrays are Value Types**

When you **assign an array to another variable**, a **copy** is made:

```go
a := [3]int{1, 2, 3}  
b := a  
b[0] = 100  
fmt.Println(a) // [1, 2, 3]  
fmt.Println(b) // [100, 2, 3]
```

> ⚠️ Important in industry: **modifying large arrays by value is costly**. Use **slices** if you need dynamic or reference behavior.

---

## **7. Multidimensional Arrays**

Arrays can have **multiple dimensions**, like matrices:

```go
var matrix [2][3]int  
matrix[0] = [3]int{1, 2, 3}  
matrix[1] = [3]int{4, 5, 6}  
  
fmt.Println(matrix[1][2]) // 6
```

- Multidimensional arrays are less common; slices are preferred for dynamic 2D arrays.
    

---

## **8. Array vs Slice**

In **modern Go**, **slices are often preferred** over arrays:

|Feature|Array|Slice|
|---|---|---|
|Size|Fixed at compile-time|Dynamic|
|Type|`[n]T`|`[]T`|
|Value semantics|Copied on assignment|Reference-like behavior|
|Common use|Small, fixed-size data|Most general use cases|

- You can convert an array to a slice:
    

```go
arr := [5]int{1,2,3,4,5}  
slice := arr[:] // slice of the whole array
```

---

## **9. Best Practices (Industry-Oriented)**

1. **Prefer slices over arrays** for most cases. Use arrays only if size is fixed and performance matters.
    
2. **Always initialize arrays** to avoid relying on default zero values unless intentional.
    
3. **Use `range` loops** for clarity and safety.
    
4. **Avoid copying large arrays**; pass pointers or slices instead.
    
5. **Multidimensional arrays** → usually better to use slices of slices for dynamic behavior.
    

---

## **10. Real-World Use Cases**

- Arrays are used when **memory layout is critical**, e.g., in **embedded systems**, **buffered data**, or **image processing**.
    
- For **general-purpose collections**, slices are preferred because of flexibility.
    
- Arrays can be **passed to functions** by value or reference using a pointer:
    

```go
func modify(arr *[3]int) {  
    arr[0] = 100  
}  
  
a := [3]int{1, 2, 3}  
modify(&a)  
fmt.Println(a) // [100, 2, 3]
```

---

Arrays in Go are foundational but usually serve as the **underlying storage for slices** in industry. Understanding their **value semantics** and **fixed size** is key to writing efficient, idiomatic Go code.