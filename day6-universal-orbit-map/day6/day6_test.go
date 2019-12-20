package day6

import "testing"

import "github.com/stretchr/testify/assert"

func TestBulkInsertIntoTree(t *testing.T) {
	// input from day 6 task
	// 	       G - H       J - K - L
	// 	      /           /
	// COM - B - C - D - E - F
	//               \
	//                I
	// inputSorted := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	inputFuzzy := []string{"B)C", "J)K", "C)D", "D)E", "E)F", "B)G", "COM)B", "G)H", "D)I", "E)J", "K)L"}
	root := BulkCreateTree(inputFuzzy)

	assert.NotNil(t, root)
	assert.NotEqual(t, Node{}, root)
	assert.Equal(t, "COM", root.Name)
	assert.Equal(t, 1, len(root.Children))

	// B
	b := root.Children[0]
	assert.Equal(t, "B", b.Name)
	assert.Equal(t, 2, len(b.Children))
	assert.Equal(t, *b.Parent, root)

	var c Node
	var g Node
	if (*b.Children[0]).Name == "C" {
		c = *b.Children[0]
		g = *b.Children[1]
	} else {
		c = *b.Children[1]
		g = *b.Children[0]
	}

	// C
	assert.Equal(t, c.Name, "C")
	assert.Equal(t, 1, len(c.Children))
	assert.Equal(t, *c.Parent, b)
	d := *c.Children[0]

	// D
	assert.Equal(t, "D", d.Name)
	assert.Equal(t, 2, len(d.Children))
	assert.Equal(t, *d.Parent, c)

	var e Node
	var i Node
	if (*d.Children[0]).Name == "E" {
		e = *d.Children[0]
		i = *d.Children[1]
	} else {
		e = *d.Children[1]
		i = *d.Children[0]
	}

	// E
	assert.Equal(t, "E", e.Name)
	assert.Equal(t, 2, len(e.Children))
	assert.Equal(t, *e.Parent, d)

	var j Node
	var f Node
	if (*e.Children[0]).Name == "J" {
		j = *e.Children[0]
		f = *e.Children[1]
	} else {
		j = *e.Children[1]
		f = *e.Children[0]
	}

	// F
	assert.Equal(t, "F", f.Name)
	assert.Equal(t, 0, len(f.Children))
	assert.Equal(t, *f.Parent, e)

	// G
	assert.Equal(t, "G", g.Name)
	assert.Equal(t, 1, len(g.Children))
	assert.Equal(t, *g.Parent, b)
	h := *g.Children[0]

	// H
	assert.Equal(t, "H", h.Name)
	assert.Equal(t, 0, len(h.Children))
	assert.Equal(t, *h.Parent, g)

	// I
	assert.Equal(t, "I", i.Name)
	assert.Equal(t, 0, len(i.Children))
	assert.Equal(t, *i.Parent, d)

	// J
	assert.Equal(t, "J", j.Name)
	assert.Equal(t, 1, len(j.Children))
	assert.Equal(t, *j.Parent, e)
	k := *j.Children[0]

	// K
	assert.Equal(t, "K", k.Name)
	assert.Equal(t, 1, len(k.Children))
	assert.Equal(t, *k.Parent, j)
	l := *k.Children[0]

	// L
	assert.Equal(t, "L", l.Name)
	assert.Equal(t, 0, len(l.Children))
	assert.Equal(t, *l.Parent, k)
}

func TestInsertIntoTreeSimple(t *testing.T) {
	root := Node{Name: "COM", Parent: nil, Children: nil}
	err := insertIntoTree(&root, "COM)AAA")
	assert.Nil(t, err)
	assert.Equal(t, len(root.Children), 1)
	assert.Equal(t, "AAA", (*root.Children[0]).Name)
	assert.Equal(t, root, *(*root.Children[0]).Parent)
}

