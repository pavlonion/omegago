package world

import "testing"

func equal(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("%v != %v", a, b)
	}
}

func TestTerrainType(t *testing.T) {
	equal(t, WaterTerrain.String(), "~")
	equal(t, GroundTerrain.String(), ".")
}

func TestLand(t *testing.T) {
	lc := NewLandColumn()
	println(lc.Y(0))

	lg := NewLandGrid()
	println(lg.X(0))
	println(lg.X(0).Y(12))
	println(lg.X(0).Y(12))
}
