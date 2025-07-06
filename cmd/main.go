package main

import (
	"baristeuer/internal/config"
	"baristeuer/internal/data"
	"baristeuer/internal/pdf"
	"baristeuer/internal/plugins"
	"baristeuer/internal/service"
	"context"
	"flag"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"os"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"
)

// loadPlugins attempts to load all Go plugins from the given directory.
// Each plugin must export a `New` function returning a plugins.Plugin
// implementation. Any errors are printed but do not stop startup.
func loadPlugins(dir string, svc *service.DataService) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".so" {
			continue
		}
		p, err := plugin.Open(filepath.Join(dir, e.Name()))
		if err != nil {
			fmt.Println("load plugin", e.Name(), "error:", err)
			continue
		}
		sym, err := p.Lookup("New")
		if err != nil {
			fmt.Println("plugin", e.Name(), "missing New symbol:", err)
			continue
		}
		newFunc, ok := sym.(func() plugins.Plugin)
		if !ok {
			fmt.Println("plugin", e.Name(), "has invalid New signature")
			continue
		}
		plg := newFunc()
		if err := plg.Init(svc); err != nil {
			fmt.Println("plugin", e.Name(), "init error:", err)
		}
	}
}

func main() {
	cfgPath := flag.String("config", "config.json", "configuration file")
	dbPath := flag.String("db", "", "SQLite database path")
	pdfDir := flag.String("pdfdir", "", "directory for generated PDFs")
	logFile := flag.String("logfile", "", "log file path")
	logLevel := flag.String("loglevel", "", "log level")
	exportPath := flag.String("exportdb", "", "export database to path and exit")
	restorePath := flag.String("restoredb", "", "restore database from path and exit")
	exportCSV := flag.String("exportcsv", "", "export project CSV as <id>:<file> and exit")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	if *dbPath != "" {
		cfg.DBPath = *dbPath
	}
	if *pdfDir != "" {
		cfg.PDFDir = *pdfDir
	}
	if *logFile != "" {
		cfg.LogFile = *logFile
	}
	if *logLevel != "" {
		cfg.LogLevel = *logLevel
	}

	if cfg.PDFDir == "" {
		cfg.PDFDir = filepath.Join(".", config.DefaultPDFDir)
	}

	if cfg.DBPath == "" {
		cfg.DBPath = "baristeuer.db"
	}

	store, err := data.NewStore(cfg.DBPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	logger, logCloser := service.NewLogger(cfg.LogFile, cfg.LogLevel, cfg.LogFormat)
	generator := pdf.NewGenerator(cfg.PDFDir, store, &cfg)
	datasvc := service.NewDataServiceFromStore(store, logger, logCloser, &cfg)
	defer datasvc.Close()

	// Load optional runtime plugins from ./plugins if available.
	loadPlugins("plugins", datasvc)

	if *exportPath != "" {
		if err := datasvc.ExportDatabase(*exportPath); err != nil {
			fmt.Println("Error exporting database:", err)
		}
		return
	}

	if *exportCSV != "" {
		parts := strings.SplitN(*exportCSV, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid -exportcsv value, expected <projectID>:<file>")
			return
		}
		id, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Println("invalid project id in -exportcsv:", err)
			return
		}
		if err := datasvc.ExportProjectCSV(context.Background(), id, parts[1]); err != nil {
			fmt.Println("Error exporting project:", err)
		}
		return
	}

	if *restorePath != "" {
		if err := datasvc.RestoreDatabase(*restorePath); err != nil {
			fmt.Println("Error restoring database:", err)
		}
		return
	}

	err = wails.Run(&options.App{
		Title:       "Baristeuer",
		AssetServer: &assetserver.Options{},
		Bind:        []interface{}{generator, datasvc},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
