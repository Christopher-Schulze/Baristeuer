package sync

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"baristeuer/internal/config"
)

func TestRemoteClient_UploadDownload(t *testing.T) {
	uploadCalled := false
	downloadContent := []byte("dbdata")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer token" {
			t.Fatalf("missing auth header")
		}
		if r.Method == http.MethodPost {
			uploadCalled = true
			body, _ := io.ReadAll(r.Body)
			if string(body) != "test" {
				t.Fatalf("unexpected upload body: %s", body)
			}
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method == http.MethodGet {
			w.Write(downloadContent)
			return
		}
	}))
	defer srv.Close()

	f := "testfile"
	os.WriteFile(f, []byte("test"), 0o644)
	defer os.Remove(f)

	cfg := &config.Config{CloudUploadURL: srv.URL, CloudDownloadURL: srv.URL, CloudToken: "token"}
	c := NewRemoteClientFromConfig(cfg)
	ctx := context.Background()
	if err := c.Upload(ctx, f); err != nil {
		t.Fatalf("upload error: %v", err)
	}
	if !uploadCalled {
		t.Fatal("upload handler not called")
	}

	dest := "out"
	defer os.Remove(dest)
	if err := c.Download(ctx, dest); err != nil {
		t.Fatalf("download error: %v", err)
	}
	data, _ := os.ReadFile(dest)
	if string(data) != string(downloadContent) {
		t.Fatalf("unexpected downloaded data: %s", data)
	}
}
