package filehandlers

import (
	"github.com/google/uuid"
	"sync"
)

type FileHandlerMgr struct {
	imageFileFormatReaders sync.Map
	imageFileFormatWriters sync.Map
}

type handlerNameDescription struct {
	name        string
	description string
}

var fileHandlerMgr *FileHandlerMgr

func init() {
	fileHandlerMgr = &FileHandlerMgr{
		sync.Map{},
		sync.Map{},
	}
}

func GetFileHandlerMgrInstance() *FileHandlerMgr {
	return fileHandlerMgr
}

func (i *FileHandlerMgr) RegisterFileType(
	reader ImageFileMetadataReader,
	writer ImageFileMetadataWriter) uuid.UUID {

	newHandlerId := uuid.New()
	if reader != nil {
		i.imageFileFormatReaders.Store(newHandlerId, reader)
	}
	if writer != nil {
		i.imageFileFormatWriters.Store(newHandlerId, writer)
	}
	return newHandlerId
}
