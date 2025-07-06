package plugins

import (
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"strings"
	"testing"

	"baristeuer/internal/service"
)

func TestLoggingPlugin(t *testing.T) {
	dir := t.TempDir()
	pluginPath := filepath.Join(dir, "logging.so")
	t.Log("plugin path", pluginPath)

	cmd := exec.Command("go", "build", "-tags=plugin", "-buildmode=plugin", "-trimpath", "-buildvcs=false", "-o", pluginPath, "./internal/plugins/logging/plugin")
	cmd.Env = append(os.Environ(), "GOFLAGS=")
	cmd.Dir = filepath.Join("..", "..")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("build plugin: %v\n%s", err, out)
	}

	logFile := filepath.Join(dir, "log.txt")
	logger, closer := service.NewLogger(logFile, "info", "text")
	defer closer.Close()

	ds, err := service.NewDataService(":memory:", logger, closer, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	p, err := plugin.Open(pluginPath)
	if err != nil {
		t.Skipf("open plugin failed: %v", err)
	}
	sym, err := p.Lookup("New")
	if err != nil {
		t.Fatalf("lookup New: %v", err)
	}
	newFunc := sym.(func() Plugin)
	plg := newFunc()
	if err := plg.Init(ds); err != nil {
		t.Fatalf("plugin init: %v", err)
	}

	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("read log: %v", err)
	}
	if !strings.Contains(string(data), "logging plugin initialized") {
		t.Fatalf("log output missing: %s", data)
	}
}
