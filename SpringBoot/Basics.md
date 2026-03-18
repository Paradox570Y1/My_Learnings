## **1. Introduction to Spring Boot**

### **1.1 What is Spring Boot?**

- **Spring Boot** is a framework built on top of the Spring Framework to simplify backend development in Java.
    
- It provides:
    
    - **Auto-configuration**: Spring Boot automatically configures beans and dependencies for you.
        
    - **Embedded servers**: Tomcat, Jetty, or Undertow are included—no external deployment required.
        
    - **Starter dependencies**: Quickly add common functionality (like REST, JPA, Security) using pre-configured “starters.”
        

**Why Spring Boot?**

- Eliminates boilerplate configuration.
    
- Promotes **industry best practices**.
    
- Provides a robust foundation for **scalable, maintainable backend systems**.

### **1.2 How Spring Boot Works (Conceptual Flow)**

1. Spring Boot scans your project for classes annotated with `@Component`, `@Service`, `@Repository`, and `@Controller`.
    
2. It auto-configures beans based on included dependencies.
    
3. Embedded server (Tomcat) starts and serves HTTP requests.
    
4. Requests flow through **Controller → Service → Repository → Database**.
    
5. Response is returned to the client.
    

**Diagram of Request Flow:**

```txt
Client ---> Controller ---> Service ---> Repository ---> Database  
```

# File Structure

```txt
com.example.project
 ├─ config            // Configuration classes (DB, Security, Web)
 ├─ controller        // REST API endpoints
 ├─ service           // Business logic interface
 ├─ service.impl      // Implementation of service interface
 ├─ repository        // Data access layer (JPA repositories)
 ├─ model             // JPA Entities
 ├─ dto               // Data Transfer Objects (avoid exposing entities)
 ├─ mapper            // Map DTOs to entities
 ├─ exception         // Custom exceptions
 ├─ advice            // Global exception handling
 └─ utils             // Utility/helper classes
```

**Explanation of Structure:**

- **Controller**: Entry point for HTTP requests. Handles request validation.
    
- **Service & Service.impl**: Business logic separated from controller. Service interface → allows **testability and decoupling**.
    
- **Repository**: Database operations using JPA/Hibernate.
    
- **Model**: Entity classes representing database tables.
    
- **DTO**: Data Transfer Objects → prevent exposing entity internals to API clients.
    
- **Mapper**: Converts DTO ↔ Entity.
    
- **Exception & Advice**: Centralized exception handling.
    
- **Utils**: Helper methods or constants.
    

**Beginner Tip:** This modular structure aligns with **SOLID principles**, which are standard in professional projects.