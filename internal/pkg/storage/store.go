package storage

import (
	"context"
	"errors"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/file"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
)

// ErrUnsupportedMimeType is returned when a file of unsupported mime type is uploaded.
var ErrUnsupportedMimeType = errors.New("unsupported_mime_type")

// UploadFile uploads the file to GCP Cloud Storage.
func UploadFile(
	ctx context.Context,
	reader io.Reader,
	mimeType string,
	bucketName string,
	storageClient *storage.Client,
	entClient *ent.Client) (*ent.File, error) {

	if err := checkIfSupportedImage(mimeType); err != nil {
		return nil, err
	}

	fileID := uuid.New()
	kind := file.KindImage

	bucket := storageClient.Bucket(bucketName)
	object := bucket.Object(UploadFileName(fileID, kind))

	writer := object.NewWriter(ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, reader); err != nil {
		return nil, logger.Errorf("could not upload file: %w", err)
	}

	fileRecord, err := entClient.File.Create().
		SetID(fileID).
		SetKind(kind).
		SetMimeType(mimeType).
		Save(ctx)
	if err != nil {
		return nil, logger.Errorf("could not save the record for uploaded file: %w", err)
	}

	return fileRecord, nil
}

// UploadFileName generates a file name for the object in the GCP bucket.
func UploadFileName(fileID uuid.UUID, kind file.Kind) string {
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
		if supported == mimeType {
			return nil
		}
	}

	return logger.Error(ErrUnsupportedMimeType)
}
