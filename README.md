# Auth Lab Go (OIDC + Keycloak)

## Features
- Full OIDC login flow (Authorization Code)
- Keycloak integration ready
- Token validation (ID Token)

## Setup (Keycloak)
1. Create Realm: auth-lab
2. Create Client:
   - Client ID: go-client
   - Access Type: confidential
   - Redirect URI: http://localhost:8080/auth/oidc/callback

## Run
go mod tidy
go run ./cmd/server

## Endpoints
- /auth/oidc/login
- /auth/oidc/callback
- /health
