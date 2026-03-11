
# **1️⃣ Purpose**

- `go build` **compiles Go code into an executable** or library.
    
- It **does not install** the binary into `$GOPATH/bin` (use `go install` for that).
    
- Works with **modules or GOPATH projects**.

# **2️⃣ Basic Syntax**

```cmd
go build [build flags] [packages]
```

- `packages` → folder(s) or package path(s) you want to compile.
    
- If no package is given, Go builds the package in the **current directory**.
    

---

# **3️⃣ Common Use Cases**

|Command|Purpose|
|---|---|
|`go build`|Build package in current folder|
|`go build ./cmd/app`|Build a specific folder/package (`main.go`)|
|`go build -o myapp ./cmd/app`|Build and output executable with custom name|
|`go build ./...`|Build **all packages recursively** in the module|
|`go build -v`|Verbose output (shows packages being compiled)|
|`go build -i`|Install dependencies in build cache for faster future builds|

---

# **4️⃣ Important Notes**

1. **Executable requirement:** Only `package main` with `func main()` produces an executable.
    
2. **Module awareness:** Imports must match **`go.mod` module path**.
    
3. **Output:** By default, builds produce an **executable in the current folder** with the folder name.
    
4. **Cross-compilation:** You can set environment variables:
    

```cmd
GOOS=linux GOARCH=amd64 go build ./cmd/app
```

- Builds a Linux AMD64 binary, even if you’re on macOS/Windows.
    

