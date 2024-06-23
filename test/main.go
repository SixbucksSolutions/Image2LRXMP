package main

import "fmt"
import "github.com/SixbucksSolutions/image2lrxmp"

func main() {
	libHandle := image2lrxmp.MakeImage2LRXMP()
	myImageBytes := make([]byte, 0)
	handlerInfo := libHandle.DetectImageType(myImageBytes, 1.0)

	for _, currMatch := range handlerInfo {
		fmt.Printf("Match for handler %s: %.0f%%\n",
			currMatch.Handler.Name(), currMatch.Confidence*100.0)
	}
}
