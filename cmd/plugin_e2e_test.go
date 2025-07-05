package main

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"baristeuer/internal/service"
)

func TestExamplePlugin(t *testing.T) {
	tmpDir := t.TempDir()
	soPath := filepath.Join(tmpDir, "example.so")

	build := exec.Command("go", "build", "-buildmode=plugin", "-o", soPath, filepath.Join("..", "plugins", "example"))
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build plugin: %v\n%s", err, out)
	}

	ds, err := service.NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()
	proj, err := ds.CreateProject(ctx, "Demo")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := ds.AddIncome(ctx, proj.ID, "donation", 10); err != nil {
		t.Fatal(err)
	}

	exportPath := filepath.Join(tmpDir, "out.json")
	os.Setenv("EXAMPLE_EXPORT_FILE", exportPath)
	defer os.Unsetenv("EXAMPLE_EXPORT_FILE")

	loadPlugins(tmpDir, ds)

	data, err := os.ReadFile(exportPath)
	if err != nil {
		t.Fatalf("read export: %v", err)
	}
	var arr []any
	if err := json.Unmarshal(data, &arr); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if len(arr) == 0 {
		t.Fatal("export empty")
	}
}
