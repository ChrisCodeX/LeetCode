package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 1 {
		return strs[0]
	}

	cm := compareTwoStrings(&strs[0], &strs[1])
	for i := 2; i < len(strs); i++ {
		cm = compareTwoStrings(&cm, &strs[i])
	}
	return cm
}

func compareTwoStrings(first *string, second *string) string {
	if *first == "" || *second == "" {
		return ""
	}

	arr := []byte{}

	var compare int

	if len(*first) < len(*second) {
		compare = len(*first)
	} else {
		compare = len(*second)
	}

	for i := 0; i < compare; i++ {
		if (*first)[i] == (*second)[i] {
			arr = append(arr, (*first)[i])
		} else {
			break
		}
	}

	s := string(arr)

	return s
}

func main() {
	strs := []string{"ab", "a"}
	lcp := longestCommonPrefix(strs)
	fmt.Println(lcp)
}
