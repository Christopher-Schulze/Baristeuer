package csvupload

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"baristeuer/internal/service"
)

func TestPluginInit_NoURL(t *testing.T) {
	dir := t.TempDir()
	logFile := filepath.Join(dir, "log.txt")
	logger, closer := service.NewLogger(logFile, "info", "text")
	defer closer.Close()
	ds, err := service.NewDataService(":memory:", logger, closer, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	plg := newWithInterval("", 0).(*Plugin)
	err = plg.Init(ds)
	if err == nil {
		t.Fatalf("expected error")
	}
	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "no endpoint configured") {
		t.Fatalf("expected log entry, got: %s", data)
	}
}

func TestUploadAll_UploadsPerProject(t *testing.T) {
	ds, err := service.NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()
	ctx := context.Background()

	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("Proj%d", i)
		if _, err := ds.CreateProject(ctx, name); err != nil {
			t.Fatal(err)
		}
	}

	var uploads int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&uploads, 1)
		body, _ := io.ReadAll(r.Body)
		if len(body) == 0 {
			t.Error("empty body")
		}
		if ct := r.Header.Get("Content-Type"); ct != "text/csv" {
			t.Errorf("unexpected content type %s", ct)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	plg := newWithInterval(server.URL, 0).(*Plugin)
	plg.uploadAll(ds)

	if uploads != 2 {
		t.Fatalf("expected 2 uploads, got %d", uploads)
	}
}
