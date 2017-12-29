package world

import (
	"math/rand"
	"time"
)

/*
 * TerrainType
 */
type TerrainType int
const (
	WaterTerrain TerrainType = 1 + iota
	GroundTerrain
)

var terrains = [...]string {"~", "."}

func (tt TerrainType) String() string {
	return terrains[tt - 1]
}

/*
 * Land
 */
type Located struct {
	X, Y int
}

const (
	LandDimention = 9
)

type View [][]string

type Land struct {
	Located
	tiles [][]string
}

func NewLand() *Land {
	var land = new(Land)
	land.Generate()
	return land
}

func (land *Land) Generate() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < LandDimention; i++ {
		land.tiles = append(land.tiles, []string{})

		for j := 0; j < LandDimention; j++ {
			land.tiles[i] = append(land.tiles[i], terrains[rand.Intn(len(terrains))])
		}
	}
}

func (land *Land) View() View {
	return land.tiles
}

func (land *Land) Update(x, y int, terrain TerrainType) {
	land.tiles[x % LandDimention][y % LandDimention] = terrain.String()
}

/*
 * World
 */
type LandColumn struct {
	storage map[int]*Land
}

func NewLandColumn() *LandColumn {
	column := new(LandColumn)
	column.storage = make(map[int]*Land)
	return column
}

func (column *LandColumn) Y(y int) *Land {
	var land *Land
	ok := false

	if land, ok = column.storage[y]; !ok {
		column.storage[y] = NewLand()
		land = column.storage[y]
	}

	return land
}

type LandGrid struct {
	storage map[int]*LandColumn
}

func NewLandGrid() *LandGrid {
	column := new(LandGrid)
	column.storage = make(map[int]*LandColumn)
	return column
}

func (grid *LandGrid) X(x int) *LandColumn {
	var column *LandColumn
	ok := false

	if column, ok = grid.storage[x]; !ok {
		grid.storage[x] = NewLandColumn()
		column = grid.storage[x]
	}

	return column
}

var lands *LandGrid

func GetLand(x, y int) *Land {
	return lands.X(x).Y(y)
}

// func GetView(x, y, dimetion int) View {
// 	landCount := dimention /
// }

func init() {
	lands = NewLandGrid()
}