func TestInsertIntoTreeAdvanced(t *testing.T) {
	// create a more complex tree
	// COM - AAA - BBB - CCC
	//           - DDD - EEE
	root := Node{Name: "COM", Parent: nil, Children: nil}
	aaa := Node{Name: "AAA", Parent: &root, Children: nil}
	bbb := Node{Name: "BBB", Parent: &aaa, Children: nil}
	ccc := Node{Name: "CCC", Parent: &bbb, Children: nil}
	ddd := Node{Name: "DDD", Parent: &aaa, Children: nil}
	eee := Node{Name: "EEE", Parent: &ddd, Children: nil}
	root.Children = []*Node{&aaa}
	aaa.Children = []*Node{&bbb, &ddd}
	bbb.Children = []*Node{&ccc}
	ddd.Children = []*Node{&eee}

	// test insertions
	// AAA - FFF
	err := insertIntoTree(&root, "AAA)FFF")
	assert.Nil(t, err)

	fff := *findNodeByName(&root, "FFF")
	assert.NotNil(t, fff)
	assert.Equal(t, fff.Parent, &aaa)

	assert.Equal(t, 3, len(aaa.Children))
	assert.Contains(t, aaa.Children, &fff)

	// BBB - GGG
	err = insertIntoTree(&root, "BBB)GGG")
	assert.Nil(t, err)

	ggg := *findNodeByName(&root, "GGG")
	assert.NotNil(t, ggg)
	assert.Equal(t, ggg.Parent, &bbb)

	assert.Equal(t, 2, len(bbb.Children))
	assert.Contains(t, bbb.Children, &ggg)

	// EEE - HHH
	err = insertIntoTree(&root, "EEE)HHH")
	assert.Nil(t, err)

	hhh := *findNodeByName(&root, "HHH")
	assert.NotNil(t, hhh)
	assert.Equal(t, hhh.Parent, &eee)

	assert.Equal(t, 1, len(eee.Children))
	assert.Contains(t, eee.Children, &hhh)
}

func TestFindNodeByNameSimple(t *testing.T) {
	// COM - AAA - BBB - CCC
	root := Node{Name: "COM", Parent: nil, Children: nil}
	aaa := Node{Name: "AAA", Parent: &root, Children: nil}
	bbb := Node{Name: "BBB", Parent: &aaa, Children: nil}
	ccc := Node{Name: "CCC", Parent: &bbb, Children: nil}
	root.Children = []*Node{&aaa}
	aaa.Children = []*Node{&bbb}
	bbb.Children = []*Node{&ccc}

	assert.Equal(t, root, *findNodeByName(&root, "COM"))
	assert.Equal(t, aaa, *findNodeByName(&root, "AAA"))
	assert.Equal(t, bbb, *findNodeByName(&root, "BBB"))
	assert.Equal(t, ccc, *findNodeByName(&root, "CCC"))
}

func TestFindNodeByNameForked(t *testing.T) {
	// COM - AAA - BBB - CCC
	//           - DDD - EEE
	root := Node{Name: "COM", Parent: nil, Children: nil}
	aaa := Node{Name: "AAA", Parent: &root, Children: nil}
	bbb := Node{Name: "BBB", Parent: &aaa, Children: nil}
	ccc := Node{Name: "CCC", Parent: &bbb, Children: nil}
	ddd := Node{Name: "DDD", Parent: &aaa, Children: nil}
	eee := Node{Name: "EEE", Parent: &ddd, Children: nil}
	root.Children = []*Node{&aaa}
	aaa.Children = []*Node{&bbb, &ddd}
	bbb.Children = []*Node{&ccc}
	ddd.Children = []*Node{&eee}

	assert.Equal(t, root, *findNodeByName(&root, "COM"))
	assert.Equal(t, aaa, *findNodeByName(&root, "AAA"))
	assert.Equal(t, bbb, *findNodeByName(&root, "BBB"))
	assert.Equal(t, ccc, *findNodeByName(&root, "CCC"))
	assert.Equal(t, ddd, *findNodeByName(&root, "DDD"))
	assert.Equal(t, eee, *findNodeByName(&root, "EEE"))
}
