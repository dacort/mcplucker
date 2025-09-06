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

MCPlucker runs as an MCP server using stdio transport

**Adding to Claude Code:**

1. Build the project:
```bash
go build -o mcplucker ./main.go
```

2. Add the MCP server to Claude Code:
```bash
claude mcp add duckdb ./mcplucker
```

3. Start using MCPlucker in your Claude Code sessions!

## Available Tools

### Core Data Management
- **`create_table_from_file`** - Create tables from CSV, JSON, or Parquet files
- **`list_tables`** - List all tables and views in the database  
- **`describe_table`** - Get detailed schema information and sample data
- **`execute_query`** - Execute SQL queries with configurable limits

### Example Usage

Once connected to an MCP client (like Claude Desktop), you can interact naturally:

**Loading data:**
```
"Hello, my friend, can you create a duckdb database from test_data/products.json?"
```

**Exploring tables:**
```
"Can you give me a quick summary of the products table?"
```

```
⏺ Products Table Summary:
  - 5 total products across 2 categories (Electronics, Furniture)
  - Price range: $25.50 - $999.99 (avg: $320.09)
  - Stock status: 4 in stock, 1 out of stock
  - Categories: Electronics (4 items), Furniture (1 item)
```

**Running queries:**
```
"Show me all electronics products that are in stock"
```

```
⏺ Electronics Products In Stock:

  | ID  | Name    | Price   |
  |-----|---------|---------|
  | 1   | Laptop  | $999.99 |
  | 2   | Mouse   | $25.50  |
  | 4   | Monitor | $299.99 |

  3 electronics products are currently in stock.
```

**Technical usage via MCP protocol:**
```json
{
  "tool": "create_table_from_file",
  "arguments": {
    "table_name": "products", 
    "file_path": "test_data/products.json"
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