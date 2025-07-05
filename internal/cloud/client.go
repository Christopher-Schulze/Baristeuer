package cloud

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Client provides methods to upload and download files via HTTPS.
type Client struct {
	UploadURL   string
	DownloadURL string
	Token       string
	HTTPClient  *http.Client
}

// NewClient creates a Client with the given endpoints and token.
func NewClient(uploadURL, downloadURL, token string) *Client {
	return &Client{
		UploadURL:   uploadURL,
		DownloadURL: downloadURL,
		Token:       token,
		HTTPClient:  &http.Client{},
	}
}

// Upload sends the file at src to the configured upload URL using HTTPS.
func (c *Client) Upload(ctx context.Context, src string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.UploadURL, f)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed: %s", string(body))
	}
	return nil
}

// Download retrieves the remote file and stores it at dest.
func (c *Client) Download(ctx context.Context, dest string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.DownloadURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("download failed: %s", string(body))
	}

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}
