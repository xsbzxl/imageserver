package jpeg

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/pierrre/imageserver"
	imageserver_processor_native "github.com/pierrre/imageserver/processor/native"
)

//Encoder encodes a native Image to a raw Image in jpeg.
type Encoder struct {
}

//Encode encodes the image.
func (e *Encoder) Encode(nim image.Image, params imageserver.Params) (*imageserver.Image, error) {
	opts, err := getOptions(params)
	if err != nil {
		return nil, err
	}
	return encodeImage(nim, opts)
}

func getOptions(params imageserver.Params) (*jpeg.Options, error) {
	quality, err := getQuality(params)
	if err != nil {
		return nil, err
	}
	return &jpeg.Options{
		Quality: quality,
	}, nil
}

func getQuality(params imageserver.Params) (int, error) {
	quality, _ := params.GetInt("quality")
	if quality == 0 {
		return jpeg.DefaultQuality, nil
	}
	if quality < 1 {
		return 0, &imageserver.ParamError{Param: "quality", Message: "must be greater than or equal to 1"}
	}
	if quality > 100 {
		return 0, &imageserver.ParamError{Param: "quality", Message: "must be less than or equal to 100"}
	}
	return quality, nil
}

func encodeImage(nativeImage image.Image, opts *jpeg.Options) (*imageserver.Image, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, nativeImage, opts)
	if err != nil {
		return nil, err
	}
	return &imageserver.Image{
		Format: "jpeg",
		Data:   buf.Bytes(),
	}, nil
}

func init() {
	imageserver_processor_native.RegisterEncoder("jpeg", &Encoder{})
}
