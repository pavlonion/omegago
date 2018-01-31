package world

import (
	"fmt"
	"math"
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
	LandDimention = 8
)

type View [][]string

func (view View) String() string {
	var result string;

	for _, row := range view {
		for _, cell := range row {
			result += cell + " "
		}

		result += "\n"
	}

	return result
}

type Land struct {
	Located
	tiles View
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
	land.tiles[y % LandDimention][x % LandDimention] = terrain.String()
}

/*
 * World
 */
type LandRow struct {
	storage map[int]*Land
}

func NewLandRow() *LandRow {
	column := new(LandRow)
	column.storage = make(map[int]*Land)
	return column
}

func (column *LandRow) X(x int) *Land {
	var land *Land
	ok := false

	if land, ok = column.storage[x]; !ok {
		column.storage[x] = NewLand()
		land = column.storage[x]
	}

	return land
}

type LandGrid struct {
	storage map[int]*LandRow
}

func NewLandGrid() *LandGrid {
	column := new(LandGrid)
	column.storage = make(map[int]*LandRow)
	return column
}

func (grid *LandGrid) Y(y int) *LandRow {
	var column *LandRow
	ok := false

	if column, ok = grid.storage[y]; !ok {
		grid.storage[y] = NewLandRow()
		column = grid.storage[y]
	}

	return column
}

var lands *LandGrid

func GetLand(x, y int) *Land {
	return lands.Y(y).X(x)
}

func ceilDiv(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}

func floorDiv(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}

func GetScreen(landX, landY, dimention int) View {
	if dimention < LandDimention {
		panic(fmt.Sprintf("dimention param should be greater then LandDimention = %d", LandDimention))
	}

	landCount := ceilDiv(dimention, LandDimention)

	totalDimention := landCount * LandDimention
	if totalDimention - dimention < 2 {
		landCount += 1
		totalDimention = landCount * LandDimention
	}

	middle := floorDiv(landCount, 2)
	totalCenter := floorDiv(totalDimention, 2)
	viewStart := totalCenter - ceilDiv(dimention, 2)
	viewEnd := totalCenter + ceilDiv(dimention, 2)

	result := make(View, dimention)

	for landShiftY := 0; landShiftY < landCount; landShiftY++ {
		for landShiftX := 0; landShiftX < landCount; landShiftX++ {
			view := GetLand(landX + landShiftY - middle, landY + landShiftX - middle).View()

			for y := 0; y < LandDimention; y++ {
				totalY := landShiftY * LandDimention + y

				for x := 0; x < LandDimention; x++ {
					totalX := landShiftX * LandDimention + x

					if totalX >= viewStart && totalX < viewEnd && totalY >= viewStart && totalY < viewEnd {
						resultY := totalY - viewStart
						result[resultY] = append(result[resultY], view[y][x])
					}
				}
			}
		}
	}

	return result;
}

func init() {
	lands = NewLandGrid()
}
