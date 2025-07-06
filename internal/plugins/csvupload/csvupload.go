package csvupload

import (
    "context"
    "fmt"
    "io"
    "net/http"
    "os"
    "time"

    "baristeuer/internal/plugins"
    "baristeuer/internal/service"
)

// Plugin periodically exports all projects as CSV files and uploads them.
type Plugin struct {
    endpoint string
    interval time.Duration
}

// New creates a plugin using the endpoint defined by the CSVUPLOAD_URL
// environment variable. The interval is one week.
func New() plugins.Plugin {
    return newWithInterval(os.Getenv("CSVUPLOAD_URL"), 7*24*time.Hour)
}

// newWithInterval creates the plugin with a custom interval. Used in tests.
func newWithInterval(endpoint string, interval time.Duration) plugins.Plugin {
    return &Plugin{endpoint: endpoint, interval: interval}
}

// Init starts the periodic upload task.
func (p *Plugin) Init(ds *service.DataService) error {
    if p.endpoint == "" {
        return fmt.Errorf("csvupload: no endpoint configured")
    }
    ticker := time.NewTicker(p.interval)
    go func() {
        for range ticker.C {
            p.uploadAll(ds)
        }
    }()
    service.Logger().Info("csvupload plugin initialized", "endpoint", p.endpoint)
    return nil
}

func (p *Plugin) uploadAll(ds *service.DataService) {
    ctx := context.Background()
    projects, err := ds.ListProjects()
    if err != nil {
        service.Logger().Error("csvupload list projects", "err", err)
        return
    }
    for _, prj := range projects {
        tmp, err := os.CreateTemp("", fmt.Sprintf("project_%d_*.csv", prj.ID))
        if err != nil {
            service.Logger().Error("csvupload temp", "err", err)
            continue
        }
        tmp.Close()
        if err := ds.ExportProjectCSV(ctx, prj.ID, tmp.Name()); err != nil {
            service.Logger().Error("csvupload export", "project", prj.ID, "err", err)
            os.Remove(tmp.Name())
            continue
        }
        if err := uploadFile(ctx, p.endpoint, tmp.Name()); err != nil {
            service.Logger().Error("csvupload upload", "project", prj.ID, "err", err)
        } else {
            service.Logger().Info("csv uploaded", "project", prj.ID)
        }
        os.Remove(tmp.Name())
    }
}

func uploadFile(ctx context.Context, url, path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()
    req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, f)
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", "text/csv")
    resp, err := http.DefaultClient.Do(req)
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

