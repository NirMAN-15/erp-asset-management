# ERP Asset Management System - Comprehensive Project Explanation

## Introduction

The ERP Asset Management System is a full-stack web application designed to manage company assets, track employee assignments, and schedule maintenance activities. This system provides a complete solution for organizations to digitize their asset management processes, ensuring better tracking, utilization, and maintenance of valuable resources.

The project is built using modern web technologies with a decoupled architecture, separating the frontend user interface from the backend business logic and database operations.

## Technologies Used

### Backend Technologies
- **Go (Golang)**: A statically typed, compiled programming language known for its performance, concurrency support, and simplicity. Used for building the REST API server.
- **Gin Framework**: A high-performance HTTP web framework for Go, used for routing, middleware, and handling HTTP requests.
- **GORM**: An Object-Relational Mapping (ORM) library for Go, providing database operations without writing raw SQL.
- **PostgreSQL**: An advanced open-source relational database system used for data persistence.

### Frontend Technologies
- **Vue.js 3**: A progressive JavaScript framework using the Composition API for building user interfaces.
- **Vite**: A fast build tool and development server that provides instant hot module replacement (HMR).
- **Pinia**: The official state management library for Vue.js, used for managing application state.
- **Axios**: A promise-based HTTP client for making API requests to the backend.

### Development Tools
- **Docker**: Used for containerizing the PostgreSQL database.
- **npm/Node.js**: Package manager and runtime for frontend dependencies.
- **Git**: Version control system.

## Project Structure

```
erp-asset-management/
├── backend/                    # Go backend application
│   ├── config/                 # Database configuration
│   ├── db/                     # Database initialization scripts
│   ├── handlers/               # HTTP request handlers (controllers)
│   ├── middleware/             # Authentication and other middleware
│   ├── models/                 # Data models (structs)
│   ├── routes/                 # API route definitions
│   ├── main.go                 # Application entry point
│   └── go.mod                  # Go module dependencies
├── frontend/                   # Vue.js frontend application
│   ├── src/
│   │   ├── api/                # API client utilities
│   │   ├── router/             # Vue Router configuration
│   │   ├── stores/             # Pinia state management
│   │   ├── views/              # Page components
│   │   ├── App.vue             # Root Vue component
│   │   └── main.js             # Application bootstrap
│   ├── public/                 # Static assets
│   ├── package.json            # Node.js dependencies
│   └── vite.config.js          # Vite configuration
└── README.md                   # Project documentation
```

## Backend Detailed Explanation

### Language: Go (Golang)

Go is chosen for the backend due to its:
- High performance and low memory footprint
- Excellent concurrency support with goroutines
- Strong standard library
- Easy deployment (single binary)
- Type safety and compilation-time error checking

### Main Components

#### 1. main.go - Application Entry Point
```go
package main

import (
    "log"
    "os"
    "github.com/NirMAN-15/erp-asset-management/backend/config"
    "github.com/NirMAN-15/erp-asset-management/backend/models"
    "github.com/NirMAN-15/erp-asset-management/backend/routes"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system env")
    }
    
    // Connect to database
    config.ConnectDatabase()
    
    // Auto-migrate database tables
    config.DB.AutoMigrate(
        &models.Asset{},
        &models.Assignment{},
        &models.Maintenance{},
    )
    log.Println("Tables migrated!")
    
    // Setup routes and start server
    router := routes.SetupRouter()
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    router.Run(":" + port)
}
```

**What it does:**
- Loads environment variables from .env file
- Establishes database connection
- Automatically creates/updates database tables (migration)
- Sets up API routes
- Starts the HTTP server on port 8080

#### 2. Models - Data Structures

**asset.go:**
```go
type Asset struct {
    ID              string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
    AssetCode       string    `gorm:"uniqueIndex;not null" json:"asset_code"`
    Name            string    `gorm:"not null" json:"name"`
    Category        string    `json:"category"`
    SerialNumber    string    `json:"serial_number"`
    PurchaseDate    time.Time `json:"purchase_date"`
    PurchaseValue   float64   `json:"purchase_value"`
    CurrentValue    float64   `json:"current_value"`
    Location        string    `json:"location"`
    Status          string    `gorm:"default:ACTIVE" json:"status"`
    PurchaseOrderID string    `json:"purchase_order_id"`
    Notes           string    `json:"notes"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

