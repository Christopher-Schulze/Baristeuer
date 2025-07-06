# CSV Upload Plugin

This plugin periodically exports all projects as CSV files and uploads them to a configurable HTTP endpoint.

## Build

```bash
# from the repository root
go build -tags=plugin -buildmode=plugin -o plugins/csvupload.so ./internal/plugins/csvupload/plugin
```

## Usage

Set the environment variable `CSVUPLOAD_URL` to the upload target and place the resulting `csvupload.so` file in the `plugins/` directory next to the application binary. The plugin will upload project CSV exports once per week.

