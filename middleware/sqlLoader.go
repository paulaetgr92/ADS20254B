package middleware

import (
	"os"
	"strings"
)

func LoadNamedQueries(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	queries := make(map[string]string)
	blocks := strings.Split(string(data), "-- name:")
	for _, block := range blocks[1:] {
		lines := strings.SplitN(block, "\n", 2)
		name := strings.TrimSpace(lines[0])
		query := strings.TrimSpace(lines[1])
		queries[name] = query
	}
	return queries, nil
}
