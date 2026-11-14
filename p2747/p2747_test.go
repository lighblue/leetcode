package p2747

import (
	"encoding/json"
	"os"
	"slices"
	"testing"
)

func TestCountServers(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		logs    [][]int
		x       int
		queries []int
		want    []int
	}{
		{name: "example 1", n: 3, logs: [][]int{{1, 3}, {2, 6}, {1, 5}}, x: 5, queries: []int{10, 11}, want: []int{1, 2}},
		{name: "example 2", n: 3, logs: [][]int{{2, 4}, {2, 1}, {1, 2}, {3, 1}}, x: 2, queries: []int{3, 4}, want: []int{0, 1}},
		{name: "single log across queries", n: 6, logs: [][]int{{1, 21}}, x: 10, queries: []int{24, 35, 28, 26, 20, 25, 16, 31, 12}, want: []int{5, 6, 5, 5, 6, 5, 6, 5, 6}},
		{name: "unsorted logs and queries", n: 4, logs: [][]int{{2, 30}, {2, 5}, {3, 9}, {4, 21}}, x: 9, queries: []int{11, 28, 16, 18}, want: []int{2, 3, 3, 3}},
		{name: "multiple servers at same time", n: 6, logs: [][]int{{2, 24}, {1, 26}, {3, 12}, {5, 12}, {6, 12}, {1, 21}}, x: 3, queries: []int{13, 25}, want: []int{3, 5}},
		{name: "repeated servers and unsorted queries", n: 3, logs: [][]int{{1, 1}, {3, 28}, {3, 32}, {2, 22}, {2, 7}, {3, 30}, {3, 9}, {1, 18}, {2, 4}, {1, 19}}, x: 6, queries: []int{15, 33, 22, 29, 34, 14, 23}, want: []int{2, 2, 1, 2, 2, 2, 1}},
		{name: "query before first full window", n: 8, logs: [][]int{{5, 23}, {2, 23}, {8, 22}, {8, 24}, {1, 12}, {4, 7}, {5, 7}, {2, 8}}, x: 6, queries: []int{10, 19}, want: []int{5, 8}},
		{name: "single log at window boundary", n: 9, logs: [][]int{{1, 6}}, x: 1, queries: []int{7, 6, 13, 20}, want: []int{8, 8, 9, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countServers(tt.n, tt.logs, tt.x, tt.queries); !slices.Equal(got, tt.want) {
				t.Errorf("countServers(%d, %v, %d, %v) = %v, want %v", tt.n, tt.logs, tt.x, tt.queries, got, tt.want)
			}
		})
	}
}

func TestCountServersLargeFixture(t *testing.T) {
	const n, x = 100000, 50000
	logs, queries := loadCountServersFixture(t)
	checkCountServersFixture(t, n, x, logs, queries)
}

func TestCountServersSmallFixture(t *testing.T) {
	const n, x = 100000, 50000
	logs, queries := loadCountServersFixture(t)
	checkCountServersFixture(t, n, x, logs[:len(logs)/5], queries[:len(queries)/5])
}

func loadCountServersFixture(t *testing.T) ([][]int, []int) {
	t.Helper()
	logsData, err := os.ReadFile("logs")
	if err != nil {
		t.Fatal(err)
	}
	var logs [][]int
	if err := json.Unmarshal(logsData, &logs); err != nil {
		t.Fatal(err)
	}

	queriesData, err := os.ReadFile("queries")
	if err != nil {
		t.Fatal(err)
	}
	var queries []int
	if err := json.Unmarshal(queriesData, &queries); err != nil {
		t.Fatal(err)
	}

	if len(logs) != 100000 || len(queries) != 100000 {
		t.Fatalf("fixture sizes: logs=%d, queries=%d; want 100000 each", len(logs), len(queries))
	}
	for i := range logs {
		if len(logs[i]) != 2 || logs[i][0] != i+1 || logs[i][1] != i+1 || queries[i] != 50001+i {
			t.Fatalf("unexpected fixture data at index %d: log=%v, query=%d", i, logs[i], queries[i])
		}
	}
	return logs, queries
}

func checkCountServersFixture(t *testing.T, n, x int, logs [][]int, queries []int) {
	t.Helper()
	got := countServers(n, logs, x, queries)
	if len(got) != len(queries) {
		t.Fatalf("result length = %d, want %d", len(got), len(queries))
	}
	for i, query := range queries {
		first := query - x
		last := min(query, len(logs))
		active := max(0, last-first+1)
		want := n - active
		if got[i] != want {
			t.Fatalf("query %d: got %d, want %d", query, got[i], want)
		}
	}
}
