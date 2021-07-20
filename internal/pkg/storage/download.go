package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"github.com/pepsighan/graftini_backend/internal/backend/config"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"golang.org/x/oauth2/google"
)

// FileURL gets the download URL for the given file.
func FileURL(ctx context.Context, file *ent.File, storageClient *storage.Client) (string, error) {
	cf, err := defaultCredentialsFile(ctx)
	if err != nil {
		return "", fmt.Errorf("could not find credentials to get file url: %w", err)
	}

	url, err := storage.SignedURL(
		config.GoogleCloudStorageBucket,
		UploadFileName(file.ID, file.Kind),
		&storage.SignedURLOptions{
			Method:         "GET",
			GoogleAccessID: cf.ClientEmail,
			PrivateKey:     []byte(cf.PrivateKey),
			// Expire the signed URL in 2 days. That should be enough time to view on the UI.
			// If it expires, the user may need to refresh the page.
			Expires: time.Now().Add(48 * time.Hour),
		},
	)

	if err != nil {
		return "", fmt.Errorf("could not get the file url: %w", err)
	}

	return url, nil
}

type credentialsFile struct {
	ClientEmail  string `json:"client_email"`
	ClientID     string `json:"client_id"`
	PrivateKey   string `json:"private_key"`
	PrivateKeyID string `json:"private_key_id"`
	ProjectID    string `json:"project_id"`
}

// defaultCredentialsFile loads the credentials file from the environment.
func defaultCredentialsFile(ctx context.Context) (*credentialsFile, error) {
	creds, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		return nil, err
	}

	cf := new(credentialsFile)
	if err := json.Unmarshal(creds.JSON, &cf); err != nil {
		return nil, err
	}

	return cf, nil
}
