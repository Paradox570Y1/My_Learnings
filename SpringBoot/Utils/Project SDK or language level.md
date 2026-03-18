### 1. **Project SDK (Software Development Kit)**

- **Definition:** The Project SDK is the set of tools and libraries that your project uses to compile, run, and debug code.
    
- **What it includes:**
    
    - Compiler (e.g., `javac` for Java)
        
    - Runtime environment (e.g., **JDK** for Java projects)
        
    - Standard libraries (collections, IO, networking, etc.)
        
    - Sometimes additional tools for building, testing, and documentation
        
- **Example:**
    
    - If you set your Project SDK to **Java 17**, the project will use **JDK 17** to compile and run code.
        
- **Why it matters:**
    
    - Using the wrong SDK can cause compilation errors or runtime issues because newer or older SDKs might support different APIs or features.
        

---

### 2. **Language Level**

- **Definition:** Language level specifies which features of a programming language your project is allowed to use.
    
- **Example (Java):**
    
    - **Language level 8** → You can use lambda expressions, streams, and default interface methods.
        
    - **Language level 17** → You can use sealed classes, records, pattern matching, and other new features introduced in Java 17.
        
- **Important distinction:**
    
    - The **SDK** provides the tools and libraries.
        
    - The **language level** controls which syntax/features the compiler will allow.
        
- **Example scenario:**
    
    - You might have **Project SDK = Java 17**, but set **language level = 11** to ensure the code can still run on Java 11 environments.

✅ **In short:**

- **Project SDK = the actual JDK/tools your project uses**
    
- **Language level = the features of the language you are allowed to write in**