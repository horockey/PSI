package main

import "fmt"

func main() {
	s1 := "accjs468invoiwepv56a3ekb7bmwb"
	s2 := "fivabnoifpemwouivoiwepv5sjdnvuwspionvwsv"

	longestSub := getLongestSusequence(s1, s2)
	fmt.Printf("%s\nlength: %d\n", longestSub, len(longestSub))
}

func getLongestSusequence(s1, s2 string) string {
	best := make([][]string, len(s1))
	for idx := 0; idx < len(s1); idx++ {
		best[idx] = make([]string, len(s2))
	}

	for n1 := 1; n1 < len(s1); n1++ {
		for n2 := 1; n2 < len(s2); n2++ {
			newRes := best[n1-1][n2]
			if len(best[n1][n2-1]) > len(newRes) {
				newRes = best[n1][n2-1]
			}

			if s1[n1] == s2[n2] {
				newRes = best[n1-1][n2-1] + string(s1[n1])
			}

			if len(newRes) > len(best[n1][n2]) {
				best[n1][n2] = newRes
			}
		}
	}

	return best[len(best)-1][len(best[0])-1]
}
