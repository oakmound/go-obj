package obj

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/oakmound/oak/alg/floatgeom"
)

// Vertex represents a OBJ Vertex
type Vertex struct {
	Index int64
	floatgeom.Point3
}

func parseVertex(items []string) (v Vertex, err error) {
	if len(items) != 3 {
		err = errors.New("item length is incorrect")
		return
	}

	var x, y, z float64
	//TODO: verify each field, merge errors
	if x, err = strconv.ParseFloat(items[0], 64); err != nil {
		err = errors.New("unable to parse X coordinate")
		return
	}
	if y, err = strconv.ParseFloat(items[1], 64); err != nil {
		err = errors.New("unable to parse Y coordinate")
		return
	}
	if z, err = strconv.ParseFloat(items[2], 64); err != nil {
		err = errors.New("unable to parse Z coordinate")
		return
	}

	v.Point3 = floatgeom.Point3{x, y, z}

	return
}

func writeVertex(v *Vertex, w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf("%f %f %f", v.X, v.Y, v.Z)))
	return err
}
