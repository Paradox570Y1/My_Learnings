# 1️⃣ What is a Package in Go?

A **package** is simply a **collection of Go files that belong together**.

Think of a package as a **folder of related code**.

Example project:

```bash
calculator/  
 ├ main.go  
 ├ math.go  
 └ utils.go
```

If all files start with:

package calculator

then they belong to the **same package**.

### Why packages exist

Packages help organize code into **modules**.

Example:

|Package|Purpose|
|---|---|
|`fmt`|printing|
|`math`|math functions|
|`net/http`|web servers|

Example import:

import "fmt"

---

# 2️⃣ Why Every Go File Starts With `package`

In Go:

> **Every `.go` file must start with a package declaration.**

Example:

package main

or

package mathutils

This tells the compiler:

- which package this file belongs to
    
- how it connects with other files
    

Without this line, the program **won’t compile**.

---

# 3️⃣ Why `main` is a Special Package

In Go there are **two main types of packages**:

|Type|Purpose|
|---|---|
|library packages|reusable code|
|main package|executable program|

The name **`main`** tells Go:

> “This program should run as an application.”

Example executable program:

package main  
  
```go
func main() {  
    println("Hello")  
}
```

When compiled:

go build

Go creates an **executable program**.

---

# 4️⃣ What Happens When You Run a Go Program

When Go sees:

package main

the compiler expects to find this function:

```go
func main()
```
Example:

```go
package main  
  
import "fmt"  
  
func main() {  
    fmt.Println("Hello world")  
}
```

Execution flow:

```txt
program starts  
     ↓  
main() function runs  
     ↓  
program ends
```

---

# 5️⃣ Library Packages vs Executable Packages

## Library package

Example:

```go
package mathutils
```

This code **cannot run directly**.

It must be imported.

Example usage:

```go
import "mathutils"
```

---

## Executable package

Example:

```go
package main
```

This creates a **program you can run**.

---

# 6️⃣ Why Go Separates Libraries and Executables

This separation helps Go:

- organize code cleanly
    
- avoid accidental execution
    
- simplify dependency management
    

Example project layout:

```bash
project/  
 ├ cmd/  
 │   └ app/  
 │       └ main.go  
 ├ internal/  
 │   └ logic.go  
 └ go.mod
```

Here:

- `cmd/app` contains **executable**
    
- `internal` contains **libraries**
    

---

# 7️⃣ What the Comment Means

Your code has this comment:

```go
// A package clause starts every source file.
```

Meaning:

Every Go source file **must begin with a package declaration**.

---

Second comment:

```go
// main is a special name declaring an executable rather than a library.
```

Meaning:

The name **`main`** signals that this code should compile into a **program you can run**.

---

# 8️⃣ Why Go Uses Packages Instead of Classes

Languages like:

- Java
    
- C++
    

organize code using **classes and namespaces**.

Go instead uses:

- packages
    
- structs
    
- functions
    

Example Go style:

```go
package mathutils  
  
func Add(a int, b int) int {  
    return a + b  
}
```

Then used in another file:

```go
import "mathutils"  
  
result := mathutils.Add(2,3)
```

---

# 9️⃣ One Important Rule

All Go files in the **same folder must have the same package name**.

Example (valid):

```bash
folder/  
 ├ a.go   -> package utils  
 ├ b.go   -> package utils
```

Invalid example:

```bash
folder/  
 ├ a.go   -> package utils  
 ├ b.go   -> package main ❌
```

---

# 🔟 Why Go Put Packages First

This design has benefits:

✔ clear project structure  
✔ easy dependency management  
✔ fast compilation  
✔ simple imports

---

# 🧠 Mental Model

Think of Go packages like this:

package = folder of related code

Example:

```bash
fmt.Println()  
│  
│  
package name
```

---

# 📦 Real Go Standard Library Example

Example from Go's standard library:

package fmt

Inside that package exists:

```cmd
fmt/  
 ├ print.go  
 ├ scan.go  
 └ format.go
```

All files belong to **package fmt**.

---

# ⚡ One Important Industry Insight

Most large Go systems structure code like this:

```cmd
project/  
 ├ cmd/         (executables)  
 ├ internal/    (private packages)  
 ├ pkg/         (public packages)  
 └ go.mod
```

This pattern is widely used in large Go systems like:

- Docker
    
- Kubernetes
    

---

# ✅ Final Simple Explanation

This line:

package main

means:

> “This file belongs to the `main` package and will build into a runnable program.”