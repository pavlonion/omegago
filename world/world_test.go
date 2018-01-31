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
	lc := NewLandRow()
	println(lc.X(0))

	lg := NewLandGrid()
	println(lg.Y(0))
	println(lg.Y(0).X(12))
	println(lg.Y(0).X(12))
}

func TestGetView(t *testing.T) {
	println(GetScreen(10, 10, 16).String());
	println(GetLand(10, 10).View().String());
	println(GetScreen(10, 10, LandDimention).String());
}
