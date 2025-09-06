# Phase 2: MVP Functionality

## Objective
Implement core MCP tools for basic DuckDB operations: file loading, table management, and querying.

## Tools to Implement

### 1. create_table_from_file
**Purpose**: Create DuckDB tables from local files (CSV, Parquet, JSON)

**Input Schema**:
```json
{
  "table_name": "string",
  "file_path": "string", 
  "file_type": "csv|parquet|json",
  "options": {
    "header": "boolean (CSV only)",
    "delimiter": "string (CSV only)", 
    "schema": "object (optional)"
  }
}
```

**Implementation**:
- Validate file exists and is readable
- Detect file type if not specified
- Use DuckDB's native file readers (CSV, Parquet, JSON)
- Handle common parsing options
- Return table creation status and row count

### 2. list_tables
**Purpose**: Show all tables and views in the current database

**Input Schema**: No input required

**Output**:
- Table/view names
- Types (table/view)  
- Row counts
- Creation timestamps if available

### 3. describe_table
**Purpose**: Get detailed schema information for a table

**Input Schema**:
```json
{
  "table_name": "string"
}
```

**Output**:
- Column names and types
- Nullable information
- Primary keys if any
- Sample data (first 5 rows)

### 4. execute_query
**Purpose**: Execute arbitrary SQL queries

**Input Schema**:
```json
{
  "query": "string",
  "limit": "number (optional, default 100)",
  "format": "table|json|csv (optional, default table)"
}
```

**Implementation**:
- Validate SQL syntax
- Execute with configurable row limits
- Format results appropriately
- Handle errors gracefully
- Support both SELECT and DDL statements

## Implementation Details

### Error Handling
- Standardized error responses with DuckDB error details
- Input validation with clear messages
- File access error handling
- SQL syntax error reporting

### Response Formatting
- Consistent JSON response structure
- Table formatting for readable output
- Proper type handling (dates, numbers, nulls)
- Truncation indicators for large results

### Testing Strategy
- Unit tests for each tool
- Integration tests with sample files
- Error condition testing
- Performance testing with larger files

## Sample Files for Testing
Create `test_data/` directory with:
- `customers.csv` - Basic CSV with headers
- `orders.parquet` - Parquet file with mixed types
- `products.json` - JSON array format

## Verification Steps
1. All tools respond correctly to valid inputs
2. Error handling works for invalid inputs
3. File loading works with different formats
4. Query execution handles various SQL types
5. Table descriptions show accurate schema info

## Success Criteria
- All 4 tools implemented and working
- Comprehensive error handling
- Clean, readable output formatting
- Ready for Phase 3 cloud connectivity