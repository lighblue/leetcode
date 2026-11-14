package p2747

import (
	"cmp"
	"slices"
)

func countServersV2(n int, logs [][]int, x int, queries []int) []int {
	slices.SortFunc(logs, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})
	SortedQueries := queries

	minT, maxT := logs[0][1], logs[len(logs)-1][1]

	logs = slim(logs)

	results := make([]int, len(queries))
	counts := count(logs, x, n)
	for index, query := range SortedQueries {
		if query-x > maxT || query < minT {
			results[index] = n
			continue
		}
		results[index] = n - counts[query-minT]
	}
	return results
}

func count(logs [][]int, x int, n int) []int {
	result := make([]int, len(logs)+x)
	hits := make([]int, n)
	count := 0
	i := 0
	for i = 0; i <= x && i < len(logs); i++ {
		for _, server := range logs[i] {
			hits[server-1]++
			if hits[server-1] == 1 {
				count++
			}
		}
		result[i] = count
	}

	if i <= x {
		for ; i <= x; i++ {
			result[i] = count
		}
	}

	for i := 1; i < len(logs)-x; i++ {
		for _, server := range logs[i+x] {
			hits[server-1]++
			if hits[server-1] == 1 {
				count++
			}
		}
		for _, server := range logs[i-1] {
			hits[server-1]--
			if hits[server-1] == 0 {
				count--
			}
		}
		result[i+x] = count
	}

	for i := max(len(logs)-x, 1); i < len(logs); i++ {
		for _, server := range logs[i-1] {
			hits[server-1]--
			if hits[server-1] == 0 {
				count--
			}
		}
		result[i+x] = count
	}
	return result
}
