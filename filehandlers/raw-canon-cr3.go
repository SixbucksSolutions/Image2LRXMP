package filehandlers

import (
	"fmt"
)

type RawCanonCr3 struct {
}

var rawCanonCr3Instance RawCanonCr3

func init() {
	//fmt.Println("Inside CR3 init")
	rawCanonCr3Instance = RawCanonCr3{}

	fileHandlerMgr = FileHandlerMgrInstance()

	fileExtensionHints := []string{"cr3"}

	if err := fileHandlerMgr.RegisterFileType(rawCanonCr3Instance, nil, fileExtensionHints); err != nil {
		fmt.Println("Error in registering RawCanonCr3: ", err)
	}
}

func (RawCanonCr3) Name() string {
	return "Raw.Canon.Cr3"
}

func (RawCanonCr3) Description() string {
	return ""
}
