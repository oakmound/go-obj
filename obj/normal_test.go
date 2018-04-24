package obj

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/oakmound/oak/alg/floatgeom"
)

var nNullIndex = int64(-1)

var normalReadTests = []struct {
	Items  stringList
	Error  string
	Normal Normal
}{
	{stringList{"1", "1", "1" /*-----------------------*/}, "" /*-------------------------------*/, Normal{nNullIndex, floatgeom.Point3{1, 1, 1}}},
	{stringList{"1", "1" /*----------------------------*/}, "item length is incorrect" /**/, Normal{nNullIndex, floatgeom.Point3{0, 0, 0}}},
	{stringList{"1.0000", "-1.0000", "-1.0000" /*------*/}, "" /*-------------------------------*/, Normal{nNullIndex, floatgeom.Point3{1, -1, -1}}},
	{stringList{"0.9999", "-1.0000", "-1.0001" /*------*/}, "" /*-------------------------------*/, Normal{nNullIndex, floatgeom.Point3{0.9999, -1, -1.0001}}},
	{stringList{"x", "-1.000000", "-1.000001" /*-------*/}, "unable to parse X coordinate" /*---*/, Normal{nNullIndex, floatgeom.Point3{0, 0, 0}}},
	{stringList{"1.0000", "y", "-1.0001" /*------------*/}, "unable to parse Y coordinate" /*---*/, Normal{nNullIndex, floatgeom.Point3{1, 0, 0}}},
	{stringList{"1.0000", "1", "z" /*------------------*/}, "unable to parse Z coordinate" /*---*/, Normal{nNullIndex, floatgeom.Point3{1, 1, 0}}},
}

func TestReadNormal(t *testing.T) {

	for _, test := range normalReadTests {

		name := fmt.Sprintf("parseNormal(%s)", test.Items)
		t.Run(name, func(t *testing.T) {
			n, err := parseNormal(test.Items)

			failed := false
			failed = failed || test.Error == "" && err != nil
			failed = failed || err != nil && test.Error != err.Error()
			failed = failed || test.Normal.Point3 != n.Point3

			if failed {
				t.Errorf("got %v, '%v', expected %v, '%v'", n, err, test.Normal, test.Error)
			}
		})
	}
}

var normalWriteTests = []struct {
	Normal Normal
	Output string
	Error  string
}{
	{Normal{nNullIndex, floatgeom.Point3{1, 1, 1}}, "1.0000 1.0000 1.0000", ""},
	{Normal{nNullIndex, floatgeom.Point3{-1, 1, 1}}, "-1.0000 1.0000 1.0000", ""},
	{Normal{nNullIndex, floatgeom.Point3{-1.0001, 0.9999, 1}}, "-1.0001 0.9999 1.0000", ""},
}

func TestWriteNormal(t *testing.T) {

	for _, test := range normalWriteTests {

		name := fmt.Sprintf("writeNormal(%v, wr)", test.Normal)
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			err := writeNormal(&test.Normal, &buf)
			body := string(buf.Bytes())

			failed := false
			failed = failed || test.Error == "" && err != nil
			failed = failed || err != nil && test.Error != err.Error()
			failed = failed || test.Output != body

			if failed {
				t.Errorf("'%v', '%v', expected '%v', '%v'", body, err, test.Output, test.Error)
			}
		})
	}

}
