// nolint: dupl
package cellwriter

import (
	"github.com/nh3000-org/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/nh3000-org/maroto/v2/pkg/core/entity"
	"github.com/nh3000-org/maroto/v2/pkg/props"
)

type borderColorStyler struct {
	stylerTemplate
	defaultColor *props.Color
}

func NewBorderColorStyler(fpdf gofpdfwrapper.Fpdf) *borderColorStyler {
	return &borderColorStyler{
		stylerTemplate: stylerTemplate{
			fpdf: fpdf,
			name: "borderColorStyler",
		},
		defaultColor: &props.BlackColor,
	}
}

func (b *borderColorStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	if prop == nil {
		b.GoToNext(width, height, config, prop)
		return
	}

	if prop.BorderColor == nil {
		b.GoToNext(width, height, config, prop)
		return
	}

	b.fpdf.SetDrawColor(prop.BorderColor.Red, prop.BorderColor.Green, prop.BorderColor.Blue)
	b.GoToNext(width, height, config, prop)
	b.fpdf.SetDrawColor(b.defaultColor.Red, b.defaultColor.Green, b.defaultColor.Blue)
}
