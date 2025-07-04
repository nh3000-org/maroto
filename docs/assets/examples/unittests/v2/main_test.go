package main_test

import (
	"testing"

	"github.com/nh3000-org/maroto/v2/pkg/components/line"

	"github.com/nh3000-org/maroto/v2"

	"github.com/nh3000-org/maroto/v2/pkg/components/code"
	"github.com/nh3000-org/maroto/v2/pkg/components/image"
	"github.com/nh3000-org/maroto/v2/pkg/components/signature"
	"github.com/nh3000-org/maroto/v2/pkg/components/text"
	"github.com/nh3000-org/maroto/v2/pkg/consts/extension"
	"github.com/nh3000-org/maroto/v2/pkg/test"
)

func TestMaroto_GetStructure(t *testing.T) {
	// Arrange
	m := maroto.New()

	m.AddRow(10,
		code.NewBarCol(4, "barcode"),
		code.NewMatrixCol(4, "matrixcode"),
		code.NewQrCol(4, "qrcode"),
	)

	m.AddRow(10,
		image.NewFromFileCol(3, "image"),
		image.NewFromBytesCol(3, []byte{0, 1, 2}, extension.Png),
		signature.NewCol(3, "signature"),
		text.NewCol(3, "text"),
	)

	m.AddRow(10, line.NewCol(12))

	// Assert
	test.New(t).Assert(m.GetStructure()).Equals("example_unit_test.json")
}
