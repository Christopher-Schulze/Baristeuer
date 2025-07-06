package cloud

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestClientUploadDownload(t *testing.T) {
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

	c := NewClient(srv.URL, srv.URL, "token")
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

func TestClientErrorsReturned(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"error":"boom"}`)
	}))
	defer srv.Close()

	f := "testfile"
	os.WriteFile(f, []byte("test"), 0o644)
	defer os.Remove(f)

	c := NewClient(srv.URL, srv.URL, "token")
	ctx := context.Background()
	if err := c.Upload(ctx, f); err == nil || err.Error() != "upload failed: boom" {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := c.Download(ctx, "out"); err == nil || err.Error() != "download failed: boom" {
		t.Fatalf("unexpected error: %v", err)
	}
}
