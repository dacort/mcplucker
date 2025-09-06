package main

import (
	"context"
	"log"

	"github.com/dacort/mcplucker/internal/server"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()

	// Create and configure the MCP server
	mcpServer, err := server.New()
	if err != nil {
		log.Fatalf("Failed to create MCP server: %v", err)
	}

	// Set up stdio transport
	transport := &mcp.StdioTransport{}

	// Run the server
	if err := mcpServer.Run(ctx, transport); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}