# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./client/...
go test ./rest/admin/v20260301/product/...

# Run a single test
go test ./client/... -run TestClient_Call

# Run tests with verbose output
go test -v ./...

# Build all packages
go build ./...

# Format code
go fmt ./...
```

## Architecture

This is the **SHOPLINE Admin API Go SDK** — a REST/GraphQL client library for the SHOPLINE e-commerce platform. It is beta-phase and actively developed.

### Core Design Patterns

**Request/Response Interface Pattern**: Every API endpoint is implemented as a struct pair. Request structs implement `client.APIRequest`, response structs embed `client.BaseAPIResponse`. The client calls `req.GetEndpoint()`, `req.GetMethod()`, `req.GetData()`, etc. to build the HTTP request.

**Aware / Dependency Injection Pattern**: Service structs embed `client.BaseService` (which holds an `IClient`) and implement the `Aware` interface (`SetClient(IClient)`). The `support` package registers all services and injects the client automatically via `MustNewClient()`.

**Functional Options**: Client configuration uses the options pattern — `WithVersion()`, `WithTimeout()`, etc.

**Generic Pagination**: `client.ListAll[T]()` is a generic helper that iterates all pages automatically.

### Key Packages

- **`client/`** — Core HTTP client (`shopline_api_client.go`), interfaces (`base_data_model.go`), OAuth (`oauth.go`), webhook decoding (`webhook_client.go`), GraphQL (`graphql_client.go`), payment client (`payment_client.go`), pagination (`pagination.go`, `lists.go`), signature verification (`sign.go`).

- **`rest/admin/v20260301/`** and **`rest/admin/v20251201/`** — API endpoint definitions organized by resource (product, order, customer, metafield, webhook, etc.). Each resource is its own package. `v20260301` is the latest version.

- **`webhook/v20260301/`** and **`webhook/v20251201/`** — Strongly-typed webhook event structs. Each event type implements `IWebhookEvent`.

- **`support/`** — `MustNewClient()` creates a client and auto-registers all services via `GetClientAwares()`.

- **`config/`** — Constants for test credentials and defaults. `unit_test_config.go` for unit test values.

- **`rest/admin/test/`** — Shared test utilities and mock data used across REST tests.

- **`server/`** — Example server demonstrating OAuth flow, webhook handling, and signature verification.

### URL Construction

```
https://{storeHandle}.myshopline.com/admin/openapi/{apiVersion}/{endpoint}
```

Default API version: `v20251201`. Override per-client with `WithVersion("v20260301")` or per-request via `GetRequestOptions()`.

### Adding a New API Endpoint

1. Create a new file in the appropriate `rest/admin/{version}/{resource}/` package.
2. Define a request struct implementing `APIRequest` (embed `client.BaseAPIRequest`).
3. Define a response struct embedding `client.BaseAPIResponse`.
4. Add a method to the resource's service struct (which embeds `client.BaseService`).
5. Mirror the implementation in the other API version package if applicable.

### Testing Pattern

Tests use `github.com/jarcoal/httpmock` to mock HTTP calls — no real API calls are made. Test setup/teardown follows this pattern:

```go
func setup() {
    cli = client.MustNewClient(app, storeHandle, accessToken)
    httpmock.ActivateNonDefault(cli.Client)
}
func teardown() {
    httpmock.DeactivateAndReset()
}
```

Mock responses are registered with `httpmock.RegisterResponder(method, url, responder)` before each test.
