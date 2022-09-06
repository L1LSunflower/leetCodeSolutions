package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"lower", "flow", "light"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"west", "flist", "test"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"lower", "fow", "iwght"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	//result := ""
	prefix := strs[0]
	for _, str := range strs {
		if len(str) < len(prefix) {
			prefix = str
		}
	}
	for _, str := range strs {
		prefix = findPrefixInString(str, prefix)
	}
	return prefix
}

func findPrefixInString(str string, prefix string) string {
	for i := 0; i < len(str)-len(prefix); i++ {
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
