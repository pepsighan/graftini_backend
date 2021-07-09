package storage

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/backend/config"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/file"
)

// UploadFile uploads the file to GCP Cloud Storage.
func UploadFile(ctx context.Context, reader io.Reader, storageClient *storage.Client, entClient *ent.Client) (*ent.File, error) {
	fileID := uuid.New()
	kind := file.KindImage

	bucket := storageClient.Bucket(config.GoogleCloudStorageBucket)
	object := bucket.Object(uploadFileName(fileID, kind))

	writer := object.NewWriter(ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, reader); err != nil {
		return nil, fmt.Errorf("could not upload file: %w", err)
	}

	return nil, nil
}

func uploadFileName(fileID uuid.UUID, kind file.Kind) string {
	return fmt.Sprintf("%v/%v", kind, fileID)
}
