package server

import (
	"context"
	"fmt"

	"github.com/dacort/mcplucker/internal/duckdb"
	"github.com/dacort/mcplucker/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Server represents the MCP server instance
type Server struct {
	server    *mcp.Server
	dbManager *duckdb.Manager
	tools     *tools.Tools
}

// New creates a new MCP server instance
func New() (*Server, error) {
	// Initialize DuckDB manager
	dbManager, err := duckdb.NewManager()
	if err != nil {
		return nil, fmt.Errorf("failed to create DuckDB manager: %w", err)
	}

	// Initialize tools
	toolsInstance := tools.New(dbManager)

	// Create MCP server
	impl := &mcp.Implementation{
		Name:    "mcplucker",
		Version: "0.1.0",
	}
	mcpServer := mcp.NewServer(impl, nil)

	// Create server instance
	server := &Server{
		server:    mcpServer,
		dbManager: dbManager,
		tools:     toolsInstance,
	}

	// Register tools
	if err := server.registerTools(); err != nil {
		_ = dbManager.Close()
		return nil, fmt.Errorf("failed to register tools: %w", err)
	}

	return server, nil
}

// registerTools registers all available tools with the MCP server
func (s *Server) registerTools() error {
	// Register ping tool
	pingTool := &mcp.Tool{
		Name:        "ping",
		Description: "Health check tool that tests server and database connectivity",
	}
	mcp.AddTool(s.server, pingTool, s.tools.Ping)

	// Register create_table_from_file tool
	createTableTool := &mcp.Tool{
		Name:        "create_table_from_file",
		Description: "Create a DuckDB table from a local file (CSV, Parquet, JSON)",
	}
	mcp.AddTool(s.server, createTableTool, s.tools.CreateTableFromFile)

	// Register list_tables tool
	listTablesTool := &mcp.Tool{
		Name:        "list_tables",
		Description: "List all tables and views in the database",
	}
	mcp.AddTool(s.server, listTablesTool, s.tools.ListTables)

	// Register describe_table tool
	describeTableTool := &mcp.Tool{
		Name:        "describe_table", 
		Description: "Get detailed schema information and sample data for a table",
	}
	mcp.AddTool(s.server, describeTableTool, s.tools.DescribeTable)

	// Register execute_query tool
	executeQueryTool := &mcp.Tool{
		Name:        "execute_query",
		Description: "Execute arbitrary SQL queries with configurable result limits",
	}
	mcp.AddTool(s.server, executeQueryTool, s.tools.ExecuteQuery)

	return nil
}

// Run starts the MCP server
func (s *Server) Run(ctx context.Context, transport mcp.Transport) error {
	defer s.dbManager.Close()
	return s.server.Run(ctx, transport)
}