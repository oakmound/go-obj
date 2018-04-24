package obj

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/oakmound/oak/alg/floatgeom"
)

// A Normal is a vertex normal
type Normal struct {
	Index int64
	floatgeom.Point3
}

func parseNormal(items []string) (n Normal, err error) {
	if len(items) != 3 {
		err = errors.New("item length is incorrect")
		return
	}

	//TODO: check all, merge error types

	if n.Point3[0], err = strconv.ParseFloat(items[0], 64); err != nil {
		err = errors.New("unable to parse X coordinate")
		return
	}
	if n.Point3[1], err = strconv.ParseFloat(items[1], 64); err != nil {
		err = errors.New("unable to parse Y coordinate")
		return
	}
	if n.Point3[2], err = strconv.ParseFloat(items[2], 64); err != nil {
		err = errors.New("unable to parse Z coordinate")
		return
	}

	return
}

func writeNormal(n *Normal, w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf("%0.4f %0.4f %0.4f", n.X(), n.Y(), n.Z())))
	return err
}
