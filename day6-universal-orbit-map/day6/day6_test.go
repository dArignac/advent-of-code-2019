package day6

import "testing"

import "github.com/stretchr/testify/assert"

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
