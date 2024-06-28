package filehandlers

import (
	"github.com/SixbucksSolutions/image2lrxmp/imagemetadata"
	"github.com/evanoberholster/imagemeta"
	"io"
)

func (cr3 RawCanonCr3) ExtractImageMetadata(reader *io.ReadSeeker) *imagemetadata.ImageMetadata {

	// Create "reader" wrapper around our bytes to meet imagemeta interface
	exifInfo, err := imagemeta.Decode(*reader)
	if err != nil {
		panic("Could not read metadata")
	}

	metadata := imagemetadata.ImageMetadata{}

	metadata.Metadata["camera-serial-number"] = exifInfo.CameraSerial

	//fmt.Println(imageInfo)

	// Check out if this package will make our lives easier vs our current
	//		key/value pair
	//	https://github.com/dsoprea/go-exif

	// These functions are interesting -- it'll find EXIF block inside arbitrary
	//		bytes.
	//		Wonder if there's anything that the XMP needs that is NOT in EXIF, but
	//		some arbitrary bullshit?
	// https://pkg.go.dev/github.com/dsoprea/go-exif/v3#pkg-functions

	// Lots of github stars and forks. Like so many
	// https://github.com/rwcarlsen/goexif9i

	// What can we populate in the key-value store we send down to XMP writer?
	return &metadata
}
