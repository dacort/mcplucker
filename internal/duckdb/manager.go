package duckdb

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/marcboeker/go-duckdb"
)

// Manager handles DuckDB connections and operations
type Manager struct {
	db   *sql.DB
	mu   sync.RWMutex
	dsn  string
}

// NewManager creates a new DuckDB manager with an in-memory database
func NewManager() (*Manager, error) {
	return NewManagerWithDSN("")
}

// NewManagerWithDSN creates a new DuckDB manager with a specific DSN
// Empty DSN creates an in-memory database
func NewManagerWithDSN(dsn string) (*Manager, error) {
	if dsn == "" {
		dsn = ":memory:"
	}

	db, err := sql.Open("duckdb", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DuckDB connection: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping DuckDB: %w", err)
	}

	return &Manager{
		db:  db,
		dsn: dsn,
	}, nil
}

// Close closes the database connection
func (m *Manager) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.db != nil {
		return m.db.Close()
	}
	return nil
}

// Query executes a query and returns the result
func (m *Manager) Query(query string, args ...interface{}) (*sql.Rows, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.db == nil {
		return nil, fmt.Errorf("database connection is closed")
	}

	return m.db.Query(query, args...)
}

// Exec executes a query without returning rows
func (m *Manager) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.db == nil {
		return nil, fmt.Errorf("database connection is closed")
	}

	return m.db.Exec(query, args...)
}

// TestConnection tests if the database connection is working
func (m *Manager) TestConnection() error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.db == nil {
		return fmt.Errorf("database connection is nil")
	}

	return m.db.Ping()
}