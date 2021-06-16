package sliceutil

import (
	"strconv"
	"strings"
)

// Reverse string slice [site user info 0] -> [0 info user site].
func Reverse(slice []string) {
	ln := len(slice)

	for i := 0; i < ln/2; i++ {
		li := ln - i - 1
		// fmt.Println(i, "<=>", li)
		slice[i], slice[li] = slice[li], slice[i]
	}
}

// StringsRemove an value form an string slice.
func StringsRemove(slice []string, str string) []string {
	var ns []string

	for _, v := range slice {
		if v != str {
			ns = append(ns, v)
		}
	}

	return ns
}

// StringsToInts string slice to int slice.
func StringsToInts(slice []string) (ints []int, err error) {
	for _, str := range slice {
		iVal, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}

	return
}

// TrimStrings trim string slice item.
func TrimStrings(slice []string, cutSet ...string) (ns []string) {
	hasCutSet := len(cutSet) > 0 && cutSet[0] != ""

	for _, str := range slice {
		if hasCutSet {
			ns = append(ns, strings.Trim(str, cutSet[0]))
		} else {
			ns = append(ns, strings.TrimSpace(str))
		}
	}

	return
}

func Intersection(slice1 []string, slice2 []string) []string {
	var res []string

	hashMap := make(map[string]bool)
	for _, s := range slice1 {
		hashMap[s] = true
	}

	for _, s := range slice2 {
		if hashMap[s] {
			res = append(res, s)
		}
	}

	return res
}

func ToSet(items []string) (ret []string) {
	setMap := make(map[string]bool)
	for _, item := range items {
		setMap[item] = true
	}

	for item := range setMap {
		ret = append(ret, item)
	}

	return
}

func LastString(ss []string) string {
	return ss[len(ss)-1]
}

func LastStringWithSplit(str, delimiter string) string {
	ss := strings.Split(str, delimiter)

	return ss[len(ss)-1]
}
