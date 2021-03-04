package main

func main() {
	minDistance("horse", "ros")
}

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	memo := make([][]int, l1)
	for i := range memo {
		memo[i] = make([]int, l2)
	}
	return subMinDist(word1, 0, word2, 0, memo)
}

func subMinDist(word1 string, index1 int, word2 string, index2 int, memo [][]int) int {
	if index1 >= len(word1) {
		return len(word2) - index2
	}
	if index2 >= len(word2) {
		return len(word1) - index1
	}
	if memo[index1][index2] > 0 {
		return memo[index1][index2]
	}
	if word1[index1] == word2[index2] {
		r := subMinDist(word1, index1+1, word2, index2+1, memo)
		memo[index1][index2] = r
		return r
	}
	replace := subMinDist(word1, index1+1, word2, index2+1, memo) + 1
	add := subMinDist(word1, index1, word2, index2+1, memo) + 1
	del := subMinDist(word1, index1+1, word2, index2, memo) + 1
	r := min(replace, min(add, del))
	memo[index1][index2] = r
	return r
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
