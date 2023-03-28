package storage

import (
	. "app/utility/logger"
	"context"
	"fmt"
	"io/ioutil"

	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseStorageConnection struct {
	bucket *storage.BucketHandle
}

func NewFirebaseStorageConnection(ctx context.Context) *FirebaseStorageConnection {
	projectID := os.Getenv("PROJECT_ID")
	var opt option.ClientOption
	if value := os.Getenv("STORAGE_EMULATOR_HOST"); value != "" {
		Log.Infof("Using Firebase Storage Emulator: %s", value)
		opt = option.WithoutAuthentication()
	} else {
		opt = option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICE_ACCOUNT_PAHT"))
	}
	conf := &firebase.Config{
		StorageBucket: fmt.Sprintf("%s.appspot.com", projectID),
	}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		Log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		Log.Fatalf("error getting Storage client: %v", err)
	}
	bucket, err := client.DefaultBucket()
	if err != nil {
		Log.Fatalf("error getting bucket: %v", err)
	}
	return &FirebaseStorageConnection{
		bucket: bucket,
	}
}

func (fsc *FirebaseStorageConnection) MakeDirectory(ctx context.Context, folderName string) error {
	folderObject := fsc.bucket.Object(folderName + "/")
	if _, err := folderObject.Attrs(ctx); err != nil {
		if err == storage.ErrObjectNotExist {
			if err := folderObject.NewWriter(ctx).Close(); err != nil {
				Log.Fatalf("error: %v", err)
				return err
			}
			return nil
		} else {
			return err
		}
	} else {
		return nil
	}
}

func (fsc *FirebaseStorageConnection) UploadFile(ctx context.Context, upFile []byte, fileName string) error {
	obj := fsc.bucket.Object(fileName)
	wc := obj.NewWriter(ctx)
	if _, err := wc.Write(upFile); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

func (fsc *FirebaseStorageConnection) DownloadFile(ctx context.Context, fileName string) ([]byte, error) {
	obj := fsc.bucket.Object(fileName)
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (fsc *FirebaseStorageConnection) GetMedialink(ctx context.Context, fileName string) (string, error) {
	obj := fsc.bucket.Object(fileName)
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		Log.Infof("GetMediaLink: %s", err)
		return "", err
	}
	url := attrs.MediaLink
	return url, nil
}
