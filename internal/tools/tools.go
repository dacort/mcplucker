package tools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dacort/mcplucker/internal/duckdb"
	"github.com/dacort/mcplucker/pkg/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Tools holds references to all tool implementations
type Tools struct {
	dbManager *duckdb.Manager
}

// New creates a new Tools instance
func New(dbManager *duckdb.Manager) *Tools {
	return &Tools{
		dbManager: dbManager,
	}
}

// Ping is a simple health check tool
func (t *Tools) Ping(ctx context.Context, req *mcp.CallToolRequest, input types.PingInput) (*mcp.CallToolResult, types.PingOutput, error) {
	// Test database connection
	if err := t.dbManager.TestConnection(); err != nil {
		return nil, types.PingOutput{}, fmt.Errorf("database connection failed: %w", err)
	}

	message := input.Message
	if message == "" {
		message = "ping"
	}

	output := types.PingOutput{
		Response: fmt.Sprintf("pong: %s", message),
		Status:   "healthy",
	}

	return nil, output, nil
}

// CreateTableFromFile creates a DuckDB table from a file
func (t *Tools) CreateTableFromFile(ctx context.Context, req *mcp.CallToolRequest, input types.CreateTableFromFileInput) (*mcp.CallToolResult, types.CreateTableFromFileOutput, error) {
	// Validate input
	if input.TableName == "" {
		return nil, types.CreateTableFromFileOutput{}, fmt.Errorf("table_name is required")
	}
	if input.FilePath == "" {
		return nil, types.CreateTableFromFileOutput{}, fmt.Errorf("file_path is required")
	}

	// Check if file exists
	if _, err := os.Stat(input.FilePath); os.IsNotExist(err) {
		return nil, types.CreateTableFromFileOutput{}, fmt.Errorf("file does not exist: %s", input.FilePath)
	}

	// Detect file type if not specified
	fileType := input.FileType
	if fileType == "" {
		ext := strings.ToLower(filepath.Ext(input.FilePath))
		switch ext {
		case ".csv":
			fileType = "csv"
		case ".parquet":
			fileType = "parquet"
		case ".json":
			fileType = "json"
		default:
			return nil, types.CreateTableFromFileOutput{}, fmt.Errorf("unsupported file extension: %s", ext)
		}
	}

	// Build CREATE TABLE AS SELECT statement
	var query string
	absPath, err := filepath.Abs(input.FilePath)
	if err != nil {
		return nil, types.CreateTableFromFileOutput{}, fmt.Errorf("failed to get absolute path: %w", err)
	}

	switch fileType {
	case "csv":
		header := true
		delimiter := ","
		if options := input.Options; options != nil {
			if h, ok := options["header"].(bool); ok {
				header = h
			}
			if d, ok := options["delimiter"].(string); ok {
				delimiter = d
			}
		}
		query = fmt.Sprintf("CREATE TABLE %s AS SELECT * FROM read_csv('%s', header=%t, delim='%s')", 
			input.TableName, absPath, header, delimiter)
	
	case "parquet":
		query = fmt.Sprintf("CREATE TABLE %s AS SELECT * FROM read_parquet('%s')", 
			input.TableName, absPath)
	
	case "json":
		query = fmt.Sprintf("CREATE TABLE %s AS SELECT * FROM read_json('%s')", 
			input.TableName, absPath)
	
	default:
		return nil, types.CreateTableFromFileOutput{}, fmt.Errorf("unsupported file type: %s", fileType)
	}

	// Execute the query
	if _, err := t.dbManager.Exec(query); err != nil {
		return nil, types.CreateTableFromFileOutput{}, fmt.Errorf("failed to create table: %w", err)
	}

	// Get row count
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", input.TableName)
	rows, err := t.dbManager.Query(countQuery)
	if err != nil {
		return nil, types.CreateTableFromFileOutput{}, fmt.Errorf("failed to count rows: %w", err)
	}
	defer rows.Close()

	var rowCount int64
	if rows.Next() {
		if err := rows.Scan(&rowCount); err != nil {
			return nil, types.CreateTableFromFileOutput{}, fmt.Errorf("failed to scan row count: %w", err)
		}
	}

	output := types.CreateTableFromFileOutput{
		TableName: input.TableName,
		RowCount:  rowCount,
		Message:   fmt.Sprintf("Successfully created table '%s' from file '%s' with %d rows", input.TableName, input.FilePath, rowCount),
	}

	return nil, output, nil
}

