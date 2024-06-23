package filehandlers

type RawCanonCr3 struct {
}

func (cr3 RawCanonCr3) Name() string {
	return "Raw.Canon.Cr3"
}

func (cr3 RawCanonCr3) Description() string {
	return ""
}

func (cr3 RawCanonCr3) RunDetector([]byte) float32 {
	return 1.0
}

func (cr3 RawCanonCr3) ExtractXmpData([]byte) string {
	return ""
}
