module baristeuer

go 1.20

require baristeuer/internal/pdf v0.0.0

require (
	baristeuer/internal/data v0.0.0 // indirect
	github.com/jung-kurt/gofpdf v1.16.2 // indirect
)

replace baristeuer/internal/pdf => ../internal/pdf

replace baristeuer/internal/data => ../internal/data
