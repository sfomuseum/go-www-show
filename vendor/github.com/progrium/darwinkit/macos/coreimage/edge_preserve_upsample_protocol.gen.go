// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure an edge preserve upsample filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample?language=objc
type PEdgePreserveUpsample interface {
	// optional
	SetSmallImage(value Image)
	HasSetSmallImage() bool

	// optional
	SmallImage() Image
	HasSmallImage() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetSpatialSigma(value float32)
	HasSetSpatialSigma() bool

	// optional
	SpatialSigma() float32
	HasSpatialSigma() bool

	// optional
	SetLumaSigma(value float32)
	HasSetLumaSigma() bool

	// optional
	LumaSigma() float32
	HasLumaSigma() bool
}

// ensure impl type implements protocol interface
var _ PEdgePreserveUpsample = (*EdgePreserveUpsampleObject)(nil)

// A concrete type for the [PEdgePreserveUpsample] protocol.
type EdgePreserveUpsampleObject struct {
	objc.Object
}

func (e_ EdgePreserveUpsampleObject) HasSetSmallImage() bool {
	return e_.RespondsToSelector(objc.Sel("setSmallImage:"))
}

// The image that the filter upsamples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228239-smallimage?language=objc
func (e_ EdgePreserveUpsampleObject) SetSmallImage(value Image) {
	objc.Call[objc.Void](e_, objc.Sel("setSmallImage:"), value)
}

func (e_ EdgePreserveUpsampleObject) HasSmallImage() bool {
	return e_.RespondsToSelector(objc.Sel("smallImage"))
}

// The image that the filter upsamples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228239-smallimage?language=objc
func (e_ EdgePreserveUpsampleObject) SmallImage() Image {
	rv := objc.Call[Image](e_, objc.Sel("smallImage"))
	return rv
}

func (e_ EdgePreserveUpsampleObject) HasSetInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228237-inputimage?language=objc
func (e_ EdgePreserveUpsampleObject) SetInputImage(value Image) {
	objc.Call[objc.Void](e_, objc.Sel("setInputImage:"), value)
}

func (e_ EdgePreserveUpsampleObject) HasInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228237-inputimage?language=objc
func (e_ EdgePreserveUpsampleObject) InputImage() Image {
	rv := objc.Call[Image](e_, objc.Sel("inputImage"))
	return rv
}

func (e_ EdgePreserveUpsampleObject) HasSetSpatialSigma() bool {
	return e_.RespondsToSelector(objc.Sel("setSpatialSigma:"))
}

// A value that specifies the influence of the input image’s spatial information on the upsampling operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228240-spatialsigma?language=objc
func (e_ EdgePreserveUpsampleObject) SetSpatialSigma(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setSpatialSigma:"), value)
}

func (e_ EdgePreserveUpsampleObject) HasSpatialSigma() bool {
	return e_.RespondsToSelector(objc.Sel("spatialSigma"))
}

// A value that specifies the influence of the input image’s spatial information on the upsampling operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228240-spatialsigma?language=objc
func (e_ EdgePreserveUpsampleObject) SpatialSigma() float32 {
	rv := objc.Call[float32](e_, objc.Sel("spatialSigma"))
	return rv
}

func (e_ EdgePreserveUpsampleObject) HasSetLumaSigma() bool {
	return e_.RespondsToSelector(objc.Sel("setLumaSigma:"))
}

// A value that specifies the influence of the input image’s luma information on the upsampling operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228238-lumasigma?language=objc
func (e_ EdgePreserveUpsampleObject) SetLumaSigma(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setLumaSigma:"), value)
}

func (e_ EdgePreserveUpsampleObject) HasLumaSigma() bool {
	return e_.RespondsToSelector(objc.Sel("lumaSigma"))
}

// A value that specifies the influence of the input image’s luma information on the upsampling operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228238-lumasigma?language=objc
func (e_ EdgePreserveUpsampleObject) LumaSigma() float32 {
	rv := objc.Call[float32](e_, objc.Sel("lumaSigma"))
	return rv
}
