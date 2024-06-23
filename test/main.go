package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)
import "github.com/SixbucksSolutions/image2lrxmp"
import "github.com/akamensky/argparse"

type programOptions struct {
	inputImageFile *os.File
	outputXmlFile  string
}

func parseArgs() programOptions {
	parser := argparse.NewParser("image2lrxmp", "Create Lightroom-friendly XMP for image files")
	inputImageFile := parser.File("i", "inputimage", os.O_RDONLY, 0400, &argparse.Options{
		Required: true,
		Help:     "Input image file",
	})
	outputXmpFile := parser.String("o", "xmpfilename", &argparse.Options{
		Required: false,
		Help:     "XMP file name if user doesn't want default",
	})
	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
	}
	var xmpFilename string
	if outputXmpFile == nil {
		xmpFilename = strings.TrimSuffix(inputImageFile.Name(), filepath.Ext(inputImageFile.Name())) + ".xmp"
	} else {
		xmpFilename = *outputXmpFile
	}

	return programOptions{
		inputImageFile: inputImageFile,
		outputXmlFile:  xmpFilename,
	}
}

func main() {
	args := parseArgs()

	inputFileStat, err := os.Stat(args.inputImageFile.Name())
	if err != nil {
		panic("Could not run stat on our open input file?!?!?")
	}

	inputFileBytes := make([]byte, inputFileStat.Size())
	_, err = args.inputImageFile.Read(inputFileBytes)
	if err != nil {
		panic("Could not read bytes from our open input file?!?!?")
	}
	fmt.Println("\tRead input file bytes successfully")

	libHandle := image2lrxmp.MakeImage2LRXMP()

	handlerInfo := libHandle.DetectImageType(inputFileBytes, 1.0)

	//for _, currMatch := range handlerInfo {
	//	fmt.Printf("Match for handler %s: %.0f%%\n",
	//		currMatch.Handler.Name(), currMatch.Confidence*100.0)
	//}

	if len(handlerInfo) != 1 {
		panic("\tFATAL: should have gotten only one file match")
	}
	imageFileHandler := handlerInfo[0].Handler

	fmt.Println("\tWas able to identify file type and get handler")

	xmpDocument, err := imageFileHandler.CreateXmpMetadataFromImageBytes(inputFileBytes)

	if err != nil {
		panic("Could not extract XMP data")
	}

	if xmpDocument != nil {
		fmt.Printf("\tGot an xmp Document back")
	} else {
		fmt.Printf("\tNot a nil pointer and no document back")
	}

}
