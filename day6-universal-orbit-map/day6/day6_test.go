package day6

import "testing"

import "github.com/stretchr/testify/assert"

func TestInsertIntoTreeSimple(t *testing.T) {
	root := Node{name: "COM", parent: nil, children: nil}
	err := insertIntoTree(&root, "COM)AAA")
	assert.Nil(t, err)
	assert.Equal(t, len(root.children), 1)
	assert.Equal(t, "AAA", (*root.children[0]).name)
	assert.Equal(t, root, *(*root.children[0]).parent)
}

func TestInsertIntoTreeAdvanced(t *testing.T) {
	// create a more complex tree
	// COM - AAA - BBB - CCC
	//           - DDD - EEE
	root := Node{name: "COM", parent: nil, children: nil}
	aaa := Node{name: "AAA", parent: &root, children: nil}
	bbb := Node{name: "BBB", parent: &aaa, children: nil}
	ccc := Node{name: "CCC", parent: &bbb, children: nil}
	ddd := Node{name: "DDD", parent: &aaa, children: nil}
	eee := Node{name: "EEE", parent: &ddd, children: nil}
	root.children = []*Node{&aaa}
	aaa.children = []*Node{&bbb, &ddd}
	bbb.children = []*Node{&ccc}
	ddd.children = []*Node{&eee}

	// test insertions
	// AAA - FFF
	err := insertIntoTree(&root, "AAA)FFF")
	assert.Nil(t, err)

	fff := *findNodeByName(&root, "FFF")
	assert.NotNil(t, fff)
	assert.Equal(t, fff.parent, &aaa)

	assert.Equal(t, 3, len(aaa.children))
	assert.Contains(t, aaa.children, &fff)

	// BBB - GGG
	err = insertIntoTree(&root, "BBB)GGG")
	assert.Nil(t, err)

	ggg := *findNodeByName(&root, "GGG")
	assert.NotNil(t, ggg)
	assert.Equal(t, ggg.parent, &bbb)

	assert.Equal(t, 2, len(bbb.children))
	assert.Contains(t, bbb.children, &ggg)

	// EEE - HHH
	err = insertIntoTree(&root, "EEE)HHH")
	assert.Nil(t, err)

	hhh := *findNodeByName(&root, "HHH")
	assert.NotNil(t, hhh)
	assert.Equal(t, hhh.parent, &eee)

	assert.Equal(t, 1, len(eee.children))
	assert.Contains(t, eee.children, &hhh)
}

func TestFindNodeByNameSimple(t *testing.T) {
	// COM - AAA - BBB - CCC
	root := Node{name: "COM", parent: nil, children: nil}
	aaa := Node{name: "AAA", parent: &root, children: nil}
	bbb := Node{name: "BBB", parent: &aaa, children: nil}
	ccc := Node{name: "CCC", parent: &bbb, children: nil}
	root.children = []*Node{&aaa}
	aaa.children = []*Node{&bbb}
	bbb.children = []*Node{&ccc}

	assert.Equal(t, root, *findNodeByName(&root, "COM"))
	assert.Equal(t, aaa, *findNodeByName(&root, "AAA"))
	assert.Equal(t, bbb, *findNodeByName(&root, "BBB"))
	assert.Equal(t, ccc, *findNodeByName(&root, "CCC"))
}

func TestFindNodeByNameForked(t *testing.T) {
	// COM - AAA - BBB - CCC
	//           - DDD - EEE
	root := Node{name: "COM", parent: nil, children: nil}
	aaa := Node{name: "AAA", parent: &root, children: nil}
	bbb := Node{name: "BBB", parent: &aaa, children: nil}
	ccc := Node{name: "CCC", parent: &bbb, children: nil}
	ddd := Node{name: "DDD", parent: &aaa, children: nil}
	eee := Node{name: "EEE", parent: &ddd, children: nil}
	root.children = []*Node{&aaa}
	aaa.children = []*Node{&bbb, &ddd}
	bbb.children = []*Node{&ccc}
	ddd.children = []*Node{&eee}

	assert.Equal(t, root, *findNodeByName(&root, "COM"))
	assert.Equal(t, aaa, *findNodeByName(&root, "AAA"))
	assert.Equal(t, bbb, *findNodeByName(&root, "BBB"))
	assert.Equal(t, ccc, *findNodeByName(&root, "CCC"))
	assert.Equal(t, ddd, *findNodeByName(&root, "DDD"))
	assert.Equal(t, eee, *findNodeByName(&root, "EEE"))
}
