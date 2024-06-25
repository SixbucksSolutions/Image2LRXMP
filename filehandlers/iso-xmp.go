package filehandlers

import (
	"fmt"
	"os"
)

type IsoXmpFileHandler struct {
	version string
}

/* Example Adobe Camera Raw (ACR)-generated XMP file:

<x:xmpmeta xmlns:x="adobe:ns:meta/" x:xmptk="Adobe XMP Core 7.0-c000 1.000000, 0000/00/00-00:00:00        ">
 <rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
  <rdf:Description rdf:about=""
    xmlns:xmp="http://ns.adobe.com/xap/1.0/"
    xmlns:tiff="http://ns.adobe.com/tiff/1.0/"
    xmlns:exif="http://ns.adobe.com/exif/1.0/"
    xmlns:dc="http://purl.org/dc/elements/1.1/"
    xmlns:aux="http://ns.adobe.com/exif/1.0/aux/"
    xmlns:exifEX="http://cipa.jp/exif/1.0/"
    xmlns:photoshop="http://ns.adobe.com/photoshop/1.0/"
    xmlns:xmpMM="http://ns.adobe.com/xap/1.0/mm/"
    xmlns:stEvt="http://ns.adobe.com/xap/1.0/sType/ResourceEvent#"
    xmlns:crd="http://ns.adobe.com/camera-raw-defaults/1.0/"
    xmlns:crs="http://ns.adobe.com/camera-raw-settings/1.0/"
   xmp:ModifyDate="2022-12-16T16:12:51.94+01:00"
   xmp:CreateDate="2022-12-16T16:12:51.94+01:00"
   xmp:MetadataDate="2023-09-23T16:47:49-04:00"
   tiff:Make="Canon"
   tiff:Model="Canon EOS R5"
   tiff:Orientation="1"
   tiff:ImageWidth="8192"
   tiff:ImageLength="5464"
   exif:ExifVersion="0231"
   exif:ExposureTime="1/1600"
   exif:ShutterSpeedValue="10643856/1000000"
   exif:FNumber="4/1"
   exif:ApertureValue="4/1"
   exif:ExposureProgram="0"
   exif:SensitivityType="2"
   exif:RecommendedExposureIndex="100"
   exif:ExposureBiasValue="0/1"
   exif:MaxApertureValue="4/1"
   exif:MeteringMode="5"
   exif:FocalLength="35/1"
   exif:CustomRendered="0"
   exif:ExposureMode="0"
   exif:WhiteBalance="0"
   exif:SceneCaptureType="0"
   exif:FocalPlaneXResolution="8192000/1419"
   exif:FocalPlaneYResolution="5464000/947"
   exif:FocalPlaneResolutionUnit="2"
   exif:DateTimeOriginal="2022-12-16T16:12:51.94+01:00"
   exif:PixelXDimension="8192"
   exif:PixelYDimension="5464"
   exif:GPSLatitude="22,52.837992N"
   exif:GPSLongitude="109,54.440808W"
   exif:GPSVersionID="2.2.0.0"
   exif:GPSAltitude="31400/10000"
   dc:format="image/x-canon-cr3"
   aux:SerialNumber="212024002620"
   aux:LensInfo="14/1 35/1 0/0 0/0"
   aux:Lens="RF14-35mm F4 L IS USM"
   aux:LensID="61182"
   aux:LensSerialNumber="0703002254"
   aux:ImageNumber="0"
   aux:ApproximateFocusDistance="992/100"
   aux:FlashCompensation="0/1"
   aux:OwnerName="JP YIM"
   aux:Firmware="1.6.0"
   exifEX:LensModel="RF14-35mm F4 L IS USM"
   photoshop:DateCreated="2022-12-16T16:12:51.94+01:00"
   photoshop:SidecarForExtension="CR3"
   photoshop:EmbeddedXMPDigest="120B7B9CF9D08EF4F483A77E5B485EA7"
   xmpMM:DocumentID="7BFEAEE81957339901D4F40FA35AB854"
   xmpMM:PreservedFileName="_A0A0001.CR3"
   xmpMM:OriginalDocumentID="7BFEAEE81957339901D4F40FA35AB854"
   xmpMM:InstanceID="xmp.iid:f5a13b37-a05c-e74e-a033-19100a76c4c8"
   crd:CameraProfile="Camera Standard"
   crd:LookName=""
   crs:CropTop="0"
   crs:CropLeft="0"
   crs:CropBottom="1"
   crs:CropRight="1"
   crs:CropAngle="0"
   crs:CropConstrainToWarp="0"
   crs:RawFileName="_A0A0001.CR3">
   <exif:ISOSpeedRatings>
    <rdf:Seq>
     <rdf:li>100</rdf:li>
    </rdf:Seq>
   </exif:ISOSpeedRatings>
   <exif:Flash
    exif:Fired="False"
    exif:Return="0"
    exif:Mode="0"
    exif:Function="False"
    exif:RedEyeMode="False"/>
   <xmpMM:History>
    <rdf:Seq>
     <rdf:li
      stEvt:action="saved"
      stEvt:instanceID="xmp.iid:f5a13b37-a05c-e74e-a033-19100a76c4c8"
      stEvt:when="2023-09-23T16:47:49-04:00"
      stEvt:softwareAgent="Adobe Photoshop Lightroom Classic 12.5 (Windows)"
      stEvt:changed="/metadata"/>
    </rdf:Seq>
   </xmpMM:History>
  </rdf:Description>
 </rdf:RDF>
</x:xmpmeta>


*/

