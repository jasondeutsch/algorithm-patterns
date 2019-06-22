package test_helpers

import "fmt"

func FailString(got, expected interface{}) string {
	return fmt.Sprintf("\nFAIL\n💣💣💣💣💣💣💣💣💣💣💣💣\nGot:\n%v\nExpected:\n%v",got, expected)
}

func IntSliceEquals(a,b []int) bool {
	if len(a) != len (b) {
		return false
	}
	for i := range a {
		if a[i] != b[i]  {
			return false
		}
	}
	return true
}