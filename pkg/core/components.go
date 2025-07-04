package core

import (
	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/nh3000-org/maroto/v2/pkg/consts/extension"
	"github.com/nh3000-org/maroto/v2/pkg/consts/fontstyle"
	"github.com/nh3000-org/maroto/v2/pkg/core/entity"
	"github.com/nh3000-org/maroto/v2/pkg/props"
)

// Math is the abstraction which deals with useful calc.
type Math interface {
	GetInnerCenterCell(inner *entity.Dimensions, outer *entity.Dimensions) *entity.Cell
	Resize(inner *entity.Dimensions, outer *entity.Dimensions, percent float64, justReferenceWidth bool) *entity.Dimensions
}

// Code is the abstraction which deals of how to add QrCodes or Barcode in a PDF.
type Code interface {
	GenQr(code string) (*entity.Image, error)
	GenDataMatrix(code string) (*entity.Image, error)
	GenBar(code string, cell *entity.Cell, prop *props.Barcode) (*entity.Image, error)
}

// Image is the abstraction which deals of how to add images in a PDF.
type Image interface {
	Add(img *entity.Image, cell *entity.Cell, margins *entity.Margins, prop *props.Rect, extension extension.Type, flow bool) error
	GetImageInfo(img *entity.Image, extension extension.Type) (*gofpdf.ImageInfoType, uuid.UUID)
}

type Line interface {
	Add(cell *entity.Cell, prop *props.Line)
}

// Text is the abstraction which deals of how to add text inside PDF.
type Text interface {
	Add(text string, cell *entity.Cell, textProp *props.Text)
	GetLinesQuantity(text string, textProp *props.Text, colWidth float64) int
}

// Font is the abstraction which deals of how to set fontstyle configurations.
type Font interface {
	SetFamily(family string)
	SetStyle(style fontstyle.Type)
	SetSize(size float64)
	SetFont(family string, style fontstyle.Type, size float64)
	GetFamily() string
	GetStyle() fontstyle.Type
	GetSize() float64
	GetFont() (string, fontstyle.Type, float64)
	GetHeight(family string, style fontstyle.Type, size float64) float64
	SetColor(color *props.Color)
	GetColor() *props.Color
}
