package filehandlers

import (
	"fmt"
	"strings"
	"sync"
)

type FileHandlerMgr struct {
	imageFileFormatReaders sync.Map
	imageFileFormatWriters sync.Map
}

var fileHandlerMgr *FileHandlerMgr

func init() {
	fileHandlerMgr = &FileHandlerMgr{
		sync.Map{},
		sync.Map{},
	}
}

func FileHandlerMgrInstance() *FileHandlerMgr {
	return fileHandlerMgr
}

func (i *FileHandlerMgr) RegisterFileType(
	reader ImageFileMetadataReader,
	writer ImageFileMetadataWriter,
	fileExtensionHints []string) error {
	if reader != nil {
		for _, fileExtension := range fileExtensionHints {
			// See if we already know about this file type?
			lcExtension := strings.ToLower(fileExtension)

			val, ok := i.imageFileFormatReaders.Load(lcExtension)

			if !ok {
				readerArray := []ImageFileMetadataReader{reader}
				i.imageFileFormatReaders.Store(fileExtension, readerArray)
			} else {
				i.imageFileFormatReaders.Store(fileExtension,
					append(val.([]ImageFileMetadataReader), reader))
			}
		}

		fmt.Printf("Successfully registered handler \"%s\" to read files with extensions %v\n",
			reader.Name(), fileExtensionHints)
	}
	if writer != nil {
		fmt.Printf("Attempted to register handler \"%s\" to write files with extensions %v\n",
			writer.Name(), fileExtensionHints)
	}

	return nil
}
