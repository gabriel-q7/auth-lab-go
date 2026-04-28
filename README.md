# Auth Lab Go (OIDC + Keycloak)

A minimal Go application demonstrating OpenID Connect (OIDC) authentication using the Authorization Code flow with Keycloak as the identity provider.

## Features
- ✅ OIDC Authorization Code Flow
- ✅ Keycloak integration
- ✅ ID Token validation
- ✅ User claims extraction
- ✅ Simple HTTP server with health check

## Prerequisites
- Go 1.22 or higher
- Docker (for running Keycloak)
- Basic understanding of OIDC/OAuth2

## Project Structure
```
auth-lab-go/
├── cmd/server/main.go           # Application entry point
├── internal/
│   ├── handlers/auth.go         # HTTP handlers
│   └── oidc/oidc.go             # OIDC client configuration
├── api-request/health.sh        # Health check script
└── go.mod                       # Go dependencies
```

## Setup

### 1. Run Keycloak (Docker)
```bash
docker run -p 8080:8080 \
  -e KEYCLOAK_ADMIN=admin \
  -e KEYCLOAK_ADMIN_PASSWORD=admin \
  quay.io/keycloak/keycloak:latest \
  start-dev
```

Access Keycloak Admin Console: http://localhost:8080

### 2. Configure Keycloak

#### Create Realm
1. Go to Keycloak Admin Console
2. Click **Create Realm**
3. Name: `auth-lab`
4. Click **Create**

#### Create Client
1. Go to **Clients** → **Create Client**
2. **General Settings:**
   - Client ID: `go-client`
   - Client authentication: **ON**
   - Authorization: **OFF**
3. Click **Next**
4. **Capability config:**
   - Standard flow: **Enabled**
   - Direct access grants: **Enabled**
5. Click **Next**
6. **Login Settings:**
   - Valid redirect URIs: `http://localhost:3000/auth/oidc/callback`
   - Web origins: `http://localhost:3000`
7. Click **Save**
8. Go to **Credentials** tab and copy the **Client Secret**

#### Create Test User
1. Go to **Users** → **Add User**
2. Username: `testuser`
3. Click **Create**
4. Go to **Credentials** tab
5. Set password: `testpass`
6. Temporary: **OFF**
7. Click **Set Password**

### 3. Configure Application

Update the configuration in [internal/oidc/oidc.go](internal/oidc/oidc.go):

```go
var (
    clientID     = "go-client"
    clientSecret = "YOUR_CLIENT_SECRET_HERE"  // From Keycloak
    redirectURL  = "http://localhost:3000/auth/oidc/callback"
    providerURL  = "http://localhost:8080/realms/auth-lab"  // Keycloak realm URL
)
```

**⚠️ Security Note:** For production, use environment variables instead of hardcoding credentials.

### 4. Install Dependencies & Run
```bash
go mod tidy
go run ./cmd/server
```

Server will start on `http://localhost:3000`

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/health` | GET | Health check endpoint |
| `/auth/oidc/login` | GET | Initiates OIDC login flow |
| `/auth/oidc/callback` | GET | Handles OIDC callback with authorization code |

## Usage

### 1. Test Health Check
```bash
curl http://localhost:3000/health
# Expected: OK
```

Or use the provided script:
```bash
chmod +x api-request/health.sh
./api-request/health.sh
```

### 2. Test OIDC Flow
1. Open browser and navigate to: `http://localhost:3000/auth/oidc/login`
2. You'll be redirected to Keycloak login page
3. Login with test credentials:
   - Username: `testuser`
   - Password: `testpass`
4. After successful login, you'll be redirected back to the callback endpoint
5. User claims will be displayed in the browser

### Expected Response
```
User info: map[email:testuser@example.com name:Test User sub:a1b2c3d4-... ...]
```

## How It Works

1. **Login Request** (`/auth/oidc/login`):
   - Generates OIDC authorization URL
   - Redirects user to Keycloak login page

2. **User Authentication**:
   - User enters credentials in Keycloak
   - Keycloak validates credentials

3. **Callback** (`/auth/oidc/callback`):
   - Keycloak redirects back with authorization code
   - App exchanges code for tokens (ID token, access token)
   - App verifies ID token signature
   - App extracts user claims from ID token

## Dependencies
- `github.com/coreos/go-oidc/v3` - OIDC client library
- `golang.org/x/oauth2` - OAuth2 client library

## Known Limitations
- ⚠️ No state parameter validation (CSRF protection)
- ⚠️ Credentials hardcoded (should use environment variables)
- ⚠️ No session management
- ⚠️ No logout functionality
- ⚠️ Basic error handling only

## Troubleshooting

### Port 8080 already in use
If you're running the Go app on port 8080, it will conflict with Keycloak. Update the Go server port in `cmd/server/main.go` to use a different port (e.g., 3000).

### Invalid redirect_uri
Ensure the redirect URI in:
- Keycloak client configuration
- `internal/oidc/oidc.go` (redirectURL variable)

matches exactly (including the port).

### Provider discovery failed
Ensure:
- Keycloak is running
- The realm name is correct
- The providerURL format is: `http://localhost:8080/realms/{realm-name}`

## Next Steps
- [ ] Add environment variable configuration
- [ ] Implement state parameter validation
- [ ] Add session management
- [ ] Implement logout endpoint
- [ ] Add refresh token support
- [ ] Add middleware for protected routes
- [ ] Improve error handling and logging
- [ ] Add unit tests

## License
MIT
