# Phase 3: Cloud Connectivity

## Objective
Add support for loading data from cloud sources and modern table formats like Iceberg.

## Tools to Implement

### 1. create_table_from_url
**Purpose**: Create tables from remote files via HTTP/HTTPS/S3/Azure/GCS

**Input Schema**:
```json
{
  "table_name": "string",
  "url": "string",
  "file_type": "csv|parquet|json (optional, auto-detect)",
  "credentials": {
    "aws_access_key_id": "string (optional)",
    "aws_secret_access_key": "string (optional)", 
    "aws_region": "string (optional)",
    "azure_storage_account": "string (optional)",
    "azure_access_key": "string (optional)"
  },
  "options": {
    "header": "boolean (CSV only)",
    "delimiter": "string (CSV only)"
  }
}
```

**Supported URL Patterns**:
- `https://example.com/data.csv` (HTTP/HTTPS)
- `s3://bucket/path/file.parquet` (S3)
- `azure://account/container/file.json` (Azure Blob)
- `gs://bucket/path/file.csv` (Google Cloud Storage)

### 2. load_iceberg_table
**Purpose**: Connect to and query Iceberg tables

**Input Schema**:
```json
{
  "table_name": "string",
  "catalog_url": "string",
  "warehouse_path": "string", 
  "table_identifier": "string (namespace.table)",
  "credentials": {
    "aws_access_key_id": "string (optional)",
    "aws_secret_access_key": "string (optional)",
    "aws_region": "string (optional)"
  }
}
```

**Implementation**:
- Use DuckDB's Iceberg extension
- Support AWS Glue, REST, and Hive catalogs
- Handle Iceberg schema evolution
- Support time travel queries

### 3. test_connection
**Purpose**: Validate cloud credentials and connectivity

**Input Schema**:
```json
{
  "connection_type": "s3|azure|gcs|http",
  "url": "string",
  "credentials": "object (same structure as above tools)"
}
```

**Output**:
- Connection status (success/failure)
- Error details if failed
- Basic metadata (file size, last modified) if successful

## Implementation Details

### Extension Management
- Auto-load required extensions:
  - `httpfs` for HTTP/S3/Azure/GCS
  - `iceberg` for Iceberg support
  - `aws` for enhanced S3 functionality
- Handle extension loading failures gracefully

### Credential Management
- Support environment variables for credentials
- Secure credential handling (no logging of secrets)
- Support for AWS credential profiles
- Azure storage account key and SAS token support

### URL Parsing and Validation
- Parse and validate different URL schemes
- Extract bucket/container names
- Handle URL encoding issues
- Support for presigned URLs

### Error Handling
- Network connectivity issues
- Authentication failures
- File not found errors
- Unsupported file formats
- Region/endpoint mismatches

### Performance Considerations
- Streaming for large files
- Parallel loading when possible
- Progress indicators for long operations
- Connection timeouts and retries

## Testing Strategy
- Mock cloud services for unit tests
- Integration tests with real cloud accounts (optional)
- Error simulation (network failures, auth errors)
- Large file handling tests

## Environment Setup
Document required environment variables:
```bash
# AWS
AWS_ACCESS_KEY_ID=your_key
AWS_SECRET_ACCESS_KEY=your_secret
AWS_DEFAULT_REGION=us-west-2

# Azure
AZURE_STORAGE_ACCOUNT=your_account
AZURE_STORAGE_KEY=your_key

# Google Cloud
GOOGLE_APPLICATION_CREDENTIALS=path/to/service-account.json
```

## Verification Steps
1. Can load CSV/Parquet/JSON from HTTP URLs
2. S3 integration works with credentials
3. Azure Blob Storage connectivity
4. Iceberg table loading and querying
5. Connection testing validates access
6. Proper error handling for auth failures

## Success Criteria
- Reliable cloud file loading
- Iceberg integration functional
- Comprehensive error handling
- Secure credential management
- Ready for Phase 4 analysis features