package db

import (
	"app/infra"
	. "app/utility/logger"
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FirestoreConnection struct {
	Client *firestore.Client
}

func NewFirestoreConnection(ctx context.Context) *FirestoreConnection {
	projectID := os.Getenv("PROJECT_ID")
	if value := os.Getenv("FIRESTORE_EMULATOR_HOST"); value != "" {
		Log.Infof("Using Firestore Emulator: %s", value)
	}
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICE_ACCOUNT_PAHT"))
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		Log.Errorf("Failed to create client: %v", err)
	}
	return &FirestoreConnection{
		Client: client,
	}
}

// Get retrieves a document by ID and returns a map of its fields.
func (db *FirestoreConnection) Get(ctx context.Context, collection, id string) (map[string]interface{}, error) {
	docRef := db.Client.Collection(collection).Doc(id)

	docSnapshot, err := docRef.Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return docSnapshot.Data(), nil
}

// Create creates a new document with the given data and returns its ID.
func (db *FirestoreConnection) Create(ctx context.Context, collection string, data map[string]interface{}) (string, error) {
	docRef := db.Client.Collection(collection).NewDoc()

	_, err := docRef.Create(ctx, data)
	if err != nil {
		return "", err
	}

	return docRef.ID, nil
}

// Update updates a document by ID with the given data.
func (db *FirestoreConnection) Update(ctx context.Context, collection, id string, data interface{}) error {
	docRef := db.Client.Collection(collection).Doc(id)

	_, err := docRef.Set(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a document by ID.
func (db *FirestoreConnection) Delete(ctx context.Context, collection, id string) error {
	docRef := db.Client.Collection(collection).Doc(id)

	_, err := docRef.Delete(ctx)
	if status.Code(err) == codes.NotFound {
		return err
	} else if err != nil {
		return err
	}

	return nil
}

// RunTransaction runs a transaction function and returns its result.
func (db *FirestoreConnection) RunTransaction(ctx context.Context, f func(tx infra.Transaction) (interface{}, error)) (interface{}, error) {
	tx := &firestoreTransaction{
		ctx:    ctx,
		client: db.Client,
		batch:  db.Client.Batch(),
	}

	result, err := f(tx)
	if err != nil {
		return nil, err
	}

	_, err = tx.batch.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// firestoreTransaction is a Firestore implementation of the Transaction interface.
type firestoreTransaction struct {
	ctx    context.Context
	client *firestore.Client
	batch  *firestore.WriteBatch
}

func (tx *firestoreTransaction) Get(ctx context.Context, collection, id string) (map[string]interface{}, error) {
	var result map[string]interface{}
	snap, err := tx.client.Collection(collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	result = snap.Data()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *firestoreTransaction) Create(ctx context.Context, collection string, data interface{}) (string, error) {
	docRef := t.client.Collection(collection).NewDoc()
	t.batch.Create(docRef, data)
	return docRef.ID, nil
}

func (t *firestoreTransaction) Update(ctx context.Context, collection, id string, data interface{}) error {
	docRef := t.client.Collection(collection).Doc(id)
	t.batch.Set(docRef, data, firestore.MergeAll)

	return nil
}

func (t *firestoreTransaction) Delete(ctx context.Context, collection, id string) error {
	docRef := t.client.Collection(collection).Doc(id)
	t.batch.Delete(docRef)

	return nil
}