func (a *Asset) BeforeCreate(tx *gorm.DB) error {
    a.ID = uuid.NewString()
    return nil
}
```

**What it does:**
- Defines the structure of Asset data
- Uses GORM tags for database mapping
- Uses JSON tags for API serialization
- Auto-generates UUID before creating records

Similar models exist for Assignment and Maintenance.

#### 3. Handlers - Business Logic

**asset.go handlers:**
- `GetAssets()`: Retrieves list of assets with optional filtering
- `CreateAsset()`: Creates new asset records
- `GetAsset()`: Retrieves single asset by ID
- `UpdateAsset()`: Updates existing asset
- `DeleteAsset()`: Soft-deletes asset (marks as DISPOSED)
- `GetDepreciation()`: Calculates asset depreciation

**How it works:**
- Uses Gin context for HTTP request/response
- Uses GORM for database operations
- Returns JSON responses
- Implements RESTful API patterns

#### 4. Routes - API Endpoints

**router.go:**
```go
func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.Use(cors.Default())
    
    api := r.Group("/api/v1")
    api.Use(middleware.AuthMiddleware())
    {
        assets := api.Group("/assets")
        {
            assets.GET("", handlers.GetAssets)
            assets.POST("", handlers.CreateAsset)
            assets.GET("/:id", handlers.GetAsset)
            assets.PUT("/:id", handlers.UpdateAsset)
            assets.DELETE("/:id", handlers.DeleteAsset)
            assets.GET("/:id/depreciation", handlers.GetDepreciation)
        }
        // Similar for assignments and maintenance
    }
    return r
}
```

**What it does:**
- Sets up Gin router
- Defines API versioning (/api/v1)
- Groups related endpoints
- Applies authentication middleware

#### 5. Middleware - Cross-cutting Concerns

**auth.go:**
Implements JWT-based authentication middleware that:
- Validates JWT tokens in request headers
- Protects API endpoints from unauthorized access
- Extracts user information from tokens

#### 6. Config - Database Connection

**database.go:**
```go
func ConnectDatabase() {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), 
        os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
    
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
}
```

**What it does:**
- Reads database credentials from environment variables
- Establishes connection to PostgreSQL using GORM
- Handles connection errors

## Frontend Detailed Explanation

### Language: JavaScript (ES6+) with Vue.js Framework

Vue.js is chosen for the frontend because:
- Reactive data binding
- Component-based architecture
- Easy learning curve
- Excellent performance
- Rich ecosystem

### Main Components

#### 1. main.js - Application Bootstrap
```javascript
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')
```

**What it does:**
- Creates Vue application instance
- Registers Pinia for state management
- Registers Vue Router for navigation
- Mounts app to DOM element with id 'app'

#### 2. App.vue - Root Component
```vue
<template>
  <div v-if="!isLogin" style="display:flex;min-height:100vh">
    <nav style="width:200px;border-right:0.5px solid #ddd;padding:20px 0;flex-shrink:0">
      <div style="padding:0 16px 20px;font-weight:500;font-size:15px">ERP Assets</div>
      <router-link v-for="l in links" :key="l.path" :to="l.path"
        style="display:block;padding:10px 16px;font-size:13px;text-decoration:none;color:inherit"
        active-class="nav-active">{{ l.label }}
      </router-link>
      <div @click="logout" style="padding:10px 16px;font-size:13px;cursor:pointer;color:#a32d2d;margin-top:auto">Logout</div>
    </nav>
    <main style="flex:1;overflow:auto"><router-view/></main>
  </div>
  <router-view v-else/>
</template>
```

**What it does:**
- Provides main application layout
- Shows navigation sidebar for authenticated users
- Handles logout functionality
- Uses Vue Router for page navigation

#### 3. Stores - State Management (Pinia)

**assetStore.js:**
```javascript
export const useAssetStore = defineStore('assets', () => {
  const assets = ref([])
  const loading = ref(false)

  async function fetchAssets(filters = {}) {
    loading.value = true
    try {
      const res = await api.get('/assets', { params: filters })
      assets.value = res.data.data || []
    } finally { loading.value = false }
  }

  async function createAsset(data) {
    const res = await api.post('/assets', data)
    assets.value.unshift(res.data.data)
    return res.data.data
  }

  return { assets, loading, fetchAssets, createAsset, updateAsset }
})
```

**What it does:**
- Manages application state for assets
- Provides reactive data (assets list, loading state)
- Handles API calls for CRUD operations
- Updates local state optimistically

#### 4. Views - Page Components

**AssetList.vue:**
- Displays list of assets in a table
- Shows dashboard statistics (total assets, active assets, etc.)
- Provides filtering by status and category
- Includes modal for adding new assets
- Handles navigation to asset details

**AssetDetail.vue:**
- Shows detailed information for a single asset
- Displays depreciation information
- Allows editing asset details

#### 5. API Client - axios.js

```javascript
import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
  headers: {
    'Content-Type': 'application/json'
  }
})

