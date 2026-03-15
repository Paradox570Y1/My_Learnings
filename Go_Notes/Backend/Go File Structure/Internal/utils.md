teach me from scratch all about database directory (in internal package in Go module) , it's significance, need, usage, responsibilities, typical code needed in that etc. Btw i have to use mysql database.


# What the `utils` Directory Is

The `utils` directory contains **small reusable helper functions** that are:

- used in **multiple parts of the application**
    
- **not tied to business logic**
    
- **not tied to a specific domain**
    

Think of them as **general-purpose tools**.


Example:

```go
internal/  
├── utils/  
│   ├── password.go  
│   ├── jwt.go  
│   ├── response.go  
│   ├── validator.go  
│   └── time.go
```

These files provide **utility helpers used across handlers, services, and middleware**.


# Why `utils` Exists

Without a utils package, developers often repeat the same code everywhere.

Example without utils:

```go
hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
```

This line might appear in **many handlers or services**.

Instead we create:

```go
utils.HashPassword(password)
```

This improves:

- **code reuse**
    
- **readability**
    
- **maintainability**


# Where `utils` Fits in Architecture

Typical Go backend architecture:

`utils` sits **beside everything**, because it’s shared.

```txt
internal/  
├── routes  
├── handlers  
├── services  
├── repository  
├── database  
└── utils
```

It provides **helper tools used by all layers**.


# Responsibilities of `utils`

Good utilities usually include:

### Security helpers

```txt
password hashing  
JWT generation  
token generation
```

### Response helpers

```txt
standard API response format  
error response helpers
```

### Validation helpers

```txt
email validation  
password validation
```

### Time helpers

```txt
formatting timestamps  
duration helpers
```

### Random generators

```txt
IDs  
tokens  
verification codes
```


# What Should NOT Go in Utils

A big mistake is putting **domain logic** here.

❌ Bad example:

utils/createUser.go

User creation belongs in **service layer**, not utils.

Rule:

```txt
If code knows about business rules → NOT utils  
If code is generic helper → utils
```


# Example Utility: Password Hashing

One of the most common utilities.

File:

`internal/utils/password.go`

Example:

```go
package utils  
  
import "golang.org/x/crypto/bcrypt"  
  
func HashPassword(password string) (string, error) {  
	hash, err := bcrypt.GenerateFromPassword(  
		[]byte(password),  
		bcrypt.DefaultCost,  
	)  
  
	return string(hash), err  
}  
  
func CheckPassword(password string, hash string) bool {  
	err := bcrypt.CompareHashAndPassword(  
		[]byte(hash),  
		[]byte(password),  
	)  
  
	return err == nil  
}
```

Now your service can do:

```go
hashed, _ := utils.HashPassword(password)
```


# Example Utility: JSON Response Helper

APIs often repeat response formatting.

Instead of writing this everywhere:

```go
c.JSON(200, gin.H{  
    "status": "success",  
    "data": user,  
})
```

We create a helper.

File:

```txt
internal/utils/response.go
```

Example:

```go
package utils  
  
import "github.com/gin-gonic/gin"  
  
func Success(c *gin.Context, data interface{}) {  
	c.JSON(200, gin.H{  
		"success": true,  
		"data": data,  
	})  
}  
  
func Error(c *gin.Context, message string) {  
	c.JSON(400, gin.H{  
		"success": false,  
		"error": message,  
	})  
}
```
Usage:

```go
utils.Success(c, user)
```

# Example Utility: JWT Generator

File:

```txt
internal/utils/jwt.go
```

Example:

```go
package utils  
  
import (  
	"time"  
	"github.com/golang-jwt/jwt/v5"  
)  
  
var secret = []byte("mysecret")  
  
func GenerateJWT(userID int) (string, error) {  
  
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{  
		"user_id": userID,  
		"exp": time.Now().Add(24 * time.Hour).Unix(),  
	})  
  
	return token.SignedString(secret)  
}
```

Used in authentication service.


# Example Utility: Random String Generator


File:

```txt
internal/utils/random.go
```

Example:

```go
package utils  
  
import (  
	"math/rand"  
	"time"  
)  
  
func RandomString(n int) string {  
  
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"  
  
	rand.Seed(time.Now().UnixNano())  
  
	b := make([]byte, n)  
  
	for i := range b {  
		b[i] = chars[rand.Intn(len(chars))]  
	}  
  
	return string(b)  
}
```

Used for:

```txt
email verification tokens  
reset password tokens  
API keys
```


# Utility for Time Formatting

Example:

```txt
internal/utils/time.go
```

```go
package utils  
  
import "time"  
  
func FormatTime(t time.Time) string {  
	return t.Format("2006-01-02 15:04:05")  
}
```

This ensures **consistent time format across API responses**.


# Typical Utils Directory in Production

A mature Go backend might have:

```txt
internal/utils/  
│  
├── password.go  
├── jwt.go  
├── response.go  
├── validator.go  
├── random.go  
├── pagination.go  
├── time.go  
└── errors.go
```

Each file contains **small focused helpers**.

# Example: Pagination Utility

Pagination is repeated everywhere.

Example helper:

```go
package utils  
  
func GetOffset(page int, limit int) int {  
	return (page - 1) * limit  
}
```

Repository usage:

```go
offset := utils.GetOffset(page, limit)
```

SQL:

```go
SELECT * FROM users LIMIT ? OFFSET ?
```

# Best Practices for Utils

### Keep functions small

Utilities should be **tiny helpers**, not large modules.

---

### Organize by topic

Instead of one giant file:

```txt
utils/helpers.go (2000 lines)
```

Use multiple files.

---

### Avoid dependency loops

Utils should **not import business packages**.

Bad:

```txt
utils importing services
```

---

### Avoid turning utils into a junk drawer

Ask yourself:

Is this generic?  
Will multiple packages use it?

If not, it probably **doesn't belong in utils**.

# Key Takeaway

The `utils` directory exists to store:

```txt
Reusable  
Generic  
Stateless  
Helper functions
```

It improves:

```txt
code reuse  
clean architecture  
readability
```

But it must **never contain business logic**.