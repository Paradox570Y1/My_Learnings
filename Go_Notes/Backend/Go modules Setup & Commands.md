### **Setting Up Go Modules**

1. **Initializing a Go Module**:
    
    To create a Go module, run the following in your project directory:
    
    `go mod init <module-name>`
    
    - This creates a `go.mod` file, which defines your module and its path.
        
2. **go.mod File**:
    
    The `go.mod` file is where all the metadata of your module and its dependencies are stored. The file will look something like this:
	```go
	module github.com/myusername/myproject  
      
    go 1.18  
      
    require (  
        github.com/sirupsen/logrus v1.8.1  
        github.com/spf13/cobra v1.1.1  
    )
	```
    
    - **module**: The module's path. This is typically the repository URL, but it can be any path.
        
    - **go**: The version of Go you're using to build the module.
        
    - **require**: Lists the dependencies and their versions.
        
3. **Adding Dependencies**:
    
    To add a dependency, you can use the `go get` command:
    
    `go get github.com/sirupsen/logrus`
    
    - This will automatically update your `go.mod` file to include the dependency and its version.
        
4. **Download Dependencies**:
    
    To download the dependencies listed in `go.mod`, run:
    
    `go mod tidy`
    
    This command will also remove any unnecessary dependencies that are no longer being used.
    

---

### **Key Commands with Go Modules**

Here are some essential Go module commands you'll use frequently:

1. **go mod init**: Initializes a Go module (creates `go.mod`).
    
2. **go mod tidy**: Cleans up the `go.mod` file and removes unused dependencies & downloads used ones.
    
3. **go get**: Adds or updates a dependency in `go.mod`.
    
4. **go list**: Displays module information, including its dependencies.
    
5. **go mod vendor**: Creates a `vendor/` directory with all dependencies, useful if you need to commit them to source control.(kind of like node modules in node)
    
6. **go build**: Uses the dependencies from the `go.mod` to build the project.