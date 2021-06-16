package sliceutil_test

import (
	"sort"
	"testing"

	"github.com/airdb/sailor/sliceutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReverse(t *testing.T) {
	Convey("Given array with few string value", t, func() {
		ss := []string{"a", "b", "c"}

		Convey("Then reverse the array", func() {
			sliceutil.Reverse(ss)
			So(ss, ShouldResemble, []string{"c", "b", "a"})
		})
	})
}

func TestStringsRemove(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		ss := []string{"a", "b", "c"}

		ns := sliceutil.StringsRemove(ss, "b")

		Convey("Then remove an item", func() {
			So(ns, ShouldContain, "c")
			So(ns, ShouldContain, "a")
			So(ns, ShouldNotContain, "b")
		})
	})
}

func TestTrimStrings(t *testing.T) {
	Convey("Given array with few string value", t, func() {
		ss := []string{" a", "b ", " c "}

		Convey("Then trim space", func() {
			bb := sliceutil.TrimStrings(ss)
			So(bb, ShouldResemble, []string{"a", "b", "c"})
		})
	})

	Convey("Given array with few string value", t, func() {
		ss := []string{",a", "b.", ",.c,"}

		Convey("Then trim characters", func() {
			bb := sliceutil.TrimStrings(ss, ",.")
			So(bb, ShouldResemble, []string{"a", "b", "c"})
		})
	})
}

func TestStringsToInts(t *testing.T) {
	Convey("Given array with few number string value", t, func() {
		ss := []string{"1", "2"}

		Convey("Then convert number string to number success", func() {
			ints, err := sliceutil.StringsToInts(ss)
			So(err, ShouldBeNil)
			So(ints, ShouldResemble, []int{1, 2})
		})
	})

	Convey("Given array with few string value", t, func() {
		ss := []string{"a", "b"}

		Convey("Then convert string to int fail", func() {
			_, err := sliceutil.StringsToInts(ss)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestIntersection(t *testing.T) {
	Convey("Given 2 array with few string value", t, func() {
		slice1 := []string{"a", "b"}
		slice2 := []string{"b", "c"}

		Convey("Then get the intersection value", func() {
			sections := sliceutil.Intersection(slice1, slice2)

			So(sections, ShouldResemble, []string{"b"})
		})
	})
}

func TestToSet(t *testing.T) {
	Convey("Given array with duplicate string value", t, func() {
		slice := []string{"a", "b", "b", "c"}
		result := []string{"a", "b", "c"}

		Convey("Then convert array the set", func() {
			sections := sliceutil.ToSet(slice)
			sort.SliceStable(sections, func(i, j int) bool {
				return sections[i] < sections[j]
			})

			So(sections, ShouldResemble, result)
		})
	})
}
