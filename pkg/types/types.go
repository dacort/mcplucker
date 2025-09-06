package types

// Common response structure for all tools
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PingInput represents input for the ping tool
type PingInput struct {
	Message string `json:"message,omitempty"`
}

// PingOutput represents output for the ping tool
type PingOutput struct {
	Response string `json:"response"`
	Status   string `json:"status"`
}

// CreateTableFromFileInput represents input for create_table_from_file tool
type CreateTableFromFileInput struct {
	TableName string                 `json:"table_name"`
	FilePath  string                 `json:"file_path"`
	FileType  string                 `json:"file_type,omitempty"` // csv, parquet, json
	Options   map[string]interface{} `json:"options,omitempty"`
}

// CreateTableFromFileOutput represents output for create_table_from_file tool
type CreateTableFromFileOutput struct {
	TableName string `json:"table_name"`
	RowCount  int64  `json:"row_count"`
	Message   string `json:"message"`
}

// ListTablesOutput represents output for list_tables tool
type ListTablesOutput struct {
	Tables []TableInfo `json:"tables"`
}

// TableInfo represents information about a table
type TableInfo struct {
	Name     string `json:"name"`
	Type     string `json:"type"` // table, view
	RowCount int64  `json:"row_count,omitempty"`
}

// DescribeTableInput represents input for describe_table tool
type DescribeTableInput struct {
	TableName string `json:"table_name"`
}

// DescribeTableOutput represents output for describe_table tool
type DescribeTableOutput struct {
	TableName   string       `json:"table_name"`
	Columns     []ColumnInfo `json:"columns"`
	SampleData  [][]string   `json:"sample_data,omitempty"`
}

// ColumnInfo represents information about a table column
type ColumnInfo struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Nullable bool   `json:"nullable"`
}

// ExecuteQueryInput represents input for execute_query tool
type ExecuteQueryInput struct {
	Query  string `json:"query"`
	Limit  int    `json:"limit,omitempty"`  // default 100
	Format string `json:"format,omitempty"` // table, json, csv
}

// ExecuteQueryOutput represents output for execute_query tool
type ExecuteQueryOutput struct {
	Columns    []string      `json:"columns"`
	Data       [][]string    `json:"data"`
	RowCount   int           `json:"row_count"`
	Truncated  bool          `json:"truncated,omitempty"`
	Message    string        `json:"message,omitempty"`
}