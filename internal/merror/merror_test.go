package merror_test

import (
	"testing"

	"github.com/nh3000-org/maroto/v2/internal/merror"

	"github.com/nh3000-org/maroto/v2/pkg/consts/fontfamily"
	"github.com/nh3000-org/maroto/v2/pkg/consts/fontstyle"
	"github.com/stretchr/testify/assert"
)

func TestDefaultErrorText(t *testing.T) {
	// Assert
	assert.Equal(t, fontfamily.Arial, merror.DefaultErrorText.Family)
	assert.Equal(t, fontstyle.Bold, merror.DefaultErrorText.Style)
	assert.Equal(t, 10.0, merror.DefaultErrorText.Size)
	assert.Equal(t, 255, merror.DefaultErrorText.Color.Red)
	assert.Equal(t, 0, merror.DefaultErrorText.Color.Green)
	assert.Equal(t, 0, merror.DefaultErrorText.Color.Blue)
}
