Writing Clean Architecture in Go requires balancing Uncle Bob's principles with Go's idioms: **explicitness, simplicity, composition, and standard library strength**. Below is a practical, battle-tested guide tailored for Go backend APIs.

---
## 🔑 Core Principles (Adapted for Go)
1. **Dependency Rule**: All dependencies point inward. Outer layers depend on inner layers. Inner layers never import outer layers.
2. **Business Logic is Framework-Agnostic**: No `net/http`, `gorm`, `gin`, or `sql` imports in your domain or use cases.
3. **Interfaces Belong Where They're Used**: Define interfaces in the layer that *consumes* them, implement them in the layer that *provides* them.
4. **Go > Theory**: Don't over-engineer. If a layer adds no value, skip it. Clean Architecture in Go should feel lightweight.

---
## 📁 Recommended Project Structure
```
cmd/
  api/
    main.go            # Entry point, DI wiring, server startup
internal/
  domain/              # Pure business logic, entities, domain interfaces/errors
  usecase/             # Application business rules, orchestrates domain + repos
  delivery/            # HTTP/gRPC handlers, DTOs, routing, request/response mapping
  repository/          # Data access implementations (DB, cache, external APIs)
infrastructure/        # Config, DB connections, logger, DI container, middleware
pkg/                   # ONLY cross-domain shared utilities (e.g., pagination, crypto)
```
> 💡 Avoid `pkg` for business logic. Keep it in `internal`. Use `delivery/http` or `delivery/grpc` if multi-protocol.

---
## 🧱 Layer-by-Layer Guidelines

### 1. `domain/`
- Pure Go structs & interfaces
- **No** external dependencies
- Define repository interfaces here
- Define domain errors (sentinels or custom types)
```go
package domain

import "errors"

var ErrUserNotFound = errors.New("user not found")

type User struct {
    ID    string
    Email string
    Name  string
}

type UserRepository interface {
    FindByID(ctx context.Context, id string) (*User, error)
    Save(ctx context.Context, u *User) error
}
```

### 2. `usecase/`
- Implements business rules
- Accepts interfaces, returns domain types/errors
- Handles transactions, orchestration, cross-entity logic
```go
package usecase

import "context"

type UserService struct {
    repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, id string) (*domain.User, error) {
    return s.repo.FindByID(ctx, id)
}
```

### 3. `delivery/` (HTTP Handlers)
- Framework-specific (`net/http`, `gin`, `echo`, `fiber`)
- Maps requests → DTOs → use cases → responses
- Handles validation, auth, middleware
```go
package http

import (
    "encoding/json"
    "net/http"
    "context"
)

type UserHandler struct {
    svc *usecase.UserService
}

func NewUserHandler(svc *usecase.UserService) *UserHandler {
    return &UserHandler{svc: svc}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    id := r.PathValue("id") // Go 1.22+ stdlib routing

    user, err := h.svc.GetUser(ctx, id)
    if err != nil {
        // map to HTTP status & JSON response
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

### 4. `repository/`
- Implements `domain` interfaces
- Talks to DB, ORM, or external services
- Maps DB rows → domain structs
```go
package postgres

import (
    "context"
    "database/sql"
    "yourapp/internal/domain"
)

type UserRepo struct {
    db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
    return &UserRepo{db: db}
}

