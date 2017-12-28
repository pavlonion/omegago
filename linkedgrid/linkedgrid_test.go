package linkedgrid

import "testing"

func equal(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("%v != %v", a, b)
	}
}

/*	5	6	7	8
	4	10		9
	3	2	1	0
 */

func TestBindGrid(t *testing.T) {
	const count = 11
	data := [count]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
	var grid [count]*LinkedGrid

	for i := 0; i < count; i++ {
		grid[i] = New(data[i])
		equal(t, grid[i].Data, data[i])
	}

	grid[1].BindTo(grid[0], LeftSide)
	grid[2].BindTo(grid[1], LeftSide)
	grid[3].BindTo(grid[2], LeftSide)
	grid[4].BindTo(grid[3], UpSide)
	grid[5].BindTo(grid[4], UpSide)
	grid[6].BindTo(grid[5], RightSide)
	grid[7].BindTo(grid[6], RightSide)
	grid[8].BindTo(grid[7], RightSide)
	grid[9].BindTo(grid[8], DownSide)
	grid[10].BindTo(grid[6], DownSide)

	equal(t, grid[0].Left(), grid[1])
	equal(t, grid[0].Up(), grid[9])
	equal(t, grid[0].Right(), Nil())
	equal(t, grid[0].Down(), Nil())

	equal(t, grid[1].Left(), grid[2])
	equal(t, grid[1].Up(), Nil())
	equal(t, grid[1].Right(), grid[0])
	equal(t, grid[1].Down(), Nil())

	equal(t, grid[2].Left(), grid[3])
	equal(t, grid[2].Up(), grid[10])
	equal(t, grid[2].Right(), grid[1])
	equal(t, grid[2].Down(), Nil())
}
