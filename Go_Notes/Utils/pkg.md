The `pkg` directory is one of the **most misunderstood parts of Go project structure**.  
It exists for a very specific reason in Go’s philosophy of **package visibility and reuse**.

# First Principle: What `pkg` Means in Go

In a Go module, the `pkg` directory contains **reusable packages that are safe to be imported by other projects**.

Think of it as:

```txt
pkg = public library code
```

While:

```txt
internal = private application code
```

So the key idea is **visibility**.

# The Go Visibility Rule

Go has a special rule for the `internal` folder.

Anything inside:
`internal/`

**cannot be imported outside the module**.

Example:
`github.com/myapp/internal/database`

Another project **cannot import it**.

But packages inside:
`pkg/`

**can be imported by other modules**.

Example:
`github.com/myapp/pkg/logger`

Any Go project can import it.


# Why `pkg` Exists

Sometimes code becomes **generic enough to reuse across multiple projects**.

Example:

- logging library
    
- custom middleware
    
- validator utilities
    
- encryption utilities
    
- retry logic
    
- rate limiting
    

Instead of copying code across projects, you place it in `pkg`.


# Real-World Analogy

Think of a project like a **company building tools**.

```txt
internal = tools only used inside company  
pkg = tools sold publicly
```

Example:

```txt
internal/payment_service  
pkg/logger  
pkg/auth
```

The logger could be used by **any project**.


# Responsibilities of the `pkg` Directory

Packages inside `pkg` should:

✔ be **reusable**  
✔ be **generic**  
✔ not depend on application business logic

They should **NOT depend on your app’s internal packages**.

Bad example:

```txt
pkg/logger importing internal/services
```

That defeats the purpose of being reusable.


# Example: Logger Package

A very common `pkg` example.

Structure:

```txt
pkg/  
└── logger/  
    └── logger.go
```

Example code:

```go
package logger  
  
import (  
	"log"  
	"os"  
)  
  
var Logger = log.New(os.Stdout, "[APP] ", log.Ldate|log.Ltime)  
  
func Info(msg string) {  
	Logger.Println("INFO:", msg)  
}  
  
func Error(msg string) {  
	Logger.Println("ERROR:", msg)  
}
```

Usage anywhere in project:

```go
import "myapp/pkg/logger"  
  
logger.Info("server started")
```

Or even another project could import it.

# Example: Middleware Package

You might write reusable middleware.

Structure:

```go
pkg/  
└── middleware/  
    └── logging.go
```

Example:

```go
package middleware  
  
import (  
	"log"  
	"time"  
  
	"github.com/gin-gonic/gin"  
)  
  func RequestLogger() gin.HandlerFunc {  
  
	return func(c *gin.Context) {  
  
		start := time.Now()  
  
		c.Next()  
  
		duration := time.Since(start)  
  
		log.Printf("%s %s %v",  
			c.Request.Method,  
			c.Request.URL.Path,  
			duration,  
		)  
	}  
}
```


Usage:

```go
router.Use(middleware.RequestLogger())
```


# Typical Beginner Mistake

People create `pkg` and put **everything inside it**.

Bad structure:

```go
pkg/  
  user_service  
  database  
  routes
```

This defeats the architecture.

Better:

```go
internal/  
  handlers  
  services  
  repository  
  database
```