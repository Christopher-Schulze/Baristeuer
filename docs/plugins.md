# Building and Using Plugins

Bari$teuer supports optional runtime plugins. Build a plugin with
`go build -buildmode=plugin` and place the resulting `.so` file in the
`plugins/` directory next to the application binary.

Example:

```bash
# from the repository root
go build -tags=plugin -buildmode=plugin -o plugins/logging.so ./internal/plugins/logging/plugin
```

At startup all `.so` files in `plugins/` are loaded automatically.
