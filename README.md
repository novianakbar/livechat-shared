# LiveChat Shared Entities

This package contains shared database entities that can be used across multiple LiveChat services.

## Purpose

This shared package allows multiple services to reuse the same entity definitions without duplicating code. This is especially useful when creating microservices that need to work with the same database schema.

## Installation

To use this package in your Go project, run:

```bash
go get github.com/novianakbar/livechat-shared
```

## Usage

Import the entities in your Go code:

```go
import "github.com/novianakbar/livechat-shared/entities"

// Example usage
func main() {
    user := entities.User{
        Email:    "user@example.com",
        Name:     "John Doe",
        Role:     "agent",
        IsActive: true,
    }
    
    // Use with GORM
    // db.Create(&user)
}
```

## Entities Included

- **User**: System users (admin, agent, etc.)
- **Department**: Agent departments
- **ChatUser**: Chat users (anonymous or logged-in)
- **ChatSession**: Chat sessions
- **ChatSessionContact**: Contact information for sessions
- **ChatMessage**: Chat messages
- **ChatLog**: Activity logs
- **ChatTag**: Chat tags
- **ChatSessionTag**: Session-tag relationships
- **AgentStatus**: Agent online status
- **ChatAnalytics**: Chat analytics data

## Usage

Add this module to your go.mod:

```bash
go mod edit -require github.com/novianakbar/livechat-shared@latest
go mod tidy
```

Then import in your Go code:

```go
import "github.com/novianakbar/livechat-shared/entities"

// Use entities
user := &entities.User{
    Email: "user@example.com",
    Name:  "John Doe",
    Role:  "agent",
}
```

## Dependencies

- github.com/google/uuid
- gorm.io/gorm

## Versioning

This package follows semantic versioning. Breaking changes to entity structures will result in a major version bump.
