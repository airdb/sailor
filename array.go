package sailor

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

var _ = fmt.Println

func StringRemoveDuplicatesAndEmpty(srcStr, delimit string) (retStr string) {
	srcArry := strings.Split(srcStr, delimit)
	retStr = strings.Join(RemoveDuplicatesAndEmpty(srcArry), delimit)
	return
}

func RemoveDuplicatesAndEmpty(array []string) (retArray []string) {
	sort.Strings(array)
	// array_len := len(array)
	for i := range array {
		// for i:=0; i < array_len; i++{
		if (i > 0 && array[i-1] == array[i]) || len(array[i]) == 0 {
			continue
		}
		retArray = append(retArray, array[i])
	}
	return
}

func Diff(slice1 []string, slice2 []string) []string {
	var diff []string

	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, s1)
		}
	}
	// Swap the slices, only if it was the first loop

	return diff
}

func DiffAll(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	// bugssss:   container null slice!!!!!
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func Contains(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not contain")
}

func IsContain(obj interface{}, target interface{}) (flag bool) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				flag = true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			flag = true
		}
	}
	return
}
