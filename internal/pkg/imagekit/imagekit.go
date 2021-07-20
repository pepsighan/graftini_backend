package imagekit

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/file"
	"github.com/pepsighan/graftini_backend/internal/pkg/storage"
)

// GetImageKitURLForFile gets the image kit url for the given file.
func GetImageKitURLForFile(imageKitURL string, fileID uuid.UUID, fileKind file.Kind) string {
	return fmt.Sprintf("%v/%v", imageKitURL, storage.UploadFileName(fileID, fileKind))
}
