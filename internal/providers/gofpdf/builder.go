package gofpdf

import (
	"github.com/jung-kurt/gofpdf"

	"github.com/nh3000-org/maroto/v2/internal/cache"
	"github.com/nh3000-org/maroto/v2/internal/code"
	"github.com/nh3000-org/maroto/v2/internal/math"
	"github.com/nh3000-org/maroto/v2/internal/providers/gofpdf/cellwriter"
	"github.com/nh3000-org/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/nh3000-org/maroto/v2/pkg/core"
	"github.com/nh3000-org/maroto/v2/pkg/core/entity"
)

// Dependencies is the dependencies provider for gofpdf
type Dependencies struct {
	Fpdf       gofpdfwrapper.Fpdf
	Font       core.Font
	Text       core.Text
	Code       core.Code
	Image      core.Image
	Line       core.Line
	Cache      cache.Cache
	CellWriter cellwriter.CellWriter
	Cfg        *entity.Config
}

// Builder is the dependencies builder for gofpdf
type Builder interface {
	Build(cfg *entity.Config, cache cache.Cache) *Dependencies
}

type builder struct{}

// NewBuilder create a new Builder
func NewBuilder() *builder {
	return &builder{}
}

// Build create a new Dependencies
func (b *builder) Build(cfg *entity.Config, cache cache.Cache) *Dependencies {
	fpdf := gofpdfwrapper.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		FontDirStr:     "",
		Size: gofpdf.SizeType{
			Wd: cfg.Dimensions.Width,
			Ht: cfg.Dimensions.Height,
		},
	})

	for _, font := range cfg.CustomFonts {
		fpdf.AddUTF8FontFromBytes(font.Family, string(font.Style), font.Bytes)
	}

	if cfg.DisableAutoPageBreak {
		fpdf.SetAutoPageBreak(false, 0)
	} else {
		fpdf.SetAutoPageBreak(true, cfg.Margins.Bottom)
	}

	fpdf.SetMargins(cfg.Margins.Left, cfg.Margins.Top, cfg.Margins.Right)
	fpdf.AddPage()

	font := NewFont(fpdf, cfg.DefaultFont.Size, cfg.DefaultFont.Family, cfg.DefaultFont.Style)
	math := math.New()
	code := code.New()
	text := NewText(fpdf, math, font)
	image := NewImage(fpdf, math)
	line := NewLine(fpdf)
	cellWriter := cellwriter.NewBuilder().
		Build(fpdf)

	return &Dependencies{
		Fpdf:       fpdf,
		Font:       font,
		Text:       text,
		Code:       code,
		Image:      image,
		Line:       line,
		CellWriter: cellWriter,
		Cfg:        cfg,
		Cache:      cache,
	}
}
