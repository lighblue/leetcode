package p1189

// Submission: 2026.11.18 | Runtime: 0 ms (beats 100%) | Memory: 4.5 MB (beats 15.38%)

func maxNumberOfBalloons(text string) int {
	cnts := make([]int, 26)
	for _, c := range text {
		cnts[c-'a']++

	}
	return min(cnts['b'-'a'], cnts['a'-'a'], cnts['l'-'a']/2, cnts['o'-'a']/2, cnts['n'-'a'])
}
