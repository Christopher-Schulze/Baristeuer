package sync

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"baristeuer/internal/cloud"
)

// Client defines upload and download operations for the database.
type Client interface {
	Upload(ctx context.Context, src string) error
	Download(ctx context.Context, dest string) error
}

// LocalClient is a placeholder implementation storing files locally.
type LocalClient struct{ BaseDir string }

// NewLocalClient returns a LocalClient using the given directory.
func NewLocalClient(dir string) *LocalClient {
	if dir == "" {
		dir = "syncdata"
	}
	return &LocalClient{BaseDir: dir}
}

func copyFile(src, dest string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

// Upload copies src into the client's base directory.
func (c *LocalClient) Upload(ctx context.Context, src string) error {
	dest := filepath.Join(c.BaseDir, filepath.Base(src))
	return copyFile(src, dest)
}

// Download copies the stored file into dest.
func (c *LocalClient) Download(ctx context.Context, dest string) error {
	src := filepath.Join(c.BaseDir, filepath.Base(dest))
	return copyFile(src, dest)
}

// RemoteClient wraps a cloud.Client implementing the Client interface.
type RemoteClient struct{ c *cloud.Client }

// NewRemoteClient constructs a RemoteClient for the given endpoints and token.
func NewRemoteClient(uploadURL, downloadURL, token string) *RemoteClient {
	return &RemoteClient{c: cloud.NewClient(uploadURL, downloadURL, token)}
}

func (c *RemoteClient) Upload(ctx context.Context, src string) error {
	return c.c.Upload(ctx, src)
}

func (c *RemoteClient) Download(ctx context.Context, dest string) error {
	return c.c.Download(ctx, dest)
}
