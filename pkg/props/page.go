package props

import (
	"fmt"
	"strings"

	"github.com/nh3000-org/maroto/v2/pkg/consts/align"
	"github.com/nh3000-org/maroto/v2/pkg/consts/breakline"
	"github.com/nh3000-org/maroto/v2/pkg/consts/fontstyle"
)

// Place is the representation of a place in a page.
type Place string

const (
	// LeftTop is the place in the left top of the page.
	LeftTop Place = "left_top"
	// Top is the place in the top of the page.
	Top Place = "top"
	// RightTop is the place in the right top of the page.
	RightTop Place = "right_top"
	// LeftBottom is the place in the left bottom of the page.
	LeftBottom Place = "left_bottom"
	// Bottom is the place in the bottom of the page.
	Bottom Place = "bottom"
	// RightBottom is the place in the right bottom of the page.
	RightBottom Place = "right_bottom"
)

// IsValid checks if the place is valid.
func (p Place) IsValid() bool {
	return p == LeftTop || p == Top || p == RightTop ||
		p == LeftBottom || p == Bottom || p == RightBottom
}

// PageNumber have attributes of page number
type PageNumber struct {
	// Pattern is the string pattern which will be used to apply the page count component.
	Pattern string
	// Place defines where the page count component will be placed.
	Place Place
	// Family defines which font family will be applied to page count.
	Family string
	// Style defines which font style will be applied to page count.
	Style fontstyle.Type
	// Size defines which font size will be applied to page count.
	Size float64
	// Color defines which will be applied to page count.
	Color *Color
}

// GetNumberTextProp returns the Text properties of the page number.
func (p *PageNumber) GetNumberTextProp(height float64) *Text {
	text := &Text{
		Family: p.Family,
		Style:  p.Style,
		Size:   p.Size,
		Color:  p.Color,
		Align:  align.Center,
	}

	if p.Place == LeftBottom || p.Place == LeftTop {
		text.Align = align.Left
	} else if p.Place == RightBottom || p.Place == RightTop {
		text.Align = align.Right
	}

	if p.Place == RightBottom || p.Place == Bottom || p.Place == LeftBottom {
		text.Top = height
	}

	text.BreakLineStrategy = breakline.EmptySpaceStrategy

	return text
}

// GetPageString returns the page string.
func (p *PageNumber) GetPageString(current, total int) string {
	pattern := strings.ReplaceAll(p.Pattern, "{current}", fmt.Sprintf("%d", current))
	return strings.ReplaceAll(pattern, "{total}", fmt.Sprintf("%d", total))
}

// WithFont apply font if not defined before.
func (p *PageNumber) WithFont(font *Font) {
	if p.Color == nil {
		p.Color = font.Color
	}

	if p.Size == 0 {
		p.Size = font.Size
	}

	if p.Style == "" {
		p.Style = font.Style
	}

	if p.Family == "" {
		p.Family = font.Family
	}
}

// AppendMap appends the font fields to a map.
func (p *PageNumber) AppendMap(m map[string]interface{}) map[string]interface{} {
	if p.Pattern != "" {
		m["page_number_pattern"] = p.Pattern
	}

	if p.Place != "" {
		m["page_number_place"] = p.Place
	}

	if p.Family != "" {
		m["page_number_family"] = p.Family
	}

	if p.Style != "" {
		m["page_number_style"] = p.Style
	}

	if p.Size != 0 {
		m["page_number_size"] = p.Size
	}

	if p.Color != nil {
		m["page_number_color"] = p.Color.ToString()
	}

	return m
}
