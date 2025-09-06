# MCPlucker - DuckDB MCP Server

## Project Overview
MCPlucker is a comprehensive Model Context Protocol (MCP) server for DuckDB, providing advanced data loading, querying, and analysis capabilities through a clean Go implementation.

## Build & Run Instructions

### Prerequisites
- Go 1.24 or later
- DuckDB extensions will be auto-loaded as needed

### Building
```bash
go mod tidy
go build -o mcplucker ./main.go
```

### Running
```bash
# Run the MCP server via stdio
./mcplucker

# Test with a simple ping
echo '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"ping","arguments":{}}}' | ./mcplucker
```

### Testing
```bash
go test ./...
```

## Development Guidelines

### Code Style
- Use simple, pragmatic, idiomatic Go code
- No complex abstractions or over-engineering
- Follow Go naming conventions and package structure
- Prefer composition over inheritance
- Handle errors explicitly, never ignore them

### Testing Requirements
- No fake data or mocks in tests
- Use real DuckDB connections for integration tests
- Create actual test files for file loading tests
- Test error conditions with real scenarios
- Include performance tests for large data operations

### Implementation Rules
- Always pause after each phase for verification before proceeding
- Each tool must have comprehensive error handling
- All responses must be structured and consistent
- Use DuckDB's native functionality wherever possible
- Secure credential handling (never log secrets)

## Phase Implementation Plan

### Current Phase: Phase 2 Complete ✅
See individual phase documents for detailed implementation plans:

- [Phase 1: Project Structure](docs/01_project_structure.md) - ✅ **COMPLETED**
- [Phase 2: MVP Functionality](docs/02_mvp_functionality.md) - ✅ **COMPLETED**
- [Phase 3: Cloud Connectivity](docs/03_cloud_connectivity.md) - **READY TO START**
- [Phase 4: Analysis Functionality](docs/04_analysis_functionality.md)
- [Phase 5: Database Connectivity](docs/05_db_connectivity.md)
- [Phase 100: Future Features](docs/100_future_features.md)

### TODO List
1. ✅ **Phase 1**: Set up Go module, project structure, basic MCP server
2. ✅ **Phase 2**: Implement core tools (create_table_from_file, list_tables, describe_table, execute_query)
3. **Phase 3**: Add cloud connectivity and Iceberg support
4. **Phase 4**: Implement query analysis and table profiling
5. **Phase 5**: Add database connectivity (PostgreSQL, MySQL, SQLite)
6. **Phase 100**: Advanced features and enterprise functionality

### Completed Features
- ✅ Basic MCP server with stdio transport
- ✅ DuckDB integration with connection management
- ✅ File loading from CSV, JSON, Parquet formats
- ✅ Table management (list, describe, create from files)
- ✅ SQL query execution with result formatting
- ✅ Comprehensive error handling
- ✅ Type-safe tool implementations

## Dependencies

### Core Dependencies
```go
require (
    github.com/marcboeker/go-duckdb/v2 v2.3.3+
    github.com/modelcontextprotocol/go-sdk latest
)
```

### DuckDB Extensions (Auto-loaded)
- Core extensions: ICU, JSON, Parquet, Autocomplete
- Phase 3: httpfs, aws, azure, iceberg  
- Phase 5: postgres, mysql, sqlite

## Architecture

### Project Structure
```
mcplucker/
├── main.go              # Entry point
├── internal/
│   ├── server/          # MCP server implementation  
│   ├── duckdb/          # Database management
│   └── tools/           # Tool implementations
├── pkg/
│   └── types/           # Shared types
└── docs/                # Phase documentation
```

### Key Components
- **MCP Server**: Handles protocol communication via stdio
- **DuckDB Manager**: Manages database connections and queries
- **Tool Registry**: Implements and registers all MCP tools
- **Type System**: Provides structured input/output schemas

## Verification Process

Each phase must pass these checks before proceeding:
1. Code compiles without warnings
2. All tests pass
3. Manual verification of phase objectives
4. Error handling tested with invalid inputs
5. Performance acceptable for intended use cases

## Security Notes
- Never log credentials or sensitive data
- Validate all user inputs
- Use parameterized queries to prevent injection
- Handle file system access safely
- Secure credential storage for database connections

## Performance Guidelines
- Stream large datasets when possible
- Use DuckDB's native bulk loading capabilities
- Implement appropriate timeouts
- Monitor memory usage for large operations
- Provide progress feedback for long-running tasks