func (r *UserRepo) FindByID(ctx context.Context, id string) (*domain.User, error) {
    var u domain.User
    err := r.db.QueryRowContext(ctx, "SELECT id, email, name FROM users WHERE id = $1", id).
        Scan(&u.ID, &u.Email, &u.Name)
    if err == sql.ErrNoRows {
        return nil, domain.ErrUserNotFound
    }
    return &u, err
}
```

---
## 🛠 Go-Specific Best Practices

| Area | Recommendation |
|------|----------------|
| **Dependency Injection** | Constructor injection. Wire in `cmd/api/main.go`. Avoid DI containers unless complex. |
| **Context** | Pass `context.Context` as first arg everywhere. Use timeouts/cancellation. |
| **Error Handling** | Define domain errors in `domain/`. Map to HTTP in `delivery/`. Use `errors.Is`/`As`. |
| **Validation** | Validate DTOs in `delivery/` (e.g., `go-playground/validator`). Enforce business rules in `usecase/`. |
| **Testing** | Table-driven tests. Mock interfaces with `mockgen` or hand-written stubs. Use `httptest`. |
| **Packages** | Keep them cohesive. Avoid god packages. Max ~100-150 files per layer. |
| **ORM** | Prefer `database/sql` + `sqlx` or lightweight mappers. If using GORM/Ent, wrap it in repo layer. |

---
## 🔄 Dependency Wiring Example (`cmd/api/main.go`)
```go
func main() {
    cfg := loadConfig()
    db := initDB(cfg.DatabaseURL)
    logger := initLogger()

    repo := postgres.NewUserRepo(db)
    svc := usecase.NewUserService(repo)
    handler := http.NewUserHandler(svc)

    mux := http.NewServeMux()
    mux.HandleFunc("GET /users/{id}", handler.GetUser)

    srv := &http.Server{
        Addr:    ":" + cfg.Port,
        Handler: loggerMiddleware(mux, logger),
    }

    log.Println("Server starting on", srv.Addr)
    log.Fatal(srv.ListenAndServe())
}
```

---
## 🧪 Testing Strategy
| Layer | Test Type | Tools |
|-------|-----------|-------|
| `usecase/` | Unit tests | Mock repositories, `context.Background()` |
| `repository/` | Integration tests | `testcontainers`, `dockertest`, real DB |
| `delivery/` | HTTP tests | `net/http/httptest`, table-driven requests/responses |
| `domain/` | Rarely needed | Only if complex value objects or rules |

Example mock generation:
```bash
go install github.com/golang/mock/mockgen@latest
mockgen -source=internal/domain/user.go -destination=internal/mocks/user_repo.go
```

---
## ⚠️ Common Pitfalls in Go
1. **Over-abstracting**: Creating interfaces for every struct. Only interface where you need to swap implementations or test.
2. **Leaking framework types**: Importing `gin.Context`, `gorm.Model`, or `echo.Echo` in `domain/` or `usecase/`.
3. **Circular dependencies**: Usually means your architecture violates the dependency rule. Extract interfaces.
4. **Ignoring `context`**: Leads to goroutine leaks, hanging DB queries, and ungraceful shutdowns.
5. **Mixing DTOs & Entities**: They serve different purposes. Map explicitly in `delivery/` and `repository/`.
6. **Global state/singletons**: Breaks testability. Pass dependencies explicitly.

---
## 📦 Recommended Tooling (2026)
- **Linting**: `golangci-lint` (enable `revive`, `staticcheck`, `errcheck`, `contextcheck`)
- **Mocking**: `go.uber.org/mock` (successor to `gomock`) or hand-written mocks
- **Validation**: `github.com/go-playground/validator/v10`
- **Routing**: Go 1.22+ `net/http` pattern matching, or `chi` if advanced middleware needed
- **Config**: `github.com/kelseyhightower/envconfig` or `github.com/caarlos0/env/v6`
- **Logging**: `log/slog` (standard since Go 1.21)
- **DI (optional)**: `github.com/google/wire` for compile-time DI, but manual is often enough

---
## ✅ Final Checklist
- [ ] Domain layer has zero external imports
- [ ] Interfaces defined in consumer layers, implemented in producer layers
- [ ] `context.Context` passed through all calls
- [ ] Errors are domain-specific, mapped at boundaries
- [ ] DTOs ≠ Domain Entities ≠ DB Models
- [ ] Tests mock interfaces, not implementations
- [ ] `main.go` only wires & starts, zero business logic
- [ ] CI runs `go vet`, `golangci-lint`, and coverage thresholds

---
## 💡 Pro Tip
Start simple. For a CRUD API, you might only need `delivery/` + `repository/` + `domain/`. Add `usecase/` when you have cross-repo transactions, caching, or complex business rules. Clean Architecture is a **spectrum**, not a rigid template. Go rewards simplicity; let your architecture grow only when complexity demands it.

Need a specific example (e.g., transactions, pagination, gRPC, auth middleware)? Tell me your use case and I'll provide a tailored snippet.