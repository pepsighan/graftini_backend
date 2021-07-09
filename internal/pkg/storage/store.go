package storage

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/backend/config"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
)

// UploadFile uploads the file to GCP Cloud Storage.
func UploadFile(ctx context.Context, file io.Reader, storageClient *storage.Client, entClient *ent.Client) (*ent.File, error) {
	fileID := uuid.New()

	bucket := storageClient.Bucket(config.GoogleCloudStorageBucket)
	object := bucket.Object(fileID.String())

	writer := object.NewWriter(ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, file); err != nil {
		return nil, fmt.Errorf("could not upload file: %w", err)
	}

	return nil, nil
}
