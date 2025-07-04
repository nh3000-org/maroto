package row_test

import (
	"testing"

	"github.com/nh3000-org/maroto/v2/internal/fixture"
	"github.com/nh3000-org/maroto/v2/mocks"
	"github.com/nh3000-org/maroto/v2/pkg/components/col"
	"github.com/nh3000-org/maroto/v2/pkg/components/row"
	"github.com/nh3000-org/maroto/v2/pkg/core"
	"github.com/nh3000-org/maroto/v2/pkg/core/entity"
	"github.com/nh3000-org/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("when there is no cols", func(t *testing.T) {
		// Act
		r := row.New(10)

		// Assert
		test.New(t).Assert(r.GetStructure()).Equals("components/rows/new_empty_col.json")
	})
	t.Run("when has component, should retrieve components", func(t *testing.T) {
		// Act
		r := row.New(12).Add(col.New(12))

		// Assert
		test.New(t).Assert(r.GetStructure()).Equals("components/rows/new_filled_col.json")
	})
	t.Run("when has prop, should apply correctly", func(t *testing.T) {
		// Act
		prop := fixture.CellProp()
		r := row.New(12).WithStyle(&prop)

		// Assert
		test.New(t).Assert(r.GetStructure()).Equals("components/rows/new_col_with_prop.json")
	})
}

func TestRow_GetHeight(t *testing.T) {
	t.Run("When a row has a column with height 5, should return 5", func(t *testing.T) {
		cell := fixture.CellEntity()

		provider := mocks.NewProvider(t)

		columns := mocks.NewCol(t)
		columns.EXPECT().GetHeight(provider, &cell).Return(5)

		// Act
		r := row.New().Add(columns)

		// Assert
		assert.Equal(t, 5.0, r.GetHeight(provider, &cell))
	})
}

func TestRow_GetColumns(t *testing.T) {
	t.Run("when GetColumns is called, should return the number of registered columns", func(t *testing.T) {
		// Act
		newCol := []core.Col{col.New(12)}

		r := row.New(10).Add(newCol[0])

		// Assert
		assert.Equal(t, newCol, r.GetColumns())
	})
}

func TestRow_GetStructure(t *testing.T) {
	t.Run("when there is no style, should call provider correctly", func(t *testing.T) {
		// Arrange
		cfg := &entity.Config{
			MaxGridSize: 12,
		}
		cell := fixture.CellEntity()

		provider := mocks.NewProvider(t)
		provider.EXPECT().CreateRow(cell.Height)

		col := mocks.NewCol(t)
		col.EXPECT().Render(provider, cell, true)
		col.EXPECT().SetConfig(cfg)
		col.EXPECT().GetSize().Return(12)

		sut := row.New(cell.Height).Add(col)
		sut.SetConfig(cfg)

		// Act
		sut.Render(provider, cell)

		// Assert
		provider.AssertNumberOfCalls(t, "CreateRow", 1)
		col.AssertNumberOfCalls(t, "Render", 1)
		col.AssertNumberOfCalls(t, "SetConfig", 1)
	})
	t.Run("when there is style, should call provider correctly", func(t *testing.T) {
		// Arrange
		cfg := &entity.Config{
			MaxGridSize: 12,
		}
		cell := fixture.CellEntity()
		prop := fixture.CellProp()

		provider := mocks.NewProvider(t)
		provider.EXPECT().CreateRow(cell.Height)
		provider.EXPECT().CreateCol(cell.Width, cell.Height, cfg, &prop)

		col := mocks.NewCol(t)
		col.EXPECT().Render(provider, cell, false)
		col.EXPECT().SetConfig(cfg)
		col.EXPECT().GetSize().Return(12)

		sut := row.New(cell.Height).Add(col).WithStyle(&prop)
		sut.SetConfig(cfg)

		// Act
		sut.Render(provider, cell)

		// Assert
		provider.AssertNumberOfCalls(t, "CreateCol", 1)
		provider.AssertNumberOfCalls(t, "CreateRow", 1)
		col.AssertNumberOfCalls(t, "Render", 1)
		col.AssertNumberOfCalls(t, "SetConfig", 1)
	})
}

func TestRow_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		sut := row.New(10)

		// Act
		sut.SetConfig(nil)
	})
}
