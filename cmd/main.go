package main

import (
	"fmt"

	"baristeuer/internal/pdf"
)

func main() {
	fmt.Println("Baristeuer CLI")

	if path, err := pdf.GenerateReport("demo"); err != nil {
		fmt.Println("error generating report:", err)
	} else {
		fmt.Println("report generated at", path)
	}
}
