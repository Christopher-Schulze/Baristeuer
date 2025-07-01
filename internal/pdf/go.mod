module baristeuer/internal/pdf

go 1.20

require (
	baristeuer/internal/data v0.0.0
	github.com/jung-kurt/gofpdf v1.16.2
)

replace baristeuer/internal/data => ../data