// Add auth token to requests
api.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

export default api
```

**What it does:**
- Creates configured Axios instance
- Sets base URL for API calls
- Automatically adds JWT token to requests
- Handles authentication headers

#### 6. Router - Navigation

**index.js:**
```javascript
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/login', component: () => import('../views/Login.vue') },
  { path: '/assets', component: () => import('../views/AssetList.vue') },
  { path: '/assets/:id', component: () => import('../views/AssetDetail.vue') },
  { path: '/assignments', component: () => import('../views/Assignments.vue') },
  { path: '/maintenance', component: () => import('../views/Maintenance.vue') }
]

export default createRouter({
  history: createWebHistory(),
  routes
})
```

**What it does:**
- Defines application routes
- Maps URLs to Vue components
- Uses lazy loading for performance
- Enables client-side navigation

## Database

### PostgreSQL Database Schema

The system uses PostgreSQL with three main tables:

#### Assets Table
- Stores asset information (ID, code, name, category, etc.)
- Tracks financial data (purchase value, current value)
- Maintains lifecycle status (ACTIVE, MAINTENANCE, DISPOSED)

#### Assignments Table
- Links assets to employees
- Tracks assignment dates and return dates
- Maintains assignment history

#### Maintenance Table
- Records maintenance activities
- Tracks maintenance schedules and completion
- Stores maintenance costs and notes

### Database Initialization
- Uses GORM auto-migration for table creation
- Includes sample data for demonstration
- Supports foreign key relationships

## How the System Works

### 1. User Authentication Flow
1. User logs in via Login.vue
2. Frontend sends credentials to /auth/login
3. Backend validates credentials and returns JWT token
4. Token stored in localStorage
5. Subsequent requests include token in Authorization header

### 2. Asset Management Flow
1. User navigates to /assets
2. AssetList.vue loads and calls fetchAssets()
3. API request sent to backend /api/v1/assets
4. Backend queries database using GORM
5. Data returned as JSON
6. Frontend displays data in reactive table

### 3. Creating New Assets
1. User clicks "Add Asset" button
2. Modal opens with form
3. User fills form and submits
4. Frontend calls createAsset() from store
5. API POST request to /api/v1/assets
6. Backend validates data and creates record
7. New asset added to local state and UI updates

### 4. Data Flow Architecture
- **Frontend**: Vue components → Pinia stores → Axios API calls
- **Backend**: Gin routes → Handlers → GORM models → Database
- **Database**: PostgreSQL with relational data structure

## Features and Capabilities

### Asset Management
- Create, read, update, delete assets
- Track asset lifecycle (active, maintenance, disposed)
- Calculate depreciation using straight-line method
- Filter and search assets by various criteria
- Track asset location and assignment

### Employee Assignment Tracking
- Assign assets to employees
- Track assignment history
- Record return dates
- Prevent double assignment of assets

### Maintenance Scheduling
- Schedule maintenance activities
- Track maintenance history
- Record maintenance costs
- Update asset status during maintenance

### User Interface Features
- Responsive design
- Real-time data updates
- Dashboard with key metrics
- Modal forms for data entry
- Navigation between different views

### API Features
- RESTful API design
- JWT authentication
- JSON data format
- Error handling and validation
- CORS support for cross-origin requests

### Security Features
- JWT-based authentication
- Protected API endpoints
- Input validation
- SQL injection prevention through ORM

## Conclusion

The ERP Asset Management System demonstrates a modern full-stack application architecture using Go for the backend and Vue.js for the frontend. The system provides comprehensive asset management capabilities with proper separation of concerns, secure authentication, and scalable database design.

Key strengths include:
- **Performance**: Go backend with Gin framework provides high-speed API responses
- **Scalability**: Modular architecture allows easy addition of new features
- **Maintainability**: Clean code structure with separation of concerns
- **User Experience**: Modern Vue.js interface with reactive updates
- **Data Integrity**: Relational database with proper constraints

This project serves as an excellent example of enterprise-grade web application development using contemporary technologies and best practices.</content>
<parameter name="filePath">d:\Data\Project\erp-asset-management\project_full_explanation.md