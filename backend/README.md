# ERP Asset Management - Backend

A Go backend service for asset and maintenance tracking using Gin framework and PostgreSQL.

## Project Structure

```
backend/
├── config/           # Database configuration
├── db/              # SQL initialization scripts
├── handlers/        # HTTP request handlers
├── middleware/      # Authentication middleware
├── models/          # Data models
├── routes/          # Route definitions
├── main.go          # Application entry point
└── go.mod           # Go module dependencies
```

## Setup & Installation

### Prerequisites
- Go 1.26.2+
- PostgreSQL 12+
- Docker (optional, for database container)

### Environment Variables
Create a `.env` file in the backend directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=erp_assets
PORT=8080
```

### Installation

```bash
# Clone and navigate to backend
cd backend

# Install dependencies
go mod download

# Run the application
go run main.go

# Or run tests
go test ./...
```

## API Endpoints

### Authentication
- `POST /api/v1/auth/login` - User login

### Assets
- `GET /api/v1/assets` - List all assets
- `POST /api/v1/assets` - Create asset
- `GET /api/v1/assets/:id` - Get asset details
- `PUT /api/v1/assets/:id` - Update asset
- `DELETE /api/v1/assets/:id` - Delete asset
- `GET /api/v1/assets/:id/depreciation` - Calculate depreciation

### Assignments
- `GET /api/v1/assignments` - List active assignments
- `POST /api/v1/assignments` - Assign asset to employee
- `PUT /api/v1/assignments/:id/return` - Return asset
- `GET /api/v1/assignments/employee/:eid` - Get employee assignments

### Maintenance
- `GET /api/v1/maintenance` - List maintenance records
- `POST /api/v1/maintenance` - Schedule maintenance
- `PUT /api/v1/maintenance/:id/complete` - Mark maintenance complete

## Fixed Issues

### Issue 1: Router Not Wired to Main Server
**Problem**: `main.go` was creating a bare Gin engine instead of using the application router.

**Solution**: Updated `main.go` to call `routes.SetupRouter()` and removed duplicate route registration.

**Impact**: All API routes now properly registered and accessible.

### Issue 2: Empty Handler Files
**Problem**: `handlers/maintenance.go` was empty (0 bytes).

**Solution**: Populated with `GetMaintenance`, `CreateMaintenance`, and `CompleteMaintenance` handlers.

**Impact**: Maintenance endpoints now functional.

### Issue 3: Import Path Placeholder
**Problem**: `handlers/asset.go` used placeholder path `github.com/YOUR_USER/...` instead of actual module path.

**Solution**: Fixed to use `github.com/NirMAN-15/erp-asset-management/backend/...`.

**Impact**: Backend compiles and builds successfully.

## Prevention Best Practices

- Always wire the main router before starting the server
- Run `go test ./...` after any route/handler changes
- Use version control to detect empty file changes
- Implement placeholder checks in CI/CD pipeline
- Add route integration tests for critical endpoints

## Database Setup

The application auto-migrates tables on startup:

```go
config.DB.AutoMigrate(
    &models.Asset{},
    &models.Assignment{},
    &models.Maintenance{},
)
```

Initialize the database with `db/init.sql` if needed.

## Authentication

All routes except `/health` and `/auth/login` require authentication via the `AuthMiddleware`.

Include the `Authorization` header with a valid JWT token:

```
Authorization: Bearer <token>
```

## Development

### Run with Hot Reload
Use `air` or similar tool for development:

```bash
go install github.com/cosmtrek/air@latest
air
```

### Format Code
```bash
go fmt ./...
```

### Lint
```bash
go vet ./...
```

## Troubleshooting

**404 on POST /api/v1/assignments**
- Verify router is initialized with `routes.SetupRouter()`
- Check handler is registered in `routes/router.go`
- Ensure all handler files exist and are not empty

**Compile Error: Expected 'package', found 'EOF'**
- Check for empty handler files (0 bytes)
- Verify all `.go` files have valid package declarations

**Database Connection Error**
- Confirm PostgreSQL is running
- Verify `.env` credentials match database
- Check `DB_HOST` and `DB_PORT` are accessible

## License

All Rights Reserved
