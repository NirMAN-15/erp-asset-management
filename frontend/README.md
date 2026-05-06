# ERP Asset Management - Frontend

A Vue 3 web application for managing assets, assignments, and maintenance records.

## Project Structure

```
frontend/
├── src/
│   ├── api/              # API client utilities
│   ├── components/       # Vue components
│   ├── router/           # Vue Router configuration
│   ├── stores/           # Pinia store (state management)
│   ├── views/            # Page components
│   ├── App.vue           # Root component
│   └── main.js           # Application entry point
├── public/               # Static assets
├── index.html            # HTML template
├── vite.config.js        # Vite configuration
└── package.json          # Dependencies
```

## Setup & Installation

### Prerequisites
- Node.js 16+
- npm or yarn

### Installation

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

## Development Server

The application runs on `http://localhost:5173` (or next available port if 5173 is in use).

- Hot Module Replacement (HMR) enabled for instant updates
- Vue DevTools available at `http://localhost:5173/__devtools__/`

## Features

### Pages
- **Login** - User authentication page
- **Asset List** - Browse and manage assets
- **Asset Detail** - View and edit individual asset information
- **Assignments** - Track asset-to-employee assignments
- **Maintenance** - Schedule and complete maintenance tasks

### Authentication
- JWT-based authentication
- Auto-redirect to login on unauthorized access
- Token stored in localStorage
- Protected routes via `meta.auth` guard

### API Integration
- Centralized API client in `src/api/`
- Axios-based HTTP requests
- Automatic bearer token injection
- Error handling and response validation

## Fixed Issues

### Issue: Views Folder Naming Error
**Problem**: Frontend folder was named `src/Veiw/` (typo) instead of `src/views/`.

**Error**: Vite couldn't resolve imports like `../views/Login.vue` because the folder had a misspelling.

**Solution**: Renamed `src/Veiw/` → `src/views/`.

**Files Affected**:
- `src/views/AssetDetail.vue`
- `src/views/AssetList.vue`
- `src/views/Assignments.vue`
- `src/views/Login.vue`
- `src/views/Maintenance.vue`

**Impact**: All Vue component imports now resolve correctly and development server can start.

## Prevention Best Practices

- Use consistent naming conventions (lowercase for directories)
- Enable spell-check in IDE/editor for folder names
- Use TypeScript or ESLint to catch import errors early
- Run `npm run dev` after significant structural changes
- Use version control to catch file/folder renames

## Building for Production

```bash
# Build optimized production bundle
npm run build

# Output goes to dist/ directory
# Ready to deploy to web server
```

## Environment Configuration

Create a `.env.local` file for development environment variables:

```env
VITE_API_BASE_URL=http://localhost:8080/api/v1
```

## Routing

Routes are defined in `src/router/index.js`:

```javascript
const routes = [
  { path: '/', redirect: '/assets' },
  { path: '/login', component: () => import('../views/Login.vue') },
  { path: '/assets', component: () => import('../views/AssetList.vue'), meta: {auth: true} },
  { path: '/assets/:id', component: () => import('../views/AssetDetail.vue'), meta: {auth: true} },
  { path: '/assignments', component: () => import('../views/Assignments.vue'), meta: {auth: true} },
  { path: '/maintenance', component: () => import('../views/Maintenance.vue'), meta: {auth: true} }
]
```

Routes with `meta.auth: true` require user authentication.

## State Management

Global state managed with Pinia in `src/stores/`:

- **User Store** - Authentication state and user data
- **Asset Store** - Asset list and details
- **Assignment Store** - Assignment records

## Component Development

### Create a new view:

```bash
touch src/views/MyPage.vue
```

### Add route:

```javascript
{ path: '/mypage', component: () => import('../views/MyPage.vue'), meta: {auth: true} }
```

## IDE Setup

Recommended: [VS Code](https://code.visualstudio.com/) + [Vue (Official)](https://marketplace.visualstudio.com/items?itemName=Vue.volar)

## Troubleshooting

**Import Error: Failed to resolve "../views/..."**
- Check folder name is `src/views/` (lowercase, correct spelling)
- Verify file exists in `src/views/` directory
- Clear node_modules and reinstall: `rm -rf node_modules && npm install`

**Development server won't start**
- Check port 5173 is not in use
- Verify Node.js version is 16+
- Run `npm install` to ensure all dependencies are installed

**Authentication issues**
- Check backend server is running on `http://localhost:8080`
- Verify API base URL matches backend in environment config
- Check browser localStorage for auth token

## License

All Rights Reserved
