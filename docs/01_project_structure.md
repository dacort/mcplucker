
# Phase 1: Project Structure

## Objective
Set up a new Go project with proper module structure and DuckDB integration.

## Tasks

### 1. Initialize Go Module
- Create `go.mod` with module name `github.com/dacort/mcplucker`
- Use Go 1.24
- Add dependency on latest `marcboeker/go-duckdb` (v2.3.3+)
- Add dependency on `github.com/modelcontextprotocol/go-sdk`

### 2. Project Structure
```
mcplucker/
├── main.go              # Entry point with MCP server setup
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── internal/
│   ├── server/          # MCP server implementation
│   │   └── server.go    # Main server logic
│   ├── duckdb/          # DuckDB connection management
│   │   └── manager.go   # Database manager
│   └── tools/           # MCP tool implementations
│       └── tools.go     # Tool definitions
├── pkg/
│   └── types/           # Shared types and schemas
│       └── types.go     # Type definitions
└── docs/                # Documentation (already created)
```

### 3. Basic Implementation
- **main.go**: Basic MCP server setup with stdio transport
- **internal/server/server.go**: Server initialization and tool registration
- **internal/duckdb/manager.go**: DuckDB connection management
- **internal/tools/tools.go**: Placeholder for tools (Phase 2)
- **pkg/types/types.go**: Common types and structs

### 4. Dependencies
```go
require (
    github.com/marcboeker/go-duckdb/v2 v2.3.3
    github.com/modelcontextprotocol/go-sdk latest
)
```

### 5. Basic Health Check
- Implement a simple "ping" tool to verify MCP server is working
- Test DuckDB connection can be established
- Ensure server can start and respond via stdio

## Verification Steps
1. `go mod tidy` runs without errors
2. `go build` compiles successfully  
3. Basic server starts and can be pinged
4. DuckDB connection can be established

## Success Criteria
- Project compiles and runs
- Basic MCP server responds to ping
- DuckDB connection works
- Ready for Phase 2 tool implementation