package test

import (
	"testing"
	// "time"

	"github.com/lanceInGz/pdfcpu/pkg/api"
)

// go test -timeout 3600s -run ^Test_PdfView$ github.com/lanceInGz/pdfcpu/gf

func Test_PdfView(t *testing.T) {
	msg := "Test_PdfView ERR::"
	inFile := "./fid_E047BD63EB0B413EBBDF3627ED331C2B.pdf"

	if _, err := api.InfoFile(inFile, nil, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}
}
