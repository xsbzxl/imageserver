package native

import (
	"testing"

	imageserver_processor "github.com/pierrre/imageserver/processor"
)

func TestProcessorInterface(t *testing.T) {
	var _ imageserver_processor.Processor = &Processor{}
}

func TestProcessorNativeFuncInterface(t *testing.T) {
	var _ ProcessorNative = ProcessorNativeFunc(nil)
}

func TestEncoderFuncInterface(t *testing.T) {
	var _ Encoder = EncoderFunc(nil)
}
