package gofpdfwrapper_test

import (
	"fmt"
	"testing"

	"github.com/jung-kurt/gofpdf"
	"github.com/nh3000-org/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/stretchr/testify/assert"
)

func TestNewCustom(t *testing.T) {
	// Act
	sut := gofpdfwrapper.NewCustom(&gofpdf.InitType{})

	// Assert
	assert.NotNil(t, "", fmt.Sprintf("%T", sut))
}
