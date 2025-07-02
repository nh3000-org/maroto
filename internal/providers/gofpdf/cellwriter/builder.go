package cellwriter

import (
	"github.com/nh3000-org/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
)

type CellWriterBuilder struct{}

func NewBuilder() *CellWriterBuilder {
	return &CellWriterBuilder{}
}

func (c *CellWriterBuilder) Build(fpdf gofpdfwrapper.Fpdf) CellWriter {
	cellCreator := NewCellWriter(fpdf)
	borderColorStyle := NewBorderColorStyler(fpdf)
	borderLineStyler := NewBorderLineStyler(fpdf)
	borderThicknessStyler := NewBorderThicknessStyler(fpdf)
	fillColorStyler := NewFillColorStyler(fpdf)

	borderThicknessStyler.SetNext(borderLineStyler)
	borderLineStyler.SetNext(borderColorStyle)
	borderColorStyle.SetNext(fillColorStyler)
	fillColorStyler.SetNext(cellCreator)

	return borderThicknessStyler
}
