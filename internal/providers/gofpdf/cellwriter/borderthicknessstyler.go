package cellwriter

import (
	"github.com/nh3000-org/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/nh3000-org/maroto/v2/pkg/consts/linestyle"
	"github.com/nh3000-org/maroto/v2/pkg/core/entity"
	"github.com/nh3000-org/maroto/v2/pkg/props"
)

type borderThicknessStyler struct {
	stylerTemplate
	defaultLineThickness float64
}

func NewBorderThicknessStyler(fpdf gofpdfwrapper.Fpdf) *borderThicknessStyler {
	return &borderThicknessStyler{
		stylerTemplate: stylerTemplate{
			fpdf: fpdf,
			name: "borderThicknessStyler",
		},
		defaultLineThickness: linestyle.DefaultLineThickness,
	}
}

func (b *borderThicknessStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	if prop == nil {
		b.GoToNext(width, height, config, prop)
		return
	}

	if prop.BorderThickness == 0 {
		b.GoToNext(width, height, config, prop)
		return
	}

	b.fpdf.SetLineWidth(prop.BorderThickness)
	b.GoToNext(width, height, config, prop)
	b.fpdf.SetLineWidth(b.defaultLineThickness)
}
