package obj

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/oakmound/oak/alg/floatgeom"
)

var vNullIndex = int64(-1)

var vertexReadTests = []struct {
	Items  stringList
	Error  string
	Vertex Vertex
}{
	{stringList{"1", "1", "1" /*-----------------------*/}, "" /*-------------------------------*/, Vertex{vNullIndex, floatgeom.Point3{1, 1, 1}}},
	{stringList{"1", "1" /*----------------------------*/}, "item length is incorrect" /**/, Vertex{vNullIndex, floatgeom.Point3{0, 0, 0}}},
	{stringList{"1.000000", "-1.000000", "-1.000000" /**/}, "" /*-------------------------------*/, Vertex{vNullIndex, floatgeom.Point3{1, -1, -1}}},
	{stringList{"0.999999", "-1.000000", "-1.000001" /**/}, "" /*-------------------------------*/, Vertex{vNullIndex, floatgeom.Point3{0.999999, -1, -1.000001}}},
	{stringList{"x", "-1.000000", "-1.000001" /*-------*/}, "unable to parse X coordinate" /*---*/, Vertex{vNullIndex, floatgeom.Point3{0, 0, 0}}},
	{stringList{"1.000000", "y", "-1.000001" /*--------*/}, "unable to parse Y coordinate" /*---*/, Vertex{vNullIndex, floatgeom.Point3{1, 0, 0}}},
	{stringList{"1.000000", "1", "z" /*----------------*/}, "unable to parse Z coordinate" /*---*/, Vertex{vNullIndex, floatgeom.Point3{1, 1, 0}}},
}

func TestReadVertex(t *testing.T) {

	for _, test := range vertexReadTests {
		name := fmt.Sprintf("parseVertex(%v)", test.Items)
		t.Run(name, func(t *testing.T) {

			v, err := parseVertex(test.Items)

			failed := false
			failed = failed || (test.Error == "" && err != nil)
			failed = failed || (err != nil && test.Error != err.Error())
			failed = failed || v.Point3 != test.Vertex.Point3

			if failed {
				t.Errorf("%v, '%v', expected %v, '%v'", v, err, test.Vertex, test.Error)
			}
		})
	}
}

var vertexWriteTests = []struct {
	Vertex Vertex
	Output string
	Error  string
}{
	{Vertex{vNullIndex, floatgeom.Point3{1, 1, 1}}, "1.000000 1.000000 1.000000", ""},
	{Vertex{vNullIndex, floatgeom.Point3{-1, 1, 1}}, "-1.000000 1.000000 1.000000", ""},
	{Vertex{vNullIndex, floatgeom.Point3{-1.000001, 0.999999, 1}}, "-1.000001 0.999999 1.000000", ""},
}

func TestWriteVertex(t *testing.T) {

	for _, test := range vertexWriteTests {
		name := fmt.Sprintf("writeVertex(%v, wr)", test.Vertex)
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			err := writeVertex(&test.Vertex, &buf)
			body := string(buf.Bytes())

			failed := false
			failed = failed || (test.Error == "" && err != nil)
			failed = failed || (err != nil && test.Error != err.Error())
			failed = failed || (test.Output != body)

			if failed {
				t.Errorf("got '%v', '%v', expected '%v', '%v'", body, err, test.Output, test.Error)
			}
		})
	}

}
