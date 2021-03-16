package sliceutil_test

import (
	"fmt"
	"testing"

	"airdb.io/airdb/sailor/sliceutil"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	ss := []string{"a", "b", "c"}

	sliceutil.Reverse(ss)

	assert.Equal(t, []string{"c", "b", "a"}, ss)
}

func TestStringsRemove(t *testing.T) {
	ss := []string{"a", "b", "c"}

	ns := sliceutil.StringsRemove(ss, "b")
	assert.Contains(t, ns, "a")
	assert.NotContains(t, ns, "b")
}

func TestTrimStrings(t *testing.T) {
	is := assert.New(t)

	// TrimStrings
	ss := sliceutil.TrimStrings([]string{" a", "b ", " c "})
	is.Equal("[a b c]", fmt.Sprint(ss))
	ss = sliceutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",.")
	is.Equal("[a b c]", fmt.Sprint(ss))
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := sliceutil.StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = sliceutil.StringsToInts([]string{"a", "b"})
	is.Error(err)
}

func TestIntersection(t *testing.T) {
	is := assert.New(t)

	slice1 := []string{"a", "b"}
	slice2 := []string{"b", "c"}

	sections := sliceutil.Intersection(slice1, slice2)
	is.Equal([]string{"b"}, sections)
}

func TestToSet(t *testing.T) {
	is := assert.New(t)

	slice := []string{"a", "b", "b", "c"}
	result := []string{"a", "b", "c"}

	set := sliceutil.ToSet(slice)
	is.Equal(result, set)
}
