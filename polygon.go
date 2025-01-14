package martinez_rueda

import (
	"fmt"
	"math"
	"strings"

	geojson "github.com/paulmach/go.geojson"
	"github.com/paulmach/orb"
)

type Polygon struct {
	contours []Contour
}

//  Get array of contours (each is array of points and each point is 2-size array)
func NewPolygon(contours_xy []Contour) *Polygon {
	pol := Polygon{
		contours: []Contour{},
	}
	for _, contour_xy := range contours_xy {
		contourPoints := []orb.Point{}
		for _, xy := range contour_xy.points {
			contourPoints = append(contourPoints, xy)
		}
		pol.pushBack(NewContour(contourPoints))
	}

	return &pol
}

func (p *Polygon) contour(index int) Contour {
	return p.contours[index]
}

func (p *Polygon) ncontours() int {
	return len(p.contours)
}

func (p *Polygon) nvertices() int {
	nv := 0
	for idx := 0; idx < len(p.contours); idx++ {
		nv += len(p.contours[idx].points)
	}

	return nv
}

// Get minimum bounding rectangle
func (p *Polygon) getBoundingBox() []orb.Point {

	minX := math.Inf(1)
	minY := math.Inf(1)
	maxX := math.Inf(-1)
	maxY := math.Inf(-1)

	for idx := 0; idx < len(p.contours); idx++ {
		box := p.contours[idx].getBoundingBox()

		minTmp := box[0]
		maxTmp := box[1]

		if minTmp.X() < minX {
			minX = minTmp.X()
		}

		if maxTmp.X() > maxX {
			maxX = maxTmp.X()
		}

		if minTmp.Y() < minY {
			minY = minTmp.Y()
		}

		if maxTmp.Y() > maxY {
			maxY = maxTmp.Y()
		}
	}

	return []orb.Point{orb.Point{minX, minY}, orb.Point{maxX, maxY}}
}

func (p *Polygon) move(x, y float64) {
	for idx := 0; idx < len(p.contours); idx++ {
		p.contours[idx].move(x, y)
	}
}

func (p *Polygon) pushBack(contour Contour) {
	p.contours = append(p.contours, contour)
}

// Pop the element off the end of array
func (p *Polygon) popBack() {
	p.contours = p.contours[:(len(p.contours) - 1)]
}

func (p *Polygon) erase(index int) {
	//         unset($this->points[$index]);
	p.contours = append(p.contours[:index], p.contours[(index+1):]...)
}

func (p *Polygon) clear() {
	p.contours = []Contour{}
}

func (p *Polygon) ToPolygonGeometry() *geojson.Geometry {
	g := geojson.Geometry{}
	multiPolygon := g.MultiPolygon

	for _, con := range p.contours {
		con.clockwise()

		line := g.LineString
		for _, point := range con.points {
			line = append(line, []float64{point[0], point[1]})
		}

		if con.cc { //外边框
			polygon := g.Polygon
			polygon = [][][]float64{line}

			multiPolygon = append(multiPolygon, polygon)
		} else { //内边框
			//理论上来说已经存在一个面了
			polygon := multiPolygon[len(multiPolygon)-1]
			polygon = append(polygon, line)

			multiPolygon = multiPolygon[:len(multiPolygon)-1]
			multiPolygon = append(multiPolygon, polygon)
		}
	}
	return geojson.NewMultiPolygonGeometry(multiPolygon...)
}

func (p *Polygon) DEBUG() {
	var data strings.Builder
	for _, con := range p.contours {
		data.WriteString("[")
		for _, point := range con.points {
			data.WriteString(fmt.Sprintf("[%v,%v],", point.Lon(), point.Lat()))
		}
		data.WriteString("]")
	}
	fmt.Println(strings.Replace(data.String(), "],]", "]]", -1))
}
