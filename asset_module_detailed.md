# 📊 DETAILED ASSET MODULE OVERVIEW

## **1. DATABASE SCHEMA (Backend)**

The Asset model is defined in [backend/models/asset.go](backend/models/asset.go) with these fields:

| Field | Type | Constraints | Purpose |
|-------|------|---|---|
| `ID` | UUID String | Primary Key | Auto-generated UUID on creation |
| `AssetCode` | String | Unique Index, Not Null | Unique identifier (e.g., ASSET-0001) |
| `Name` | String | Not Null | Asset name (e.g., "Dell Laptop XPS 15") |
| `Category` | String | - | Asset category (IT, Furniture, Vehicle, Equipment) |
| `SerialNumber` | String | - | Manufacturer serial number |
| `PurchaseDate` | DateTime | - | Date asset was purchased |
| `PurchaseValue` | Float64 | - | Original purchase cost (in LKR) |
| `CurrentValue` | Float64 | - | Current depreciated value |
| `Location` | String | - | Physical location (e.g., "IT Department - Floor 2") |
| `Status` | String | Default: ACTIVE | Asset lifecycle status |
| `PurchaseOrderID` | String | - | Link to procurement PO |
| `Notes` | String | - | Additional notes |
| `CreatedAt` | DateTime | - | Auto-set creation timestamp |
| `UpdatedAt` | DateTime | - | Auto-set update timestamp |

**Database Creation**: Tables auto-migrate via GORM in [main.go](backend/main.go)

---

## **2. API ENDPOINTS**

All endpoints are protected with **AuthMiddleware** and follow the pattern `/api/v1/`:

| Method | Endpoint | Handler | Description |
|--------|----------|---------|---|
| `GET` | `/assets` | `GetAssets()` | List all assets with optional filters |
| `POST` | `/assets` | `CreateAsset()` | Create new asset |
| `GET` | `/assets/:id` | `GetAsset()` | Get specific asset by ID |
| `PUT` | `/assets/:id` | `UpdateAsset()` | Update asset details |
| `DELETE` | `/assets/:id` | `DeleteAsset()` | Soft-delete (mark as DISPOSED) |
| `GET` | `/assets/:id/depreciation` | `GetDepreciation()` | Calculate depreciation |

**Query Parameters for GET /assets:**
- `status`: Filter by status (ACTIVE, MAINTENANCE, DISPOSED)
- `category`: Filter by category

---

## **3. KEY HANDLER LOGIC** ([backend/handlers/asset.go](backend/handlers/asset.go))

**GetAssets**: 
- Filters by status and category if provided
- Orders results by creation date (newest first)
- Returns count with data array

**CreateAsset**:
- Accepts JSON payload
- Auto-sets `CurrentValue = PurchaseValue` for new assets
- Returns 201 with created asset object

**UpdateAsset**:
- Finds asset by ID, updates all fields from JSON
- Returns updated asset

**DeleteAsset**:
- Soft delete: Changes status to "DISPOSED" instead of hard delete
- Preserves data for audit trail

**GetDepreciation**:
- Uses **Straight-Line Depreciation** method
- Assumes 5-year useful life
- Formula: `Annual Depreciation = PurchaseValue / 5`
- Returns depreciation metrics

---

## **4. FRONTEND - ASSET LIST VIEW** ([frontend/src/views/AssetList.vue](frontend/src/views/AssetList.vue))

**Features:**
- **Dashboard Stats**: Shows 4 metrics at top
  - Total assets count
  - Active assets count
  - Assets in maintenance count
  - Total portfolio value in LKR

- **Filtering**: Dropdown filters for Status & Category
  
- **Asset Table**: Displays columns
  - Asset Code (monospace, blue text)
  - Name
  - Category
  - Location
  - Status (color-coded badge)
  - Current Value (formatted with commas)

- **Add Asset Modal**: Form popup with fields
  - Name, Category, Serial Number, Location
  - Purchase Date, Purchase Value, Asset Code
  - Purchase Order ID

- **Row Click**: Navigate to detail page via router

**Status Badge Colors:**
- ACTIVE: Green (#eaf3de background, #3b6d11 text)
- MAINTENANCE: Orange (#faeeda background, #854f0b text)
- DISPOSED: Red (#fcebeb background, #a32d2d text)

---

## **5. FRONTEND - ASSET DETAIL VIEW** ([frontend/src/views/AssetDetail.vue](frontend/src/views/AssetDetail.vue))

**Features:**
- Back button navigation
- Display asset header with status badge
- **Detail table** showing:
  - Asset Code
  - Category
  - Serial Number
  - Location
  - Purchase Value
  - Current Value
  - Purchase Order ID

- **Action Buttons**:
  - Edit: Enable edit mode
  - Dispose: Soft delete with confirmation dialog

---

## **6. VUEX PINIA STORE** ([frontend/src/stores/assetStore.js](frontend/src/stores/assetStore.js))

**State:**
- `assets`: Array of asset objects
- `loading`: Boolean for async operations

**Actions:**
- `fetchAssets(filters = {})`: GET /assets with optional filters
  - Accepts `{status, category}` objects
  
- `createAsset(data)`: POST /assets
  - Adds new asset to beginning of array
  - Returns created asset object
  
- `updateAsset(id, data)`: PUT /assets/:id
  - Updates local array if asset found
  - Returns updated object

---

## **7. SAMPLE DATA** ([backend/db/init.sql](backend/db/init.sql))

4 demo assets pre-populated:

1. **Dell Laptop XPS 15** (IT)
   - Code: ASSET-0001 | Value: LKR 250,000 → 200,000
   - Status: ACTIVE

2. **Office Chair - Ergonomic** (Furniture)
   - Code: ASSET-0002 | Value: LKR 45,000 → 40,000
   - Status: ACTIVE

3. **Toyota Corolla 2023** (Vehicle)
   - Code: ASSET-0003 | Value: LKR 7,500,000 → 6,800,000
   - Status: ACTIVE

4. **HP LaserJet Pro** (IT)
   - Code: ASSET-0004 | Value: LKR 85,000 → 75,000
   - Status: MAINTENANCE

---

## **8. ASSET LIFECYCLE STATES**

| Status | Description | Can Transition To |
|--------|---|---|
| ACTIVE | In use and operational | MAINTENANCE, DISPOSED |
| MAINTENANCE | Currently under repair/maintenance | ACTIVE, DISPOSED |
| DISPOSED | Decommissioned or sold | (None - final state) |

---

## **9. INTEGRATION POINTS**

- **Assignments Module**: Assets can be assigned to employees
- **Maintenance Module**: Tracks maintenance records for assets
- **Procurement Module**: Links to Purchase Order IDs for audit trail
- **Authentication**: All routes protected by JWT middleware

---

## **10. TECH STACK**

- **Backend**: Go, Gin Framework, GORM ORM, PostgreSQL
- **Frontend**: Vue 3, Pinia Store, Axios HTTP Client
- **Database**: PostgreSQL with auto-migrations