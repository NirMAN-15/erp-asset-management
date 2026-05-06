# ERP Asset Management System

A full-stack application for managing company assets, employee assignments, and maintenance scheduling.

## Overview

This system consists of two main components:

- **Backend**: Go REST API with Gin framework and PostgreSQL database
- **Frontend**: Vue 3 single-page application with Vite bundler

## Quick Start

### Prerequisites
- Go 1.26.2+
- Node.js 16+
- PostgreSQL 12+
- Docker (optional)

### Backend Setup

```bash
cd backend

# Create .env file
cat > .env << EOF
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=erp_assets
PORT=8080
EOF

# Start PostgreSQL (if using Docker)
docker run --name asset-db -e POSTGRES_PASSWORD=your_password -p 5432:5432 -d postgres:15

# Run the backend
go run main.go
```

Backend runs on: `http://localhost:8080`

### Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

Frontend runs on: `http://localhost:5173`

## Project Structure

```
erp-asset-management/
├── backend/                 # Go REST API
│   ├── config/             # Database config
│   ├── db/                 # SQL scripts
│   ├── handlers/           # HTTP handlers
│   ├── middleware/         # Auth middleware
│   ├── models/             # Data models
│   ├── routes/             # Route definitions
│   ├── main.go
│   ├── go.mod
│   └── README.md           # Backend docs
│
└── frontend/               # Vue 3 SPA
    ├── src/
    │   ├── api/           # API client
    │   ├── components/    # Vue components
    │   ├── router/        # Vue Router config
    │   ├── stores/        # Pinia stores
    │   ├── views/         # Page components
    │   ├── App.vue
    │   └── main.js
    ├── package.json
    ├── vite.config.js
    └── README.md          # Frontend docs
```

## Core Features

### Asset Management
- Create, read, update, delete assets
- Track asset status (ACTIVE, ASSIGNED, RETIRED)
- Calculate depreciation over time
- Filter by status and category

### Asset Assignment
- Assign assets to employees
- Track assignment history
- Return assets from employees
- View employee's current assignments

### Maintenance Tracking
- Schedule maintenance tasks
- Track maintenance history
- Mark maintenance as completed
- Calculate maintenance costs

### Authentication
- JWT-based user authentication
- Protected API endpoints
- Secure session management

## API Endpoints

See [backend/README.md](backend/README.md#api-endpoints) for complete API documentation.

### Summary
- **Auth**: `POST /api/v1/auth/login`
- **Assets**: `GET/POST/PUT/DELETE /api/v1/assets`
- **Assignments**: `GET/POST/PUT /api/v1/assignments`
- **Maintenance**: `GET/POST/PUT /api/v1/maintenance`

## Fixed Issues & Solutions

### Issue 1: POST /api/v1/assignments Returns 404
**Root Cause**: Router not wired to main server. `main.go` created bare Gin engine instead of using `routes.SetupRouter()`.

**Fix**: Updated `main.go` to call `routes.SetupRouter()` and removed duplicate route registration.

**Lessons**: Always wire main entry point to route handler before starting server.

### Issue 2: Empty Handler Files
**Root Cause**: `handlers/maintenance.go` was empty (0 bytes).

**Fix**: Populated with required handler functions.

**Lessons**: Use version control and run compile tests to catch empty files.

### Issue 3: Import Path Placeholders
**Root Cause**: `handlers/asset.go` used `github.com/YOUR_USER/...` instead of actual module path.

**Fix**: Updated all imports to `github.com/NirMAN-15/erp-asset-management/backend/...`.

**Lessons**: Avoid placeholder paths in imports; catch them early in CI/CD.

### Issue 4: Views Folder Typo
**Root Cause**: Frontend folder named `src/Veiw/` instead of `src/views/`.

**Fix**: Renamed folder to correct spelling.

**Lessons**: Use consistent naming (lowercase) and enable IDE spell-check.

## Development Workflow

### Backend Development

```bash
cd backend

# Run tests
go test ./...

# Run with hot reload
go install github.com/cosmtrek/air@latest
air

# Build
go build -o erp-backend main.go
```

### Frontend Development

```bash
cd frontend

# Run dev server with HMR
npm run dev

# Build for production
npm run build

# Preview build
npm run preview
```

## Database

### Auto-Migration
Tables are automatically created on startup:
- `assets` - Asset inventory
- `assignments` - Asset-to-employee assignments
- `maintenance` - Maintenance records

### Manual Setup
If needed, run:
```bash
psql -U postgres -d erp_assets -f backend/db/init.sql
```

## Deployment

### Backend
```bash
cd backend
go build -o erp-backend main.go

# Run
./erp-backend
```

### Frontend
```bash
cd frontend
npm run build

# Deploy dist/ folder to web server
# (Nginx, Apache, Vercel, etc.)
```

## Environment Configuration

### Backend (.env)
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=erp_assets
PORT=8080
```

### Frontend (.env.local)
```env
VITE_API_BASE_URL=http://localhost:8080/api/v1
```

## Troubleshooting

### Backend Issues

**404 on POST /api/v1/assignments**
- Verify router initialization in `main.go`
- Check handler registration in `routes/router.go`
- Ensure handler file exists and is not empty

**Compile Error: "expected 'package', found 'EOF'"**
- Check for empty `.go` files (0 bytes)
- Run `go test ./...` to catch issues

**Database Connection Error**
- Verify PostgreSQL is running
- Check `.env` credentials
- Confirm database exists

### Frontend Issues

**Import Error: Failed to resolve "../views/...""**
- Check folder is `src/views/` (lowercase, no typos)
- Verify file exists in correct location
- Clear cache: `rm -rf node_modules && npm install`

**Development server won't start**
- Check Node.js version is 16+
- Verify port 5173 is available
- Run `npm install`

**API calls return 401 Unauthorized**
- Check backend is running on port 8080
- Verify authentication token in browser localStorage
- Check `VITE_API_BASE_URL` in environment config

## Performance Tips

- Backend: Use connection pooling for database
- Frontend: Enable HTTP/2, gzip compression on web server
- Both: Implement caching strategies
- Database: Add indices on frequently queried columns

## Security Best Practices

- Keep dependencies updated: `go mod tidy`, `npm audit`
- Use HTTPS in production
- Validate all user inputs
- Implement rate limiting on API
- Use environment variables for sensitive data
- Rotate JWT secrets regularly

## License

All Rights Reserved

## Support

For issues or questions, refer to:
- [Backend README](backend/README.md)
- [Frontend README](frontend/README.md)
- Project documentation files
