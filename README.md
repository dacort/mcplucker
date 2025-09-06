# MCPlucker 🦆

A comprehensive Model Context Protocol (MCP) server for DuckDB, providing advanced data loading, querying, and analysis capabilities.

## Quick Start

### Prerequisites
- Go 1.24 or later

### Installation
```bash
git clone https://github.com/dacort/mcplucker.git
cd mcplucker
go build -o mcplucker ./main.go
```

### Usage
MCPlucker runs as an MCP server using stdio transport:

```bash
./mcplucker
```

## Available Tools

### Core Data Management
- **`create_table_from_file`** - Create tables from CSV, JSON, or Parquet files
- **`list_tables`** - List all tables and views in the database  
- **`describe_table`** - Get detailed schema information and sample data
- **`execute_query`** - Execute SQL queries with configurable limits

### Example Usage via MCP Client
```json
{
  "tool": "create_table_from_file",
  "arguments": {
    "table_name": "customers", 
    "file_path": "data/customers.csv"
  }
}
```

## Features

✅ **File Format Support**: CSV, JSON, Parquet  
✅ **DuckDB Integration**: Native DuckDB performance and SQL compatibility  
✅ **Type Safety**: Structured input/output with comprehensive validation  
✅ **Error Handling**: Clear error messages and validation  
✅ **MCP Compliant**: Full Model Context Protocol compatibility  

## Development Status

- ✅ **Phase 1**: Project structure and basic MCP server
- ✅ **Phase 2**: MVP functionality (file loading, querying)  
- 🚧 **Phase 3**: Cloud connectivity (S3, Azure, Iceberg)
- 📋 **Phase 4**: Analysis functionality (query optimization, profiling)
- 📋 **Phase 5**: Database connectivity (PostgreSQL, MySQL, SQLite)

## Contributing

See [CLAUDE.md](CLAUDE.md) for detailed development guidelines and build instructions.

## License

MIT License