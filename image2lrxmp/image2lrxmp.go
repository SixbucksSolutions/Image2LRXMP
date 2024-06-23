package image2lrxmp

import (
	"fmt"
	"github.com/google/uuid"
	"image2lrxmp/image2lrxmp/filehandlers"
)

type Image2LRXMP struct {
	imageFileFormatHandlers map[uuid.UUID]ImageFileFormatHandler
}

type ImageFileFormatHandler interface {
	Name() string
	Description() string
	RunDetector([]byte) float32
	RunParser([]byte) string
}

type HandlerMatch struct {
	Handler    ImageFileFormatHandler
	Confidence float32
}

func MakeImage2LRXMP() Image2LRXMP {
	myLib := Image2LRXMP{
		imageFileFormatHandlers: make(map[uuid.UUID]ImageFileFormatHandler),
	}

	// Add known handlers -- how can we scan for these?
	myLib.imageFileFormatHandlers[uuid.New()] = filehandlers.RawCanonCr3{}
	return myLib
}
func (i Image2LRXMP) DetectImageType(bytes []byte, minConfidenceForMatch float32) []HandlerMatch {

	fmt.Println("DetectImageType")
	var matchingHandlers []HandlerMatch

	for _, handler := range i.imageFileFormatHandlers {
		// See how well we match
		thisMatch := handler.RunDetector(bytes)
		if thisMatch >= minConfidenceForMatch {
			matchingHandlers = append(matchingHandlers, HandlerMatch{handler, thisMatch})
		}
	}

	return matchingHandlers
}
