package p2914

// Submission: 2026.11.18 | Runtime: 0 ms (beats 100%) | Memory: 7 MB (beats 37.50%)

func minChanges(s string) (cnt int) {
	for i := 0; i < len(s); i += 2 {
		cnt += int(s[i] ^ s[i+1])
	}
	return
}
