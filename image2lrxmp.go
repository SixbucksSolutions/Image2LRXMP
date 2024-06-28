package image2lrxmp

import (
	"github.com/SixbucksSolutions/image2lrxmp/filehandlers"
	"github.com/SixbucksSolutions/image2lrxmp/imagemetadata"
	"os"
)

type Image2LRXMP struct {
	fileHandlerMgr *filehandlers.FileHandlerMgr
}

func NewImage2LRXMP() *Image2LRXMP {
	return &Image2LRXMP{
		fileHandlerMgr: filehandlers.FileHandlerMgrInstance(),
	}
}

func (Image2LRXMP) ReadImageMetadata(filePath string) *imagemetadata.ImageMetadata {
	// Try to read an image
	imageFile, err := os.OpenFile(filePath, os.O_RDONLY, 0400)
	if err != nil {
		panic("Could not open image file")
	}
	defer imageFile.Close()

	filehandlers.FileHandlerMgrInstance()

	return nil
}
