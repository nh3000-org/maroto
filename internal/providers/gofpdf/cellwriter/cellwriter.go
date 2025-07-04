package cellwriter

import (
	"github.com/nh3000-org/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/nh3000-org/maroto/v2/pkg/consts/border"
	"github.com/nh3000-org/maroto/v2/pkg/core/entity"
	"github.com/nh3000-org/maroto/v2/pkg/props"
)

type CellWriter interface {
	SetNext(next CellWriter)
	GetNext() CellWriter
	GetName() string
	Apply(width, height float64, config *entity.Config, prop *props.Cell)
}

type cellWriter struct {
	stylerTemplate
	defaultColor *props.Color
}

func NewCellWriter(fpdf gofpdfwrapper.Fpdf) *cellWriter {
	return &cellWriter{
		stylerTemplate: stylerTemplate{
			fpdf: fpdf,
			name: "cellWriter",
		},
		defaultColor: &props.BlackColor,
	}
}

func (c *cellWriter) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	if prop == nil {
		bd := border.None
		if config.Debug {
			bd = border.Full
		}

		c.fpdf.CellFormat(width, height, "", string(bd), 0, "C", false, 0, "")
		return
	}

	bd := prop.BorderType
	if config.Debug {
		bd = border.Full
	}

	c.fpdf.CellFormat(width, height, "", string(bd), 0, "C", prop.BackgroundColor != nil, 0, "")
}
