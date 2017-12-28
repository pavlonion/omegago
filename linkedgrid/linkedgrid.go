package linkedgrid

type Side int
const (
	LeftSide = iota + 1
	RightSide
	UpSide
	DownSide
)


func Oposide(side Side) Side {
	oposideMap := map[Side]Side{
		LeftSide: RightSide,
		RightSide: LeftSide,
		UpSide: DownSide,
		DownSide: UpSide,
	}
	return oposideMap[side]
}

type LinkedGrid struct {
	Data interface{}
	neighbors map[Side]*LinkedGrid
}

func Nil() *LinkedGrid {
	return nil
}

func New(data interface{}) *LinkedGrid {
	var grid = new(LinkedGrid)
	grid.Data = data
	initNeiboursMap(grid)
	return grid
}

func initNeiboursMap(grid *LinkedGrid) {
	if grid.neighbors == nil {
		grid.neighbors = make(map[Side]*LinkedGrid)
	}
}

func (grid *LinkedGrid) Left() *LinkedGrid {
	return grid.neighbors[LeftSide]
}

func (grid *LinkedGrid) Right() *LinkedGrid {
	return grid.neighbors[RightSide]
}

func (grid *LinkedGrid) Up() *LinkedGrid {
	return grid.neighbors[UpSide]
}

func (grid *LinkedGrid) Down() *LinkedGrid {
	return grid.neighbors[DownSide]
}

func (grid *LinkedGrid) BindTo(other *LinkedGrid, side Side) {
	if other == nil {
		return
	}

	initNeiboursMap(grid)
	initNeiboursMap(other)
	other.neighbors[side] = grid
	grid.neighbors[Oposide(side)] = other

	switch side {
		case LeftSide:
			if other.Up() != nil {
				grid.neighbors[UpSide] = other.Up().Right()
			}

			if other.Down() != nil {
				grid.neighbors[DownSide] = other.Down().Right()
			}

		case RightSide:
			if other.Up() != nil {
				grid.neighbors[UpSide] = other.Up().Left()
			}

			if other.Down() != nil {
				grid.neighbors[DownSide] = other.Down().Left()
			}

		case UpSide:
			if other.Right() != nil {
				grid.neighbors[RightSide] = other.Right().Down()
			}

			if other.Left() != nil {
				grid.neighbors[LeftSide] = other.Left().Down()
			}

		case DownSide:
			if other.Right() != nil {
				grid.neighbors[RightSide] = other.Right().Up()
			}

			if other.Left() != nil {
				grid.neighbors[LeftSide] = other.Left().Up()
			}
	}
}