type Tag_x_xmpmeta struct {
	Attr_xmlns_X string
	Attr_x_xmptk string

	Tag_rdf Tag_rdf_RDF
}

type Tag_rdf_RDF struct {
}

var requiredMetadataKeys = []string{
	"xmp:CreateDate",         // XMP Line 29
	"tiff:Make",              // XMP Line 31
	"tiff:Model",             // XMP Line 32
	"tiff:Orientation",       // XMP Line 33
	"tiff:ImageWidth",        // XMP Line 34
	"tiff:ImageLength",       // XMP Line 35
	"exif:ExifVersion",       // XMP Line 36
	"exif:ExposureTime",      // XMP Line 37
	"exif:ShutterSpeedValue", // XMP Line 38
	"exif:FNumber",           // XMP Line 39
	"exif:ApertureValue",     // XMP Line 40
	"exif:ExposureProgram",   // XMP Line 41
}

func (xg IsoXmpFileHandler) validateMetadataKeyList(imageMetadata map[string]string) error {
	// Create a COPY of the incoming map as we're gonna remove values to make our lives easier
	metadataCopy := make(map[string]string)
	for k, v := range imageMetadata {
		metadataCopy[k] = v
	}

	// Test required keys
	var missingKeys []string
	for _, currRequiredKey := range requiredMetadataKeys {
		if _, ok := metadataCopy[currRequiredKey]; !ok {
			missingKeys = append(missingKeys, currRequiredKey)
			continue
		}

		// We know this key was found, so remove it from copy
		delete(metadataCopy, currRequiredKey)
	}
	if len(missingKeys) > 0 {
		return fmt.Errorf("Metadata is missing one or more required metadata keys: %v", missingKeys)
	}

	// Make sure no EXTRA keys
	if len(metadataCopy) > 0 {
		var keysThatShouldNotExist []string
		for key := range metadataCopy {
			keysThatShouldNotExist = append(keysThatShouldNotExist, key)
		}

		return fmt.Errorf("Metadata has metadata keys it should not: %v", keysThatShouldNotExist)
	}

	fmt.Printf("\tImage metadata had exact set of required keys needed!")

	return nil
}

func (xg IsoXmpFileHandler) validateMetadata(imageMetadata map[string]string) error {
	fmt.Println("\tValidating incoming metadata is correct before writing")
	if err := xg.validateMetadataKeyList(imageMetadata); err != nil {
		return err
	}
	return nil
}

func (xg IsoXmpFileHandler) WriteImageMetadataToXmp(imageMetadata map[string]string, outputFile *os.File) error {
	fmt.Println("\tAsked to write image metadata to XMP")
	if err := xg.validateMetadata(imageMetadata); err != nil {
		return err
	}

	return nil
}
