package pattern

import (
	"reflect"
	"testing"
)

func TestMaterialIterator(t *testing.T) {
	p1 := plan{walls: matHedge}
	p2 := plan{walls: matHedge}
	p3 := plan{walls: matBrick}
	p4 := plan{walls: matCorn}
	p5 := plan{walls: matHedge}
	p6 := plan{walls: matCorn}

	plans := []*plan{&p1, &p2, &p3, &p4, &p5, &p6}

	var c materialCollection

	itr := c.createIterator(plans)

	expected := []material{matHedge, matHedge, matHedge, matCorn, matCorn, matBrick}

	var got []material
	for i := 0; i < len(expected); i++ {
		got = append(got, itr.next().walls)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("got %v but expected %v", got, expected)
	}
}
