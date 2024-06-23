package filehandlers

import (
	"bytes"
	"github.com/evanoberholster/imagemeta"
	"github.com/trimmer-io/go-xmp/xmp"
)

type RawCanonCr3 struct {
}

func (cr3 RawCanonCr3) Name() string {
	return "Raw.Canon.Cr3"
}

func (cr3 RawCanonCr3) Description() string {
	return ""
}

func (cr3 RawCanonCr3) RunDetector([]byte) float32 {
	return 1.0
}

func (cr3 RawCanonCr3) CreateXmpMetadataFromImageBytes(imageBytes []byte) (*xmp.Document, error) {
	// Create "reader" wrapper around our bytes to meet imagemeta interface
	bytesReader := bytes.NewReader(imageBytes)
	_, err := imagemeta.Decode(bytesReader)

	myXmp := xmp.Document{}

	return &myXmp, err
}
