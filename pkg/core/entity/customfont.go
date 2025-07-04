package entity

import "github.com/nh3000-org/maroto/v2/pkg/consts/fontstyle"

// CustomFont is the representation of a font that can be added to the pdf.
type CustomFont struct {
	Family string
	Style  fontstyle.Type
	File   string
	Bytes  []byte
}
