package infra

import "context"

// Database is an interface for database operations.
type Database interface {
	// Get retrieves a document by ID and returns a map of its fields.
	Get(ctx context.Context, collection, id string) (map[string]interface{}, error)

	// Create creates a new document with the given data and returns its ID.
	Create(ctx context.Context, collection string, data map[string]interface{}) (string, error)

	// Update updates a document by ID with the given data.
	Update(ctx context.Context, collection, id string, data map[string]interface{}) error

	// Delete deletes a document by ID.
	Delete(ctx context.Context, collection, id string) error

	// RunTransaction runs a transaction function and returns its result.
	// The transaction function should call the given Transaction object's methods to read and write data within the transaction.
	RunTransaction(ctx context.Context, f func(tx Transaction) (interface{}, error)) (interface{}, error)
}

// Transaction is an interface for transaction operations.
type Transaction interface {
	// Get retrieves a document by ID and returns a map of its fields.
	Get(ctx context.Context, collection, id string) (map[string]interface{}, error)

	// Create creates a new document with the given data and returns its ID.
	Create(ctx context.Context, collection string, data interface{}) (string, error)

	// Update updates a document by ID with the given data.
	Update(ctx context.Context, collection, id string, data interface{}) error

	// Delete deletes a document by ID.
	Delete(ctx context.Context, collection, id string) error
}

type Query interface {
	Where(field string, op string, value interface{}) Query
	OrderBy(field string, direction string) Query
	Limit(limit int) Query
	Offset(offset int) Query
}
