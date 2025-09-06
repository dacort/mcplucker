# Phase 5: Database Connectivity

## Objective
Enable connections to external databases (PostgreSQL, MySQL, SQLite) for data federation and cross-database queries.

## Tools to Implement

### 1. connect_database
**Purpose**: Establish connections to external databases and make them available for querying

**Input Schema**:
```json
{
  "connection_name": "string (unique identifier)",
  "database_type": "postgres|mysql|sqlite",
  "connection_params": {
    "host": "string (postgres/mysql)",
    "port": "number (postgres/mysql)", 
    "database": "string",
    "username": "string (postgres/mysql)",
    "password": "string (postgres/mysql)",
    "file_path": "string (sqlite only)",
    "ssl_mode": "disable|require|verify-full (postgres only)",
    "charset": "string (mysql only, default utf8mb4)"
  },
  "options": {
    "read_only": "boolean (default true)",
    "connection_timeout": "number (seconds, default 30)",
    "query_timeout": "number (seconds, default 300)"
  }
}
```

**Implementation**:
- Use DuckDB's postgres, mysql, and sqlite extensions
- Validate connection parameters
- Test connectivity before storing
- Support connection pooling
- Handle SSL/TLS configuration

### 2. list_database_connections
**Purpose**: Show all active database connections and their status

**Output**:
```json
{
  "connections": [
    {
      "name": "string",
      "type": "postgres|mysql|sqlite", 
      "status": "connected|disconnected|error",
      "last_used": "timestamp",
      "table_count": "number",
      "error_message": "string (if status=error)"
    }
  ]
}
```

### 3. list_remote_tables
**Purpose**: List tables available in a connected database

**Input Schema**:
```json
{
  "connection_name": "string",
  "schema_filter": "string (optional, postgres only)",
  "table_pattern": "string (optional, SQL LIKE pattern)"
}
```

**Output**:
- Table names with schema (postgres)
- Table types (table, view, materialized view)
- Row count estimates
- Last modified timestamps where available

### 4. describe_remote_table
**Purpose**: Get detailed schema information from external database tables

**Input Schema**:
```json
{
  "connection_name": "string", 
  "table_name": "string",
  "include_sample": "boolean (optional, default true)"
}
```

**Output**:
- Column definitions with native types
- DuckDB equivalent types
- Constraints (primary keys, foreign keys, unique)
- Indexes information
- Sample data (first 5 rows)

### 5. sync_table
**Purpose**: Copy external table data into local DuckDB for faster querying

**Input Schema**:
```json
{
  "connection_name": "string",
  "source_table": "string", 
  "local_table": "string",
  "sync_mode": "full|incremental",
  "incremental_column": "string (required for incremental)",
  "where_clause": "string (optional filter)"
}
```

**Implementation**:
- Full table copy for complete refresh
- Incremental sync based on timestamp/ID columns
- Batch processing for large tables
- Progress reporting for long operations
- Conflict resolution strategies

## Implementation Details

### Connection Management
- Secure credential storage (no plaintext passwords)
- Connection pooling and reuse
- Automatic reconnection on failures
- Health checks for active connections
- Graceful connection cleanup

### Cross-Database Queries
Enable queries across multiple databases:
```sql
-- Query joining local DuckDB data with remote PostgreSQL
SELECT l.*, r.additional_info 
FROM local_table l
JOIN postgres_connection.public.remote_table r ON l.id = r.local_id
```

### Type Mapping
Handle type conversions between database systems:
- PostgreSQL → DuckDB type mapping
- MySQL → DuckDB type mapping  
- SQLite → DuckDB type mapping
- Handle special types (JSON, arrays, enums)
- Timezone handling for timestamp columns

### Performance Optimization
- Query pushdown to remote databases
- Parallel data loading
- Compression during data transfer
- Smart caching strategies
- Query result pagination

### Security Considerations
- Encrypted credential storage
- Connection string sanitization
- SQL injection prevention
- Role-based access control
- Audit logging for sensitive operations

## Extension Requirements
Ensure these DuckDB extensions are available:
- `postgres` - PostgreSQL connectivity
- `mysql` - MySQL connectivity  
- `sqlite` - SQLite file access

## Error Handling
- Network connectivity issues
- Authentication failures
- Permission denied errors
- Timeout handling
- Schema mismatch detection
- Large result set warnings

## Testing Strategy
- Docker containers for test databases
- Connection pooling stress tests
- Large data sync testing
- Cross-database query validation
- Network failure simulation
- Security vulnerability testing

## Configuration Management
Support for connection configuration files:
```yaml
# connections.yaml
connections:
  prod_postgres:
    type: postgres
    host: db.example.com
    port: 5432
    database: production
    username: ${POSTGRES_USER}
    password: ${POSTGRES_PASSWORD}
    
  analytics_mysql:
    type: mysql
    host: mysql.example.com
    database: analytics
    username: ${MYSQL_USER}
    password: ${MYSQL_PASSWORD}
```

## Verification Steps
1. Can connect to PostgreSQL, MySQL, and SQLite databases
2. Table listing works for each database type
3. Schema description shows accurate information
4. Data synchronization completes successfully
5. Cross-database queries execute correctly
6. Connection management handles failures gracefully

## Success Criteria
- Reliable multi-database connectivity
- Efficient data synchronization
- Cross-database query capabilities
- Secure credential management
- Comprehensive error handling
- Ready for advanced features in Phase 100