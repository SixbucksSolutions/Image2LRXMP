package filehandlers

import (
	"fmt"
	"github.com/google/uuid"
)

type RawCanonCr3 struct {
	handlerUuid uuid.UUID
}

func init() {
	fmt.Println("Inside CR3 init")

	fileHandlerMgr = GetFileHandlerMgrInstance()

	fileHandlerMgr.RegisterFileType(RawCanonCr3{}, nil)
}

func (i RawCanonCr3) Name() string {
	return "Raw.Canon.Cr3"
}

func (i RawCanonCr3) Description() string {
	return ""
}
