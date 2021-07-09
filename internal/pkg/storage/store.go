package storage

import (
	"context"
	"errors"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/backend/config"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/file"
)

// ErrUnsupportedMimeType is returned when a file of unsupported mime type is uploaded.
var ErrUnsupportedMimeType = errors.New("unsupported_mime_type")

// UploadFile uploads the file to GCP Cloud Storage.
func UploadFile(ctx context.Context, reader io.Reader, mimeType string, storageClient *storage.Client, entClient *ent.Client) (*ent.File, error) {
	if err := checkIfSupportedImage(mimeType); err != nil {
		return nil, err
	}

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

// uploadFileName generates a file name for the object in the GCP bucket.
func uploadFileName(fileID uuid.UUID, kind file.Kind) string {
	return fmt.Sprintf("%v/%v", kind, fileID)
}

var supportedImageTypes = []string{
	"image/jpeg",
	"image/png",
	"image/webp",
}

// checkIfSupportedImage checks whether the uploaded file is actually an image and
// a supported one.
func checkIfSupportedImage(mimeType string) error {
	for _, supported := range supportedImageTypes {
		if supported != mimeType {
			return ErrUnsupportedMimeType
		}
	}

	return nil
}
