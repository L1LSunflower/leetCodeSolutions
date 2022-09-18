package main

import "fmt"

func main() {
	//strs := []string{"flower", "flow", "flight"}
	//fmt.Println(longestCommonPrefix(strs))
	//strs = []string{"lower", "flow", "light"}
	//fmt.Println(longestCommonPrefix(strs))
	//strs = []string{"west", "flist", "test"}
	//fmt.Println(longestCommonPrefix(strs))
	//strs = []string{"lower", "fow", "iwght"}
	//fmt.Println(longestCommonPrefix(strs))
	//strs = []string{""}
	//fmt.Println(longestCommonPrefix(strs))
	strs := []string{"a"}
	fmt.Println(longestCommonPrefix(strs))
	//strs = []string{}
	//fmt.Println(longestCommonPrefix(strs))
	//strs = []string{"dog", "racecar", "car"}
	//fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs {
		if len(str) < len(prefix) {
			prefix = str
		}
	}
	fmt.Println(prefix)
	for _, str := range strs {
		prefix = findPrefixInString(str, prefix)
	}
	fmt.Println(prefix)
	return prefix
}

func findPrefixInString(str string, prefix string) string {
	for i := 0; i < len(str)-len(prefix); i++ {
		//fmt.Println(str, prefix, str[i:i+len(prefix)])
		if str[i:i+len(prefix)] == prefix {
			return prefix
		}
	}

	if len(findPrefixInString(str, prefix[:len(prefix)-1])) > len(findPrefixInString(str, prefix[1:])) {
		return findPrefixInString(str, prefix[:len(prefix)-1])
	} else {
		return findPrefixInString(str, prefix[1:])
	}
}
