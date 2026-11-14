package p2747

// Submission: 2026.11.14 | Runtime: 26 ms (beats 42.80%) | Memory: 23.7 MB (beats 26.27%)

import (
	"cmp"
	"slices"
)

func countServers(n int, logs [][]int, x int, queries []int) []int {
	slices.SortFunc(logs, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})
	SortedQueries := convertQueries(queries)

	minT, maxT := logs[0][1], logs[len(logs)-1][1]

	logs = slim(logs)

	results := make([]int, len(queries))
	hits := make([]int, n)
	lastLoopLeft, lastLoopRight, count := 0, 0, n
	for _, query := range SortedQueries {
		if query[0]-x > maxT || query[0] < minT {
			results[query[1]] = n
			continue
		}
		loopLeft := max(query[0]-x, minT) - minT
		loopRight := min(query[0], maxT) - minT + 1
		for _, log := range logs[max(lastLoopRight, loopLeft):loopRight] {
			for _, server := range log {
				hits[server-1]++
				if hits[server-1] == 1 {
					count--
				}
			}
		}
		for _, log := range logs[lastLoopLeft:min(lastLoopRight, loopLeft)] {
			for _, server := range log {
				hits[server-1]--
				if hits[server-1] == 0 {
					count++
				}
			}
		}
		lastLoopLeft, lastLoopRight = loopLeft, loopRight

		results[query[1]] = count
	}
	return results
}

func slim(logs [][]int) [][]int {
	minT := logs[0][1]
	maxT := logs[len(logs)-1][1]
	size := maxT - minT + 1
	newLogs := make([][]int, size)
	t := minT
	newLogs[0] = logs[0][:1]
	for _, log := range logs[1:] {
		if log[1] == t {
			newLogs[t-minT] = append(newLogs[t-minT], log[0])
			continue
		}

		t = log[1]
		newLogs[t-minT] = log[:1]
	}
	return newLogs
}

func convertQueries(queries []int) [][2]int {
	newQueries := make([][2]int, len(queries))
	for i := range queries {
		newQueries[i] = [2]int{queries[i], i}
	}

	slices.SortFunc(newQueries, func(a, b [2]int) int { return cmp.Compare(a[0], b[0]) })
	return newQueries
}
