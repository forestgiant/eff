package grid

import (
	"math"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/util"
)

// Grid eff.Shape container that places children in a grid pattern
type Grid struct {
	eff.Shape
	rows       int
	cols       int
	padding    int
	cellHeight int
}

func (g *Grid) rectForIndex(index int) eff.Rect {
	roundDivide := func(v1 int, v2 int) int {
		return util.RoundToInt(float64(v1) / float64(v2))
	}

	r := eff.Rect{}
	if g.rows == 0 || g.cols == 0 {
		return r
	}

	row := index / g.cols
	col := index % g.cols
	cellWidth := roundDivide(g.Rect().W-(g.padding*(g.cols+1)), g.cols)
	cellHeight := roundDivide(g.Rect().H-(g.padding*(g.rows+1)), g.rows)
	if g.cellHeight > 0 {
		cellHeight = g.cellHeight
	}
	if cellWidth <= 0 {
		cellWidth = 1
	}

	if cellHeight <= 0 {
		cellHeight = 1
	}

	r.X = col*cellWidth + (g.padding * (col + 1))
	r.Y = row*cellHeight + (g.padding * (row + 1))
	r.W = cellWidth
	r.H = cellHeight

	return r
}

func (g *Grid) updateGrid() {
	if len(g.Children()) == 0 || g.rows == 0 || g.cols == 0 {
		return
	}

	for i, c := range g.Children() {
		c.SetRect(g.rectForIndex(i))
	}
}

func (g *Grid) updateRect() {
	if g.cellHeight > 0 && g.rows > 0 {
		rowCount := int(math.Ceil(float64((len(g.Children()) + 1)) / float64(g.cols)))
		newHeight := rowCount*(g.cellHeight) + ((rowCount + 1) * g.padding)

		g.SetRect(eff.Rect{
			X: g.Rect().X,
			Y: g.Rect().Y,
			W: g.Rect().W,
			H: newHeight,
		})
		g.RedrawChildren()
	}
}

// AddChild adds a child to the grid and places it in the correct spot
func (g *Grid) AddChild(c eff.Drawable) error {
	g.updateRect()
	c.SetRect(g.rectForIndex(len(g.Children())))
	return g.Shape.AddChild(c)
}

// RemoveChild removes a child from the grid
func (g *Grid) RemoveChild(c eff.Drawable) error {
	err := g.Shape.RemoveChild(c)
	if err != nil {
		return err
	}
	g.updateRect()
	g.updateGrid()

	return nil
}

// SetRect sets the rectangle for the grid
func (g *Grid) SetRect(r eff.Rect) {
	if r.W != g.Rect().W || r.H != g.Rect().H {
		g.updateGrid()
	}

	g.Shape.SetRect(r)
}

// Rows returns the row count of the grid
func (g *Grid) Rows() int {
	return g.rows
}

// SetRows sets the row count for the grid, updates all children
func (g *Grid) SetRows(r int) {
	g.rows = r
	g.updateGrid()
}

// Cols returns the column count of the grid
func (g *Grid) Cols() int {
	return g.cols
}

// SetCols sets the column count for the grid, updates all children
func (g *Grid) SetCols(c int) {
	g.cols = c
	g.updateGrid()
}

// NewGrid creates a new grid instance, the cellHeight is an optional override of the derived cellHeight (grid.Rect().H/grid.Rows()), this will grow the grid to hold all children
func NewGrid(rows int, cols int, padding int, cellHeight int) *Grid {
	g := &Grid{}
	g.rows = rows
	g.cols = cols
	g.padding = padding
	g.cellHeight = cellHeight
	g.SetBackgroundColor(eff.White())
	return g
}
