package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFileContent(t *testing.T) {
	content, err := LoadFileContent()
	assert.Nil(t, err)
	assert.Equal(t, []string{"line0", "line1"}, content)
}

func TestConvertStringArrayToIntArray(t *testing.T) {
	values, err := ConvertStringArrayToIntArray([]string{})
	assert.Nil(t, err)
	assert.Equal(t, []int(nil), values)

	values, err = ConvertStringArrayToIntArray([]string{"1", "2", "576182"})
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 576182}, values)

	values, err = ConvertStringArrayToIntArray([]string{"1", "A", "576182"})
	assert.NotNil(t, err)
	assert.Nil(t, values)
}

func TestSplitStringToIntArray(t *testing.T) {
	r1, e1 := SplitStringToIntArray("0,1,2,3")
	assert.Nil(t, e1)
	assert.Equal(t, r1, []int{0, 1, 2, 3})

	r2, e2 := SplitStringToIntArray("7,129,18,44")
	assert.Nil(t, e2)
	assert.Equal(t, r2, []int{7, 129, 18, 44})

	r3, e3 := SplitStringToIntArray("")
	assert.NotNil(t, e3)
	assert.Equal(t, []int(nil), r3)
}