// ListTables lists all tables and views in the database
func (t *Tools) ListTables(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, types.ListTablesOutput, error) {
	query := `
		SELECT table_name, table_type 
		FROM information_schema.tables 
		WHERE table_schema = 'main' 
		ORDER BY table_name
	`

	rows, err := t.dbManager.Query(query)
	if err != nil {
		return nil, types.ListTablesOutput{}, fmt.Errorf("failed to list tables: %w", err)
	}
	defer rows.Close()

	var tables []types.TableInfo
	for rows.Next() {
		var name, tableType string
		if err := rows.Scan(&name, &tableType); err != nil {
			return nil, types.ListTablesOutput{}, fmt.Errorf("failed to scan table info: %w", err)
		}

		// Get row count for tables (not views)
		var rowCount int64
		if strings.ToUpper(tableType) == "BASE TABLE" {
			countQuery := fmt.Sprintf("SELECT COUNT(*) FROM \"%s\"", name)
			countRows, err := t.dbManager.Query(countQuery)
			if err == nil {
				defer countRows.Close()
				if countRows.Next() {
					countRows.Scan(&rowCount)
				}
			}
		}

		tables = append(tables, types.TableInfo{
			Name:     name,
			Type:     strings.ToLower(tableType),
			RowCount: rowCount,
		})
	}

	output := types.ListTablesOutput{
		Tables: tables,
	}

	return nil, output, nil
}

// DescribeTable provides detailed schema information for a table
func (t *Tools) DescribeTable(ctx context.Context, req *mcp.CallToolRequest, input types.DescribeTableInput) (*mcp.CallToolResult, types.DescribeTableOutput, error) {
	if input.TableName == "" {
		return nil, types.DescribeTableOutput{}, fmt.Errorf("table_name is required")
	}

	// Get column information
	query := `
		SELECT column_name, data_type, is_nullable 
		FROM information_schema.columns 
		WHERE table_name = ? AND table_schema = 'main'
		ORDER BY ordinal_position
	`

	rows, err := t.dbManager.Query(query, input.TableName)
	if err != nil {
		return nil, types.DescribeTableOutput{}, fmt.Errorf("failed to describe table: %w", err)
	}
	defer rows.Close()

	var columns []types.ColumnInfo
	for rows.Next() {
		var name, dataType, nullable string
		if err := rows.Scan(&name, &dataType, &nullable); err != nil {
			return nil, types.DescribeTableOutput{}, fmt.Errorf("failed to scan column info: %w", err)
		}

		columns = append(columns, types.ColumnInfo{
			Name:     name,
			Type:     dataType,
			Nullable: strings.ToUpper(nullable) == "YES",
		})
	}

	if len(columns) == 0 {
		return nil, types.DescribeTableOutput{}, fmt.Errorf("table '%s' not found", input.TableName)
	}

	// Get sample data (first 5 rows)
	sampleQuery := fmt.Sprintf("SELECT * FROM \"%s\" LIMIT 5", input.TableName)
	sampleRows, err := t.dbManager.Query(sampleQuery)
	if err != nil {
		return nil, types.DescribeTableOutput{}, fmt.Errorf("failed to get sample data: %w", err)
	}
	defer sampleRows.Close()

	var sampleData [][]string
	for sampleRows.Next() {
		// Create a slice to hold the values
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := sampleRows.Scan(valuePtrs...); err != nil {
			return nil, types.DescribeTableOutput{}, fmt.Errorf("failed to scan sample data: %w", err)
		}

		// Convert to strings
		row := make([]string, len(columns))
		for i, val := range values {
			if val == nil {
				row[i] = "NULL"
			} else {
				row[i] = fmt.Sprintf("%v", val)
			}
		}
		sampleData = append(sampleData, row)
	}

	output := types.DescribeTableOutput{
		TableName:  input.TableName,
		Columns:    columns,
		SampleData: sampleData,
	}

	return nil, output, nil
}

// ExecuteQuery executes an arbitrary SQL query
func (t *Tools) ExecuteQuery(ctx context.Context, req *mcp.CallToolRequest, input types.ExecuteQueryInput) (*mcp.CallToolResult, types.ExecuteQueryOutput, error) {
	if input.Query == "" {
		return nil, types.ExecuteQueryOutput{}, fmt.Errorf("query is required")
	}

	// Set default limit
	limit := input.Limit
	if limit <= 0 {
		limit = 100
	}

	// Execute query
	rows, err := t.dbManager.Query(input.Query)
	if err != nil {
		return nil, types.ExecuteQueryOutput{}, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	// Get column names
	columnNames, err := rows.Columns()
	if err != nil {
		return nil, types.ExecuteQueryOutput{}, fmt.Errorf("failed to get columns: %w", err)
	}

	// Read data
	var data [][]string
	rowCount := 0
	truncated := false

	for rows.Next() && rowCount < limit {
		// Create a slice to hold the values
		values := make([]interface{}, len(columnNames))
		valuePtrs := make([]interface{}, len(columnNames))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, types.ExecuteQueryOutput{}, fmt.Errorf("failed to scan row: %w", err)
		}

		// Convert to strings
		row := make([]string, len(columnNames))
		for i, val := range values {
			if val == nil {
				row[i] = "NULL"
			} else {
				row[i] = fmt.Sprintf("%v", val)
			}
		}
		data = append(data, row)
		rowCount++
	}

	// Check if there are more rows
	if rows.Next() {
		truncated = true
	}

	var message string
	if truncated {
		message = fmt.Sprintf("Results limited to %d rows", limit)
	}

	output := types.ExecuteQueryOutput{
		Columns:   columnNames,
		Data:      data,
		RowCount:  rowCount,
		Truncated: truncated,
		Message:   message,
	}

	return nil, output, nil
}