package image2lrxmp

import (
	"fmt"
	"github.com/SixbucksSolutions/image2lrxmp/filehandlers"
)

type Image2LRXMP struct {
	fileHandlerMgr *filehandlers.FileHandlerMgr
}

func NewImage2LRXMP() *Image2LRXMP {
	return &Image2LRXMP{
		fileHandlerMgr: filehandlers.GetFileHandlerMgrInstance(),
	}
}

func (Image2LRXMP) PrintHello() {
	fmt.Println("Hello")
}
