package martinez_rueda

import (
	"reflect"
	"testing"

	"github.com/paulmach/orb"
)

func TestCompute(t *testing.T) {

	// point0s := []orb.Point{{0, 0}, {2, 0}, {2, 2}, {0, 2}, {0, 0}}
	// contour0 := NewContour(point0s)
	// polygon0 := NewPolygon([]Contour{contour0})

	// point1s := []orb.Point{{1, 1}, {3, 1}, {3, 4}, {1, 4}, {1, 1}}
	// contour1 := NewContour(point1s)
	// polygon1 := NewPolygon([]Contour{contour1})

	point2s := []orb.Point{{0, 0}, {2, 0}, {2, 1}, {3, 1}, {3, 4}, {1, 4}, {1, 2}, {0, 2}, {0, 0}}
	contour2 := NewContour(point2s)
	polygon2 := NewPolygon([]Contour{contour2})

	//

	point01s := []orb.Point{
		{
			111.4892578125,
			35.65729624809628},
		{
			113.818359375,
			35.65729624809628},
		{
			113.818359375,
			37.71859032558816},
		{
			111.4892578125,
			37.71859032558816},
		{
			111.4892578125,
			35.65729624809628}}
	contour0 := NewContour(point01s)
	polygon0 := NewPolygon([]Contour{contour0})

	point1s := []orb.Point{
		{
			115.77392578125,
			35.746512259918504}, {
			118.30078125,
			35.746512259918504}, {
			118.30078125,
			37.64903402157866}, {
			115.77392578125,
			37.64903402157866}, {
			115.77392578125,
			35.746512259918504}}
	contour1 := NewContour(point1s)
	polygon1 := NewPolygon([]Contour{contour1})

	type args struct {
		subject   *Polygon
		clipping  *Polygon
		operation OPERATION
	}

	argStruct := args{
		subject:   polygon0,
		clipping:  polygon1,
		operation: OP_UNION,
	}

	tests := []struct {
		name       string
		args       args
		wantResult *Polygon
	}{{
		name:       "测试union",
		args:       argStruct,
		wantResult: polygon2,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Compute(tt.args.subject, tt.args.clipping, tt.args.operation); !reflect.DeepEqual(gotResult, tt.wantResult) {
				byte, err := gotResult.ToPolygonGeometry().MarshalJSON()
				if err != nil {
					t.Log("报错")
				} else {
					t.Errorf("Compute() = %v, want %v", gotResult, tt.wantResult)
					geojson := string(byte)
					t.Log(geojson)
				}
			}
		})
	}
}
