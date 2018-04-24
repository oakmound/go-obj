package obj

import "fmt"

// shared test code

type stringList []string

func (sl stringList) ToByteList() (arr [][]byte) {
	for _, s := range sl {
		arr = append(arr, []byte(s))
	}
	return
}

func comparePoints(p1 *Point, p2 *Point) bool {

	if p1 == p2 {
		return true
	}

	if (p1 != nil && p2 == nil) || (p1 == nil && p2 != nil) {
		return false
	}

	eq := true
	eq = eq && compareVertices(p1.Vertex, p2.Vertex)
	eq = eq && compareNormals(p1.Normal, p2.Normal)
	eq = eq && compareTextureCoords(p1.Texture, p2.Texture)

	return eq
}

func compareVertices(v1 *Vertex, v2 *Vertex) bool {
	if v1 == nil && v2 != nil || v1 != nil && v2 == nil {
		return false
	}
	return v1 == nil && v2 == nil || v1.Index == v2.Index && v1.Point3 == v2.Point3
}

func compareNormals(n1 *Normal, n2 *Normal) bool {
	if n1 == nil && n2 != nil || n1 != nil && n2 == nil {
		return false
	}
	return n1 == nil && n2 == nil || n1.Index == n2.Index && n1.Point3 == n2.Point3
}

func compareTextureCoords(t1 *TextureCoord, t2 *TextureCoord) bool {
	if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) {
		return false
	}
	return t1 == nil && t2 == nil || t1.Index == t2.Index && t1.U == t2.U && t1.V == t2.V && t1.W == t2.W
}

func compareErrors(err error, expected string) bool {
	if expected == "" && err != nil {
		return false
	} else if err != nil && expected != err.Error() {
		return false
	}
	return true
}

func (v *Vertex) String() string {
	if v == nil {
		return "Vertex{nil}"
	}

	return fmt.Sprintf("Vertex{%d %f %f %f}", v.Index, v.X(), v.Y(), v.Z())
}

func (n *Normal) String() string {
	if n == nil {
		return "Normal{nil}"
	}

	return fmt.Sprintf("Normal{%d %f %f %f}", n.Index, n.X(), n.Y(), n.Z())
}

func (t *TextureCoord) String() string {
	if t == nil {
		return "TextureCoord{nil}"
	}

	return fmt.Sprintf("TextureCoord{%d %f %f %f}", t.Index, t.U, t.V, t.W)
}
