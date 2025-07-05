# Plugin Development Guide

Baristeuer supports optional runtime plugins written in Go. Plugins are built as
shared objects and loaded from the `plugins/` directory next to the executable on
startup.

## Writing a Plugin

A plugin must implement the `plugins.Plugin` interface located in
`internal/plugins`:

```go
package plugins

type Plugin interface {
    Init(*service.DataService) error
}
```

Create a new module in `plugins/<name>` with a `New` function returning your
implementation. The exported `Init` method receives a `DataService` instance and
can register additional functionality or perform tasks during start up.

Example structure:

```text
plugins/example/
├─ go.mod
├─ example.go
```

Build the plugin using the Go compiler with `-buildmode=plugin`:

```bash
go build -buildmode=plugin -o plugins/example/example.so ./plugins/example
```

After placing the compiled `.so` file in the `plugins/` directory, Baristeuer
loads it automatically on the next start.

## Example Plugin

The repository contains a simple exporter plugin in `plugins/example` which
writes all projects, incomes and expenses to a JSON file. The output path can be
configured via the `EXAMPLE_EXPORT_FILE` environment variable.
